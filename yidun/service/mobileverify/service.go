package mobileverify

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/auth"
	"github.com/yidun/yidun-golang-sdk/yidun/core/client"
)

// MobileNumberClient 客户端
type MobileNumberClient struct {
	Client client.ExecuteClient
}

// NewMobileNumberClient 创建客户端
func NewMobileNumberClient(profile *client.ClientProfile) *MobileNumberClient {
	return &MobileNumberClient{
		Client: client.NewDefaultClient(profile),
	}
}

// NewMobileNumberClientWithAccessKey 创建客户端
func NewMobileNumberClientWithAccessKey(accessKeyId string, accessKeySecret string) *MobileNumberClient {
	credential := auth.NewCredentials(accessKeyId, accessKeySecret)
	return NewMobileNumberClient(client.NewClientProfile(credential))
}
