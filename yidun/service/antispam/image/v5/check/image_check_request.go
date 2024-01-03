package check

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/yidun/yidun-golang-sdk/yidun/core/validation"
)

type ImageV5CheckRequest struct {
	*ImageCheckSceneRequest
	PublishTime *int64              `json:"publishTime,omitempty"` // timestamp milliseconds
	Images      *[]ImageBeanRequest `json:"images"`                // 图片检测列表参数
	Email       *string             `json:"email,omitempty"`       // 邮箱(反作弊用)
	DataType    *string             `json:"dataType,omitempty"`    // 数据类型
	Extension   *string             `json:"extension,omitempty"`   // 扩展字段(反作弊用)
	CheckLabels *[]int              `json:"checkLabels,omitempty"` // 图片过检分类
	ExtStr1     *string             `json:"extStr1,omitempty"`
	ExtStr2     *string             `json:"extStr2,omitempty"`
	ExtLon1     *int64              `json:"extLon1,omitempty"`
	ExtLon2     *int64              `json:"extLon2,omitempty"`
	CensorExt   *string             `json:"censorExt,omitempty"` // 人审扩展字段，用于人审调度中心的规则匹配
	SubProduct  *string             `json:"subProduct,omitempty"`
}

type ImageBeanRequest struct {
	DataId      *string `json:"dataId,omitempty"`      // 客户图片唯一标识
	Name        *string `json:"name"`                  // 图片名称
	Type        *int    `json:"type"`                  // 图片data字段的类型，包括url，base64编码等，具体定义见{@link ImageTypeEnum}
	Data        *string `json:"data"`                  // 图片内容，根据type字段传不同的值
	CallbackUrl *string `json:"callbackUrl,omitempty"` // 主动回调url
	TaskId      *string `json:"taskId,omitempty"`      // taskId
}

// For ImageV5CheckRequest
func NewImageV5CheckRequest(businessId string) *ImageV5CheckRequest {
	var checkSceneRequest = NewImageCheckSceneRequest(businessId)
	var request = &ImageV5CheckRequest{
		ImageCheckSceneRequest: checkSceneRequest,
	}
	checkSceneRequest.SetProductCode("imageCheck")
	checkSceneRequest.SetUriPattern("/v5/image/check")
	checkSceneRequest.SetVersion("v5.2")
	return request
}

func (r *ImageV5CheckRequest) SetImageCheckSceneRequest(imageCheckSceneRequest *ImageCheckSceneRequest) {
	r.ImageCheckSceneRequest = imageCheckSceneRequest
}

func (r *ImageV5CheckRequest) SetPublishTime(publishTime int64) {
	r.PublishTime = &publishTime
}

func (r *ImageV5CheckRequest) SetImages(images []ImageBeanRequest) {
	r.Images = &images
}

func (r *ImageV5CheckRequest) SetEmail(email string) {
	r.Email = &email
}

func (r *ImageV5CheckRequest) SetDataType(dataType string) {
	r.DataType = &dataType
}

func (r *ImageV5CheckRequest) SetExtension(extension string) {
	r.Extension = &extension
}

func (r *ImageV5CheckRequest) SetCheckLabels(checkLabels []int) {
	r.CheckLabels = &checkLabels
}

func (r *ImageV5CheckRequest) SetExtStr1(extStr1 string) {
	r.ExtStr1 = &extStr1
}

func (r *ImageV5CheckRequest) SetExtStr2(extStr2 string) {
	r.ExtStr2 = &extStr2
}

func (r *ImageV5CheckRequest) SetExtLon1(extLon1 int64) {
	r.ExtLon1 = &extLon1
}

func (r *ImageV5CheckRequest) SetExtLon2(extLon2 int64) {
	r.ExtLon2 = &extLon2
}

func (r *ImageV5CheckRequest) SetCensorExt(censorExt string) {
	r.CensorExt = &censorExt
}

func (r *ImageV5CheckRequest) SetSubProduct(subProduct string) {
	r.SubProduct = &subProduct
}

// For ImageBeanRequest
func NewImageBeanRequest() *ImageBeanRequest {
	return &ImageBeanRequest{}
}

func (r *ImageBeanRequest) SetDataId(dataId string) {
	r.DataId = &dataId
}

func (r *ImageBeanRequest) SetName(name string) {
	r.Name = &name
}

func (r *ImageBeanRequest) SetType(_type int) {
	r.Type = &_type
}

func (r *ImageBeanRequest) SetData(data string) {
	r.Data = &data
}

func (r *ImageBeanRequest) SetCallbackUrl(callbackUrl string) {
	r.CallbackUrl = &callbackUrl
}

func (r *ImageBeanRequest) SetTaskId(taskId string) {
	r.TaskId = &taskId
}

func (r *ImageV5CheckRequest) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "ImageCheckRequest"}
	if r.Images == nil {
		invalidParams.Add(validation.NewErrParamRequired("Images"))
	}

	for _, image := range *r.Images {
		if image.Name == nil || *image.Name == "" {
			invalidParams.Add(validation.NewErrParamRequired("图片名字不能为空"))
		}

		if image.Type == nil || *image.Type < 1 || *image.Type > 2 {
			invalidParams.Add(validation.NewErrParamRequired("图片数据类型不合法"))
		}

		if image.Data == nil || *image.Data == "" {
			invalidParams.Add(validation.NewErrParamRequired("图片数据不能为空"))
		}
	}

	if r.CensorExt != nil && len(*r.CensorExt) > 1024 {
		invalidParams.Add(validation.NewErrParamMaxLen("人审扩展字段长度不能超过1024", 1024, strconv.Itoa(len(*r.CensorExt))))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (r *ImageV5CheckRequest) GetBusinessCustomSignParams() map[string]string {
	result := r.ImageCheckSceneRequest.GetBusinessCustomSignParams()

	if r.PublishTime != nil {
		result["publishTime"] = strconv.FormatInt(*r.PublishTime, 10)
	}
	if r.Email != nil {
		result["email"] = *r.Email
	}
	if r.DataType != nil {
		result["dataType"] = *r.DataType
	}
	if r.Extension != nil {
		result["extension"] = *r.Extension
	}
	if r.ExtStr1 != nil {
		result["extStr1"] = *r.ExtStr1
	}
	if r.ExtStr2 != nil {
		result["extStr2"] = *r.ExtStr2
	}
	if r.ExtLon1 != nil {
		result["extLon1"] = strconv.FormatInt(*r.ExtLon1, 10)
	}
	if r.ExtLon2 != nil {
		result["extLon2"] = strconv.FormatInt(*r.ExtLon2, 10)
	}
	if r.CensorExt != nil {
		result["censorExt"] = *r.CensorExt
	}
	if r.SubProduct != nil {
		result["subProduct"] = *r.SubProduct
	}
	if r.Images != nil {
		imagesJson, _ := json.Marshal(r.Images)
		result["images"] = string(imagesJson)
	}
	if r.CheckLabels != nil {
		checkLabelsJson, _ := json.Marshal(r.CheckLabels)
		result["checkLabels"] = strings.Trim(string(checkLabelsJson), "[]")
	}

	return result
}
