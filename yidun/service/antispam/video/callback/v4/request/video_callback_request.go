package request

import "github.com/yidun/yidun-golang-sdk/yidun/core/types"

type VideoCallbackV4Request struct {
	*types.BizPostFormRequest
	YidunRequestId *string `json:"yidunRequestId,omitempty"` // 请求的唯一ID
}

func (r *VideoCallbackV4Request) SetYidunRequestId(yidunRequestId string) {
	r.YidunRequestId = &yidunRequestId
}

// NewVideoCallbackV4Request 初始化VideoCallbackV4Request对象
func NewVideoCallbackV4Request(businessId string) *VideoCallbackV4Request {
	var request = &VideoCallbackV4Request{
		BizPostFormRequest: types.NewBizPostFormRequest(businessId),
	}
	request.SetProductCode("videoCommon")
	request.SetUriPattern("/v4/video/callback/results")
	request.SetVersion("v4.1")
	return request
}
