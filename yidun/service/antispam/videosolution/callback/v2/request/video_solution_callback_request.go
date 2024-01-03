package request

import "github.com/yidun/yidun-golang-sdk/yidun/core/types"

// VideoSolutionCallbackV2Request 视频解决方案回调请求
type VideoSolutionCallbackV2Request struct {
	*types.PostFormRequest
	YidunRequestId *string `json:"yidunRequestId,omitempty"` // 请求的唯一ID
}

func (r *VideoSolutionCallbackV2Request) SetYidunRequestId(yidunRequestId string) {
	r.YidunRequestId = &yidunRequestId
}

func NewVideoSolutionCallbackV2Request() *VideoSolutionCallbackV2Request {
	var request = &VideoSolutionCallbackV2Request{
		PostFormRequest: types.NewPostFormRequest(),
	}
	request.SetProductCode("videoSolutionCommon")
	request.SetUriPattern("/v2/videosolution/callback/results")
	request.SetVersion("v2.1")
	return request
}
