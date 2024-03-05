package submit

import "github.com/yidun/yidun-golang-sdk/yidun/core/types"

// 公众号检测提交接口响应对象v1.0
type OfficialAccountsSubmitV1Response struct {
	*types.CommonResponse
	Result *OfficialAccountsSubmitResult `json:"result,omitempty"`
}

type OfficialAccountsSubmitResult struct {
	JobId  *int64  `json:"jobId,omitempty"`
	TaskId *string `json:"taskId,omitempty"`
	DataId *string `json:"dataId,omitempty"`
}
