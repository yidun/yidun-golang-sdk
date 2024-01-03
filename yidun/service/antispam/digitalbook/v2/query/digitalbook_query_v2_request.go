package query

import (
	"encoding/json"
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
)

/**
 * 数字阅读检测结果查询
 */
type DigitalBookQueryRequestV2 struct {
	*types.BizPostFormRequest
	TaskIds *[]string `json:"taskIds,omitempty"` // taskId列表
}

// Getter for TaskIds
func (m *DigitalBookQueryRequestV2) GetTaskIds() *[]string {
	return m.TaskIds
}

// Setter for TaskIds
func (m *DigitalBookQueryRequestV2) SetTaskIds(taskIds *[]string) {
	m.TaskIds = taskIds
}

func NewDigitalBookQueryRequest() *DigitalBookQueryRequestV2 {
	var request = &DigitalBookQueryRequestV2{
		BizPostFormRequest: types.NewBizPostFormRequestWithoutBizId(),
	}
	request.SetProductCode("digitalBook")
	request.SetUriPattern("/v2/digital/callback/query")
	request.SetVersion("v2")
	return request
}

// GetBusinessCustomSignParams 获取具体业务中特有的需要参与签名计算的参数
func (f *DigitalBookQueryRequestV2) GetBusinessCustomSignParams() map[string]string {
	result := f.PostFormRequest.GetBusinessCustomSignParams()

	if f.TaskIds != nil {
		// convert to json
		taskIds, _ := json.Marshal(f.TaskIds)
		result["taskIds"] = string(taskIds)
	}

	return result
}
