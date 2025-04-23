package request

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	response2 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/audio/callback/v4/response"
	response3 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/video/callback/v4/response"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/videosolution/callback/v2/response"
	response4 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/videosolution/callback/v2/response"
)

// VideoSolutionQueryTaskV2Request 视频解决方案查询任务响应
type VideoSolutionQueryTaskV2Response struct {
	*types.CommonResponse
	Result *[]VideoSolutionQueryTaskV2Result `json:"result,omitempty"`
}
type VideoSolutionQueryTaskV2Result struct {
	TaskID   *string                                          `json:"taskId,omitempty"`
	DataID   *string                                          `json:"dataId,omitempty"`
	Status   *int                                             `json:"status,omitempty"`
	Antispam response.VideoSolutionAntispamCallbackV2Response `json:"antispam,omitempty"`
	Language response2.AudioLanguageCallbackV4Response        `json:"language,omitempty"`
	Voice    response2.AudioVoiceCallbackV4Response           `json:"voice,omitempty"`
	Asr      response2.AudioAsrCallbackV4Response             `json:"asr,omitempty"`
	Ocr      response3.VideoCallbackOcrV4Response             `json:"ocr,omitempty"`
	Discern  response3.VideoCallbackDiscernV4Response         `json:"discern,omitempty"`
	Logo     response3.VideoCallbackLogoV4Response            `json:"logo,omitempty"`
	Face     response3.VideoCallbackFaceV4Response            `json:"face,omitempty"`
	Aigc     response3.VideoCallbackAigcV4Response            `json:"aigc,omitempty"`
	Quality  response4.VideoSolutionQualityCallbackV2Response `json:"quality,omitempty"`
	Ad       response4.VideoSolutionAdCallbackV4Response      `json:"ad,omitempty"`
}
