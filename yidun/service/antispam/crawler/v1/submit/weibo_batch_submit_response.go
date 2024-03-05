package submit

import "github.com/yidun/yidun-golang-sdk/yidun/core/types"

// WeiboSubmitV1Response 微博检测提交接口响应对象v1.0
type WeiboBatchSubmitV1Response struct {
	*types.CommonResponse
	Result *[]WeiboSubmitResult `json:"result,omitempty"`
}
