package livevideosolution

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/auth"
	"github.com/yidun/yidun-golang-sdk/yidun/core/client"
)

type LiveVideoSolutionClient struct {
	Client client.ExecuteClient
}

func NewLiveVideoSolutionClient(profile *client.ClientProfile) *LiveVideoSolutionClient {
	return &LiveVideoSolutionClient{
		Client: client.NewDefaultClient(profile),
	}
}

func NewLiveVideoSolutionClientWithAccessKey(accessKeyId string, accessKeySecret string) *LiveVideoSolutionClient {
	credential := auth.NewCredentials(accessKeyId, accessKeySecret)
	return NewLiveVideoSolutionClient(client.NewClientProfile(credential))
}
