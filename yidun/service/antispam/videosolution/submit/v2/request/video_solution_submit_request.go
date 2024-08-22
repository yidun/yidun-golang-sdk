package request

import (
	"encoding/json"
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/core/validation"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/video/check/v4/request"
	"strconv"
)

// VideoSolutionSubmitV2Request 视频解决方案提交请求
type VideoSolutionSubmitV2Request struct {
	*types.PostFormRequest
	// 用户IP地址
	IP *string `json:"ip,omitempty"`
	// 用户账号
	Account *string `json:"account,omitempty"`
	// 用户设备id
	DeviceID *string `json:"deviceId,omitempty"`
	// 用户设备类型
	DeviceType *string `json:"deviceType,omitempty"`
	// 数据唯一标识，能够根据该值定位到该条数据，如对数据结果有异议，可以发送该值给客户经理查询
	DataID *string `json:"dataId,omitempty"`
	// 文件标题
	Title *string `json:"title,omitempty"`
	// 图片检测列表参数
	Images *[]ImageBeanRequest `json:"images,omitempty"`
	// 音视频url
	URL *string `json:"url,omitempty"`
	// 回调数据
	Callback *string `json:"callback,omitempty"`
	// 回调地址
	CallbackURL *string `json:"callbackUrl,omitempty"`
	// 截图帧数，默认5
	ScFrequency *string `json:"scFrequency,omitempty"`
	// 高级截图频率配置，结构是json结构
	AdvancedFrequency *request.AdvancedFrequencyRequest `json:"advancedFrequency,omitempty"`
	// 简介内容
	Content *string `json:"content,omitempty"`
	// 去重key
	UniqueKey *string `json:"uniqueKey,omitempty"`
	// 指定过检测类型
	DetectType *int `json:"detectType,omitempty"`
	// 固定截图数
	CheckFrameCount *int `json:"checkFrameCount,omitempty"`
	// 用于聚合的id
	RelationID *string `json:"relationId,omitempty"`
	// 关联查询id
	RelationSearchID *string `json:"relationSearchId,omitempty"`
	// 自定义排序字段
	PriorityID *int64 `json:"priorityId,omitempty"`
	// 用户自定义标签
	Tags *[]string `json:"tags,omitempty"`
	// 指定语言检测音频内容，需以易盾标准填入
	CheckLanguageCode *string `json:"checkLanguageCode,omitempty"`
	Nickname          *string `json:"nickname,omitempty"`
	SubProduct        *string `json:"subProduct,omitempty"`
	Extension         *string `json:"extension,omitempty"`
}

type ImageBeanRequest struct {
	// 图片唯一标识
	Name *string `json:"name,omitempty"`
	// 图片data字段的类型，包括url，base64编码等，具体定义见
	// @see com.netease.yidun.sdk.antispam.image.v5.enums.ImageTypeEnum
	Type *int `json:"type,omitempty"`
	// 图片内容，根据type字段传不同的值
	Data *string `json:"data,omitempty"`
}

func NewVideoSolutionSubmitV2Req() *VideoSolutionSubmitV2Request {
	var request = &VideoSolutionSubmitV2Request{
		PostFormRequest: types.NewPostFormRequest(),
	}
	request.SetProductCode("videoSolutionCheck")
	request.SetUriPattern("/v2/videosolution/submit")
	request.SetVersion("v2")
	return request
}
func (r *VideoSolutionSubmitV2Request) GetBusinessCustomSignParams() map[string]string {
	parentParams := r.PostFormRequest.GetBusinessCustomSignParams()
	if r.URL != nil {
		parentParams["url"] = *r.URL
	}
	if r.IP != nil {
		parentParams["ip"] = *r.IP
	}

	if r.Account != nil {
		parentParams["account"] = *r.Account
	}

	if r.DeviceID != nil {
		parentParams["deviceId"] = *r.DeviceID
	}

	if r.DeviceType != nil {
		parentParams["deviceType"] = *r.DeviceType
	}

	if r.DataID != nil {
		parentParams["dataId"] = *r.DataID
	}

	if r.Title != nil {
		parentParams["title"] = *r.Title
	}

	if r.Images != nil {
		imagesJSON, _ := json.Marshal(r.Images)
		parentParams["images"] = string(imagesJSON)
	}

	if r.URL != nil {
		parentParams["url"] = *r.URL
	}

	if r.Callback != nil {
		parentParams["callback"] = *r.Callback
	}

	if r.CallbackURL != nil {
		parentParams["callbackUrl"] = *r.CallbackURL
	}

	if r.ScFrequency != nil {
		parentParams["scFrequency"] = *r.ScFrequency
	}

	if r.AdvancedFrequency != nil {
		advFreqJSON, _ := json.Marshal(r.AdvancedFrequency)
		parentParams["advancedFrequency"] = string(advFreqJSON)
	}

	if r.Content != nil {
		parentParams["content"] = *r.Content
	}

	if r.UniqueKey != nil {
		parentParams["uniqueKey"] = *r.UniqueKey
	}

	if r.DetectType != nil {
		parentParams["detectType"] = strconv.Itoa(*r.DetectType)
	}

	if r.CheckFrameCount != nil {
		parentParams["checkFrameCount"] = strconv.Itoa(*r.CheckFrameCount)
	}

	if r.RelationID != nil {
		parentParams["relationId"] = *r.RelationID
	}

	if r.RelationSearchID != nil {
		parentParams["relationSearchId"] = *r.RelationSearchID
	}

	if r.PriorityID != nil {
		parentParams["priorityId"] = strconv.FormatInt(*r.PriorityID, 10)
	}

	if r.Tags != nil {
		tagsJSON, _ := json.Marshal(r.Tags)
		parentParams["tags"] = string(tagsJSON)
	}

	if r.CheckLanguageCode != nil {
		parentParams["checkLanguageCode"] = *r.CheckLanguageCode
	}

	if r.Nickname != nil {
		parentParams["nickname"] = *r.Nickname
	}

	if r.SubProduct != nil {
		parentParams["subProduct"] = *r.SubProduct
	}
	if r.Extension != nil {
		parentParams["extension"] = *r.Extension
	}
	return parentParams
}

func (r *VideoSolutionSubmitV2Request) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "VideoSolutionSubmitV2Request"}
	if r == nil {
		invalidParams.Add(validation.NewErrParamRequired("VideoSolutionSubmitV2Request"))
	}
	if r.URL == nil {
		invalidParams.Add(validation.NewErrParamRequired("URL"))
	}

	if r.URL != nil && len(*r.URL) > 2048 {
		invalidParams.Add(validation.NewErrParamMaxLen("URL", 2048, strconv.Itoa(len(*r.URL))))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (r *ImageBeanRequest) SetName(name string) {
	r.Name = &name
}
func (r *ImageBeanRequest) SetType(t int) {
	r.Type = &t
}
func (r *ImageBeanRequest) SetData(data string) {
	r.Data = &data
}

func (r *VideoSolutionSubmitV2Request) SetIP(ip string) {
	r.IP = &ip
}

func (r *VideoSolutionSubmitV2Request) SetAccount(account string) {
	r.Account = &account
}

func (r *VideoSolutionSubmitV2Request) SetDeviceID(deviceID string) {
	r.DeviceID = &deviceID
}

func (r *VideoSolutionSubmitV2Request) SetDeviceType(deviceType string) {
	r.DeviceType = &deviceType
}

func (r *VideoSolutionSubmitV2Request) SetDataID(dataID string) {
	r.DataID = &dataID
}

func (r *VideoSolutionSubmitV2Request) SetTitle(title string) {
	r.Title = &title
}

func (r *VideoSolutionSubmitV2Request) SetImages(images []ImageBeanRequest) {
	r.Images = &images
}

func (r *VideoSolutionSubmitV2Request) SetURL(url string) {
	r.URL = &url
}

func (r *VideoSolutionSubmitV2Request) SetCallback(callback string) {
	r.Callback = &callback
}

func (r *VideoSolutionSubmitV2Request) SetCallbackURL(callbackURL string) {
	r.CallbackURL = &callbackURL
}

func (r *VideoSolutionSubmitV2Request) SetScFrequency(scFrequency string) {
	r.ScFrequency = &scFrequency
}

func (r *VideoSolutionSubmitV2Request) SetAdvancedFrequency(advancedFrequency request.AdvancedFrequencyRequest) {
	r.AdvancedFrequency = &advancedFrequency
}

func (r *VideoSolutionSubmitV2Request) SetContent(content string) {
	r.Content = &content
}

func (r *VideoSolutionSubmitV2Request) SetUniqueKey(uniqueKey string) {
	r.UniqueKey = &uniqueKey
}

func (r *VideoSolutionSubmitV2Request) SetDetectType(detectType int) {
	r.DetectType = &detectType
}

func (r *VideoSolutionSubmitV2Request) SetCheckFrameCount(checkFrameCount int) {
	r.CheckFrameCount = &checkFrameCount
}

func (r *VideoSolutionSubmitV2Request) SetRelationID(relationID string) {
	r.RelationID = &relationID
}

func (r *VideoSolutionSubmitV2Request) SetRelationSearchID(relationSearchID string) {
	r.RelationSearchID = &relationSearchID
}

func (r *VideoSolutionSubmitV2Request) SetPriorityID(priorityID int64) {
	r.PriorityID = &priorityID
}

func (r *VideoSolutionSubmitV2Request) SetTags(tags []string) {
	r.Tags = &tags
}

func (r *VideoSolutionSubmitV2Request) SetCheckLanguageCode(checkLanguageCode string) {
	r.CheckLanguageCode = &checkLanguageCode
}

func (r *VideoSolutionSubmitV2Request) SetNickname(nickname string) {
	r.Nickname = &nickname
}

func (r *VideoSolutionSubmitV2Request) SetSubProduct(subProduct string) {
	r.SubProduct = &subProduct
}
func (r *VideoSolutionSubmitV2Request) SetExtension(extension string) {
	r.Extension = &extension
}
