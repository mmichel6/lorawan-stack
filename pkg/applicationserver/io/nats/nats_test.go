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

package nats_test

import (
	"context"
	"fmt"
	"testing"
	"time"

	nats_server "github.com/nats-io/nats-server/test"
	nats_client "github.com/nats-io/nats.go"
	"github.com/smartystreets/assertions"
	"go.thethings.network/lorawan-stack/pkg/applicationserver/io/mock"
	. "go.thethings.network/lorawan-stack/pkg/applicationserver/io/nats"
	"go.thethings.network/lorawan-stack/pkg/applicationserver/io/pubsub"
	"go.thethings.network/lorawan-stack/pkg/jsonpb"
	"go.thethings.network/lorawan-stack/pkg/log"
	"go.thethings.network/lorawan-stack/pkg/ttnpb"
	"go.thethings.network/lorawan-stack/pkg/unique"
	"go.thethings.network/lorawan-stack/pkg/util/test"
	"go.thethings.network/lorawan-stack/pkg/util/test/assertions/should"
)

var (
	registeredApplicationID = ttnpb.ApplicationIdentifiers{ApplicationID: "test-app"}
	registeredDeviceID      = ttnpb.EndDeviceIdentifiers{
		ApplicationIdentifiers: registeredApplicationID,
		DeviceID:               "test-device",
	}

	timeout = 100 * test.Delay
)

func TestTraffic(t *testing.T) {
	a := assertions.New(t)

	ctx := log.NewContext(test.Context(), test.GetLogger(t))
	ctx, cancelCtx := context.WithCancel(ctx)
	defer cancelCtx()

	as := mock.NewServer()
	natsServer := nats_server.RunDefaultServer()
	go natsServer.Start()
	defer natsServer.Shutdown()
	sub, err := Start(ctx, as, pubsub.JSON, nats_client.DefaultURL)
	a.So(err, should.BeNil)

	client, err := nats_client.Connect(nats_client.DefaultURL)
	a.So(err, should.BeNil)
	defer client.Close()

	t.Run("Upstream", func(t *testing.T) {
		for _, tc := range []struct {
			Topic   string
			Message *ttnpb.ApplicationUp
			OK      bool
		}{
			{
				Topic: ">",
				Message: &ttnpb.ApplicationUp{
					EndDeviceIdentifiers: registeredDeviceID,
					Up: &ttnpb.ApplicationUp_UplinkMessage{
						UplinkMessage: &ttnpb.ApplicationUplink{FRMPayload: []byte{0x1, 0x1, 0x1}},
					},
				},
				OK: true,
			},
			{
				Topic: fmt.Sprintf("v3.%v.devices.%v.up", unique.ID(ctx, registeredDeviceID.ApplicationIdentifiers), registeredDeviceID.DeviceID),
				Message: &ttnpb.ApplicationUp{
					EndDeviceIdentifiers: registeredDeviceID,
					Up: &ttnpb.ApplicationUp_UplinkMessage{
						UplinkMessage: &ttnpb.ApplicationUplink{FRMPayload: []byte{0x2, 0x2, 0x2}},
					},
				},
				OK: true,
			},
			{
				Topic: fmt.Sprintf("v3.%v.devices.%v.join", unique.ID(ctx, registeredDeviceID.ApplicationIdentifiers), registeredDeviceID.DeviceID),
				Message: &ttnpb.ApplicationUp{
					EndDeviceIdentifiers: registeredDeviceID,
					Up: &ttnpb.ApplicationUp_UplinkMessage{
						UplinkMessage: &ttnpb.ApplicationUplink{FRMPayload: []byte{0x3, 0x3, 0x3}},
					},
				},
				OK: false, // Invalid topic
			},
			{
				Topic: fmt.Sprintf("v3.%v.devices.%v.up", "invalid-application", "invalid-device"),
				Message: &ttnpb.ApplicationUp{
					EndDeviceIdentifiers: registeredDeviceID,
					Up: &ttnpb.ApplicationUp_UplinkMessage{
						UplinkMessage: &ttnpb.ApplicationUplink{FRMPayload: []byte{0x4, 0x4, 0x4}},
					},
				},
				OK: false, // Invalid application ID
			},
		} {
			t.Run(tc.Topic, func(t *testing.T) {
				a := assertions.New(t)

				upCh := make(chan *ttnpb.ApplicationUp)
				handler := func(msg *nats_client.Msg) {
					up := &ttnpb.ApplicationUp{}
					err := jsonpb.TTN().Unmarshal(msg.Data, up)
					a.So(err, should.BeNil)
					upCh <- up
				}

				if sub, err := client.Subscribe(tc.Topic, handler); !a.So(err, should.BeNil) {
					t.FailNow()
				} else {
					defer sub.Unsubscribe()
				}

				err := sub.SendUp(tc.Message)
				if !a.So(err, should.BeNil) {
					t.FailNow()
				}

				select {
				case up := <-upCh:
					if tc.OK {
						a.So(up, should.Resemble, tc.Message)
					} else {
						t.Fatalf("Expected no upstream message but have %v", up)
					}
				case <-time.After(timeout):
					if tc.OK {
						t.Fatal("Receive expected upstream timeout")
					}
				}
			})
		}
	})

	t.Run("Downstream", func(t *testing.T) {
		for _, tc := range []struct {
			Topic    string
			IDs      ttnpb.EndDeviceIdentifiers
			Message  *ttnpb.ApplicationDownlinks
			Expected []*ttnpb.ApplicationDownlink
		}{
			{
				Topic: fmt.Sprintf("v3.%v.devices.%v.down.push", unique.ID(ctx, registeredDeviceID.ApplicationIdentifiers), registeredDeviceID.DeviceID),
				IDs:   registeredDeviceID,
				Message: &ttnpb.ApplicationDownlinks{
					Downlinks: []*ttnpb.ApplicationDownlink{
						{
							FPort:      42,
							FRMPayload: []byte{0x1, 0x1, 0x1},
						},
					},
				},
				Expected: []*ttnpb.ApplicationDownlink{
					{
						FPort:      42,
						FRMPayload: []byte{0x1, 0x1, 0x1},
					},
				},
			},
			{
				Topic: fmt.Sprintf("v3.%v.devices.%v.down.replace", unique.ID(ctx, registeredDeviceID.ApplicationIdentifiers), registeredDeviceID.DeviceID),
				IDs:   registeredDeviceID,
				Message: &ttnpb.ApplicationDownlinks{
					Downlinks: []*ttnpb.ApplicationDownlink{
						{
							FPort:      42,
							FRMPayload: []byte{0x2, 0x2, 0x2},
						},
					},
				},
				Expected: []*ttnpb.ApplicationDownlink{
					{
						FPort:      42,
						FRMPayload: []byte{0x2, 0x2, 0x2},
					},
				},
			},
			{
				Topic: fmt.Sprintf("v3.%v.devices.%v.down.push", unique.ID(ctx, registeredDeviceID.ApplicationIdentifiers), "invalid-device"),
				IDs:   registeredDeviceID,
				Message: &ttnpb.ApplicationDownlinks{
					Downlinks: []*ttnpb.ApplicationDownlink{
						{
							FPort:      42,
							FRMPayload: []byte{0x3, 0x3, 0x3},
						},
					},
				},
				Expected: []*ttnpb.ApplicationDownlink{
					{
						FPort:      42,
						FRMPayload: []byte{0x2, 0x2, 0x2}, // Do not expect a change.
					},
				},
			},
			{
				Topic: fmt.Sprintf("v3.%v.devices.%v.down.push", "invalid-application", "invalid-device"),
				IDs:   registeredDeviceID,
				Message: &ttnpb.ApplicationDownlinks{
					Downlinks: []*ttnpb.ApplicationDownlink{
						{
							FPort:      42,
							FRMPayload: []byte{0x4, 0x4, 0x4},
						},
					},
				},
				Expected: []*ttnpb.ApplicationDownlink{
					{
						FPort:      42,
						FRMPayload: []byte{0x2, 0x2, 0x2}, // Do not expect a change.
					},
				},
			},
		} {
			tcok := t.Run(tc.Topic, func(t *testing.T) {
				a := assertions.New(t)
				buf, err := jsonpb.TTN().Marshal(tc.Message)
				a.So(err, should.BeNil)
				if err := client.Publish(tc.Topic, buf); !a.So(err, should.BeNil) {
					t.FailNow()
				}
				<-time.After(timeout)
				res, err := as.DownlinkQueueList(ctx, tc.IDs)
				a.So(err, should.BeNil)
				a.So(res, should.Resemble, tc.Expected)
			})
			if !tcok {
				t.FailNow()
			}
		}
	})
}
