package request

import (
	"strconv"
	"strings"

	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
)

type LiveAudioSubmitV4Req struct {
	*types.BizPostFormRequest
	// 直播语音url
	Url *string `json:"url,omitempty"`
	// 直播音频标题
	Title *string `json:"title,omitempty"`
	// 用户IP地址
	Ip *string `json:"ip,omitempty"`
	// 用户唯一标识，如果无需登录则为空
	Account *string `json:"account,omitempty"`
	// 用户设备 id
	DeviceId *string `json:"deviceId,omitempty"`
	// 用户设备类型
	DeviceType *int `json:"deviceType,omitempty"`
	// 数据回调参数，调用方根据业务情况自行设计
	Callback *string `json:"callback,omitempty"`
	// 异步检测结果回调通知的URL
	CallbackUrl *string `json:"callbackUrl,omitempty"`
	// 主播房间号
	RoomNo *string `json:"roomNo,omitempty"`
	// 主播昵称
	AccountName *string `json:"accountName,omitempty"`
	// 主播头像
	Photo *string `json:"photo,omitempty"`
	// 背景图
	BackgroundImage *string `json:"backgroundImage,omitempty"`
	// 封面
	Cover *string `json:"cover,omitempty"`
	// dataId
	DataId *string `json:"dataId,omitempty"`
	// 主播等级
	AccountLevel *string `json:"accountLevel,omitempty"`
	// uniqueKey
	UniqueKey *string `json:"uniqueKey,omitempty"`
	// 指定语言检测音频内容，需以易盾标准填入
	CheckLanguageCode *string `json:"checkLanguageCode,omitempty"`
	// 发布时间
	PublishTime *int64 `json:"publishTime,omitempty"`
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
	// 主播所属工会
	LabourUnion *string `json:"labourUnion,omitempty"`
	// 运营管理者
	OperationManager *string `json:"operationManager,omitempty"`
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
	// 聊天室编号
	RoomId *string `json:"roomId,omitempty"`
	// 指定监听必审列表范围内的数据 speakerId
	CheckSpeakerIds *[]string `json:"checkSpeakerIds,omitempty"`
	// 指定不监听信任用户列表范围内的数据 speakerId
	NoCheckSpeakerIds *[]string `json:"noCheckSpeakerIds,omitempty"`
	// 扩展字段
	ExtLon1 *int64  `json:"extLon1,omitempty"`
	ExtLon2 *int64  `json:"extLon2,omitempty"`
	ExtStr1 *string `json:"extStr1,omitempty"`
	ExtStr2 *string `json:"extStr2,omitempty"`
	// 业务扩展字段,json字符串
	Extension *string `json:"extension,omitempty"`
	// 子产品
	SubProduct *string `json:"subProduct,omitempty"`
}

// NewLiveAudioBarrageV1Req 是一个创建 LiveAudioBarrageV1Req 的函数
func NewLiveAudioSubmitV4Req(businessId string) *LiveAudioSubmitV4Req {
	var request = &LiveAudioSubmitV4Req{
		BizPostFormRequest: types.NewBizPostFormRequest(businessId),
	}
	request.SetProductCode("liveAudio")
	request.SetUriPattern("/v4/liveaudio/check")
	request.SetVersion("v4")
	return request
}

// GetCustomSignParams 获取自定义签名参数
func (l *LiveAudioSubmitV4Req) GetBusinessCustomSignParams() map[string]string {
	result := l.BizPostFormRequest.GetBusinessCustomSignParams()
	if l.Url != nil {
		result["url"] = *l.Url
	}
	if l.Title != nil {
		result["title"] = *l.Title
	}
	if l.Ip != nil {
		result["ip"] = *l.Ip
	}
	if l.Account != nil {
		result["account"] = *l.Account
	}
	if l.DeviceId != nil {
		result["deviceId"] = *l.DeviceId
	}
	if l.DeviceType != nil {
		result["deviceType"] = strconv.Itoa(*l.DeviceType)
	}
	if l.Callback != nil {
		result["callback"] = *l.Callback
	}
	if l.CallbackUrl != nil {
		result["callbackUrl"] = *l.CallbackUrl
	}
	if l.RoomNo != nil {
		result["roomNo"] = *l.RoomNo
	}
	if l.AccountName != nil {
		result["accountName"] = *l.AccountName
	}
	if l.Photo != nil {
		result["photo"] = *l.Photo
	}
	if l.BackgroundImage != nil {
		result["backgroundImage"] = *l.BackgroundImage
	}
	if l.Cover != nil {
		result["cover"] = *l.Cover
	}
	if l.DataId != nil {
		result["dataId"] = *l.DataId
	}
	if l.AccountLevel != nil {
		result["accountLevel"] = *l.AccountLevel
	}
	if l.UniqueKey != nil {
		result["uniqueKey"] = *l.UniqueKey
	}
	if l.CheckLanguageCode != nil {
		result["checkLanguageCode"] = *l.CheckLanguageCode
	}
	if l.PublishTime != nil {
		result["publishTime"] = strconv.FormatInt(*l.PublishTime, 10)
	}
	if l.Nickname != nil {
		result["nickname"] = *l.Nickname
	}
	if l.Gender != nil {
		result["gender"] = strconv.Itoa(*l.Gender)
	}
	if l.Age != nil {
		result["age"] = strconv.Itoa(*l.Age)
	}
	if l.Level != nil {
		result["level"] = strconv.Itoa(*l.Level)
	}
	if l.RegisterTime != nil {
		result["registerTime"] = strconv.FormatInt(*l.RegisterTime, 10)
	}
	if l.FriendNum != nil {
		result["friendNum"] = strconv.FormatInt(*l.FriendNum, 10)
	}
	if l.FansNum != nil {
		result["fansNum"] = strconv.FormatInt(*l.FansNum, 10)
	}
	if l.IsPremiumUse != nil {
		result["isPremiumUse"] = strconv.Itoa(*l.IsPremiumUse)
	}
	if l.Role != nil {
		result["role"] = *l.Role
	}
	if l.LabourUnion != nil {
		result["labourUnion"] = *l.LabourUnion
	}
	if l.OperationManager != nil {
		result["operationManager"] = *l.OperationManager
	}
	if l.Phone != nil {
		result["phone"] = *l.Phone
	}
	if l.Token != nil {
		result["token"] = *l.Token
	}
	if l.Mac != nil {
		result["mac"] = *l.Mac
	}
	if l.Imei != nil {
		result["imei"] = *l.Imei
	}
	if l.Idfa != nil {
		result["idfa"] = *l.Idfa
	}
	if l.Idfv != nil {
		result["idfv"] = *l.Idfv
	}
	if l.AppVersion != nil {
		result["appVersion"] = *l.AppVersion
	}
	if l.RoomId != nil {
		result["roomId"] = *l.RoomId
	}
	checkSpeakerIds := l.CheckSpeakerIds
	if checkSpeakerIds != nil && len(*checkSpeakerIds) > 0 {
		result["checkSpeakerIds"] = strings.Join(*l.CheckSpeakerIds, ",")
	}
	// Convert NoCheckSpeakerIds to comma-separated string
	noCheckSpeakerIds := l.NoCheckSpeakerIds
	if noCheckSpeakerIds != nil && len(*noCheckSpeakerIds) > 0 {
		result["noCheckSpeakerIds"] = strings.Join(*l.NoCheckSpeakerIds, ",")
	}
	if l.ExtLon1 != nil {
		result["extLon1"] = strconv.FormatInt(*l.ExtLon1, 10)
	}
	if l.ExtLon2 != nil {
		result["extLon2"] = strconv.FormatInt(*l.ExtLon2, 10)
	}
	if l.ExtStr1 != nil {
		result["extStr1"] = *l.ExtStr1
	}
	if l.ExtStr2 != nil {
		result["extStr2"] = *l.ExtStr2
	}
	if l.Extension != nil {
		result["extension"] = *l.Extension
	}
	if l.SubProduct != nil {
		result["subProduct"] = *l.SubProduct
	}
	return result
}

// SetUrl sets the Url field.
func (l *LiveAudioSubmitV4Req) SetUrl(url string) {
	l.Url = &url
}

// SetTitle sets the Title field.
func (l *LiveAudioSubmitV4Req) SetTitle(title string) {
	l.Title = &title
}

// SetIp sets the Ip field.
func (l *LiveAudioSubmitV4Req) SetIp(ip string) {
	l.Ip = &ip
}

// SetAccount sets the Account field.
func (l *LiveAudioSubmitV4Req) SetAccount(account string) {
	l.Account = &account
}

// SetDeviceId sets the DeviceId field.
func (l *LiveAudioSubmitV4Req) SetDeviceId(deviceId string) {
	l.DeviceId = &deviceId
}

// SetDeviceType sets the DeviceType field.
func (l *LiveAudioSubmitV4Req) SetDeviceType(deviceType int) {
	l.DeviceType = &deviceType
}

// SetCallback sets the Callback field.
func (l *LiveAudioSubmitV4Req) SetCallback(callback string) {
	l.Callback = &callback
}

// SetCallbackUrl sets the CallbackUrl field.
func (l *LiveAudioSubmitV4Req) SetCallbackUrl(callbackUrl string) {
	l.CallbackUrl = &callbackUrl
}

// SetRoomNo sets the RoomNo field.
func (l *LiveAudioSubmitV4Req) SetRoomNo(roomNo string) {
	l.RoomNo = &roomNo
}

// SetAccountName sets the AccountName field.
func (l *LiveAudioSubmitV4Req) SetAccountName(accountName string) {
	l.AccountName = &accountName
}

// SetPhoto sets the Photo field.
func (l *LiveAudioSubmitV4Req) SetPhoto(photo string) {
	l.Photo = &photo
}

// SetBackgroundImage sets the BackgroundImage field.
func (l *LiveAudioSubmitV4Req) SetBackgroundImage(backgroundImage string) {
	l.BackgroundImage = &backgroundImage
}

// SetCover sets the Cover field.
func (l *LiveAudioSubmitV4Req) SetCover(cover string) {
	l.Cover = &cover
}

// SetDataId sets the DataId field.
func (l *LiveAudioSubmitV4Req) SetDataId(dataId string) {
	l.DataId = &dataId
}

// SetAccountLevel sets the AccountLevel field.
func (l *LiveAudioSubmitV4Req) SetAccountLevel(accountLevel string) {
	l.AccountLevel = &accountLevel
}

// SetUniqueKey sets the UniqueKey field.
func (l *LiveAudioSubmitV4Req) SetUniqueKey(uniqueKey string) {
	l.UniqueKey = &uniqueKey
}

// SetCheckLanguageCode sets the CheckLanguageCode field.
func (l *LiveAudioSubmitV4Req) SetCheckLanguageCode(checkLanguageCode string) {
	l.CheckLanguageCode = &checkLanguageCode
}

// SetPublishTime sets the PublishTime field.
func (l *LiveAudioSubmitV4Req) SetPublishTime(publishTime int64) {
	l.PublishTime = &publishTime
}

// SetNickname sets the Nickname field.
func (l *LiveAudioSubmitV4Req) SetNickname(nickname string) {
	l.Nickname = &nickname
}

// SetGender sets the Gender field.
func (l *LiveAudioSubmitV4Req) SetGender(gender int) {
	l.Gender = &gender
}

// SetAge sets the Age field.
func (l *LiveAudioSubmitV4Req) SetAge(age int) {
	l.Age = &age
}

// SetLevel sets the Level field.
func (l *LiveAudioSubmitV4Req) SetLevel(level int) {
	l.Level = &level
}

// SetRegisterTime sets the RegisterTime field.
func (l *LiveAudioSubmitV4Req) SetRegisterTime(registerTime int64) {
	l.RegisterTime = &registerTime
}

// SetFriendNum sets the FriendNum field.
func (l *LiveAudioSubmitV4Req) SetFriendNum(friendNum int64) {
	l.FriendNum = &friendNum
}

// SetFansNum sets the FansNum field.
func (l *LiveAudioSubmitV4Req) SetFansNum(fansNum int64) {
	l.FansNum = &fansNum
}

// SetIsPremiumUse sets the IsPremiumUse field.
func (l *LiveAudioSubmitV4Req) SetIsPremiumUse(isPremiumUse int) {
	l.IsPremiumUse = &isPremiumUse
}

// SetRole sets the Role field.
func (l *LiveAudioSubmitV4Req) SetRole(role string) {
	l.Role = &role
}

// SetLabourUnion sets the LabourUnion
// SetLabourUnion sets the LabourUnion field.
func (l *LiveAudioSubmitV4Req) SetLabourUnion(labourUnion string) {
	l.LabourUnion = &labourUnion
}

// SetOperationManager sets the OperationManager field.
func (l *LiveAudioSubmitV4Req) SetOperationManager(operationManager string) {
	l.OperationManager = &operationManager
}

// SetPhone sets the Phone field.
func (l *LiveAudioSubmitV4Req) SetPhone(phone string) {
	l.Phone = &phone
}

// SetToken sets the Token field.
func (l *LiveAudioSubmitV4Req) SetToken(token string) {
	l.Token = &token
}

// SetMac sets the Mac field.
func (l *LiveAudioSubmitV4Req) SetMac(mac string) {
	l.Mac = &mac
}

// SetImei sets the Imei field.
func (l *LiveAudioSubmitV4Req) SetImei(imei string) {
	l.Imei = &imei
}

// SetIdfa sets the Idfa field.
func (l *LiveAudioSubmitV4Req) SetIdfa(idfa string) {
	l.Idfa = &idfa
}

// SetIdfv sets the Idfv field.
func (l *LiveAudioSubmitV4Req) SetIdfv(idfv string) {
	l.Idfv = &idfv
}

// SetAppVersion sets the AppVersion field.
func (l *LiveAudioSubmitV4Req) SetAppVersion(appVersion string) {
	l.AppVersion = &appVersion
}

// SetRoomId sets the RoomId field.
func (l *LiveAudioSubmitV4Req) SetRoomId(roomId string) {
	l.RoomId = &roomId
}

// SetCheckSpeakerIds sets the CheckSpeakerIds field.
func (l *LiveAudioSubmitV4Req) SetCheckSpeakerIds(checkSpeakerIds []string) {
	l.CheckSpeakerIds = &checkSpeakerIds
}

// SetNoCheckSpeakerIds sets the NoCheckSpeakerIds field.
func (l *LiveAudioSubmitV4Req) SetNoCheckSpeakerIds(noCheckSpeakerIds []string) {
	l.NoCheckSpeakerIds = &noCheckSpeakerIds
}

// SetExtLon1 sets the ExtLon1 field.
func (l *LiveAudioSubmitV4Req) SetExtLon1(extLon1 int64) {
	l.ExtLon1 = &extLon1
}

// SetExtLon2 sets the ExtLon2 field.
func (l *LiveAudioSubmitV4Req) SetExtLon2(extLon2 int64) {
	l.ExtLon2 = &extLon2
}

// SetExtStr1 sets the ExtStr1 field.
func (l *LiveAudioSubmitV4Req) SetExtStr1(extStr1 string) {
	l.ExtStr1 = &extStr1
}

// SetExtStr2 sets the ExtStr2 field.
func (l *LiveAudioSubmitV4Req) SetExtStr2(extStr2 string) {
	l.ExtStr2 = &extStr2
}

// SetExtension sets the Extension field.
func (l *LiveAudioSubmitV4Req) SetExtension(extension string) {
	l.Extension = &extension
}

// SetSubProduct sets the SubProduct field.
func (l *LiveAudioSubmitV4Req) SetSubProduct(subProduct string) {
	l.SubProduct = &subProduct
}
