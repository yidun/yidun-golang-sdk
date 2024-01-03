package callback

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
)

/**
 * 融媒体检测结果获取（轮询模式）
 */
type MediaCallBackRequestV2 struct {
	*types.BizPostFormRequest
	YidunRequestId *string `json:"yidunRequestId,omitempty"` // 请求的唯一ID
}

func (r *MediaCallBackRequestV2) SetYidunRequestId(yidunRequestId string) {
	r.YidunRequestId = &yidunRequestId
}

func NewMediaCallBackRequest() *MediaCallBackRequestV2 {
	var request = &MediaCallBackRequestV2{
		BizPostFormRequest: types.NewBizPostFormRequestWithoutBizId(),
	}
	request.SetProductCode("media")
	request.SetUriPattern("/v2/mediasolution/callback/results")
	request.SetVersion("v2.1")
	return request
}

// GetBusinessCustomSignParams 获取具体业务中特有的需要参与签名计算的参数
func (f *MediaCallBackRequestV2) GetBusinessCustomSignParams() map[string]string {
	result := f.PostFormRequest.GetBusinessCustomSignParams()

	if f.YidunRequestId != nil {
		result["yidunRequestId"] = *f.YidunRequestId
	}

	return result
}
