package feedback

import (
	"encoding/json"

	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/core/validation"
)

type ImageV5FeedbackRequest struct {
	*types.BizPostFormRequest
	Feedbacks *[]ImageFeedbackBeanRequest `json:"feedbacks,omitempty"`
}

func (r *ImageV5FeedbackRequest) SetFeedbacks(feedbacks []ImageFeedbackBeanRequest) {
	r.Feedbacks = &feedbacks
}

// 构建request
func NewImageFeedbackRequest(businessId string) *ImageV5FeedbackRequest {
	var request = &ImageV5FeedbackRequest{
		BizPostFormRequest: types.NewBizPostFormRequest(businessId),
	}
	request.SetProductCode("imageCommon")
	request.SetUriPattern("/v2/image/feedback")
	request.SetVersion("v2")
	return request
}

// 获取业务自定义参数
func (r *ImageV5FeedbackRequest) GetBusinessCustomSignParams() map[string]string {
	params := r.BizPostFormRequest.GetBusinessCustomSignParams()
	if r.Feedbacks != nil {
		feedbackJson, _ := json.Marshal(r.Feedbacks)
		params["feedbacks"] = string(feedbackJson)
	}

	return params
}

// 参数校验
func (r *ImageV5FeedbackRequest) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "ImageV5FeedBackRequest"}
	if r == nil {
		invalidParams.Add(validation.NewErrParamRequired("ImageV5FeedBackRequest"))
		return invalidParams
	}
	if r.Feedbacks == nil || len(*r.Feedbacks) == 0 {
		invalidParams.Add(validation.NewErrParamRequired("Feedbacks"))
		return invalidParams
	}
	return nil
}

type ImageFeedbackBeanRequest struct {
	// taskId不能为空
	TaskId *string `json:"taskId,omitempty"`
	// level不能为空
	Level *int `json:"level,omitempty"`
	Label *int `json:"label,omitempty"`
}

func (r *ImageFeedbackBeanRequest) SetTaskId(taskId string) {
	r.TaskId = &taskId
}

func (r *ImageFeedbackBeanRequest) SetLevel(level int) {
	r.Level = &level
}

func (r *ImageFeedbackBeanRequest) SetLabel(label int) {
	r.Label = &label
}
