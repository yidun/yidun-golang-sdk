package fingerprint

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/auth"
	"github.com/yidun/yidun-golang-sdk/yidun/core/client"
)

type FingerprintClient struct {
	Client client.ExecuteClient
}

func NewAuthClient(profile *client.ClientProfile) *FingerprintClient {
	return &FingerprintClient{
		Client: client.NewDefaultClient(profile),
	}
}

func NewAuthClientWithAccessKey(accessKeyId string, accessKeySecret string) *FingerprintClient {
	credential := auth.NewCredentials(accessKeyId, accessKeySecret)
	return NewAuthClient(client.NewClientProfile(credential))
}
