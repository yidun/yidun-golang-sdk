package push

import (
	"strconv"

	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
)

/**
 * AIGC流式检测解决方案检测提交请求体
 */
type AigcStreamPushRequestV1 struct {
	*types.BizPostFormRequest
	SessionId string `json:"sessionId"` // 会话开启接口返回的Id，用于关联整个流式检测会话，用户也可以直接传一个自定义会话id, 系统接收到之后会自动创建一个新会话
	Type      int    `json:"type"`      // 推送的事件类型，1：流式输出检测  2：输入检测  3：会话结束，默认值为：1
	DataId    string `json:"dataId"`    // 上传数据唯一标识，能够根据该值定位到该条数据，如对数据结果有异议，可以发送该值给策略经理查询，建议客户传递
	Content   string `json:"content"`   // type = 1：content必传，流式检测内容片段，可对应aigc场景流式输出的tokens，检测片段传入AIGC-输出文本中，最大长度200
	// type = 2：content必传，流式检测场景下输入内容，建议对内容中json、表情符、HTML标签、UBB标签等做过滤，只传递纯文本，以减少误判概率，对应传入AIGC-输入文本中，最大长度10000
	// type = 3：content不需要传
	PublishTime int64  `json:"publishTime"` // 时间戳，数据片段生成的时间，可以不传，传了之后拼接的内容会基于时间戳进行排序，UNIX 时间戳（毫秒值）
	Callback    string `json:"callback"`    // 数据回调参数，调用方根据业务情况自行设计，当调用回调接口时，该接口会原样返回该字段
	CallbackUrl string `json:"callbackUrl"` // 异步及离线策略检测结果回调通知到客户的URL。主动回调数据接口超时时间设置为2s，为了保证顺利接收数据，需保证接收接口性能稳定并且保证幂等性
	CheckLabels string `json:"checkLabels"` // 业务过检分类,如果没有勾选分类提交返回参数错误，多个垃圾类别以逗号分隔（"100,200"）
	Account     string `json:"account"`     // 用户唯一标识
	Ip          string `json:"ip"`          // 用户IP地址，建议抄送，辅助机审策略精准调优
	DeviceId    string `json:"deviceId"`    // 用户设备 id
	ExtStr1     string `json:"extStr1"`     // 自定义扩展参数1
	ExtStr2     string `json:"extStr2"`     // 自定义扩展参数2
	ExtLon1     int64  `json:"extLon1"`     // 自定义扩展参数
	ExtLon2     int64  `json:"extLon2"`     // 自定义扩展参数
	Extension   string `json:"extension"`   // 自定义扩展参数，JSON字符串格式。如："{"keyName1":"value1","keyName2":"value2"}"
}

func NewAigcStreamPushRequestV1() *AigcStreamPushRequestV1 {
	var request = &AigcStreamPushRequestV1{
		BizPostFormRequest: types.NewBizPostFormRequestWithoutBizId(),
	}
	request.SetProductCode("aigc-stream-check-api")
	request.SetUriPattern("/v1/stream/push")
	request.SetVersion("v1.0")
	return request
}

/**
 * 获取参与验签自定义解析字段
 */
func (m *AigcStreamPushRequestV1) GetBusinessCustomSignParams() map[string]string {
	params := m.PostFormRequest.GetBusinessCustomSignParams()
	setStringIfNotNil(&params, "sessionId", &m.SessionId)
	setIntToStringIfNotNil(&params, "type", &m.Type)
	setStringIfNotNil(&params, "dataId", &m.DataId)
	setStringIfNotNil(&params, "content", &m.Content)
	setInt64ToStringIfNotNil(&params, "publishTime", &m.PublishTime)
	setStringIfNotNil(&params, "callback", &m.Callback)
	setStringIfNotNil(&params, "callbackUrl", &m.CallbackUrl)
	setStringIfNotNil(&params, "checkLabels", &m.CheckLabels)
	setStringIfNotNil(&params, "account", &m.Account)
	setStringIfNotNil(&params, "ip", &m.Ip)
	setStringIfNotNil(&params, "deviceId", &m.DeviceId)
	setStringIfNotNil(&params, "extStr1", &m.ExtStr1)
	setStringIfNotNil(&params, "extStr2", &m.ExtStr2)
	setInt64ToStringIfNotNil(&params, "extLon1", &m.ExtLon1)
	setInt64ToStringIfNotNil(&params, "extLon2", &m.ExtLon2)
	setStringIfNotNil(&params, "extension", &m.Extension)
	return params
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
