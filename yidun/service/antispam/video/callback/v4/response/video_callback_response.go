package response

import "github.com/yidun/yidun-golang-sdk/yidun/core/types"

// VideoCallbackV4Response 点播视频回调响应
type VideoCallbackV4Response struct {
	*types.CommonResponse
	Result *[]VideoCallbackV4Result `json:"result"`
}

// VideoCallbackV4Result 点播视频回调结果
type VideoCallbackV4Result struct {
	Antispam *VideoCallbackAntispamV4Response `json:"antispam,omitempty"` // 视频内容安全检测结果
	Ocr      *VideoCallbackOcrV4Response      `json:"ocr,omitempty"`      // 视频文本识别结果
	Discern  *VideoCallbackDiscernV4Response  `json:"discern,omitempty"`  // 视频画面识别结果
	Logo     *VideoCallbackLogoV4Response     `json:"logo,omitempty"`     // 视频logo识别结果
	Quality  *VideoCallbackQualityV4Response  `json:"quality,omitempty"`  // 视频质量检测结果
	Face     *VideoCallbackFaceV4Response     `json:"face,omitempty"`     // 视频人脸识别结果
	Ad       *VideoCallbackAdV4Response       `json:"ad,omitempty"`       // 视频广告检测结果
}
