package request

import (
	"encoding/json"
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
)

type LiveWallSolutionFeedbackV1Req struct {
	*types.PostFormRequest
	realTimeInfoList []*LiveWallSolutionFeedback
}

func NewLiveWallSolutionFeedbackV1Req() *LiveWallSolutionFeedbackV1Req {
	var request = &LiveWallSolutionFeedbackV1Req{
		PostFormRequest: types.NewPostFormRequest(),
	}
	request.SetProductCode("liveVideoSolutionCommon")
	request.SetUriPattern("/v1/livewallsolution/feedback")
	request.SetVersion("v1")
	return request
}

func (req *LiveWallSolutionFeedbackV1Req) SetFeedbacks(feedbacks []*LiveWallSolutionFeedback) {
	req.realTimeInfoList = feedbacks
}

// 定义 GetCustomSignParams() 方法的功能
func (req *LiveWallSolutionFeedbackV1Req) GetBusinessCustomSignParams() map[string]string {
	params := req.PostFormRequest.GetBusinessCustomSignParams()
	if req.realTimeInfoList != nil {
		// 转换为 json 的操作
		feedbacksJson, _ := json.Marshal(req.realTimeInfoList)
		params["realTimeInfoList"] = string(feedbacksJson)
	}
	return params
}
