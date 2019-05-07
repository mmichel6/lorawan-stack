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

import (
	"github.com/TheThingsIndustries/mystique/pkg/topic"
	"go.thethings.network/lorawan-stack/pkg/applicationserver/io/pubsub/topics"
)

const topicV3 = "v3"

type v3 struct {
	topics.V3
}

func (v3) AcceptedTopic(applicationUID string, requested []string) ([]string, bool) {
	// Rewrite # to v3/uid/#
	if requested[0] == topic.Wildcard {
		return []string{topicV3, applicationUID, topic.Wildcard}, true
	}
	if requested[0] != topicV3 || len(requested) < 2 {
		return nil, false
	}
	switch requested[1] {
	case topic.Wildcard:
		// Rewrite v3/# to v3/uid/#
		return []string{topicV3, applicationUID, topic.Wildcard}, true
	case topic.PartWildcard:
		// Rewrite v3/+/... to v3/uid/...
		requested[1] = applicationUID
		return requested, true
	case applicationUID:
		return requested, true
	}
	return nil, false
}

// Default is the default topic layout.
var Default topics.Layout = &v3{}
