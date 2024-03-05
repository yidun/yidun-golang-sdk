package submit

import "github.com/yidun/yidun-golang-sdk/yidun/core/types"

// WeiboSubmitV1Response 微博检测提交接口响应对象v1.0
type WeiboSubmitV1Response struct {
	*types.CommonResponse
	Result *WeiboSubmitResult `json:"result,omitempty"`
}

// WeiboSubmitResult contains the result of a Weibo submission in WeiboSubmitV1Response.
type WeiboSubmitResult struct {
	JobId  *int64  `json:"jobId,omitempty"`
	TaskId *string `json:"taskId,omitempty"`
}
