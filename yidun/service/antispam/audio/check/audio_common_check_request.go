package check

import (
	"encoding/json"
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"strconv"
)

// AudioCommonCheckRequest 音频检测请求
type AudioCommonCheckRequest struct {
	*types.BizPostFormRequest
	// 点播语音url
	URL *string `json:"url,omitempty"`
	// 点播语音标题
	Title *string `json:"title,omitempty"`
	// 用户IP地址
	IP *string `json:"ip,omitempty"`
	// 用户唯一标识，如果无需登录则为空
	Account *string `json:"account,omitempty"`
	// 用户设备 id
	DeviceID *string `json:"deviceId,omitempty"`
	// 用户设备类型
	DeviceType *string `json:"deviceType,omitempty"`
	// 数据回调参数，调用方根据业务情况自行设计
	Callback *string `json:"callback,omitempty"`
	// 异步检测结果回调通知的URL
	CallbackUrl *string `json:"callbackUrl,omitempty"`
	// 音频标签
	Tags *[]string `json:"tags,omitempty"`
	// dataId
	DataID *string `json:"dataId,omitempty"`
	// 上下文关联key列表
	RelatedKeys *[]string `json:"relatedKeys,omitempty"`
	// 指定语言检测音频内容，需以易盾标准填入
	CheckLanguageCode *string `json:"checkLanguageCode,omitempty"`
	// 发布时间
	PublishTime *int64 `json:"publishTime,omitempty"`
	// 去重key
	UniqueKey *string `json:"uniqueKey,omitempty"`
	// 客户提交的扩展字段，需要严格按照 json 格式
	Extension *string `json:"extension,omitempty"`
	// 自定义业务标识
	SubProduct *string `json:"subProduct,omitempty"`
	// 昵称
	Nickname *string `json:"nickname,omitempty"`
	// 性别
	Gender *int `json:"gender,omitempty"`
	// 年龄
	Age *int `json:"age,omitempty"`
	// 用户等级
	Level *int `json:"level,omitempty"`
	// 注册时间，UNIX 时间戳(毫秒值)
	RegisterTime *int64 `json:"registerTime,omitempty"`
	// 好友数
	FriendNum *int64 `json:"friendNum,omitempty"`
	// 粉丝数
	FansNum *int64 `json:"fansNum,omitempty"`
	// 是否付费
	IsPremiumUse *int `json:"isPremiumUse,omitempty"`
	// 用户类型
	Role *string `json:"role,omitempty"`
	// 手机号
	Phone *string `json:"phone,omitempty"`
	// token
	Token *string `json:"token,omitempty"`
	// 设备mac地址
	Mac *string `json:"mac,omitempty"`
	// Android设备imei信息
	Imei *string `json:"imei,omitempty"`
	// ios设备idfa信息
	Idfa *string `json:"idfa,omitempty"`
	// ios设备idfv信息
	Idfv *string `json:"idfv,omitempty"`
	// app版本号
	AppVersion *string `json:"appVersion,omitempty"`
	// 接收人id
	ReceiveUid *string `json:"receiveUid,omitempty"`
	// 好友关系，1接收人关注发送人，2发送人关注接收人，3互相关注，4互未关注
	Relationship *int `json:"relationship,omitempty"`
	// 群id
	GroupId *string `json:"groupId,omitempty"`
	// 聊天室编号
	RoomId *string `json:"roomId,omitempty"`
	// 动态id
	Topic *string `json:"topic,omitempty"`
	// 主评论id
	CommentId *string `json:"commentId,omitempty"`
	// 商品id
	CommodityId *string `json:"commodityId,omitempty"`
}

func NewAudioCommonCheckRequest(businessId string) *AudioCommonCheckRequest {
	request := &AudioCommonCheckRequest{
		BizPostFormRequest: types.NewBizPostFormRequest(businessId),
	}
	return request
}

func (r *AudioCommonCheckRequest) GetBusinessCustomSignParams() map[string]string {
	parentParams := r.BizPostFormRequest.GetBusinessCustomSignParams()
	if r.URL != nil {
		parentParams["url"] = *r.URL
	}
	if r.Title != nil {
		parentParams["title"] = *r.Title
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
	if r.Callback != nil {
		parentParams["callback"] = *r.Callback
	}
	if r.CallbackUrl != nil {
		parentParams["callbackUrl"] = *r.CallbackUrl
	}
	if r.Tags != nil {
		tagsJson, _ := json.Marshal(r.Tags)
		parentParams["tags"] = string(tagsJson)
	}
	if r.DataID != nil {
		parentParams["dataId"] = *r.DataID
	}
	if r.RelatedKeys != nil {
		relatedKeysJson, _ := json.Marshal(r.RelatedKeys)
		parentParams["relatedKeys"] = string(relatedKeysJson)
	}
	if r.CheckLanguageCode != nil {
		parentParams["checkLanguageCode"] = *r.CheckLanguageCode
	}
	if r.PublishTime != nil {
		parentParams["publishTime"] = strconv.FormatInt(*r.PublishTime, 10)
	}
	if r.UniqueKey != nil {
		parentParams["uniqueKey"] = *r.UniqueKey
	}
	if r.Extension != nil {
		parentParams["extension"] = *r.Extension
	}
	if r.SubProduct != nil {
		parentParams["subProduct"] = *r.SubProduct
	}
	if r.Nickname != nil {
		parentParams["nickname"] = *r.Nickname
	}
	if r.Gender != nil {
		parentParams["gender"] = strconv.Itoa(*r.Gender)
	}
	if r.Age != nil {
		parentParams["age"] = strconv.Itoa(*r.Age)
	}
	if r.Level != nil {
		parentParams["level"] = strconv.Itoa(*r.Level)
	}
	if r.RegisterTime != nil {
		parentParams["registerTime"] = strconv.FormatInt(*r.RegisterTime, 10)
	}
	if r.FriendNum != nil {
		parentParams["friendNum"] = strconv.FormatInt(*r.FriendNum, 10)
	}
	if r.FansNum != nil {
		parentParams["fansNum"] = strconv.FormatInt(*r.FansNum, 10)
	}
	if r.IsPremiumUse != nil {
		parentParams["isPremiumUse"] = strconv.Itoa(*r.IsPremiumUse)
	}
	if r.Role != nil {
		parentParams["role"] = *r.Role
	}
	if r.Phone != nil {
		parentParams["phone"] = *r.Phone
	}
	if r.Token != nil {
		parentParams["token"] = *r.Token
	}
	if r.Mac != nil {
		parentParams["mac"] = *r.Mac
	}
	if r.Imei != nil {
		parentParams["imei"] = *r.Imei
	}
	if r.Idfa != nil {
		parentParams["idfa"] = *r.Idfa
	}
	if r.Idfv != nil {
		parentParams["idfv"] = *r.Idfv
	}
	if r.AppVersion != nil {
		parentParams["appVersion"] = *r.AppVersion
	}
	if r.ReceiveUid != nil {
		parentParams["receiveUid"] = *r.ReceiveUid
	}
	if r.Relationship != nil {
		parentParams["relationship"] = strconv.Itoa(*r.Relationship)
	}
	if r.GroupId != nil {
		parentParams["groupId"] = *r.GroupId
	}
	if r.RoomId != nil {
		parentParams["roomId"] = *r.RoomId
	}
	if r.Topic != nil {
		parentParams["topic"] = *r.Topic
	}
	if r.CommentId != nil {
		parentParams["commentId"] = *r.CommentId
	}
	if r.CommodityId != nil {
		parentParams["commodityId"] = *r.CommodityId
	}
	return parentParams
}

// TODO
func (r *AudioCommonCheckRequest) ValidateParam() error {
	return nil
}

// 生成所有属性的 set 方法
func (r *AudioCommonCheckRequest) SetURL(value string) *AudioCommonCheckRequest {
	r.URL = &value
	return r
}
func (r *AudioCommonCheckRequest) SetTitle(value string) *AudioCommonCheckRequest {
	r.Title = &value
	return r
}
func (r *AudioCommonCheckRequest) SetIP(value string) *AudioCommonCheckRequest {
	r.IP = &value
	return r
}
func (r *AudioCommonCheckRequest) SetAccount(value string) *AudioCommonCheckRequest {
	r.Account = &value
	return r
}
func (r *AudioCommonCheckRequest) SetDeviceID(value string) *AudioCommonCheckRequest {
	r.DeviceID = &value
	return r
}
func (r *AudioCommonCheckRequest) SetDeviceType(value string) *AudioCommonCheckRequest {
	r.DeviceType = &value
	return r
}
func (r *AudioCommonCheckRequest) SetCallback(value string) *AudioCommonCheckRequest {
	r.Callback = &value
	return r
}
func (r *AudioCommonCheckRequest) SetCallbackUrl(value string) *AudioCommonCheckRequest {
	r.CallbackUrl = &value
	return r
}
func (r *AudioCommonCheckRequest) SetTags(value []string) *AudioCommonCheckRequest {
	r.Tags = &value
	return r
}
func (r *AudioCommonCheckRequest) SetDataID(value string) *AudioCommonCheckRequest {
	r.DataID = &value
	return r
}
func (r *AudioCommonCheckRequest) SetRelatedKeys(value []string) *AudioCommonCheckRequest {
	r.RelatedKeys = &value
	return r
}
func (r *AudioCommonCheckRequest) SetCheckLanguageCode(value string) *AudioCommonCheckRequest {
	r.CheckLanguageCode = &value
	return r
}
func (r *AudioCommonCheckRequest) SetPublishTime(value int64) *AudioCommonCheckRequest {
	r.PublishTime = &value
	return r
}
func (r *AudioCommonCheckRequest) SetUniqueKey(value string) *AudioCommonCheckRequest {
	r.UniqueKey = &value
	return r
}
func (r *AudioCommonCheckRequest) SetExtension(value string) *AudioCommonCheckRequest {
	r.Extension = &value
	return r
}
func (r *AudioCommonCheckRequest) SetSubProduct(value string) *AudioCommonCheckRequest {
	r.SubProduct = &value
	return r
}
func (r *AudioCommonCheckRequest) SetNickname(value string) *AudioCommonCheckRequest {
	r.Nickname = &value
	return r
}
func (r *AudioCommonCheckRequest) SetGender(value int) *AudioCommonCheckRequest {
	r.Gender = &value
	return r
}
func (r *AudioCommonCheckRequest) SetAge(value int) *AudioCommonCheckRequest {
	r.Age = &value
	return r
}
func (r *AudioCommonCheckRequest) SetLevel(value int) *AudioCommonCheckRequest {
	r.Level = &value
	return r

}
func (r *AudioCommonCheckRequest) SetRegisterTime(value int64) *AudioCommonCheckRequest {
	r.RegisterTime = &value
	return r
}
func (r *AudioCommonCheckRequest) SetFriendNum(value int64) *AudioCommonCheckRequest {
	r.FriendNum = &value
	return r
}
func (r *AudioCommonCheckRequest) SetFansNum(value int64) *AudioCommonCheckRequest {
	r.FansNum = &value
	return r
}
func (r *AudioCommonCheckRequest) SetIsPremiumUse(value int) *AudioCommonCheckRequest {
	r.IsPremiumUse = &value
	return r
}
func (r *AudioCommonCheckRequest) SetRole(value string) *AudioCommonCheckRequest {
	r.Role = &value
	return r
}
func (r *AudioCommonCheckRequest) SetPhone(value string) *AudioCommonCheckRequest {
	r.Phone = &value
	return r
}
func (r *AudioCommonCheckRequest) SetToken(value string) *AudioCommonCheckRequest {
	r.Token = &value
	return r
}
func (r *AudioCommonCheckRequest) SetMac(value string) *AudioCommonCheckRequest {
	r.Mac = &value
	return r
}
func (r *AudioCommonCheckRequest) SetImei(value string) *AudioCommonCheckRequest {
	r.Imei = &value
	return r
}
func (r *AudioCommonCheckRequest) SetIdfa(value string) *AudioCommonCheckRequest {
	r.Idfa = &value
	return r
}
func (r *AudioCommonCheckRequest) SetIdfv(value string) *AudioCommonCheckRequest {
	r.Idfv = &value
	return r
}

func (r *AudioCommonCheckRequest) SetAppVersion(value string) *AudioCommonCheckRequest {
	r.AppVersion = &value
	return r
}
func (r *AudioCommonCheckRequest) SetReceiveUid(value string) *AudioCommonCheckRequest {
	r.ReceiveUid = &value
	return r
}
func (r *AudioCommonCheckRequest) SetRelationship(value int) *AudioCommonCheckRequest {
	r.Relationship = &value
	return r
}
func (r *AudioCommonCheckRequest) SetGroupId(value string) *AudioCommonCheckRequest {
	r.GroupId = &value
	return r
}
func (r *AudioCommonCheckRequest) SetRoomId(value string) *AudioCommonCheckRequest {
	r.RoomId = &value
	return r
}
func (r *AudioCommonCheckRequest) SetTopic(value string) *AudioCommonCheckRequest {
	r.Topic = &value
	return r
}
func (r *AudioCommonCheckRequest) SetCommentId(value string) *AudioCommonCheckRequest {
	r.CommentId = &value
	return r
}
func (r *AudioCommonCheckRequest) SetCommodityId(value string) *AudioCommonCheckRequest {
	r.CommodityId = &value
	return r
}
