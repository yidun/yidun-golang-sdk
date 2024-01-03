package v2

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/auth"
	"github.com/yidun/yidun-golang-sdk/yidun/core/client"
)

// 数字阅读解决方案检测提交客户端
type DigitalBookClient struct {
	Client client.ExecuteClient
}

// NewDigitalBookClient creates a new client for DigitalBook.
func NewDigitalBookClient(profile *client.ClientProfile) *DigitalBookClient {
	return &DigitalBookClient{
		Client: client.NewDefaultClient(profile),
	}
}

// NewDigitalBookClientWithAccessKey creates a new client for DigitalBook
func NewDigitalBookClientWithAccessKey(accessKeyId string, accessKeySecret string) *DigitalBookClient {
	credential := auth.NewCredentials(accessKeyId, accessKeySecret)
	return NewDigitalBookClient(client.NewClientProfile(credential))
}
