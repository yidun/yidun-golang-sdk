package liveaudio

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/auth"
	"github.com/yidun/yidun-golang-sdk/yidun/core/client"
)

type LiveAudioClient struct {
	Client client.ExecuteClient
}

func NewLiveAudioClient(profile *client.ClientProfile) *LiveAudioClient {
	return &LiveAudioClient{
		Client: client.NewDefaultClient(profile),
	}
}

func NewLiveAudioClientWithAccessKey(accessKeyId string, accessKeySecret string) *LiveAudioClient {
	credential := auth.NewCredentials(accessKeyId, accessKeySecret)
	return NewLiveAudioClient(client.NewClientProfile(credential))
}
