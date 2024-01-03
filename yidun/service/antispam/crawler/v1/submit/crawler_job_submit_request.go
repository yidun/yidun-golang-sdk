package submit

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
)

type CrawlerJobSubmitV1Request struct {
	*types.BizPostFormRequest
	// 数据唯一标识，能够根据该值定位到该条数据，如对数据结果有异议，可以发送该值给客户经理查询
	DataId *string `json:"dataId,omitempty"`
	// 循环爬虫时间区间--开始时间
	SliceStartTime *int64 `json:"sliceStartTime,omitempty"`
	// 循环爬虫时间区间--结束时间
	SliceEndTime *int64 `json:"sliceEndTime,omitempty"`
	// 主站URL
	SiteUrl *string `json:"siteUrl,omitempty"`
	// 爬虫深度/网站层级
	Level *int `json:"level,omitempty"`
	// 检测频率/多久爬取一次，单位毫秒
	Frequency *int64 `json:"frequency,omitempty"`
	// 单次任务周期内爬取页面的最大数量
	MaxResourceAmount *int `json:"maxResourceAmount,omitempty"`
	// 任务类型；0：循环任务；1：单次任务
	Type *int `json:"type,omitempty"`
	// 主动回调地址
	CallbackUrl *string `json:"callbackUrl,omitempty"`
	// 网站名称
	SiteName *string `json:"siteName,omitempty"`
	// 检测策略；当循环任务时，此配置生效。1：全量页面；2：首次全量，后续增量
	CheckStrategy *int `json:"checkStrategy,omitempty"`
	// 网站爬取配置
	Config *string `json:"config,omitempty"`
	// user agent 匹配规则
	UserAgentMatchType *int `json:"userAgentMatchType,omitempty"`
	// 浏览器user agent
	UserAgent *string `json:"userAgent,omitempty"`
	// 重点关注资源
	FocusList *[]string `json:"focusList,omitempty"`
	// URL过滤条件集
	UrlFilters *[]CrawlerUrlFilter `json:"urlFilters,omitempty"`
	// 检测内容 默认为1和2。1-检测文本，2-检测图片，4-检测点播音频，5-检测文档附件，6-检测点播音视频
	CheckFlags *[]int `json:"checkFlag,omitempty"`
}

// URL过滤条件
type CrawlerUrlFilter struct {
	// 过滤类型
	Type *int    `json:"type,omitempty"`
	Url  *string `json:"url,omitempty"`
}

// For ImageV5CheckRequest
func NewCrawlerJobSubmitV1Request() *CrawlerJobSubmitV1Request {
	var request = &CrawlerJobSubmitV1Request{
		BizPostFormRequest: types.NewBizPostFormRequestWithoutBizId(),
	}
	request.SetProductCode("crawler")
	request.SetUriPattern("/v1/crawler/job/submit")
	request.SetMethod(http.HttpMethodPost)
	request.SetVersion("v1.0")
	return request
}

func (c *CrawlerJobSubmitV1Request) SetDataId(dataId string) {
	c.DataId = &dataId
}

func (c *CrawlerJobSubmitV1Request) SetSliceStartTime(sliceStartTime int64) {
	c.SliceStartTime = &sliceStartTime
}

func (c *CrawlerJobSubmitV1Request) SetSliceEndTime(sliceEndTime int64) {
	c.SliceEndTime = &sliceEndTime
}

func (c *CrawlerJobSubmitV1Request) SetSiteUrl(siteUrl string) {
	c.SiteUrl = &siteUrl
}

func (c *CrawlerJobSubmitV1Request) SetLevel(level int) {
	c.Level = &level
}

func (c *CrawlerJobSubmitV1Request) SetFrequency(frequency int64) {
	c.Frequency = &frequency
}

func (c *CrawlerJobSubmitV1Request) SetMaxResourceAmount(maxResourceAmount int) {
	c.MaxResourceAmount = &maxResourceAmount
}

func (c *CrawlerJobSubmitV1Request) SetType(typeCode int) {
	c.Type = &typeCode
}

func (c *CrawlerJobSubmitV1Request) SetCallbackUrl(callbackUrl string) {
	c.CallbackUrl = &callbackUrl
}

func (c *CrawlerJobSubmitV1Request) SetSiteName(siteName string) {
	c.SiteName = &siteName
}

func (c *CrawlerJobSubmitV1Request) SetCheckStrategy(checkStrategy int) {
	c.CheckStrategy = &checkStrategy
}

func (c *CrawlerJobSubmitV1Request) SetConfig(config string) {
	c.Config = &config
}

func (c *CrawlerJobSubmitV1Request) SetUserAgentMatchType(userAgentMatchType int) {
	c.UserAgentMatchType = &userAgentMatchType
}

func (c *CrawlerJobSubmitV1Request) SetUserAgent(userAgent string) {
	c.UserAgent = &userAgent
}

func (c *CrawlerJobSubmitV1Request) SetFocusList(focusList []string) {
	c.FocusList = &focusList
}

func (c *CrawlerJobSubmitV1Request) SetUrlFilters(urlFilters []CrawlerUrlFilter) {
	c.UrlFilters = &urlFilters
}

func (c *CrawlerUrlFilter) SetType(typeCode int) {
	c.Type = &typeCode
}

func (c *CrawlerUrlFilter) SetUrl(url string) {
	c.Url = &url
}
func (c *CrawlerJobSubmitV1Request) SetCheckFlags(checkFlag []int) {
	c.CheckFlags = &checkFlag
}

func (c *CrawlerJobSubmitV1Request) GetBusinessCustomSignParams() map[string]string {
	data := c.PostFormRequest.GetBusinessCustomSignParams()

	if c.DataId != nil {
		data["dataId"] = *c.DataId
	}
	if c.SliceStartTime != nil {
		data["sliceStartTime"] = strconv.FormatInt(*c.SliceStartTime, 10)
	}
	if c.SliceEndTime != nil {
		data["sliceEndTime"] = strconv.FormatInt(*c.SliceEndTime, 10)
	}
	if c.SiteUrl != nil {
		data["siteUrl"] = *c.SiteUrl
	}
	if c.Level != nil {
		data["level"] = strconv.Itoa(*c.Level)
	}
	if c.Frequency != nil {
		data["frequency"] = strconv.FormatInt(*c.Frequency, 10)
	}
	if c.MaxResourceAmount != nil {
		data["maxResourceAmount"] = strconv.Itoa(*c.MaxResourceAmount)
	}
	if c.Type != nil {
		data["type"] = strconv.Itoa(*c.Type)
	}
	if c.CallbackUrl != nil {
		data["callbackUrl"] = *c.CallbackUrl
	}
	if c.SiteName != nil {
		data["siteName"] = *c.SiteName
	}
	if c.CheckStrategy != nil {
		data["checkStrategy"] = strconv.Itoa(*c.CheckStrategy)
	}
	if c.Config != nil {
		data["config"] = *c.Config
	}
	if c.UserAgentMatchType != nil {
		data["userAgentMatchType"] = strconv.Itoa(*c.UserAgentMatchType)
	}
	if c.UserAgent != nil {
		data["userAgent"] = *c.UserAgent
	}
	if c.FocusList != nil {
		data["focusList"] = strings.Join(*c.FocusList, ",")
	}
	if c.UrlFilters != nil {
		urlFiltersJson, _ := json.Marshal(*c.UrlFilters)
		data["urlFilters"] = string(urlFiltersJson)
	}
	if c.CheckFlags != nil {
		checkFlagsJson, _ := json.Marshal(c.CheckFlags)
		data["checkFlags"] = strings.Trim(string(checkFlagsJson), "[]")
	}

	return data
}
