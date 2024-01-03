package v5

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/auth"
	"github.com/yidun/yidun-golang-sdk/yidun/core/client"
)

type TextClient struct {
	Client client.ExecuteClient
}
// NewTextClient creates a new client for Text.
func NewTextClient(profile *client.ClientProfile) *TextClient {
	return &TextClient{
		Client: client.NewDefaultClient(profile),
	}
}
// NewTextClientWithAccessKey creates a new client for Text
func NewTextClientWithAccessKey(accessKeyId string, accessKeySecret string) *TextClient {
	credential := auth.NewCredentials(accessKeyId, accessKeySecret)
	return NewTextClient(client.NewClientProfile(credential))
}
