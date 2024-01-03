package videosolution

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/auth"
	"github.com/yidun/yidun-golang-sdk/yidun/core/client"
)

// VideoSolutionClient 视频解决方案服务
type VideoSolutionClient struct {
	Client client.ExecuteClient
}

// NewVideoSolutionClient 实例化一个视频解决方案服务客户端
func NewVideoSolutionClient(profile *client.ClientProfile) *VideoSolutionClient {
	return &VideoSolutionClient{
		Client: client.NewDefaultClient(profile),
	}
}

// NewVideoSolutionClientWithAccessKey 实例化一个视频解决方案服务客户端
func NewVideoSolutionClientWithAccessKey(accessKeyId string, accessKeySecret string) *VideoSolutionClient {
	credential := auth.NewCredentials(accessKeyId, accessKeySecret)
	return NewVideoSolutionClient(client.NewClientProfile(credential))
}
