package callback

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
)

/**
 * 文本检测结果获取（轮询模式）
 */
type TextCallBackRequest struct {
    *types.BizPostFormRequest
    YidunRequestId *string `json:"yidunRequestId,omitempty"` // 请求的唯一ID
}
// 设置请求的唯一ID
func (r *TextCallBackRequest) SetYidunRequestId(yidunRequestId string) {
	r.YidunRequestId = &yidunRequestId

}
// 构建request
func NewTextCallBackRequest(businessId string) *TextCallBackRequest {
	var request = &TextCallBackRequest{
		BizPostFormRequest: types.NewBizPostFormRequest(businessId),
	}
	request.SetProductCode("text-api")
	request.SetUriPattern("/v5/text/callback/results")
	request.SetVersion("v5.2")
	return request
}