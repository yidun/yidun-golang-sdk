package audio

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/auth"
	"github.com/yidun/yidun-golang-sdk/yidun/core/client"
)

type AudioClient struct {
	Client client.ExecuteClient
}

// NewAudioClient 初始化AudioClient对象
func NewAudioClient(profile *client.ClientProfile) *AudioClient {
	return &AudioClient{
		Client: client.NewDefaultClient(profile),
	}
}

// NewAudioClientWithAccessKey 初始化AudioClient对象
func NewAudioClientWithAccessKey(accessKeyId string, accessKeySecret string) *AudioClient {
	credential := auth.NewCredentials(accessKeyId, accessKeySecret)
	return NewAudioClient(client.NewClientProfile(credential))
}
