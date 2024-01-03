package callback

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
)

/**
 * 数字阅读检测结果获取（轮询模式）
 */
type DigitalBookCallBackRequestV2 struct {
	*types.BizPostFormRequest
	YidunRequestId *string `json:"yidunRequestId,omitempty"` // 请求的唯一ID
}

func (r *DigitalBookCallBackRequestV2) SetYidunRequestId(yidunRequestId string) {
	r.YidunRequestId = &yidunRequestId
}

func NewDigitalBookCallBackRequest() *DigitalBookCallBackRequestV2 {
	var request = &DigitalBookCallBackRequestV2{
		BizPostFormRequest: types.NewBizPostFormRequestWithoutBizId(),
	}
	request.SetProductCode("digitalBook")
	request.SetUriPattern("/v2/digital/callback/results")
	request.SetVersion("v2")
	return request
}

// GetBusinessCustomSignParams 获取具体业务中特有的需要参与签名计算的参数
func (f *DigitalBookCallBackRequestV2) GetBusinessCustomSignParams() map[string]string {
	result := f.PostFormRequest.GetBusinessCustomSignParams()

	if f.YidunRequestId != nil {
		result["yidunRequestId"] = *f.YidunRequestId
	}

	return result
}
