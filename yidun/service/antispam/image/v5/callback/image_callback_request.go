package callback

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/core/validation"
)

type ImageV5CallbackRequest struct {
	*types.BizPostFormRequest
	YidunRequestId *string `json:"yidunRequestId,omitempty"`
}

func (r *ImageV5CallbackRequest) SetYidunRequestId(yidunRequestId string) {
	r.YidunRequestId = &yidunRequestId
}

// 构建request
func NewImageCallbackRequest(businessId string) *ImageV5CallbackRequest {
	var request = &ImageV5CallbackRequest{
		BizPostFormRequest: types.NewBizPostFormRequest(businessId),
	}
	request.SetProductCode("imageCommon")
	request.SetUriPattern("/v5/image/callback/results")
	request.SetVersion("v5.2")
	return request
}

// 获取业务自定义参数
func (r *ImageV5CallbackRequest) GetBusinessCustomSignParams() map[string]string {
	params := r.BizPostFormRequest.GetBusinessCustomSignParams()
	if r.YidunRequestId != nil {
		params["yidunRequestId"] = *r.YidunRequestId
	}

	return params
}

// 参数校验
func (r *ImageV5CallbackRequest) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "ImageV5CallbackRequest"}
	if r == nil {
		invalidParams.Add(validation.NewErrParamRequired("ImageV5CallbackRequest"))
		return invalidParams
	}

	return nil
}
