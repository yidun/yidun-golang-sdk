package file

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/auth"
	"github.com/yidun/yidun-golang-sdk/yidun/core/client"
)

type FileClient struct {
	Client client.ExecuteClient
}

func NewFileClient(profile *client.ClientProfile) *FileClient {
	return &FileClient{
		Client: client.NewDefaultClient(profile),
	}
}

func NewFileClientWithAccessKey(accessKeyId string, accessKeySecret string) *FileClient {
	credential := auth.NewCredentials(accessKeyId, accessKeySecret)
	return NewFileClient(client.NewClientProfile(credential))
}
