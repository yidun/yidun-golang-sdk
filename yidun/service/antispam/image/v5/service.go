package image

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/auth"
	"github.com/yidun/yidun-golang-sdk/yidun/core/client"
)

type ImageClient struct {
	Client client.ExecuteClient
}

func NewImageClient(profile *client.ClientProfile) *ImageClient {
	return &ImageClient{
		Client: client.NewDefaultClient(profile),
	}
}

func NewImageClientWithAccessKey(accessKeyId string, accessKeySecret string) *ImageClient {
	credential := auth.NewCredentials(accessKeyId, accessKeySecret)
	return NewImageClient(client.NewClientProfile(credential))
}
