package video

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/auth"
	"github.com/yidun/yidun-golang-sdk/yidun/core/client"
)

type VideoClient struct {
	Client client.ExecuteClient
}

func NewVideoClient(profile *client.ClientProfile) *VideoClient {
	return &VideoClient{
		Client: client.NewDefaultClient(profile),
	}
}

func NewVideoClientWithAccessKey(accessKeyId string, accessKeySecret string) *VideoClient {
	credential := auth.NewCredentials(accessKeyId, accessKeySecret)
	return NewVideoClient(client.NewClientProfile(credential))
}
