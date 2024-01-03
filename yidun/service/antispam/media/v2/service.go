package v2

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/auth"
	"github.com/yidun/yidun-golang-sdk/yidun/core/client"
)

// 融媒体解决方案检测提交客户端
type MediaClient struct {
	Client client.ExecuteClient
}

// NewMediaClient creates a new client for Media.
func NewMediaClient(profile *client.ClientProfile) *MediaClient {
	return &MediaClient{
		Client: client.NewDefaultClient(profile),
	}
}

// NewMediaClientWithAccessKey creates a new client for Media
func NewMediaClientWithAccessKey(accessKeyId string, accessKeySecret string) *MediaClient {
	credential := auth.NewCredentials(accessKeyId, accessKeySecret)
	return NewMediaClient(client.NewClientProfile(credential))
}
