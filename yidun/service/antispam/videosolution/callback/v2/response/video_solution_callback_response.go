package response

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/audio/callback/v4/response"
	response2 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/video/callback/v4/response"
)

// VideoSolutionCallbackV2Response 视频解决方案回调响应
type VideoSolutionCallbackV2Response struct {
	*types.CommonResponse
	Result *[]VideoSolutionCallbackV2Result `json:"result,omitempty"`
}

type VideoSolutionCallbackV2Result struct {
	Antispam *VideoSolutionAntispamCallbackV2Response  `json:"antispam,omitempty"`
	Language *response.AudioLanguageCallbackV4Response `json:"language,omitempty"`
	Voice    *response.AudioVoiceCallbackV4Response    `json:"voice,omitempty"`
	Asr      *response.AudioAsrCallbackV4Response      `json:"asr,omitempty"`
	Ocr      *response2.VideoCallbackOcrV4Response     `json:"ocr,omitempty"`
	Discern  *response2.VideoCallbackDiscernV4Response `json:"discern,omitempty"`
	Logo     *response2.VideoCallbackLogoV4Response    `json:"logo,omitempty"`
	Face     *response2.VideoCallbackFaceV4Response    `json:"face,omitempty"`
	Quality  *VideoSolutionQualityCallbackV2Response   `json:"quality,omitempty"`
}

type VideoSolutionQualityCallbackV2Response struct {
	Audio *response.AudioQualityCallbackV4Response  `json:"audio,omitempty"`
	Video *response2.VideoCallbackQualityV4Response `json:"video,omitempty"`
}
