package submit

import (
	"encoding/json"
	"fmt"
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
)

/**
 * 融媒体解决方案检测提交请求体
 */
type MediaSubmitRequestV2 struct {
	*types.BizPostFormRequest
	Ip                    *string                `json:"ip,omitempty"`                    // 用户IP地址
	Account               *string                `json:"account,omitempty"`               // 用户账号
	DeviceID              *string                `json:"deviceId,omitempty"`              // 用户设备id
	DeviceType            *string                `json:"deviceType,omitempty"`            // 用户设备类型
	DataID                *string                `json:"dataId,omitempty"`                // 数据唯一标识，能够根据该值定位到该条数据，如对数据结果有异议，可以发送该值给客户经理查询
	Title                 *string                `json:"title,omitempty"`                 // 标题
	Content               []*DataItem            `json:"content,omitempty"`               // 支持多种类型内容同时过检，包括文本，图片，点播语音,点播音视频，文档
	CustomParseFieldMap   map[string][]*DataItem `json:"customParseFieldMap,omitempty"`   // 用户自定义解析字段
	CustomUnParseFieldMap map[string]string      `json:"customUnParseFieldMap,omitempty"` // 用户自定义不解析解析字段
	Callback              *string                `json:"callback,omitempty"`              // 回调数据
	CallbackUrl           *string                `json:"callbackUrl,omitempty"`           // 回调地址
	PublishTime           *int64                 `json:"publishTime,omitempty"`           // publishTime
	CheckLanguageCode     *string                `json:"checkLanguageCode,omitempty"`     // 语种
	Token                 *string                `json:"token,omitempty"`                 // 反作弊的 token，当开通反作弊时，抄送到反作弊服务
	CensorExt             *string                `json:"censorExt,omitempty"`             // 人审标签扩展字段
}

type DataItem struct {
	Type   *string           `json:"type,omitempty"`   // type
	Data   *string           `json:"data,omitempty"`   // data
	DataID *string           `json:"dataId,omitempty"` // dataId
	Config map[string]string `json:"config,omitempty"` // 子数据数据配置，主要是针对于过检数据的特定参数设置
}

func NewMediaSubmitRequestV2() *MediaSubmitRequestV2 {
	var request = &MediaSubmitRequestV2{
		BizPostFormRequest: types.NewBizPostFormRequestWithoutBizId(),
	}
	request.SetProductCode("mediaSubmit")
	request.SetUriPattern("/v2/mediasolution/submit")
	request.SetVersion("v2.1")
	return request
}

func NewDataItem() *DataItem {
	var dataItem = &DataItem{}
	return dataItem
}

func (m *DataItem) SetType(stype string) {
	m.Type = &stype
}

func (m *DataItem) SetData(data string) {
	m.Data = &data
}

func (m *DataItem) SetDataID(dataId string) {
	m.DataID = &dataId
}

func (m *DataItem) SetConfig(config map[string]string) {
	m.Config = config
}

func (m *MediaSubmitRequestV2) SetIp(ip string) {
	m.Ip = &ip
}

func (m *MediaSubmitRequestV2) SetAccount(account string) {
	m.Account = &account
}

func (m *MediaSubmitRequestV2) SetDeviceID(deviceId string) {
	m.DeviceID = &deviceId
}

func (m *MediaSubmitRequestV2) SetDeviceType(deviceType string) {
	m.DeviceType = &deviceType
}

func (m *MediaSubmitRequestV2) SetDataID(dataID string) {
	m.DataID = &dataID
}

func (m *MediaSubmitRequestV2) SetTitle(title string) {
	m.Title = &title
}

func (m *MediaSubmitRequestV2) SetContent(content []*DataItem) {
	m.Content = content
}

func (m *MediaSubmitRequestV2) SetCustomParseFieldMap(customParseFieldMap map[string][]*DataItem) {
	m.CustomParseFieldMap = customParseFieldMap
}

func (m *MediaSubmitRequestV2) SetCustomUnParseFieldMap(customUnParseFieldMap map[string]string) {
	m.CustomUnParseFieldMap = customUnParseFieldMap
}

func (m *MediaSubmitRequestV2) SetCallback(callback string) {
	m.Callback = &callback
}

func (m *MediaSubmitRequestV2) SetCallbackUrl(callbackUrl string) {
	m.CallbackUrl = &callbackUrl
}

func (m *MediaSubmitRequestV2) SetPublishTime(publishTime int64) {
	m.PublishTime = &publishTime
}

func (m *MediaSubmitRequestV2) SetCheckLanguageCode(checkLanguageCode string) {
	m.CheckLanguageCode = &checkLanguageCode
}

func (m *MediaSubmitRequestV2) SetToken(token string) {
	m.Token = &token
}

func (m *MediaSubmitRequestV2) SetCensorExt(censorExt string) {
	m.CensorExt = &censorExt
}

/**
 * 获取参与验签自定义解析字段
 */
func (m *MediaSubmitRequestV2) GetBusinessCustomSignParams() map[string]string {
	params := m.PostFormRequest.GetBusinessCustomSignParams()
	if m.Ip != nil {
		params["ip"] = *m.Ip
	}
	if m.Account != nil {
		params["account"] = *m.Account
	}
	if m.DeviceID != nil {
		params["deviceId"] = *m.DeviceID
	}
	if m.DeviceType != nil {
		params["deviceType"] = *m.DeviceType
	}
	if m.DataID != nil {
		params["dataId"] = *m.DataID
	}
	if m.Title != nil {
		params["title"] = *m.Title
	}
	if m.Callback != nil {
		params["callback"] = *m.Callback
	}
	if m.CallbackUrl != nil {
		params["callbackUrl"] = *m.CallbackUrl
	}
	if m.PublishTime != nil {
		params["publishTime"] = fmt.Sprintf("%d", *m.PublishTime)
	}
	if m.CheckLanguageCode != nil {
		params["checkLanguageCode"] = *m.CheckLanguageCode
	}
	if m.Token != nil {
		params["token"] = *m.Token
	}
	if m.CensorExt != nil {
		params["censorExt"] = *m.CensorExt
	}
	if m.Content != nil {
		contentJson, _ := json.Marshal(m.Content)
		params["content"] = string(contentJson)
	}
	if m.CustomParseFieldMap != nil {
		for key, value := range m.CustomParseFieldMap {
			if key != "" && value != nil {
				valueJson, _ := json.Marshal(value)
				params[key] = string(valueJson)
			}
		}
	}
	if m.CustomUnParseFieldMap != nil {
		for key, value := range m.CustomUnParseFieldMap {
			if key != "" && value != "" {
				params[key] = value
			}
		}
	}
	return params
}
