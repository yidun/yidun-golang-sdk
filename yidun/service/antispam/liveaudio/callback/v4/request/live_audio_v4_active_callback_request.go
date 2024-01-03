package request

import "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/callback"

type LiveAudioV4ActiveCallbackRequest struct {
	*callback.ActiveCallbackRequest
}

func NewLiveAudioV4ActiveCallbackRequest(params map[string]string) *LiveAudioV4ActiveCallbackRequest {

	return &LiveAudioV4ActiveCallbackRequest{callback.NewActiveCallbackRequest(params)}
}
