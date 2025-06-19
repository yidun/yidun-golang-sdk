package textfeature

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/core/validation"
)

// TextFeatureDeleteRequest 文本特征删除请求
type TextFeatureDeleteRequest struct {
	*types.BizPostFormRequest
	Uuids *string `json:"uuids,omitempty"`
}

func (r *TextFeatureDeleteRequest) SetUuids(uuids string) {
	r.Uuids = &uuids
}

func NewTextFeatureDeleteRequest(businessId string) *TextFeatureDeleteRequest {
	request := &TextFeatureDeleteRequest{
		BizPostFormRequest: types.NewBizPostFormRequest(businessId),
	}
	request.SetProductCode("text-api")
	request.SetUriPattern("/v1/text-feature/delete")
	request.SetVersion("v1")
	return request
}

func (r *TextFeatureDeleteRequest) GetBusinessCustomSignParams() map[string]string {
	params := r.BizPostFormRequest.GetBusinessCustomSignParams()
	if r.Uuids != nil {
		params["uuids"] = *r.Uuids
	}
	return params
}

func (r *TextFeatureDeleteRequest) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "TextFeatureDeleteRequest"}
	if r == nil {
		invalidParams.Add(validation.NewErrParamRequired("TextFeatureDeleteRequest"))
		return invalidParams
	}
	if r.Uuids == nil || len(*r.Uuids) == 0 {
		invalidParams.Add(validation.NewErrParamRequired("Uuids"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}
