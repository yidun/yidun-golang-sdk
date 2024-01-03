package request

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
)

type LiveWallSolutionCallbackV3Req struct {
	*types.PostFormRequest
	YidunRequestId *string `json:"yidunRequestId,omitempty"` // 请求的唯一ID
}

// NewLiveAudioCallbackV4Req 创建一个 LiveAudioCallbackV4Req 实例
func NewLiveWallSolutionCallbackV3Req() *LiveWallSolutionCallbackV3Req {
	var request = &LiveWallSolutionCallbackV3Req{
		PostFormRequest: types.NewPostFormRequest(),
	}
	request.SetProductCode("liveVideoSolutionCommon")
	request.SetUriPattern("/v3/livewallsolution/callback/results")
	request.SetVersion("v3.1")
	return request
}

// SetYidunRequestId 设置请求的唯一ID
func (r *LiveWallSolutionCallbackV3Req) SetYidunRequestId(yidunRequestId string) {
	r.YidunRequestId = &yidunRequestId
}

// GetCustomSignParams 获取自定义签名参数
func (r *LiveWallSolutionCallbackV3Req) GetBusinessCustomSignParams() map[string]string {
	params := r.PostFormRequest.GetBusinessCustomSignParams()
	if r.YidunRequestId != nil {
		params["yidunRequestId"] = *r.YidunRequestId
	}
	return params
}
