package auth

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/auth"
	"github.com/yidun/yidun-golang-sdk/yidun/core/client"
)

// AuthClient 客户端
type AuthClient struct {
	Client client.ExecuteClient
}

// NewAuthClient 创建客户端
func NewAuthClient(profile *client.ClientProfile) *AuthClient {
	return &AuthClient{
		Client: client.NewDefaultClient(profile),
	}
}

// NewAuthClientWithAccessKey 创建客户端
func NewAuthClientWithAccessKey(accessKeyId string, accessKeySecret string) *AuthClient {
	credential := auth.NewCredentials(accessKeyId, accessKeySecret)
	return NewAuthClient(client.NewClientProfile(credential))
}
