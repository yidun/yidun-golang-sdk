package feedback

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/core/validation"
)

type TextFeedbackRequest struct {
    *types.BizPostFormRequest
    Feedbacks          *string            `json:"feedbacks,omitempty"` //文本反馈数据，json格式对象数组，1-100条文本反馈数据
}

// 设置反馈数据
func (r *TextFeedbackRequest) SetFeedbacks(feedbacks string) {
	r.Feedbacks = &feedbacks

}
// 构建request
func NewTextCallBackRequest(businessId string) *TextFeedbackRequest {
	var request = &TextFeedbackRequest{
		BizPostFormRequest: types.NewBizPostFormRequest(businessId),
	}
	request.SetProductCode("text-api")
	request.SetUriPattern("/v2/text/feedback")
	request.SetVersion("v2")
	return request
}
//获取业务自定义参数
func (r *TextFeedbackRequest) GetBusinessCustomSignParams() map[string]string {
	params := r.BizPostFormRequest.GetBusinessCustomSignParams()
	if r.Feedbacks != nil {
		params["feedbacks"] = *r.Feedbacks
	}

	return params
}
//参数校验
func (r *TextFeedbackRequest) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "TextFeedbackRequest"}
	if r == nil {
		invalidParams.Add(validation.NewErrParamRequired("TextFeedbackRequest"))
		return invalidParams
	}
	if (r.Feedbacks == nil || len(*r.Feedbacks) == 0) {
		invalidParams.Add(validation.NewErrParamRequired("Feedbacks"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}