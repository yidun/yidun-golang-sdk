package single

import (
	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/core/validation"
)

/**
 * 文本检测请求基类
 */
type TextCheckSceneRequest struct {
	*types.BizPostFormRequest
	*TextSceneRequest
}

// 构建request
func NewTextCheckSceneRequest(businessId string) *TextCheckSceneRequest {
	var checkSceneRequest = &TextCheckSceneRequest{
		BizPostFormRequest: types.NewBizPostFormRequest(businessId),
		TextSceneRequest:   NewTextSceneRequest(),
	}

	return checkSceneRequest
}

// 参数校验
func (r *TextCheckSceneRequest) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "TextCheckSceneRequest"}

	validateField := func(fieldName string, fieldValue *string, maxLen int) {
		if fieldValue == nil {
			return
		}
		if utf8.RuneCountInString(*fieldValue) > maxLen {
			invalidParams.Add(validation.NewErrParamMaxLen(fieldName, maxLen, strconv.Itoa(utf8.RuneCountInString(*fieldValue))))
		}
	}

	if r == nil {
		invalidParams.Add(validation.NewErrParamRequired("TextCheckSceneRequest"))
	} else {
		if r.DataId == nil {
			invalidParams.Add(validation.NewErrParamRequired("DataID"))
		} else {
			validateField("DataId", r.DataId, 128)
		}

		if r.Content == nil {
			invalidParams.Add(validation.NewErrParamRequired("Content"))
		} else {
			validateField("Content", r.Content, 10000)
		}

		validateField("Title", r.Title, 512)
		validateField("Callback", r.Callback, 65535)
		validateField("CallbackUrl", r.CallbackUrl, 256)
		validateField("Category", r.Category, 128)
		validateField("Account", r.Account, 128)
		validateField("Phone", r.Phone, 64)
		validateField("Nickname", r.Nickname, 128)
		validateField("Role", r.Role, 32)
		validateField("DeviceId", r.DeviceId, 128)
		validateField("Mac", r.Mac, 64)
		validateField("Imei", r.Imei, 64)
		validateField("Idfa", r.Idfa, 64)
		validateField("Idfv", r.Idfv, 64)
		validateField("AppVersion", r.AppVersion, 32)
		validateField("ReceiveUid", r.ReceiveUid, 64)
		validateField("GroupId", r.GroupId, 32)
		validateField("RoomId", r.RoomId, 32)
		validateField("Topic", r.Topic, 128)
		validateField("CommentId", r.CommentId, 32)
		validateField("CommodityId", r.CommodityId, 32)
		validateField("Ip", r.Ip, 128)
		validateField("ReceiveIp", r.ReceiveIp, 128)
		validateField("ExtStr1", r.ExtStr1, 128)
		validateField("ExtStr2", r.ExtStr2, 128)
		validateField("CensorExt", r.CensorExt, 128)
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// 获取业务自定义参数
func (req *TextCheckSceneRequest) GetBusinessCustomSignParams() map[string]string {
	result := req.BizPostFormRequest.GetBusinessCustomSignParams()
	setStringIfNotNil(&result, "account", req.Account)
	setIntToStringIfNotNil(&result, "age", req.Age)
	setStringIfNotNil(&result, "appVersion", req.AppVersion)
	setStringIfNotNil(&result, "callback", req.Callback)
	setStringIfNotNil(&result, "callbackUrl", req.CallbackUrl)
	setStringIfNotNil(&result, "category", req.Category)
	setStringIfNotNil(&result, "censorExt", req.CensorExt)
	setStringIfNotNil(&result, "commentId", req.CommentId)
	setStringIfNotNil(&result, "commodityId", req.CommodityId)
	setStringIfNotNil(&result, "content", req.Content)
	setStringIfNotNil(&result, "dataId", req.DataId)
	setIntToStringIfNotNil(&result, "dataType", req.DataType)
	setStringIfNotNil(&result, "deviceId", req.DeviceId)
	setIntToStringIfNotNil(&result, "deviceType", req.DeviceType)
	setInt64ToStringIfNotNil(&result, "extLon1", req.ExtLon1)
	setInt64ToStringIfNotNil(&result, "extLon2", req.ExtLon2)
	setStringIfNotNil(&result, "extStr1", req.ExtStr1)
	setStringIfNotNil(&result, "extStr2", req.ExtStr2)
	setInt64ToStringIfNotNil(&result, "fansNum", req.FansNum)
	setInt64ToStringIfNotNil(&result, "friendNum", req.FriendNum)
	setIntToStringIfNotNil(&result, "gender", req.Gender)
	setStringIfNotNil(&result, "groupId", req.GroupId)
	setStringIfNotNil(&result, "idfa", req.Idfa)
	setStringIfNotNil(&result, "idfv", req.Idfv)
	setStringIfNotNil(&result, "imei", req.Imei)
	setStringIfNotNil(&result, "ip", req.Ip)
	setStringIfNotNil(&result, "receiveIp", req.ReceiveIp)
	setIntToStringIfNotNil(&result, "isPremiumUse", req.IsPremiumUse)
	setIntToStringIfNotNil(&result, "level", req.Level)
	setStringIfNotNil(&result, "mac", req.Mac)
	setStringIfNotNil(&result, "nickname", req.Nickname)
	setStringIfNotNil(&result, "phone", req.Phone)
	setInt64ToStringIfNotNil(&result, "publishTime", req.PublishTime)
	setStringIfNotNil(&result, "receiveUid", req.ReceiveUid)
	setInt64ToStringIfNotNil(&result, "registerTime", req.RegisterTime)
	setStringSliceToStringIfNotNil(&result, "relatedKeys", req.RelatedKeys)
	setIntToStringIfNotNil(&result, "relationship", req.Relationship)
	setStringIfNotNil(&result, "riskControlBusinessId", req.RiskControlBusinessId)
	setStringIfNotNil(&result, "riskControlToken", req.RiskControlToken)
	setStringIfNotNil(&result, "role", req.Role)
	setStringIfNotNil(&result, "roomId", req.RoomId)
	setStringIfNotNil(&result, "title", req.Title)
	setStringIfNotNil(&result, "topic", req.Topic)
	setStringIfNotNil(&result, "extension", req.Extension)
	setStringIfNotNil(&result, "subProduct", req.SubProduct)
	return result
}

// 设置字段值
func setStringIfNotNil(result *map[string]string, key string, value *string) {
	if value != nil {
		(*result)[key] = *value
	}
}

// 设置字段值
func setIntToStringIfNotNil(result *map[string]string, key string, value *int) {
	if value != nil {
		(*result)[key] = strconv.Itoa(*value)
	}
}

// 设置字段值
func setInt64ToStringIfNotNil(result *map[string]string, key string, value *int64) {
	if value != nil {
		(*result)[key] = strconv.FormatInt(*value, 10)
	}
}

// 设置字段值
func setStringSliceToStringIfNotNil(result *map[string]string, key string, value *[]string) {
	if value != nil {
		(*result)[key] = strings.Join(*value, ",")
	}
}
