package crawler

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/auth"
	"github.com/yidun/yidun-golang-sdk/yidun/core/client"
)

type CrawlerClient struct {
	Client client.ExecuteClient
}

func NewCrawlerClient(profile *client.ClientProfile) *CrawlerClient {
	return &CrawlerClient{
		Client: client.NewDefaultClient(profile),
	}
}

func NewCrawlerClientWithAccessKey(accessKeyId string, accessKeySecret string) *CrawlerClient {
	credential := auth.NewCredentials(accessKeyId, accessKeySecret)
	return NewCrawlerClient(client.NewClientProfile(credential))
}
