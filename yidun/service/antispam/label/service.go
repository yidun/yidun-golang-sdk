package label

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/auth"
	"github.com/yidun/yidun-golang-sdk/yidun/core/client"
)

type LabelClient struct {
	Client client.ExecuteClient
}

func NewLabelClient(profile *client.ClientProfile) *LabelClient {
	return &LabelClient{
		Client: client.NewOpenapiClient(profile),
	}
}

func NewLabelClientWithAccessKey(accessKeyId string, accessKeySecret string) *LabelClient {
	credential := auth.NewCredentials(accessKeyId, accessKeySecret)
	return NewLabelClient(client.NewClientProfile(credential))
}