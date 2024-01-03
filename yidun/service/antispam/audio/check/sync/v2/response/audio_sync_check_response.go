package response

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/audio/callback/v4/response"
)

// AudioSyncCheckResponse 音频同步检测提交结果
type AudioSyncCheckResponse struct {
	*types.CommonResponse
	Result *AudioSyncCheckV2Result `json:"result,omitempty"` // 音频检测结果
}

type AudioSyncCheckV2Result struct {
	Antispam response.AudioAntispamCallbackV4Response `json:"antispam,omitempty"` // 音频内容安全检测结果
	Asr      response.AudioAsrCallbackV4Response      `json:"asr,omitempty"`      // 音频语音识别结果
	Language response.AudioLanguageCallbackV4Response `json:"language,omitempty"` // 音频语种检测结果
	Voice    response.AudioVoiceCallbackV4Response    `json:"voice,omitempty"`    // 音频语音检测结果
}
