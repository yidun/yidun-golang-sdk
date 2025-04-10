package response

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
)

// AudioCallBackV4Response 音频回调结果
type AudioCallBackV4Response struct {
	*types.CommonResponse
	Result []*AudioCallbackV4Result `json:"result,omitempty"` // 音频检测结果
}

type AudioCallbackV4Result struct {
	Antispam *AudioAntispamCallbackV4Response `json:"antispam,omitempty"` // 音频内容安全检测结果
	Asr      *AudioAsrCallbackV4Response      `json:"asr,omitempty"`      // 音频语音识别结果
	Language *AudioLanguageCallbackV4Response `json:"language,omitempty"` // 音频语种检测结果
	Quality  *AudioQualityCallbackV4Response  `json:"quality,omitempty"`  // 音频质量检测结果
	Voice    *AudioVoiceCallbackV4Response    `json:"voice,omitempty"`    // 音频语音检测结果
	Ad       *AudioAdCallbackV4Response       `json:"ad,omitempty"`       // 音频广告检测结果
}
