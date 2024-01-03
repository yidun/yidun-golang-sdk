package update

import (
	"strconv"

	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/core/validation"
)

// ImageListUpdateRequest 图片名单更新请求
type ImageListUpdateRequest struct {
	*types.BizPostFormRequest
	// 图片名单唯一标识
	Uuid *string `json:"uuid,omitempty"`
	// 名单库（黑白或者md5）ImageStorageSwitchEnum
	Type *int `json:"type,omitempty"`
	// 图片名单状态,ImageListStatus
	Status *int `json:"status,omitempty"`
	// 图片md5，uuid，url和md5必须传一个,且如果只传了md5或者url必须传businessId
	Md5 *string `json:"md5,omitempty"`
	// 图片url，uuid，url和md5必须传一个,且如果只传了md5或者url必须传businessId
	Url *string `json:"url,omitempty"`
	// 业务id
	BusinessId *string `json:"businessId,omitempty"`
}

// 参数校验方法
func (r *ImageListUpdateRequest) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "ImageListUpdateRequest"}
	if r.Type == nil {
		invalidParams.Add(validation.NewErrParamRequired("Type"))
	}
	if r.Status == nil {
		invalidParams.Add(validation.NewErrParamRequired("Status"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// NewImageListUpdateRequest 初始化 ImageListSubmitRequest 对象
func NewImageListUpdateRequest(businessId string) *ImageListUpdateRequest {
	var request = &ImageListUpdateRequest{
		BizPostFormRequest: types.NewBizPostFormRequest(businessId),
	}
	request.SetProductCode("list")
	request.SetUriPattern("/v1/image/list/update")
	request.SetMethod(http.HttpMethodPost)
	request.SetVersion("v1")
	return request
}

// GetCustomSignParams returns the custom sign parameters for the request
func (req *ImageListUpdateRequest) GetBusinessCustomSignParams() map[string]string {
	params := req.PostFormRequest.GetBusinessCustomSignParams()
	if req.Uuid != nil {
		params["uuid"] = *req.Uuid
	}
	if req.Type != nil {
		params["type"] = strconv.Itoa(*req.Type)
	}
	if req.Status != nil {
		params["status"] = strconv.Itoa(*req.Status)
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
func (req *ImageListUpdateRequest) SetUuid(uuid string) {
	req.Uuid = &uuid
}

// SetType sets the Type field.
func (req *ImageListUpdateRequest) SetType(typ int) {
	req.Type = &typ
}

// SetStatus sets the Status field.
func (req *ImageListUpdateRequest) SetStatus(status int) {
	req.Status = &status
}

// SetMd5 sets the Md5 field.
func (req *ImageListUpdateRequest) SetMd5(md5 string) {
	req.Md5 = &md5
}

// SetUrl sets the Url field.
func (req *ImageListUpdateRequest) SetUrl(url string) {
	req.Url = &url
}

// SetBusinessId sets the BusinessId field.
func (req *ImageListUpdateRequest) SetBusinessId(businessId string) {
	req.BusinessId = &businessId
}
