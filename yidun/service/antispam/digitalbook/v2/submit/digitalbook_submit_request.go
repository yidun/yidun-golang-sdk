package submit

import (
	"encoding/json"
	"fmt"
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"strconv"
)

type DigitalBookSubmitRequestV2 struct {
	*types.BizPostFormRequest
	Ip                    *string                `json:"ip,omitempty"`                    // 用户IP地址
	Account               *string                `json:"account,omitempty"`               // 用户账号
	DeviceID              *string                `json:"deviceId,omitempty"`              // 用户设备id
	DeviceType            *string                `json:"deviceType,omitempty"`            // 用户设备类型
	DataId                *string                `json:"dataId,omitempty"`                // 数据唯一标识，能够根据该值定位到该条数据，如对数据结果有异议，可以发送该值给客户经理查询
	Title                 *string                `json:"title,omitempty"`                 // 标题
	Content               []*DataItem            `json:"content,omitempty"`               // 支持多种类型内容同时过检，包括文本，图片，点播语音,点播音视频，文档
	CustomParseFieldMap   map[string][]*DataItem `json:"customParseFieldMap,omitempty"`   // 用户自定义解析字段
	CustomUnParseFieldMap map[string]string      `json:"customUnParseFieldMap,omitempty"` // 用户自定义不解析解析字段
	Callback              *string                `json:"callback,omitempty"`              // 回调数据
	CallbackUrl           *string                `json:"callbackUrl,omitempty"`           // 回调地址
	PublishTime           *int64                 `json:"publishTime,omitempty"`           // publishTime
	Token                 *string                `json:"token,omitempty"`                 // 反作弊的 token，当开通反作弊时，抄送到反作弊服务
	BookID                *string                `json:"bookId,omitempty"`                // 书籍id
	ChapterNumber         *int                   `json:"chapterNumber,omitempty"`         // 章节号
	TotalChapters         *int                   `json:"totalChapters,omitempty"`         // 章节总数
	BookName              []*DataItem            `json:"bookName,omitempty"`              // 作品名称
	BookCover             []*DataItem            `json:"bookCover,omitempty"`             // 作品封面图片
	BookInformation       []*DataItem            `json:"bookInformation,omitempty"`       // 作品信息
	AuthorName            *string                `json:"authorName,omitempty"`            // 作者名
	TotalBooks            *int                   `json:"totalBooks,omitempty"`            // 作者作品总数
	TotalFans             *int                   `json:"totalFans,omitempty"`             // 作者粉丝数
	TotalCreationDays     *int                   `json:"totalCreationDays,omitempty"`     // 作者创作天数
	AuthorRank            *string                `json:"authorRank,omitempty"`            // 作者等级
	Type                  *string                `json:"type,omitempty"`                  // 场景分类 1 提交只包含书的一章，2 表示提交包含整本书内容

}

type DataItem struct {
	// type最长10个字符
	Type *string `json:"type,omitempty"`
	// data最长5000个字符
	Data *string `json:"data,omitempty"`
	// dataId最长128个字符
	DataId *string `json:"dataId,omitempty"`
}

func NewDigitalBookSubmitRequestV2() *DigitalBookSubmitRequestV2 {
	var request = &DigitalBookSubmitRequestV2{
		BizPostFormRequest: types.NewBizPostFormRequestWithoutBizId(),
	}
	request.SetProductCode("digitalBook")
	request.SetUriPattern("/v2/digital/submit")
	request.SetVersion("v2")
	return request
}

func NewDataItem() *DataItem {
	var dataItem = &DataItem{}
	return dataItem
}

// Setters
func (m *DigitalBookSubmitRequestV2) SetIp(ip string) {
	m.Ip = &ip
}

func (d *DigitalBookSubmitRequestV2) SetAccount(account string) {
	d.Account = &account
}

func (d *DigitalBookSubmitRequestV2) SetDeviceID(deviceID string) {
	d.DeviceID = &deviceID
}

func (d *DigitalBookSubmitRequestV2) SetDeviceType(deviceType string) {
	d.DeviceType = &deviceType
}

func (d *DigitalBookSubmitRequestV2) SetDataId(dataID string) {
	d.DataId = &dataID
}

func (d *DigitalBookSubmitRequestV2) SetTitle(title string) {
	d.Title = &title
}

func (d *DigitalBookSubmitRequestV2) SetContent(content []*DataItem) {
	d.Content = content
}

func (m *DigitalBookSubmitRequestV2) SetCustomParseFieldMap(customParseFieldMap map[string][]*DataItem) {
	m.CustomParseFieldMap = customParseFieldMap
}

func (d *DigitalBookSubmitRequestV2) SetCustomUnParseFieldMap(customUnParseFieldMap map[string]string) {
	d.CustomUnParseFieldMap = customUnParseFieldMap
}

func (d *DigitalBookSubmitRequestV2) SetCallback(callback string) {
	d.Callback = &callback
}

func (d *DigitalBookSubmitRequestV2) SetCallbackUrl(callbackUrl string) {
	d.CallbackUrl = &callbackUrl
}

func (d *DigitalBookSubmitRequestV2) SetPublishTime(publishTime int64) {
	d.PublishTime = &publishTime
}

func (d *DigitalBookSubmitRequestV2) SetToken(token string) {
	d.Token = &token
}

func (d *DigitalBookSubmitRequestV2) SetBookID(bookID string) {
	d.BookID = &bookID
}

func (d *DigitalBookSubmitRequestV2) SetChapterNumber(chapterNumber int) {
	d.ChapterNumber = &chapterNumber
}

func (d *DigitalBookSubmitRequestV2) SetTotalChapters(totalChapters int) {
	d.TotalChapters = &totalChapters
}

func (d *DigitalBookSubmitRequestV2) SetBookName(bookName []*DataItem) {
	d.BookName = bookName
}

func (d *DigitalBookSubmitRequestV2) SetBookCover(bookCover []*DataItem) {
	d.BookCover = bookCover
}

func (d *DigitalBookSubmitRequestV2) SetBookInformation(bookInformation []*DataItem) {
	d.BookInformation = bookInformation
}

func (d *DigitalBookSubmitRequestV2) SetAuthorName(authorName string) {
	d.AuthorName = &authorName
}

func (d *DigitalBookSubmitRequestV2) SetTotalBooks(totalBooks int) {
	d.TotalBooks = &totalBooks
}

func (d *DigitalBookSubmitRequestV2) SetTotalFans(totalFans int) {
	d.TotalFans = &totalFans
}

func (d *DigitalBookSubmitRequestV2) SetTotalCreationDays(totalCreationDays int) {
	d.TotalCreationDays = &totalCreationDays
}

func (d *DigitalBookSubmitRequestV2) SetAuthorRank(authorRank string) {
	d.AuthorRank = &authorRank
}

func (d *DigitalBookSubmitRequestV2) SetType(typ string) {
	d.Type = &typ
}

func (d *DataItem) SetType(typ string) {
	d.Type = &typ
}

func (d *DataItem) SetData(data string) {
	d.Data = &data
}

func (d *DataItem) SetDataId(dataID string) {
	d.DataId = &dataID
}

/**
 * 获取参与验签自定义解析字段
 */
func (d *DigitalBookSubmitRequestV2) GetBusinessCustomSignParams() map[string]string {
	params := d.PostFormRequest.GetBusinessCustomSignParams()
	if d.Ip != nil {
		params["ip"] = *d.Ip
	}
	if d.Account != nil {
		params["account"] = *d.Account
	}
	if d.DeviceID != nil {
		params["deviceId"] = *d.DeviceID
	}
	if d.DeviceType != nil {
		params["deviceType"] = *d.DeviceType
	}
	if d.DataId != nil {
		params["dataId"] = *d.DataId
	}
	if d.Title != nil {
		params["title"] = *d.Title
	}
	if d.Callback != nil {
		params["callback"] = *d.Callback
	}
	if d.CallbackUrl != nil {
		params["callbackUrl"] = *d.CallbackUrl
	}
	if d.PublishTime != nil {
		params["publishTime"] = fmt.Sprintf("%d", *d.PublishTime)
	}
	if d.Token != nil {
		params["token"] = *d.Token
	}
	if d.BookID != nil {
		params["bookId"] = *d.BookID
	}
	if d.AuthorName != nil {
		params["authorName"] = *d.AuthorName
	}
	if d.AuthorRank != nil {
		params["authorRank"] = *d.AuthorRank
	}
	if d.Type != nil {
		params["type"] = *d.Type
	}
	if d.ChapterNumber != nil {
		params["chapterNumber"] = strconv.Itoa(*d.ChapterNumber)
	}
	if d.TotalChapters != nil {
		params["totalChapters"] = strconv.Itoa(*d.TotalChapters)
	}
	if d.TotalBooks != nil {
		params["totalBooks"] = strconv.Itoa(*d.TotalBooks)
	}
	if d.TotalFans != nil {
		params["totalFans"] = strconv.Itoa(*d.TotalFans)
	}
	if d.TotalCreationDays != nil {
		params["totalCreationDays"] = strconv.Itoa(*d.TotalCreationDays)
	}
	if d.Content != nil && len(d.Content) > 0 {
		contentBytes, _ := json.Marshal(d.Content)
		params["content"] = string(contentBytes)
	}
	if d.BookCover != nil && len(d.BookCover) > 0 {
		bookCoverBytes, _ := json.Marshal(d.BookCover)
		params["bookCover"] = string(bookCoverBytes)
	}
	if d.BookName != nil && len(d.BookName) > 0 {
		bookNameBytes, _ := json.Marshal(d.BookName)
		params["bookName"] = string(bookNameBytes)
	}
	if d.BookInformation != nil && len(d.BookInformation) > 0 {
		bookInformationBytes, _ := json.Marshal(d.BookInformation)
		params["bookInformation"] = string(bookInformationBytes)
	}
	if d.CustomParseFieldMap != nil {
		for key, value := range d.CustomParseFieldMap {
			if key != "" && value != nil {
				valueJson, _ := json.Marshal(value)
				params[key] = string(valueJson)
			}
		}
	}
	if d.CustomUnParseFieldMap != nil {
		for key, value := range d.CustomUnParseFieldMap {
			if key != "" && value != "" {
				params[key] = value
			}
		}
	}
	return params
}
