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

package topics

const topicV3 = "v3"

// V3 is the generic PubSub topic layout.
type V3 struct{}

// AcceptedTopic implements the topics.Layout interface.
func (V3) AcceptedTopic(applicationUID string, requested []string) ([]string, bool) {
	return requested, true
}

// UplinkTopic implements the topics.Layout interface.
func (V3) UplinkTopic(applicationUID, deviceID string) []string {
	return []string{topicV3, applicationUID, "devices", deviceID, "up"}
}

// JoinAcceptTopic implements the topics.Layout interface.
func (V3) JoinAcceptTopic(applicationUID, deviceID string) []string {
	return []string{topicV3, applicationUID, "devices", deviceID, "join"}
}

// DownlinkAckTopic implements the topics.Layout interface.
func (V3) DownlinkAckTopic(applicationUID, deviceID string) []string {
	return []string{topicV3, applicationUID, "devices", deviceID, "down", "ack"}
}

// DownlinkNackTopic implements the topics.Layout interface.
func (V3) DownlinkNackTopic(applicationUID, deviceID string) []string {
	return []string{topicV3, applicationUID, "devices", deviceID, "down", "nack"}
}

// DownlinkSentTopic implements the topics.Layout interface.
func (V3) DownlinkSentTopic(applicationUID, deviceID string) []string {
	return []string{topicV3, applicationUID, "devices", deviceID, "down", "sent"}
}

// DownlinkFailedTopic implements the topics.Layout interface.
func (V3) DownlinkFailedTopic(applicationUID, deviceID string) []string {
	return []string{topicV3, applicationUID, "devices", deviceID, "down", "failed"}
}

// DownlinkQueuedTopic implements the topics.Layout interface.
func (V3) DownlinkQueuedTopic(applicationUID, deviceID string) []string {
	return []string{topicV3, applicationUID, "devices", deviceID, "down", "queued"}
}

// LocationSolvedTopic implements the topics.Layout interface.
func (V3) LocationSolvedTopic(applicationUID, deviceID string) []string {
	return []string{topicV3, applicationUID, "devices", deviceID, "location", "solved"}
}

// DownlinkPushTopic implements the topics.Layout interface.
func (V3) DownlinkPushTopic(applicationUID, deviceID string) []string {
	return []string{topicV3, applicationUID, "devices", deviceID, "down", "push"}
}

// IsDownlinkPushTopic implements the topics.Layout interface.
func (V3) IsDownlinkPushTopic(parts []string) bool {
	return len(parts) == 6 && parts[0] == topicV3 && parts[2] == "devices" && parts[4] == "down" && parts[5] == "push"
}

// ParseDownlinkPushTopic implements the topics.Layout interface.
func (V3) ParseDownlinkPushTopic(parts []string) (deviceID string) {
	return parts[3]
}

// DownlinkReplaceTopic implements the topics.Layout interface.
func (V3) DownlinkReplaceTopic(applicationUID, deviceID string) []string {
	return []string{topicV3, applicationUID, "devices", deviceID, "down", "replace"}
}

// IsDownlinkReplaceTopic implements the topics.Layout interface.
func (V3) IsDownlinkReplaceTopic(parts []string) bool {
	return len(parts) == 6 && parts[0] == topicV3 && parts[2] == "devices" && parts[4] == "down" && parts[5] == "replace"
}

// ParseDownlinkReplaceTopic implements the topics.Layout interface.
func (V3) ParseDownlinkReplaceTopic(parts []string) (deviceID string) {
	return parts[3]
}

// Default is the default topic layout.
var Default Layout = &V3{}
