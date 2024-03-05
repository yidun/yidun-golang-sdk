package submit

import "github.com/yidun/yidun-golang-sdk/yidun/core/types"

// 公众号批量检测提交接口响应对象v1.0
type OfficialAccountsBatchSubmitV1Response struct {
	*types.CommonResponse
	Result *[]OfficialAccountsSubmitResult `json:"result,omitempty"`
}
