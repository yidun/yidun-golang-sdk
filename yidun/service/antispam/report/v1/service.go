package v1

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/auth"
	"github.com/yidun/yidun-golang-sdk/yidun/core/client"
)

// 投诉举报解决方案检测提交客户端
type ReportClient struct {
	Client client.ExecuteClient
}

// NewReportClient creates a new client for Report.
func NewReportClient(profile *client.ClientProfile) *ReportClient {
	return &ReportClient{
		Client: client.NewDefaultClient(profile),
	}
}

// NewReportClientWithAccessKey creates a new client for Report
func NewReportClientWithAccessKey(accessKeyId string, accessKeySecret string) *ReportClient {
	credential := auth.NewCredentials(accessKeyId, accessKeySecret)
	return NewReportClient(client.NewClientProfile(credential))
}
