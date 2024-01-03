package request

import (
	"encoding/json"
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
)

type LiveAudioFeedbackV1Req struct {
	*types.BizPostFormRequest
	feedbacks []*LiveAudioFeedback
}

func NewLiveAudioFeedbackV1Req(businessId string) *LiveAudioFeedbackV1Req {
	var request = &LiveAudioFeedbackV1Req{
		BizPostFormRequest: types.NewBizPostFormRequest(businessId),
	}
	request.SetProductCode("liveAudio")
	request.SetUriPattern("/v1/liveaudio/feedback")
	request.SetVersion("v1")
	return request
}

func (req *LiveAudioFeedbackV1Req) SetFeedbacks(feedbacks []*LiveAudioFeedback) {
	req.feedbacks = feedbacks
}

// 定义 GetCustomSignParams() 方法的功能
func (req *LiveAudioFeedbackV1Req) GetBusinessCustomSignParams() map[string]string {
	params := req.BizPostFormRequest.GetBusinessCustomSignParams()
	if req.feedbacks != nil {
		// 转换为 json 的操作
		feedbacksJson, _ := json.Marshal(req.feedbacks)
		params["feedbacks"] = string(feedbacksJson)
	}
	return params
}
