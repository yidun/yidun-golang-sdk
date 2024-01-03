package add

import (
	"strconv"

	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/core/validation"
)

// PretreatmentAddRequest 自定义忽略词添加请求
type PretreatmentAddRequest struct {
	*types.BizPostFormRequest
	// 添加描述
	Description *string `json:"description,omitempty"`
	// 忽略词列表，逗号分隔
	Entitys *string `json:"entitys,omitempty"`
	// 业务id
	BusinessId *string `json:"businessId,omitempty"`
}

// 参数校验方法
func (r *PretreatmentAddRequest) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "PretreatmentAddRequest"}
	if r.Description != nil && len(*r.Description) > 128 {
		invalidParams.Add(validation.NewErrParamMaxLen("描述不能超过128", 128, strconv.Itoa(len(*r.Description))))
	}
	if r.Entitys == nil {
		invalidParams.Add(validation.NewErrParamRequired("Entitys"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// NewPretreatmentAddRequest 初始化 PretreatmentAddRequest 对象
func NewPretreatmentAddRequest(businessId string) *PretreatmentAddRequest {
	var request = &PretreatmentAddRequest{
		BizPostFormRequest: types.NewBizPostFormRequest(businessId),
	}
	request.SetProductCode("pretreatment")
	request.SetUriPattern("/v2/pretreatment/add")
	request.SetMethod(http.HttpMethodPost)
	request.SetVersion("v2")
	return request
}

// GetCustomSignParams returns the custom sign parameters for the request
func (req *PretreatmentAddRequest) GetBusinessCustomSignParams() map[string]string {
	params := req.PostFormRequest.GetBusinessCustomSignParams()

	if req.Description != nil {
		params["description"] = *req.Description
	}
	if req.Entitys != nil {
		params["entitys"] = *req.Entitys
	}
	if req.BusinessId != nil {
		params["businessId"] = *req.BusinessId
	}
	return params
}

// SetDescription sets the Description field.
func (req *PretreatmentAddRequest) SetDescription(description string) {
	req.Description = &description
}

// SetEntitys sets the Entitys field.
func (req *PretreatmentAddRequest) SetEntitys(entitys string) {
	req.Entitys = &entitys
}

// SetBusinessId sets the BusinessId field.
func (req *PretreatmentAddRequest) SetBusinessId(businessId string) {
	req.BusinessId = &businessId
}
