package request

import (
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/audio/check"
)

type AudioSyncCheckRequest struct {
	*check.AudioCommonCheckRequest
}

// NewAudioSyncCheckRequest 初始化AudioSyncCheckRequest对象
func NewAudioSyncCheckRequest(businessId string) *AudioSyncCheckRequest {
	var request = &AudioSyncCheckRequest{
		AudioCommonCheckRequest: check.NewAudioCommonCheckRequest(businessId),
	}
	request.SetProductCode("audioCheck")
	request.SetUriPattern("/v2/audio/check")
	request.SetVersion("v2.2")
	return request
}

// GetCustomSignParams 获取具体业务中特有的需要参与签名计算的参数
func (r *AudioSyncCheckRequest) GetBusinessCustomSignParams() map[string]string {
	params := r.AudioCommonCheckRequest.GetBusinessCustomSignParams()
	return params
}
