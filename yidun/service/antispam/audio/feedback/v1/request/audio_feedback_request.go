package request

import (
	"encoding/json"
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/core/validation"
)

type AudioFeedbackRequest struct {
	*types.BizPostFormRequest
	Feedbacks []AudioFeedback `json:"feedbacks"` // 反馈信息
}

// 反馈信息
type AudioFeedback struct {
	TaskId *string `json:"taskId"`          // 任务ID
	Level  *int    `json:"level,omitempty"` // 反馈结果，1：通过，2：不通过
	Label  *int    `json:"label,omitempty"` // 反馈结果对应的分类
}

// SetFeedbacks 设置反馈信息
func (r *AudioFeedbackRequest) SetFeedbacks(feedbacks []AudioFeedback) {
	r.Feedbacks = feedbacks
}

func (a *AudioFeedback) SetTaskId(taskId string) {
	a.TaskId = &taskId
}
func (a *AudioFeedback) SetLevel(level int) {
	a.Level = &level
}
func (a *AudioFeedback) SetLabel(label int) {
	a.Label = &label
}

// NewAudioFeedbackRequest 初始化AudioFeedbackRequest对象
func NewAudioFeedbackRequest(businessId string) *AudioFeedbackRequest {
	var request = &AudioFeedbackRequest{
		BizPostFormRequest: types.NewBizPostFormRequest(businessId),
	}
	request.SetProductCode("audioCommon")
	request.SetUriPattern("/v1/audio/feedback")
	request.SetVersion("v1")
	return request
}

// GetCustomSignParams 获取具体业务中特有的需要参与签名计算的参数
func (r *AudioFeedbackRequest) GetBusinessCustomSignParams() map[string]string {
	params := r.BizPostFormRequest.GetBusinessCustomSignParams()
	if r.Feedbacks != nil {
		feedbackJson, _ := json.Marshal(r.Feedbacks)
		params["feedbacks"] = string(feedbackJson)
	}

	return params
}

// ValidateParam 检查参数是否正确
func (r *AudioFeedbackRequest) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "AudioFeedbackRequest"}
	if r == nil {
		invalidParams.Add(validation.NewErrParamRequired("AudioFeedbackRequest"))
	}
	if r.Feedbacks == nil || len(r.Feedbacks) == 0 {
		invalidParams.Add(validation.NewErrParamRequired("Feedbacks"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}
