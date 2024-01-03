package request

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
)

/**
 * 音频检测结果获取（轮询模式）
 */
type AudioCallBackRequest struct {
	*types.BizPostFormRequest
	YidunRequestId *string `json:"yidunRequestId,omitempty"` // 请求的唯一ID
}

func (r *AudioCallBackRequest) SetYidunRequestId(yidunRequestId string) {
	r.YidunRequestId = &yidunRequestId

}

// NewAudioCallBackRequest 初始化AudioCallBackRequest对象
func NewAudioCallBackRequest(businessId string) *AudioCallBackRequest {
	var request = &AudioCallBackRequest{
		BizPostFormRequest: types.NewBizPostFormRequest(businessId),
	}
	request.SetProductCode("audioCommon")
	request.SetUriPattern("/v4/audio/callback/results")
	request.SetVersion("v4.1")
	return request
}
