package request

import (
	"encoding/json"
	"fmt"
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/core/validation"
	"strconv"
)

// VideoCheckRequest 点播视频提交检测请求
type VideoCheckRequest struct {
	*types.BizPostFormRequest
	// 链接地址
	URL *string `json:"url,omitempty"`
	// 数据唯一标识
	DataID *string `json:"dataId,omitempty"`
	// 点播名称/标题
	Title *string `json:"title,omitempty"`
	// ip
	IP *string `json:"ip,omitempty"`
	// 账号
	Account *string `json:"account,omitempty"`
	// 设备类型
	DeviceType *string `json:"deviceType,omitempty"`
	// 设备id
	DeviceID *string `json:"deviceId,omitempty"`
	// 回调地址
	Callback *string `json:"callback,omitempty"`
	// 主动回调地址
	CallbackUrl *string `json:"callbackUrl,omitempty"`
	// 截图间隔
	ScFrequency *float32 `json:"scFrequency,omitempty"`
	// 高级截图频率配置，结构是json结构
	AdvancedFrequency *AdvancedFrequencyRequest `json:"advancedFrequency,omitempty"` // Assuming it's a JSON string
	// unique key
	UniqueKey *string `json:"uniqueKey,omitempty"`
	// 昵称
	Nickname *string `json:"nickname,omitempty"`
	// 手机号
	Phone *string `json:"phone,omitempty"`
	// 反作弊token
	Token *string `json:"token,omitempty"`
	// 性别， 0: 未知，1: 女性，2: 男性 非必填字段
	Gender *int `json:"gender,omitempty"`
	// 年龄，0: 未知 非必填字段
	Age *int `json:"age,omitempty"`
	// 固定截图数
    CheckFrameCount *int `json:"checkFrameCount,omitempty"`
	// 用户等级，0: 未知，1: 低，2: 中，3: 高 非必填字段
	Level *int `json:"level,omitempty"`
	// 注册时间 非必填字段
	RegisterTime *int64 `json:"registerTime,omitempty"`
	// 好友数 非必填字段
	FriendNum *int64 `json:"friendNum,omitempty"`
	// 粉丝数 非必填字段
	FansNum *int64 `json:"fansNum,omitempty"`
	// 是否付费用户，1:优质账号，0:默认 非必填字段
	IsPremiumUse *int `json:"isPremiumUse,omitempty"`
	// 用户角色 非必填字段
	Role *string `json:"role,omitempty"`
	// 设备唯一标识mac 非必填字段
	Mac *string `json:"mac,omitempty"`
	// 设备唯一标识imei 非必填字段
	Imei *string `json:"imei,omitempty"`
	// 设备唯一标识idfa 非必填字段
	Idfa *string `json:"idfa,omitempty"`
	// 设备唯一标识idfv 非必填字段
	Idfv *string `json:"idfv,omitempty"`
	// 文章/主贴/动态id 非必填字段
	Topic *string `json:"topic,omitempty"`
	// 私聊接收用户uid 非必填字段
	ReceiveUid *string `json:"receiveUid,omitempty"`
	// 群聊id 非必填字段
	GroupId *string `json:"groupId,omitempty"`
	// 房间id 非必填字段
	RoomId *string `json:"roomId,omitempty"`
	// 商品id 非必填字段
	CommodityId *string `json:"commodityId,omitempty"`
	// 主评论id 非必填字段
	CommentId *string `json:"commentId,omitempty"`
	// app版本号 非必填字段
	AppVersion *string `json:"appVersion,omitempty"`
	// 好友关系，1接收人关注发送人，2发送人关注接收人，3互相关注，4互未关注 非必填字段
	Relationship *int `json:"relationship,omitempty"`
	// 客户提交的额外信息
	Extension *string `json:"extension,omitempty"`
	// 自定义业务标识
	SubProduct *string `json:"subProduct,omitempty"`
}
type AdvancedFrequencyRequest struct {
	DurationPoints *[]int64   `json:"durationPoints,omitempty"`
	Frequencies    *[]float64 `json:"frequencies,omitempty"`
}

func (r *AdvancedFrequencyRequest) NewAdvancedFrequencyRequest(durationPoints []int64, frequencies []float64) *AdvancedFrequencyRequest {
	var req = &AdvancedFrequencyRequest{
		DurationPoints: &durationPoints,
		Frequencies:    &frequencies,
	}
	return req
}

func NewVideoCheckRequest(businessId string) *VideoCheckRequest {
	var request = &VideoCheckRequest{
		BizPostFormRequest: types.NewBizPostFormRequest(businessId),
	}
	request.SetProductCode("videoCheck")
	request.SetUriPattern("/v4/video/submit")
	request.SetVersion("v4.1")
	return request
}

func (r *VideoCheckRequest) GetBusinessCustomSignParams() map[string]string {
	parentParams := r.BizPostFormRequest.GetBusinessCustomSignParams()
	if r.URL != nil {
		parentParams["url"] = *r.URL
	}
	if r.DataID != nil {
		parentParams["dataId"] = *r.DataID
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
	if r.DeviceType != nil {
		parentParams["deviceType"] = *r.DeviceType
	}
	if r.DeviceID != nil {
		parentParams["deviceId"] = *r.DeviceID
	}
	if r.Callback != nil {
		parentParams["callback"] = *r.Callback
	}
	if r.CallbackUrl != nil {
		parentParams["callbackUrl"] = *r.CallbackUrl
	}
	if r.ScFrequency != nil {
		parentParams["scFrequency"] = fmt.Sprintf("%f", *r.ScFrequency)
	}
	if r.AdvancedFrequency != nil {
		jsonString, _ := json.Marshal(r.AdvancedFrequency)
		parentParams["advancedFrequency"] = string(jsonString)
	}
	if r.UniqueKey != nil {
		parentParams["uniqueKey"] = *r.UniqueKey
	}
	if r.Nickname != nil {
		parentParams["nickname"] = *r.Nickname
	}
	if r.Phone != nil {
		parentParams["phone"] = *r.Phone
	}
	if r.Token != nil {
		parentParams["token"] = *r.Token
	}
	if r.Gender != nil {
		parentParams["gender"] = strconv.Itoa(*r.Gender)
	}
	if r.Age != nil {
		parentParams["age"] = strconv.Itoa(*r.Age)
	}
	if r.CheckFrameCount != nil {
        parentParams["checkFrameCount"] = strconv.Itoa(*r.CheckFrameCount)
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
	if r.Topic != nil {
		parentParams["topic"] = *r.Topic
	}
	if r.ReceiveUid != nil {
		parentParams["receiveUid"] = *r.ReceiveUid
	}
	if r.GroupId != nil {
		parentParams["groupId"] = *r.GroupId
	}
	if r.RoomId != nil {
		parentParams["roomId"] = *r.RoomId
	}
	if r.CommodityId != nil {
		parentParams["commodityId"] = *r.CommodityId
	}
	if r.CommentId != nil {
		parentParams["commentId"] = *r.CommentId
	}
	if r.AppVersion != nil {
		parentParams["appVersion"] = *r.AppVersion
	}
	if r.Relationship != nil {
		parentParams["relationship"] = strconv.Itoa(*r.Relationship)
	}
	if r.Extension != nil {
		parentParams["extension"] = *r.Extension
	}
	if r.SubProduct != nil {
		parentParams["subProduct"] = *r.SubProduct
	}
	return parentParams
}

// TODO
func (r *VideoCheckRequest) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "VideoCheckRequest"}
	if r == nil {
		invalidParams.Add(validation.NewErrParamRequired("VideoCheckRequest"))
	}
	if r.URL == nil {
		invalidParams.Add(validation.NewErrParamRequired("URL"))
	}
	if r.DataID == nil {
		invalidParams.Add(validation.NewErrParamRequired("DataID"))
	}
	if r.URL != nil && len(*r.URL) > 2048 {
		invalidParams.Add(validation.NewErrParamMaxLen("URL", 2048, strconv.Itoa(len(*r.URL))))
	}
	if r.DataID != nil && len(*r.DataID) > 128 {
		invalidParams.Add(validation.NewErrParamMaxLen("DataID", 128, strconv.Itoa(len(*r.DataID))))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (r *VideoCheckRequest) SetURL(url string) {
	r.URL = &url
}

func (r *VideoCheckRequest) SetDataID(dataId string) {
	r.DataID = &dataId
}

func (r *VideoCheckRequest) SetTitle(title string) {
	r.Title = &title
}

func (r *VideoCheckRequest) SetIP(ip string) {
	r.IP = &ip
}

func (r *VideoCheckRequest) SetAccount(account string) {
	r.Account = &account
}

func (r *VideoCheckRequest) SetDeviceType(deviceType string) {
	r.DeviceType = &deviceType
}

func (r *VideoCheckRequest) SetDeviceID(deviceId string) {
	r.DeviceID = &deviceId
}

func (r *VideoCheckRequest) SetCallback(callback string) {
	r.Callback = &callback
}

func (r *VideoCheckRequest) SetCallbackUrl(callbackUrl string) {
	r.CallbackUrl = &callbackUrl
}

func (r *VideoCheckRequest) SetScFrequency(scFrequency float32) {
	r.ScFrequency = &scFrequency
}

func (r *VideoCheckRequest) SetAdvancedFrequency(advancedFrequency AdvancedFrequencyRequest) {
	r.AdvancedFrequency = &advancedFrequency
}

func (r *VideoCheckRequest) SetUniqueKey(uniqueKey string) {
	r.UniqueKey = &uniqueKey
}

func (r *VideoCheckRequest) SetNickname(nickname string) {
	r.Nickname = &nickname
}

func (r *VideoCheckRequest) SetPhone(phone string) {
	r.Phone = &phone
}

func (r *VideoCheckRequest) SetToken(token string) {
	r.Token = &token
}

func (r *VideoCheckRequest) SetGender(gender int) {
	r.Gender = &gender
}

func (r *VideoCheckRequest) SetAge(age int) {
	r.Age = &age
}
func (r *VideoCheckRequest) SetCheckFrameCount(checkFrameCount int) {
	r.CheckFrameCount = &checkFrameCount
}

func (r *VideoCheckRequest) SetLevel(level int) {
	r.Level = &level
}

func (r *VideoCheckRequest) SetRegisterTime(registerTime int64) {
	r.RegisterTime = &registerTime
}

func (r *VideoCheckRequest) SetFriendNum(friendNum int64) {
	r.FriendNum = &friendNum
}

func (r *VideoCheckRequest) SetFansNum(fansNum int64) {
	r.FansNum = &fansNum
}

func (r *VideoCheckRequest) SetIsPremiumUse(isPremiumUse int) {
	r.IsPremiumUse = &isPremiumUse
}

func (r *VideoCheckRequest) SetRole(role string) {
	r.Role = &role
}

func (r *VideoCheckRequest) SetMac(mac string) {
	r.Mac = &mac
}

func (r *VideoCheckRequest) SetImei(imei string) {
	r.Imei = &imei
}

func (r *VideoCheckRequest) SetIdfa(idfa string) {
	r.Idfa = &idfa
}

func (r *VideoCheckRequest) SetIdfv(idfv string) {
	r.Idfv = &idfv
}

func (r *VideoCheckRequest) SetTopic(topic string) {
	r.Topic = &topic
}

func (r *VideoCheckRequest) SetReceiveUid(receiveUid string) {
	r.ReceiveUid = &receiveUid
}

func (r *VideoCheckRequest) SetGroupId(groupId string) {
	r.GroupId = &groupId
}

func (r *VideoCheckRequest) SetRoomId(roomId string) {
	r.RoomId = &roomId
}

func (r *VideoCheckRequest) SetCommodityId(commodityId string) {
	r.CommodityId = &commodityId
}

func (r *VideoCheckRequest) SetCommentId(commentId string) {
	r.CommentId = &commentId
}

func (r *VideoCheckRequest) SetAppVersion(appVersion string) {
	r.AppVersion = &appVersion
}

func (r *VideoCheckRequest) SetRelationship(relationship int) {
	r.Relationship = &relationship
}

func (r *VideoCheckRequest) SetExtension(extension string) {
	r.Extension = &extension
}

func (r *VideoCheckRequest) SetSubProduct(subProduct string) {
	r.SubProduct = &subProduct
}
