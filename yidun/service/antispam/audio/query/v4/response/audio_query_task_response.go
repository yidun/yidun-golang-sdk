package response

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/audio/callback/v4/response"
)

// AudioQueryTaskV4Response 音频检测结果查询结果
type AudioQueryTaskV4Response struct {
	*types.CommonResponse
	Result *[]AudioQueryTaskV4Result `json:"result,omitempty"` // 音频检测结果
}

type AudioQueryTaskV4Result struct {
	TaskId   *string                                  `json:"taskId,omitempty"`   // 任务id，64位字符串
	Status   *int                                     `json:"status,omitempty"`   // 任务状态：0:正常，1:过期，2:不存在，3:检测中
	Antispam response.AudioAntispamCallbackV4Response `json:"antispam,omitempty"` // 音频内容安全检测结果
	Language response.AudioLanguageCallbackV4Response `json:"language,omitempty"` // 音频语种检测结果
	Asr      response.AudioAsrCallbackV4Response      `json:"asr,omitempty"`      // 音频语音识别结果
	Voice    response.AudioVoiceCallbackV4Response    `json:"voice,omitempty"`    // 音频语音检测结果
	Quality  response.AudioQualityCallbackV4Response  `json:"quality,omitempty"`  // 音频质量检测结果
}
