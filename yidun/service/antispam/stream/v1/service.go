package v1

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/auth"
	"github.com/yidun/yidun-golang-sdk/yidun/core/client"
)

// AIGC流式检测解决方案检测提交客户端
type AigcStreamClient struct {
	Client client.ExecuteClient
}

// NewAigcStreamClient creates a new client for AigcStream.
func NewAigcStreamClient(profile *client.ClientProfile) *AigcStreamClient {
	return &AigcStreamClient{
		Client: client.NewDefaultClient(profile),
	}
}

// NewAigcStreamClientWithAccessKey creates a new client for AigcStream
func NewAigcStreamClientWithAccessKey(accessKeyId string, accessKeySecret string) *AigcStreamClient {
	credential := auth.NewCredentials(accessKeyId, accessKeySecret)
	return NewAigcStreamClient(client.NewClientProfile(credential))
}
