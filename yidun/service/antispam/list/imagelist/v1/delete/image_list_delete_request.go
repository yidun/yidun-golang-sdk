package delete

import (
	"strconv"

	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/core/validation"
)

// ImageListDeleteRequest 图片名单删除请求
type ImageListDeleteRequest struct {
	*types.BizPostFormRequest
	// 图片名单唯一标识
	Uuid *string `json:"uuid,omitempty"`
	// 名单库（黑白或者md5）ImageStorageSwitchEnum
	Type *int `json:"type,omitempty"`
	// 图片md5，uuid，url和md5必须传一个,且如果只传了md5或者url必须传businessId
	Md5 *string `json:"md5,omitempty"`
	// 图片url，uuid，url和md5必须传一个,且如果只传了md5或者url必须传businessId
	Url *string `json:"url,omitempty"`
	// 业务id
	BusinessId *string `json:"businessId,omitempty"`
}

// 参数校验方法
func (r *ImageListDeleteRequest) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "ImageListDeleteRequest"}
	if r.Type == nil {
		invalidParams.Add(validation.NewErrParamRequired("Type"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// NewImageListDeleteRequest 初始化 ImageListDeleteRequest 对象
func NewImageListDeleteRequest(businessId string) *ImageListDeleteRequest {
	var request = &ImageListDeleteRequest{
		BizPostFormRequest: types.NewBizPostFormRequest(businessId),
	}
	request.SetProductCode("list")
	request.SetUriPattern("/v1/image/list/delete")
	request.SetMethod(http.HttpMethodPost)
	request.SetVersion("v1")
	return request
}

// GetCustomSignParams returns the custom sign parameters for the request
func (req *ImageListDeleteRequest) GetBusinessCustomSignParams() map[string]string {
	params := req.PostFormRequest.GetBusinessCustomSignParams()
	if req.Uuid != nil {
		params["uuid"] = *req.Uuid
	}
	if req.Type != nil {
		params["type"] = strconv.Itoa(*req.Type)
	}
	if req.Md5 != nil {
		params["md5"] = *req.Md5
	}
	if req.Url != nil {
		params["url"] = *req.Url
	}
	if req.BusinessId != nil {
		params["businessId"] = *req.BusinessId
	}
	return params
}

// SetUuid sets the Uuid field.
func (i *ImageListDeleteRequest) SetUuid(Uuid string) {
	i.Uuid = &Uuid
}

// SetType sets the Type field.
func (i *ImageListDeleteRequest) SetType(Type int) {
	i.Type = &Type
}

// SetMd5 sets the Md5 field.
func (i *ImageListDeleteRequest) SetMd5(Md5 string) {
	i.Md5 = &Md5
}

// SetUrl sets the Url field.
func (i *ImageListDeleteRequest) SetUrl(Url string) {
	i.Url = &Url
}

// SetBusinessId sets the BusinessId field.
func (i *ImageListDeleteRequest) SetBusinessId(BusinessId string) {
	i.BusinessId = &BusinessId
}
