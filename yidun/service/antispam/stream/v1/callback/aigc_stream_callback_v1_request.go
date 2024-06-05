package callback

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
)

/**
 * AIGC流式检测检测结果获取（轮询模式）
 */
type AigcStreamCallBackRequestV1 struct {
	*types.BizPostFormRequest
	YidunRequestId *string `json:"yidunRequestId,omitempty"` // 请求的唯一ID
}

func (r *AigcStreamCallBackRequestV1) SetYidunRequestId(yidunRequestId string) {
	r.YidunRequestId = &yidunRequestId
}

func NewAigcStreamCallBackRequest() *AigcStreamCallBackRequestV1 {
	var request = &AigcStreamCallBackRequestV1{
		BizPostFormRequest: types.NewBizPostFormRequestWithoutBizId(),
	}
	request.SetProductCode("aigc-stream-api")
	request.SetUriPattern("/v1/stream/callback/results")
	request.SetVersion("v1.0")
	return request
}

// GetBusinessCustomSignParams 获取具体业务中特有的需要参与签名计算的参数
func (f *AigcStreamCallBackRequestV1) GetBusinessCustomSignParams() map[string]string {
	result := f.PostFormRequest.GetBusinessCustomSignParams()

	if f.YidunRequestId != nil {
		result["yidunRequestId"] = *f.YidunRequestId
	}

	return result
}
