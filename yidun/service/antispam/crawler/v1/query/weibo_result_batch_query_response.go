package query

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
)

// WeiboBatchQueryV1Response 网站任务检测批量查询接口响应对象v1.0
type WeiboBatchQueryV1Response struct {
	*types.CommonResponse
	Result *[]WeiboBatchQueryV1Result `json:"result,omitempty"`
}

// WeiboBatchQueryV1Result
type WeiboBatchQueryV1Result struct {
	// 本次查询的任务id
	JobId *int64 `json:"jobId,omitempty"`
	// 任务状态
	Status *int `json:"status,omitempty"`
}
