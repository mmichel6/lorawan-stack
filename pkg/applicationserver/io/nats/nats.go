// Copyright © 2019 The Things Network Foundation, The Things Industries B.V.
//
// Licensed under the Apache License, Version 2.0 (the "License");
// you may not use this file except in compliance with the License.
// You may obtain a copy of the License at
//
//     http://www.apache.org/licenses/LICENSE-2.0
//
// Unless required by applicable law or agreed to in writing, software
// distributed under the License is distributed on an "AS IS" BASIS,
// WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
// See the License for the specific language governing permissions and
// limitations under the License.

// Package nats implements the NATS frontend.
package nats

import (
	"context"
	"fmt"
	"strings"

	nats "github.com/nats-io/nats.go"
	"go.thethings.network/lorawan-stack/pkg/applicationserver/io"
	iopubsub "go.thethings.network/lorawan-stack/pkg/applicationserver/io/pubsub"
	"go.thethings.network/lorawan-stack/pkg/log"
	"go.thethings.network/lorawan-stack/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/pkg/unique"
)

const (
	wildcard         = "*"
	tailWildcard     = ">"
	subjectSeparator = "."
)

type srv struct {
	ctx    context.Context
	server io.Server
	format iopubsub.Format
	conn   *nats.Conn
	subs   []*nats.Subscription
}

// Start starts the nats frontend.
func Start(ctx context.Context, server io.Server, format iopubsub.Format, natsServer string) (sub *io.Subscription, err error) {
	ctx = log.NewContextWithField(ctx, "namespace", "applicationserver/io/nats")
	conn, err := nats.Connect(natsServer)
	if err != nil {
		return nil, err
	}
	s := &srv{
		ctx:    ctx,
		server: server,
		format: format,
		conn:   conn,
	}
	go func() {
		<-ctx.Done()
		for _, sub := range s.subs {
			sub.Unsubscribe()
		}
		s.conn.Close()
	}()

	sub = io.NewSubscription(s.ctx, "nats", nil)
	// Publish upstream
	go func() {
		logger := log.FromContext(s.ctx)
		for {
			select {
			case <-sub.Context().Done():
				logger.WithError(sub.Context().Err()).Debug("Done sending upstream messages")
				return
			case up := <-sub.Up():
				appID := up.ApplicationIdentifiers
				var subjectParts []string
				switch up.Up.(type) {
				case *ttnpb.ApplicationUp_UplinkMessage:
					subjectParts = s.format.UplinkTopic(unique.ID(sub.Context(), appID), up.DeviceID)
				case *ttnpb.ApplicationUp_JoinAccept:
					subjectParts = s.format.JoinAcceptTopic(unique.ID(sub.Context(), appID), up.DeviceID)
				case *ttnpb.ApplicationUp_DownlinkAck:
					subjectParts = s.format.DownlinkAckTopic(unique.ID(sub.Context(), appID), up.DeviceID)
				case *ttnpb.ApplicationUp_DownlinkNack:
					subjectParts = s.format.DownlinkNackTopic(unique.ID(sub.Context(), appID), up.DeviceID)
				case *ttnpb.ApplicationUp_DownlinkSent:
					subjectParts = s.format.DownlinkSentTopic(unique.ID(sub.Context(), appID), up.DeviceID)
				case *ttnpb.ApplicationUp_DownlinkFailed:
					subjectParts = s.format.DownlinkFailedTopic(unique.ID(sub.Context(), appID), up.DeviceID)
				case *ttnpb.ApplicationUp_DownlinkQueued:
					subjectParts = s.format.DownlinkQueuedTopic(unique.ID(sub.Context(), appID), up.DeviceID)
				case *ttnpb.ApplicationUp_LocationSolved:
					subjectParts = s.format.LocationSolvedTopic(unique.ID(sub.Context(), appID), up.DeviceID)
				}
				if subjectParts == nil {
					continue
				}
				buf, err := s.format.FromUp(up)
				if err != nil {
					log.WithError(err).Warn("Failed to marshal upstream message")
					continue
				}
				subject := createSubject(subjectParts)
				logger.Infof("Publish upstream message to subject %v", subject)
				s.conn.Publish(subject, buf)
			}
		}
	}()

	// Subscribe downstream
	for _, subjectParts := range [][]string{
		s.format.DownlinkPushTopic(wildcard, wildcard),
		s.format.DownlinkReplaceTopic(wildcard, wildcard),
	} {
		nsub, err := conn.Subscribe(createSubject(subjectParts), func(m *nats.Msg) {
			logger := log.FromContext(s.ctx).WithField("subject", m.Subject)
			subjectParts := createSubjectParts(m.Subject)
			var applicationUID, deviceID string
			var op func(io.Server, context.Context, ttnpb.EndDeviceIdentifiers, []*ttnpb.ApplicationDownlink) error
			switch {
			case s.format.IsDownlinkPushTopic(subjectParts):
				applicationUID, deviceID = s.format.ParseDownlinkPushTopic(subjectParts)
				op = io.Server.DownlinkQueuePush
			case s.format.IsDownlinkReplaceTopic(subjectParts):
				applicationUID, deviceID = s.format.ParseDownlinkReplaceTopic(subjectParts)
				op = io.Server.DownlinkQueueReplace
			default:
				panic(fmt.Errorf("invalid subject: %v", m.Subject))
			}
			items, err := s.format.ToDownlinks(m.Data)
			if err != nil {
				logger.WithError(err).Warn("Failed to decode downlink messages")
				return
			}
			appIDs, err := unique.ToApplicationID(applicationUID)
			if err != nil {
				logger.WithError(err).Warn("Failed to decode application ID")
				return
			}
			ids := ttnpb.EndDeviceIdentifiers{
				ApplicationIdentifiers: appIDs,
				DeviceID:               deviceID,
			}
			logger.WithFields(log.Fields(
				"device_uid", unique.ID(s.ctx, ids),
				"count", len(items.Downlinks),
			)).Debug("Handle downlink messages")
			if err := op(s.server, s.ctx, ids, items.Downlinks); err != nil {
				logger.WithError(err).Warn("Failed to handle downlink messages")
			}
		})
		if err != nil {
			return nil, err
		}
		s.subs = append(s.subs, nsub)
	}

	return sub, nil
}

func createSubject(parts []string) string {
	return strings.Join(parts, subjectSeparator)
}

func createSubjectParts(subject string) []string {
	return strings.Split(subject, subjectSeparator)
}
