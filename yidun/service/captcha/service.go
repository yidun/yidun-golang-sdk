package captcha

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/auth"
	"github.com/yidun/yidun-golang-sdk/yidun/core/client"
)

// CaptchaVerifyClient 客户端
type CaptchaVerifyClient struct {
	Client client.ExecuteClient
}

// NewCaptchaVerifyClient 创建客户端
func NewCaptchaVerifyClient(profile *client.ClientProfile) *CaptchaVerifyClient {
	return &CaptchaVerifyClient{
		Client: client.NewDefaultClient(profile),
	}
}

// NewCaptchaVerifyClientWithAccessKey 创建客户端
func NewCaptchaVerifyClientWithAccessKey(accessKeyId string, accessKeySecret string) *CaptchaVerifyClient {
	credential := auth.NewCredentials(accessKeyId, accessKeySecret)
	return NewCaptchaVerifyClient(client.NewClientProfile(credential))
}
