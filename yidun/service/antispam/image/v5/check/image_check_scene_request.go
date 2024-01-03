package check

import (
	"strconv"

	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
)

type ImageCheckSceneRequest struct {
	*types.BizPostFormRequest
	Nickname              *string `json:"nickname,omitempty"`              // 用户昵称nickname
	Account               *string `json:"account,omitempty"`               // 用户账号
	Gender                *int    `json:"gender,omitempty"`                // 性别
	Age                   *int    `json:"age,omitempty"`                   // 年龄
	Level                 *int    `json:"level,omitempty"`                 // 用户等级
	RegisterTime          *int64  `json:"registerTime,omitempty"`          // 注册时间，UNIX 时间戳(毫秒值)
	FriendNum             *int64  `json:"friendNum,omitempty"`             // 好友数
	FansNum               *int64  `json:"fansNum,omitempty"`               // 粉丝数
	IsPremiumUse          *int    `json:"isPremiumUse,omitempty"`          // 是否付费
	Role                  *string `json:"role,omitempty"`                  // 用户类型
	Phone                 *string `json:"phone,omitempty"`                 // 手机号
	Ip                    *string `json:"ip,omitempty"`                    // 用户IP地址
	DeviceId              *string `json:"deviceId,omitempty"`              // 用户设备id
	DeviceType            *int    `json:"deviceType,omitempty"`            // 用户设备类型
	Token                 *string `json:"token,omitempty"`                 // 反作弊token
	Mac                   *string `json:"mac,omitempty"`                   // 设备mac地址
	Imei                  *string `json:"imei,omitempty"`                  // Android设备imei信息
	Idfa                  *string `json:"idfa,omitempty"`                  // ios设备idfa信息
	Idfv                  *string `json:"idfv,omitempty"`                  // ios设备idfv信息
	AppVersion            *string `json:"appVersion,omitempty"`            // app版本号
	ReceiveUid            *string `json:"receiveUid,omitempty"`            // 接收人id
	Relationship          *int    `json:"relationship,omitempty"`          // 好友关系，1接收人关注发送人，2发送人关注接收人，3互相关注，4互未关注
	GroupId               *string `json:"groupId,omitempty"`               // 群id
	RoomId                *string `json:"roomId,omitempty"`                // 聊天室编号
	Topic                 *string `json:"topic,omitempty"`                 // 动态id
	CommentId             *string `json:"commentId,omitempty"`             // 主评论id
	CommodityId           *string `json:"commodityId,omitempty"`           // 商品id
	RiskControlToken      *string `json:"riskControlToken,omitempty"`      // 智能风控token
	RiskControlBusinessId *string `json:"riskControlBusinessId,omitempty"` // 智能风控businessId
}

// 构建request
func NewImageCheckSceneRequest(businessId string) *ImageCheckSceneRequest {
	var checkSceneRequest = &ImageCheckSceneRequest{
		BizPostFormRequest: types.NewBizPostFormRequest(businessId),
	}

	return checkSceneRequest
}

func (r *ImageCheckSceneRequest) SetNickname(nickname string) {
	r.Nickname = &nickname
}

func (r *ImageCheckSceneRequest) SetAccount(account string) {
	r.Account = &account
}

func (r *ImageCheckSceneRequest) SetGender(gender int) {
	r.Gender = &gender
}

func (r *ImageCheckSceneRequest) SetAge(age int) {
	r.Age = &age
}

func (r *ImageCheckSceneRequest) SetLevel(level int) {
	r.Level = &level
}

func (r *ImageCheckSceneRequest) SetRegisterTime(registerTime int64) {
	r.RegisterTime = &registerTime
}

func (r *ImageCheckSceneRequest) SetFriendNum(friendNum int64) {
	r.FriendNum = &friendNum
}

func (r *ImageCheckSceneRequest) SetFansNum(fansNum int64) {
	r.FansNum = &fansNum
}

func (r *ImageCheckSceneRequest) SetIsPremiumUse(isPremiumUse int) {
	r.IsPremiumUse = &isPremiumUse
}

func (r *ImageCheckSceneRequest) SetRole(role string) {
	r.Role = &role
}

func (r *ImageCheckSceneRequest) SetPhone(phone string) {
	r.Phone = &phone
}

func (r *ImageCheckSceneRequest) SetIp(ip string) {
	r.Ip = &ip
}

func (r *ImageCheckSceneRequest) SetDeviceId(deviceId string) {
	r.DeviceId = &deviceId
}

func (r *ImageCheckSceneRequest) SetDeviceType(deviceType int) {
	r.DeviceType = &deviceType
}

func (r *ImageCheckSceneRequest) SetToken(token string) {
	r.Token = &token
}

func (r *ImageCheckSceneRequest) SetMac(mac string) {
	r.Mac = &mac
}

func (r *ImageCheckSceneRequest) SetImei(imei string) {
	r.Imei = &imei
}

func (r *ImageCheckSceneRequest) SetIdfa(idfa string) {
	r.Idfa = &idfa
}

func (r *ImageCheckSceneRequest) SetIdfv(idfv string) {
	r.Idfv = &idfv
}

func (r *ImageCheckSceneRequest) SetAppVersion(appVersion string) {
	r.AppVersion = &appVersion
}

func (r *ImageCheckSceneRequest) SetReceiveUid(receiveUid string) {
	r.ReceiveUid = &receiveUid
}

func (r *ImageCheckSceneRequest) SetRelationship(relationship int) {
	r.Relationship = &relationship
}

func (r *ImageCheckSceneRequest) SetGroupId(groupId string) {
	r.GroupId = &groupId
}

func (r *ImageCheckSceneRequest) SetRoomId(roomId string) {
	r.RoomId = &roomId
}

func (r *ImageCheckSceneRequest) SetTopic(topic string) {
	r.Topic = &topic
}

func (r *ImageCheckSceneRequest) SetCommentId(commentId string) {
	r.CommentId = &commentId
}

func (r *ImageCheckSceneRequest) SetCommodityId(commodityId string) {
	r.CommodityId = &commodityId
}

func (r *ImageCheckSceneRequest) SetRiskControlToken(riskControlToken string) {
	r.RiskControlToken = &riskControlToken
}

func (r *ImageCheckSceneRequest) SetRiskControlBusinessId(riskControlBusinessId string) {
	r.RiskControlBusinessId = &riskControlBusinessId
}

func (r *ImageCheckSceneRequest) GetBusinessCustomSignParams() map[string]string {
	result := r.BizPostFormRequest.GetBusinessCustomSignParams()

	if r.Nickname != nil {
		result["nickname"] = *r.Nickname
	}
	if r.Account != nil {
		result["account"] = *r.Account
	}
	if r.Gender != nil {
		result["gender"] = strconv.Itoa(*r.Gender)
	}
	if r.Age != nil {
		result["age"] = strconv.Itoa(*r.Age)
	}
	if r.Level != nil {
		result["level"] = strconv.Itoa(*r.Level)
	}
	if r.RegisterTime != nil {
		result["registerTime"] = strconv.FormatInt(*r.RegisterTime, 10)
	}
	if r.FriendNum != nil {
		result["friendNum"] = strconv.FormatInt(*r.FriendNum, 10)
	}
	if r.FansNum != nil {
		result["fansNum"] = strconv.FormatInt(*r.FansNum, 10)
	}
	if r.IsPremiumUse != nil {
		result["isPremiumUse"] = strconv.Itoa(*r.IsPremiumUse)
	}
	if r.Role != nil {
		result["role"] = *r.Role
	}
	if r.Phone != nil {
		result["phone"] = *r.Phone
	}
	if r.Ip != nil {
		result["ip"] = *r.Ip
	}
	if r.DeviceId != nil {
		result["deviceId"] = *r.DeviceId
	}
	if r.DeviceType != nil {
		result["deviceType"] = strconv.Itoa(*r.DeviceType)
	}
	if r.Token != nil {
		result["token"] = *r.Token
	}
	if r.Mac != nil {
		result["mac"] = *r.Mac
	}
	if r.Imei != nil {
		result["imei"] = *r.Imei
	}
	if r.Idfa != nil {
		result["idfa"] = *r.Idfa
	}
	if r.Idfv != nil {
		result["idfv"] = *r.Idfv
	}
	if r.AppVersion != nil {
		result["appVersion"] = *r.AppVersion
	}
	if r.ReceiveUid != nil {
		result["receiveUid"] = *r.ReceiveUid
	}
	if r.Relationship != nil {
		result["relationship"] = strconv.Itoa(*r.Relationship)
	}
	if r.GroupId != nil {
		result["groupId"] = *r.GroupId
	}
	if r.RoomId != nil {
		result["roomId"] = *r.RoomId
	}
	if r.Topic != nil {
		result["topic"] = *r.Topic
	}
	if r.CommentId != nil {
		result["commentId"] = *r.CommentId
	}
	if r.CommodityId != nil {
		result["commodityId"] = *r.CommodityId
	}
	if r.RiskControlToken != nil {
		result["riskControlToken"] = *r.RiskControlToken
	}
	if r.RiskControlBusinessId != nil {
		result["riskControlBusinessId"] = *r.RiskControlBusinessId
	}

	return result
}
