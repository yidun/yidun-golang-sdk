package pretreatment

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/auth"
	"github.com/yidun/yidun-golang-sdk/yidun/core/client"
)

type PretreatmentClient struct {
	Client client.ExecuteClient
}

// NewListClient 实例化一个名单服务客户端
func NewPretreatmentClient(profile *client.ClientProfile) *PretreatmentClient {
	return &PretreatmentClient{
		Client: client.NewDefaultClient(profile),
	}
}

func NewPretreatmentClientWithAccessKey(accessKeyId string, accessKeySecret string) *PretreatmentClient {
	credential := auth.NewCredentials(accessKeyId, accessKeySecret)
	return NewPretreatmentClient(client.NewClientProfile(credential))
}
