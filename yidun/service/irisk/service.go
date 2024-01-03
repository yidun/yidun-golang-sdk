package irisk

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/auth"
	"github.com/yidun/yidun-golang-sdk/yidun/core/client"
)

// IRiskClient is the client to invoke APIs of irisk service
type IRiskClient struct {
	Client client.ExecuteClient
}

// NewIRiskClient returns a new irisk client
func NewIRiskClient(profile *client.ClientProfile) *IRiskClient {
	return &IRiskClient{
		Client: client.NewDefaultClient(profile),
	}
}

func NewIRiskClientWithAccessKey(accessKeyId string, accessKeySecret string) *IRiskClient {
	credential := auth.NewCredentials(accessKeyId, accessKeySecret)
	return NewIRiskClient(client.NewClientProfile(credential))
}
