package submit

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
)

type CrawlerJobBatchSubmitV1Request struct {
	*types.BizPostFormRequest
	// 网站提交公共参数
	CrawlerJobBaseSubmitV1Request
	// URL过滤条件集
	Websites *[]CrawlerJobWebSite `json:"websites,omitempty"`
}

type CrawlerJobWebSite struct {
	// 数据唯一标识，能够根据该值定位到该条数据，如对数据结果有异议，可以发送该值给客户经理查询
	DataId *string `json:"dataId,omitempty"`
	// 主站URL
	SiteUrl *string `json:"siteUrl,omitempty"`
	// 网站名称
	SiteName *string `json:"siteName,omitempty"`
	// 重点关注资源
	FocusList *[]string `json:"focusList,omitempty"`
	// URL过滤条件集
	UrlFilters *[]CrawlerUrlFilter `json:"urlFilters,omitempty"`
	// 用户账号
	Account *string `json:"account,omitempty"`
}

// For ImageV5CheckRequest
func NewCrawlerJobBatchSubmitV1Request() *CrawlerJobBatchSubmitV1Request {
	var request = &CrawlerJobBatchSubmitV1Request{
		BizPostFormRequest: types.NewBizPostFormRequestWithoutBizId(),
	}
	request.SetProductCode("crawler")
	request.SetUriPattern("/v1/crawler/job/batch-submit")
	request.SetMethod(http.HttpMethodPost)
	request.SetVersion("v1.0")
	return request
}

func (c *CrawlerJobBatchSubmitV1Request) SetWebsites(websites []CrawlerJobWebSite) {
	c.Websites = &websites
}

func (c *CrawlerJobWebSite) SetDataId(dataId string) {
	c.DataId = &dataId
}

func (c *CrawlerJobWebSite) SetSiteUrl(siteUrl string) {
	c.SiteUrl = &siteUrl
}

func (c *CrawlerJobWebSite) SetSiteName(siteName string) {
	c.SiteName = &siteName
}

func (c *CrawlerJobWebSite) SetFocusList(focusList []string) {
	c.FocusList = &focusList
}

func (c *CrawlerJobWebSite) SetUrlFilters(urlFilters []CrawlerUrlFilter) {
	c.UrlFilters = &urlFilters
}

// 设置用户账号
func (c *CrawlerJobWebSite) SetAccount(account string) {
	c.Account = &account
}

func (c *CrawlerJobBatchSubmitV1Request) GetBusinessCustomSignParams() map[string]string {
	data := c.PostFormRequest.GetBusinessCustomSignParams()

	if c.SliceStartTime != nil {
		data["sliceStartTime"] = strconv.FormatInt(*c.SliceStartTime, 10)
	}
	if c.SliceEndTime != nil {
		data["sliceEndTime"] = strconv.FormatInt(*c.SliceEndTime, 10)
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
	if c.CheckFlags != nil {
		checkFlagsJson, _ := json.Marshal(c.CheckFlags)
		data["checkFlags"] = strings.Trim(string(checkFlagsJson), "[]")
	}
	if c.Websites != nil {
		websitesJson, _ := json.Marshal(*c.Websites)
		data["websites"] = string(websitesJson)
	}

	return data
}
