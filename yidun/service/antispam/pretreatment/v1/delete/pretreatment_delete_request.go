package delete

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/core/validation"
)

// ImageListDeleteRequest 图片名单删除请求
type PretreatmentDeleteRequest struct {
	*types.BizPostFormRequest
	// 忽略词id列表，逗号分隔
	Ids *string `json:"ids,omitempty"`
	// 业务id
	BusinessId *string `json:"businessId,omitempty"`
}

// 参数校验方法
func (r *PretreatmentDeleteRequest) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "PretreatmentDeleteRequest"}
	if r.Ids == nil {
		invalidParams.Add(validation.NewErrParamRequired("Ids"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// NewPretreatmentDeleteRequest 初始化 PretreatmentDeleteRequest 对象
func NewPretreatmentDeleteRequest(businessId string) *PretreatmentDeleteRequest {
	var request = &PretreatmentDeleteRequest{
		BizPostFormRequest: types.NewBizPostFormRequest(businessId),
	}
	request.SetProductCode("pretreatment")
	request.SetUriPattern("/v1/pretreatment/batchDelete")
	request.SetMethod(http.HttpMethodPost)
	request.SetVersion("v1")
	return request
}

// GetCustomSignParams returns the custom sign parameters for the request
func (req *PretreatmentDeleteRequest) GetBusinessCustomSignParams() map[string]string {
	params := req.PostFormRequest.GetBusinessCustomSignParams()

	if req.Ids != nil {
		params["ids"] = *req.Ids
	}
	if req.BusinessId != nil {
		params["businessId"] = *req.BusinessId
	}
	return params
}

// SetDescription sets the Description field.
func (req *PretreatmentDeleteRequest) SetIds(ids string) {
	req.Ids = &ids
}

// SetBusinessId sets the BusinessId field.
func (req *PretreatmentDeleteRequest) SetBusinessId(businessId string) {
	req.BusinessId = &businessId
}
