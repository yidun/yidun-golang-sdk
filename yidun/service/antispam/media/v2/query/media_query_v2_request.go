package query

import (
	"encoding/json"
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
)

/**
 * 融媒体检测结果查询
 */
type MediaQueryRequestV2 struct {
	*types.BizPostFormRequest
	TaskIds *[]string `json:"taskIds,omitempty"` // taskId列表
}

// Getter for TaskIds
func (m *MediaQueryRequestV2) GetTaskIds() *[]string {
	return m.TaskIds
}

// Setter for TaskIds
func (m *MediaQueryRequestV2) SetTaskIds(taskIds *[]string) {
	m.TaskIds = taskIds
}

func NewMediaQueryRequestV2() *MediaQueryRequestV2 {
	var request = &MediaQueryRequestV2{
		BizPostFormRequest: types.NewBizPostFormRequestWithoutBizId(),
	}
	request.SetProductCode("media")
	request.SetUriPattern("/v2/mediasolution/callback/query")
	request.SetVersion("v2.1")
	return request
}

// GetBusinessCustomSignParams 获取具体业务中特有的需要参与签名计算的参数
func (f *MediaQueryRequestV2) GetBusinessCustomSignParams() map[string]string {
	result := f.PostFormRequest.GetBusinessCustomSignParams()

	if f.TaskIds != nil {
		// convert to json
		taskIds, _ := json.Marshal(f.TaskIds)
		result["taskIds"] = string(taskIds)
	}

	return result
}
