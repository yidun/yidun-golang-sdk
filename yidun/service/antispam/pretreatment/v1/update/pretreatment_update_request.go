package update

import (
	"strconv"

	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/core/validation"
)

// PretreatmentUpdateRequest 自定义忽略词修改请求
type PretreatmentUpdateRequest struct {
	*types.BizPostFormRequest
	// 添加描述
	Description *string `json:"description,omitempty"`
	// 忽略词id
	Id *int64 `json:"id,omitempty"`
	// 忽略词内容
	Entity *string `json:"entity,omitempty"`
	// 业务id
	BusinessId *string `json:"businessId,omitempty"`
}

// 参数校验方法
func (r *PretreatmentUpdateRequest) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "PretreatmentUpdateRequest"}
	if r.Entity == nil {
		invalidParams.Add(validation.NewErrParamRequired("Entity"))
	}
	if r.Entity != nil && len(*r.Entity) > 128 {
		invalidParams.Add(validation.NewErrParamMaxLen("entity不能超过128", 128, strconv.Itoa(len(*r.Entity))))
	}
	if r.Id == nil {
		invalidParams.Add(validation.NewErrParamRequired("Id"))
	}
	if r.Description != nil && len(*r.Description) > 128 {
		invalidParams.Add(validation.NewErrParamMaxLen("描述不能超过128", 128, strconv.Itoa(len(*r.Description))))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// NewPretreatmentUpdateRequest 初始化 PretreatmentUpdateRequest 对象
func NewPretreatmentUpdateRequest(businessId string) *PretreatmentUpdateRequest {
	var request = &PretreatmentUpdateRequest{
		BizPostFormRequest: types.NewBizPostFormRequest(businessId),
	}
	request.SetProductCode("pretreatment")
	request.SetUriPattern("/v1/pretreatment/update")
	request.SetMethod(http.HttpMethodPost)
	request.SetVersion("v1")
	return request
}

// GetCustomSignParams returns the custom sign parameters for the request
func (req *PretreatmentUpdateRequest) GetBusinessCustomSignParams() map[string]string {
	params := req.PostFormRequest.GetBusinessCustomSignParams()
	if req.Description != nil {
		params["description"] = *req.Description
	}
	if req.Id != nil {
		params["id"] = strconv.FormatInt(*req.Id, 10)
	}
	if req.Entity != nil {
		params["entity"] = *req.Entity
	}
	if req.BusinessId != nil {
		params["businessId"] = *req.BusinessId
	}
	return params
}

// SetDescription sets the Description field of the PretreatmentUpdateRequest.
func (req *PretreatmentUpdateRequest) SetDescription(description string) {
	req.Description = &description
}

// SetId sets the Id field of the PretreatmentUpdateRequest.
func (req *PretreatmentUpdateRequest) SetId(id int64) {
	req.Id = &id
}

// SetEntity sets the Entity field of the PretreatmentUpdateRequest.
func (req *PretreatmentUpdateRequest) SetEntity(entity string) {
	req.Entity = &entity
}

// SetBusinessId sets the BusinessId field.
func (req *PretreatmentUpdateRequest) SetBusinessId(businessId string) {
	req.BusinessId = &businessId
}
