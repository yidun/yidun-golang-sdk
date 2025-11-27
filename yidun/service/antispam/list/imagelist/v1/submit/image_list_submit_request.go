package submit

import (
	"strconv"

	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/core/validation"
)

// ImageListSubmitRequest 图片名单添加请求
type ImageListSubmitRequest struct {
	*types.BizPostFormRequest
	// 图片url或者base64，json数组格式，最多200条
	Images *string `json:"images,omitempty"`
	// 名单分类，1: 白名单，2: 黑名单
	ListType *int `json:"listType,omitempty"`
	// 名单分类（加黑名单时传），100：色情，110：性感，200：广告，210：二维码，300：暴恐，400：违禁，500：涉政
	ImageLabel *int `json:"imageLabel,omitempty"`
	// 名单二级分类
	ImageSubLabel *string `json:"imageSubLabel,omitempty"`
	// 名单类型，0: 相似名单，1: md5精确名单
	Type *int `json:"type,omitempty"`
	// 图片类型，包括url，base64编码等
	ImageType *int `json:"imageType,omitempty"`
	// 描述
	Description *string `json:"description,omitempty"`
	// 易盾分组
	TagGroup *string `json:"tagGroup,omitempty"`
	// 子标签
	TagName *string `json:"tagName,omitempty"`
	// 业务id
	BusinessId *string `json:"businessId,omitempty"`
	// 嫌疑级别
	SuspectLevel *int `json:"suspectLevel,omitempty"`
}

// 参数校验方法
func (r *ImageListSubmitRequest) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "ImageListSubmitRequest"}
	if r.Images == nil {
		invalidParams.Add(validation.NewErrParamRequired("Images"))
	}
	if r.ListType == nil {
		invalidParams.Add(validation.NewErrParamRequired("ListType"))
	}
	if r.ImageLabel == nil {
		invalidParams.Add(validation.NewErrParamRequired("ImageLabel"))
	}
	if r.Type == nil {
		invalidParams.Add(validation.NewErrParamRequired("Type"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// NewImageListSubmitRequest 初始化ImageListSubmitRequest对象
func NewImageListSubmitRequest(businessId string) *ImageListSubmitRequest {
	var request = &ImageListSubmitRequest{
		BizPostFormRequest: types.NewBizPostFormRequest(businessId),
	}
	request.SetProductCode("list")
	request.SetUriPattern("/v1/image/list/submit")
	request.SetMethod(http.HttpMethodPost)
	request.SetVersion("v1")
	return request
}

// SetImages sets the Images field.
func (req *ImageListSubmitRequest) SetImages(images string) {
	req.Images = &images
}

// SetListType sets the ListType field.
func (req *ImageListSubmitRequest) SetListType(listType int) {
	req.ListType = &listType
}

// SetImageLabel sets the ImageLabel field.
func (req *ImageListSubmitRequest) SetImageLabel(imageLabel int) {
	req.ImageLabel = &imageLabel
}

// SetImageSubLabel sets the ImageSubLabel field.
func (req *ImageListSubmitRequest) SetImageSubLabel(imageSubLabel string) {
	req.ImageSubLabel = &imageSubLabel
}

// SetType sets the Type field.
func (req *ImageListSubmitRequest) SetType(typ int) {
	req.Type = &typ
}

// SetImageType sets the ImageType field.
func (req *ImageListSubmitRequest) SetImageType(imageType int) {
	req.ImageType = &imageType
}

// SetDescription sets the Description field.
func (req *ImageListSubmitRequest) SetDescription(description string) {
	req.Description = &description
}

// SetTagGroup sets the TagGroup field.
func (req *ImageListSubmitRequest) SetTagGroup(tagGroup string) {
	req.TagGroup = &tagGroup
}

// SetTagName sets the TagName field.
func (req *ImageListSubmitRequest) SetTagName(tagName string) {
	req.TagName = &tagName
}

// SetBusinessId sets the BusinessId field.
func (req *ImageListSubmitRequest) SetBusinessId(businessId string) {
	req.BusinessId = &businessId
}

// SetSuspectLevel sets the SuspectLevel field.
func (req *ImageListSubmitRequest) SetSuspectLevel(suspectLevel int) {
	req.SuspectLevel = &suspectLevel
}

// GetCustomSignParams 获取具体业务中特有的需要参与签名计算的参数
func (req *ImageListSubmitRequest) GetBusinessCustomSignParams() map[string]string {
	params := req.PostFormRequest.GetBusinessCustomSignParams()

	if req.Images != nil {
		params["images"] = *req.Images
	}
	if req.ListType != nil {
		params["listType"] = strconv.Itoa(*req.ListType)
	}
	if req.ImageLabel != nil {
		params["imageLabel"] = strconv.Itoa(*req.ImageLabel)
	}
	if req.ImageSubLabel != nil {
		params["imageSubLabel"] = *req.ImageSubLabel
	}
	if req.Type != nil {
		params["type"] = strconv.Itoa(*req.Type)
	}
	if req.ImageType != nil {
		params["imageType"] = strconv.Itoa(*req.ImageType)
	}
	if req.Description != nil {
		params["description"] = *req.Description
	}
	if req.TagGroup != nil {
		params["tagGroup"] = *req.TagGroup
	}
	if req.TagName != nil {
		params["tagName"] = *req.TagName
	}
	if req.BusinessId != nil {
		params["businessId"] = *req.BusinessId
	}
	if req.SuspectLevel != nil {
		params["suspectLevel"] = strconv.Itoa(*req.SuspectLevel)
	}
	return params
}
