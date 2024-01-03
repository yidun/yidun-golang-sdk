package request

import (
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/audio/check"
)

type AudioAsyncCheckRequest struct {
	*check.AudioCommonCheckRequest
}

// NewAudioAsyncCheckRequest 初始化AudioAsyncCheckRequest对象
func NewAudioAsyncCheckRequest(businessId string) *AudioAsyncCheckRequest {
	var request = &AudioAsyncCheckRequest{
		AudioCommonCheckRequest: check.NewAudioCommonCheckRequest(businessId),
	}
	request.SetProductCode("audioCheck")
	request.SetUriPattern("/v4/audio/submit")
	request.SetVersion("v4.1")
	return request
}

// GetCustomSignParams 获取具体业务中特有的需要参与签名计算的参数
func (r *AudioAsyncCheckRequest) GetBusinessCustomSignParams() map[string]string {
	params := r.AudioCommonCheckRequest.GetBusinessCustomSignParams()
	return params
}
