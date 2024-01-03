package response

import "github.com/yidun/yidun-golang-sdk/yidun/core/types"

type LiveAudioCallbackV4Resp struct {
	*types.CommonResponse
	Result []*LiveAudioCallbackV4Result `json:"result"`
}

func NewLiveAudioCallbackV4Resp() *LiveAudioCallbackV4Resp {
	return &LiveAudioCallbackV4Resp{}
}
