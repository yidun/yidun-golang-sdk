package request

import (
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/callback"
)

type AudioV4ActiveCallbackRequest struct {
	*callback.ActiveCallbackRequest
}

// 音频主动回调
func NewAudioV4ActiveCallbackRequest(params map[string]string) *AudioV4ActiveCallbackRequest {

	return &AudioV4ActiveCallbackRequest{callback.NewActiveCallbackRequest(params)}
}
