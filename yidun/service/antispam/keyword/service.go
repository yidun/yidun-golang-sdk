package keyword

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/auth"
	"github.com/yidun/yidun-golang-sdk/yidun/core/client"
)

type KeywordClient struct {
	Client client.ExecuteClient
}

// NewKeywordClient 实例化一个关键词服务客户端
func NewKeywordClient(profile *client.ClientProfile) *KeywordClient {
	return &KeywordClient{
		Client: client.NewDefaultClient(profile),
	}
}

func NewKeywordClientWithAccessKey(accessKeyId string, accessKeySecret string) *KeywordClient {
	credential := auth.NewCredentials(accessKeyId, accessKeySecret)
	return NewKeywordClient(client.NewClientProfile(credential))
}
