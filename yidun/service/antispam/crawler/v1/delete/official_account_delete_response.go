package delete

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
)

// OfficialAccountsDeleteV1Response 网站任务检测批量查询接口响应对象v1.0
type OfficialAccountsDeleteV1Response struct {
	*types.CommonResponse
	Result *OfficialAccountsDeleteV1Result `json:"result,omitempty"`
}

// OfficialAccountsDeleteV1Result
type OfficialAccountsDeleteV1Result struct {
	// 本次查询的任务id
	DeletedJobIds *[]int64 `json:"deletedJobIds,omitempty"`
}
