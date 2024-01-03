package list

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/auth"
	"github.com/yidun/yidun-golang-sdk/yidun/core/client"
)

type ListClient struct {
	Client client.ExecuteClient
}

// NewListClient 实例化一个名单服务客户端
func NewListClient(profile *client.ClientProfile) *ListClient {
	return &ListClient{
		Client: client.NewDefaultClient(profile),
	}
}

func NewListClientWithAccessKey(accessKeyId string, accessKeySecret string) *ListClient {
	credential := auth.NewCredentials(accessKeyId, accessKeySecret)
	return NewListClient(client.NewClientProfile(credential))
}
