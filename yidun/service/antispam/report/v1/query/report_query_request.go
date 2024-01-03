package query

import (
	"encoding/json"
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
)

/**
 * 投诉举报检测结果查询
 */
type ReportQueryRequestV1 struct {
	*types.BizPostFormRequest
	TaskIds *[]string `json:"taskIds,omitempty"` // taskId列表
}

// Getter for TaskIds
func (m *ReportQueryRequestV1) GetTaskIds() *[]string {
	return m.TaskIds
}

// Setter for TaskIds
func (m *ReportQueryRequestV1) SetTaskIds(taskIds *[]string) {
	m.TaskIds = taskIds
}

func NewReportQueryRequestV1() *ReportQueryRequestV1 {
	var request = &ReportQueryRequestV1{
		BizPostFormRequest: types.NewBizPostFormRequestWithoutBizId(),
	}
	request.SetProductCode("report")
	request.SetUriPattern("/v1/report/callback/query")
	request.SetVersion("v1")
	return request
}

// GetBusinessCustomSignParams 获取具体业务中特有的需要参与签名计算的参数
func (f *ReportQueryRequestV1) GetBusinessCustomSignParams() map[string]string {
	result := f.PostFormRequest.GetBusinessCustomSignParams()

	if f.TaskIds != nil {
		// convert to json
		taskIds, _ := json.Marshal(f.TaskIds)
		result["taskIds"] = string(taskIds)
	}

	return result
}
