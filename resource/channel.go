package resource

import (
	"github.com/WebmiRU/YudolePlatformPackages/channel"
)

type Channel struct {
	Payload *channel.Channel `json:"payload"`
}

type ChannelIndex struct {
	Payload map[string]*channel.Channel `json:"payload"`
}

type ChannelConfig struct {
	Type    string           `json:"type"`
	Payload *channel.Channel `json:"payload"`
}
