package request

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
)

type LiveAudioCallbackV4Req struct {
	*types.BizPostFormRequest
	YidunRequestId *string `json:"yidunRequestId,omitempty"` // 请求的唯一ID
}

// NewLiveAudioCallbackV4Req 创建一个 LiveAudioCallbackV4Req 实例
func NewLiveAudioCallbackV4Req(businessId string) *LiveAudioCallbackV4Req {
	var request = &LiveAudioCallbackV4Req{
		BizPostFormRequest: types.NewBizPostFormRequest(businessId),
	}
	request.SetProductCode("liveAudio")
	request.SetUriPattern("/v4/liveaudio/callback/results")
	request.SetVersion("v4.1")
	return request
}

// SetYidunRequestId 设置请求的唯一ID
func (r *LiveAudioCallbackV4Req) SetYidunRequestId(yidunRequestId string) {
	r.YidunRequestId = &yidunRequestId
}

// GetCustomSignParams 获取自定义签名参数
func (r *LiveAudioCallbackV4Req) GetBusinessCustomSignParams() map[string]string {
	params := r.BizPostFormRequest.GetBusinessCustomSignParams()
	if r.YidunRequestId != nil {
		params["yidunRequestId"] = *r.YidunRequestId
	}
	return params
}
