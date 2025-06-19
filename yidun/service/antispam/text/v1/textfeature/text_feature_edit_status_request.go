package textfeature

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/core/validation"
	"strconv"
)

// TextFeatureEditStatusRequest 文本特征状态编辑请求

type TextFeatureEditStatusRequest struct {
	*types.BizPostFormRequest
	Uuids  *string `json:"uuids,omitempty"`  // JSON数组字符串
	Status *int    `json:"status,omitempty"` // 1: 有效, 2: 失效
}

func (r *TextFeatureEditStatusRequest) SetUuids(uuids string) {
	r.Uuids = &uuids
}

func (r *TextFeatureEditStatusRequest) SetStatus(status int) {
	r.Status = &status
}

func NewTextFeatureEditStatusRequest(businessId string) *TextFeatureEditStatusRequest {
	request := &TextFeatureEditStatusRequest{
		BizPostFormRequest: types.NewBizPostFormRequest(businessId),
	}
	request.SetProductCode("text-api")
	request.SetUriPattern("/v1/text-feature/edit-status")
	request.SetVersion("v1")
	return request
}

func (r *TextFeatureEditStatusRequest) GetBusinessCustomSignParams() map[string]string {
	params := r.BizPostFormRequest.GetBusinessCustomSignParams()
	if r.Uuids != nil {
		params["uuids"] = *r.Uuids
	}
	if r.Status != nil {
		params["status"] = strconv.Itoa(*r.Status)
	}
	return params
}

func (r *TextFeatureEditStatusRequest) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "TextFeatureEditStatusRequest"}
	if r == nil {
		invalidParams.Add(validation.NewErrParamRequired("TextFeatureEditStatusRequest"))
		return invalidParams
	}
	if r.Uuids == nil || len(*r.Uuids) == 0 {
		invalidParams.Add(validation.NewErrParamRequired("Uuids"))
	}
	if r.Status == nil || (*r.Status != 1 && *r.Status != 2) {
		invalidParams.Add(validation.NewErrParamRequired("Status(1|2)"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}
