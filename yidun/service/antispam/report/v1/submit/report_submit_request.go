package submit

import (
	"encoding/json"
	"fmt"
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
)

/**
 * 投诉举报解决方案检测提交请求体
 */
type ReportSubmitRequestV1 struct {
	*types.BizPostFormRequest
	Ip                    *string                `json:"ip,omitempty"`                    // 用户IP地址
	Account               *string                `json:"account,omitempty"`               // 用户账号
	DeviceID              *string                `json:"deviceId,omitempty"`              // 用户设备id
	DeviceType            *string                `json:"deviceType,omitempty"`            // 用户设备类型
	DataID                *string                `json:"dataId,omitempty"`                // 数据唯一标识，能够根据该值定位到该条数据，如对数据结果有异议，可以发送该值给客户经理查询
	Content               []*DataItem            `json:"content,omitempty"`               // 投诉举报内容,支持多种类型内容同时过检，包括文本，图片，点播语音,点播音视频，文档
	Evidence              []*DataItem            `json:"evidence,omitempty"`              // 证据信息,支持多种类型内容同时过检，包括文本，图片，点播语音,点播音视频，文档
	CustomParseFieldMap   map[string][]*DataItem `json:"customParseFieldMap,omitempty"`   // 用户自定义解析字段
	CustomUnParseFieldMap map[string]string      `json:"customUnParseFieldMap,omitempty"` // 用户自定义不解析解析字段
	Callback              *string                `json:"callback,omitempty"`              // 回调数据
	CallbackUrl           *string                `json:"callbackUrl,omitempty"`           // 回调地址
	PublishTime           *int64                 `json:"publishTime,omitempty"`           // publishTime
	Token                 *string                `json:"token,omitempty"`                 // 反作弊的 token，当开通反作弊时，抄送到反作弊服务
	ReportedId            *string                `json:"reportedId,omitempty"`            // 被投诉举报人
	Scenarios             *string                `json:"scenarios,omitempty"`             // 投诉举报场景
	ReportType            *string                `json:"reportType,omitempty"`            // 投诉举报类型
	RoomId                *string                `json:"roomId,omitempty"`                // 房间id
	CheckLanguageCode     *string                `json:"checkLanguageCode,omitempty"`     // 语种
}

type DataItem struct {
	Type   *string `json:"type,omitempty"`
	Data   *string `json:"data,omitempty"`
	DataID *string `json:"dataId,omitempty"`
}

func NewReportSubmitRequestV1() *ReportSubmitRequestV1 {
	var request = &ReportSubmitRequestV1{
		BizPostFormRequest: types.NewBizPostFormRequestWithoutBizId(),
	}
	request.SetProductCode("report")
	request.SetUriPattern("/v1/report/submit")
	request.SetVersion("v1")
	return request
}

func NewDataItem() *DataItem {
	var dataItem = &DataItem{}
	return dataItem
}

func (r *ReportSubmitRequestV1) GetBusinessCustomSignParams() map[string]string {
	params := r.PostFormRequest.GetBusinessCustomSignParams()
	if r.Ip != nil {
		params["ip"] = *r.Ip
	}

	if r.Account != nil {
		params["account"] = *r.Account
	}

	if r.DeviceID != nil {
		params["deviceId"] = *r.DeviceID
	}

	if r.DeviceType != nil {
		params["deviceType"] = *r.DeviceType
	}

	if r.DataID != nil {
		params["dataId"] = *r.DataID
	}

	if r.Callback != nil {
		params["callback"] = *r.Callback
	}

	if r.CallbackUrl != nil {
		params["callbackUrl"] = *r.CallbackUrl
	}

	if r.PublishTime != nil {
		params["publishTime"] = fmt.Sprintf("%d", *r.PublishTime)
	}

	if r.Token != nil {
		params["token"] = *r.Token
	}

	if r.ReportedId != nil {
		params["reportedId"] = *r.ReportedId
	}

	if r.Scenarios != nil {
		params["scenarios"] = *r.Scenarios
	}

	if r.ReportType != nil {
		params["reportType"] = *r.ReportType
	}

	if r.RoomId != nil {
		params["roomId"] = *r.RoomId
	}

	if r.CheckLanguageCode != nil {
		params["checkLanguageCode"] = *r.CheckLanguageCode
	}

	if r.Content != nil && len(r.Content) > 0 {
		contentJson, _ := json.Marshal(r.Content)
		params["content"] = string(contentJson)
	}

	if r.Evidence != nil && len(r.Evidence) > 0 {
		evidenceJson, _ := json.Marshal(r.Evidence)
		params["evidence"] = string(evidenceJson)
	}

	if r.CustomParseFieldMap != nil && len(r.CustomParseFieldMap) > 0 {
		for key, dataItemList := range r.CustomParseFieldMap {
			if key != "" && dataItemList != nil && len(dataItemList) > 0 {
				dataItemListJson, _ := json.Marshal(dataItemList)
				params[key] = string(dataItemListJson)
			}
		}
	}

	if r.CustomUnParseFieldMap != nil && len(r.CustomUnParseFieldMap) > 0 {
		for key, data := range r.CustomUnParseFieldMap {
			if key != "" && data != "" {
				params[key] = data
			}
		}
	}

	return params
}

// SetIP sets the Ip field of ReportSubmitRequestV1.
func (r *ReportSubmitRequestV1) SetIP(ip string) {
	r.Ip = &ip
}

// SetAccount sets the Account field of ReportSubmitRequestV1.
func (r *ReportSubmitRequestV1) SetAccount(account string) {
	r.Account = &account
}

// SetDeviceID sets the DeviceID field of ReportSubmitRequestV1.
func (r *ReportSubmitRequestV1) SetDeviceID(deviceID string) {
	r.DeviceID = &deviceID
}

// SetDeviceType sets the DeviceType field of ReportSubmitRequestV1.
func (r *ReportSubmitRequestV1) SetDeviceType(deviceType string) {
	r.DeviceType = &deviceType
}

// SetDataID sets the DataID field of ReportSubmitRequestV1.
func (r *ReportSubmitRequestV1) SetDataID(dataID string) {
	r.DataID = &dataID
}

// SetContent sets the Content field of ReportSubmitRequestV1.
func (r *ReportSubmitRequestV1) SetContent(content []*DataItem) {
	r.Content = content
}

// SetEvidence sets the Evidence field of ReportSubmitRequestV1.
func (r *ReportSubmitRequestV1) SetEvidence(evidence []*DataItem) {
	r.Evidence = evidence
}

// SetCustomParseFieldMap sets the CustomParseFieldMap field of ReportSubmitRequestV1.
func (r *ReportSubmitRequestV1) SetCustomParseFieldMap(customParseFieldMap map[string][]*DataItem) {
	r.CustomParseFieldMap = customParseFieldMap
}

// SetCustomUnParseFieldMap sets the CustomUnParseFieldMap field of ReportSubmitRequestV1.
func (r *ReportSubmitRequestV1) SetCustomUnParseFieldMap(customUnParseFieldMap map[string]string) {
	r.CustomUnParseFieldMap = customUnParseFieldMap
}

// SetCallback sets the Callback field of ReportSubmitRequestV1.
func (r *ReportSubmitRequestV1) SetCallback(callback string) {
	r.Callback = &callback
}

// SetCallbackUrl sets the CallbackUrl field of ReportSubmitRequestV1.
func (r *ReportSubmitRequestV1) SetCallbackUrl(callbackUrl string) {
	r.CallbackUrl = &callbackUrl
}

// SetPublishTime sets the PublishTime field of ReportSubmitRequestV1.
func (r *ReportSubmitRequestV1) SetPublishTime(publishTime int64) {
	r.PublishTime = &publishTime
}

// SetToken sets the Token field of ReportSubmitRequestV1.
func (r *ReportSubmitRequestV1) SetToken(token string) {
	r.Token = &token
}

// SetReportedId sets the ReportedId field of ReportSubmitRequestV1.
func (r *ReportSubmitRequestV1) SetReportedId(reportedId string) {
	r.ReportedId = &reportedId
}

// SetScenarios sets the Scenarios field of ReportSubmitRequestV1.
func (r *ReportSubmitRequestV1) SetScenarios(scenarios string) {
	r.Scenarios = &scenarios
}

// SetReportType sets the ReportType field of ReportSubmitRequestV1.
func (r *ReportSubmitRequestV1) SetReportType(reportType string) {
	r.ReportType = &reportType
}

// SetRoomId sets the RoomId field of ReportSubmitRequestV1.
func (r *ReportSubmitRequestV1) SetRoomId(roomId string) {
	r.RoomId = &roomId
}

// SetCheckLanguageCode sets the CheckLanguageCode field of ReportSubmitRequestV1.
func (r *ReportSubmitRequestV1) SetCheckLanguageCode(checkLanguageCode string) {
	r.CheckLanguageCode = &checkLanguageCode
}

// SetType sets the Type field of DataItem.
func (d *DataItem) SetType(typ string) {
	d.Type = &typ
}

// SetData sets the Data field of DataItem.
func (d *DataItem) SetData(data string) {
	d.Data = &data
}

// SetDataID sets the DataID field of DataItem.
func (d *DataItem) SetDataID(dataID string) {
	d.DataID = &dataID
}
