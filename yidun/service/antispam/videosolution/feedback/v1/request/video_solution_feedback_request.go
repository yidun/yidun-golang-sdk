package request

import (
	"encoding/json"
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/core/validation"
)

// VideoSolutionFeedbackV1Request 视频解决方案反馈请求
type VideoSolutionFeedbackV1Request struct {
	*types.PostFormRequest
	Feedbacks *[]VideoSolutionFeedback `json:"feedbacks"` // 反馈信息
}

type VideoSolutionFeedback struct {
	TaskId *string   `json:"taskId"`          // 任务ID
	Tags   *[]string `json:"tags,omitempty"`  // 标签
	Level  *int      `json:"level,omitempty"` // 反馈结果，1：通过，2：不通过
	Label  *int      `json:"label,omitempty"` // 反馈结果对应的分类
}

func (r *VideoSolutionFeedbackV1Request) SetFeedbacks(feedbacks []VideoSolutionFeedback) {
	r.Feedbacks = &feedbacks
}

func (a *VideoSolutionFeedback) SetTaskId(taskId string) {
	a.TaskId = &taskId
}
func (a *VideoSolutionFeedback) SetTags(tags []string) {
	a.Tags = &tags
}
func (a *VideoSolutionFeedback) SetLevel(level int) {
	a.Level = &level
}
func (a *VideoSolutionFeedback) SetLabel(label int) {
	a.Label = &label
}

func NewVideoSolutionFeedbackRequest() *VideoSolutionFeedbackV1Request {
	var request = &VideoSolutionFeedbackV1Request{
		PostFormRequest: types.NewPostFormRequest(),
	}
	request.SetProductCode("videoSolutionCommon")
	request.SetUriPattern("/v1/videosolution/feedback")
	request.SetVersion("v1")
	return request
}
func (r *VideoSolutionFeedbackV1Request) GetBusinessCustomSignParams() map[string]string {
	params := r.PostFormRequest.GetBusinessCustomSignParams()
	if r.Feedbacks != nil {
		feedbackJson, _ := json.Marshal(r.Feedbacks)
		params["feedbacks"] = string(feedbackJson)
	}

	return params
}

func (r *VideoSolutionFeedbackV1Request) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "VideoSolutionFeedbackV1Request"}
	if r == nil {
		invalidParams.Add(validation.NewErrParamRequired("VideoSolutionFeedbackV1Request"))
	}
	if r.Feedbacks == nil || len(*r.Feedbacks) == 0 {
		invalidParams.Add(validation.NewErrParamRequired("Feedbacks"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}
