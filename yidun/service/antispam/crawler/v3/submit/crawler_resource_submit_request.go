package submit

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/core/validation"
)

type CrawlerResourceSubmitV3Request struct {
	*types.BizPostFormRequest
	// 资源URL
	Url *string `json:"url,omitempty"`
	// 数据唯一标识，能够根据该值定位到该条数据，如对数据结果有异议，可以发送该值给客户经理查询
	DataId *string `json:"dataId,omitempty"`
	// 用户传过来的
	Callback *string `json:"callback,omitempty"`
	// 爬虫检测控制，控制需要爬取页面的哪些内容；1：检测文本；2：检测图片；4：检测点播音频；5：检测文档附件；6：检测点播音视频；
	CheckFlags *[]int `json:"checkFlags,omitempty"`
	// 回调接口地址
	CallbackUrl *string `json:"callbackUrl,omitempty"`
	// 网站名称
	SiteName *string `json:"siteName,omitempty"`
	// HTML内容
	Content *string `json:"content,omitempty"`
	// 类型检测数据配置
	Config *string `json:"config,omitempty"`
	// 业务指定过检策略组id
	CheckStrategyGroupIds *[]int64 `json:"checkStrategyGroupIds,omitempty"`
	// 自定义扩展参数，JSON字符串格式。如："{"keyName1":"value1","keyName2":"value2"}"
	Extension *string `json:"extension,omitempty"`
}

// For ImageV5CheckRequest
func NewCrawlerResourceV3SubmitRequest() *CrawlerResourceSubmitV3Request {
	var request = &CrawlerResourceSubmitV3Request{
		BizPostFormRequest: types.NewBizPostFormRequestWithoutBizId(),
	}
	request.SetProductCode("crawler")
	request.SetUriPattern("/v3/crawler/submit")
	request.SetMethod(http.HttpMethodPost)
	request.SetVersion("v3.0")
	return request
}

func (c *CrawlerResourceSubmitV3Request) SetUrl(url string) {
	c.Url = &url
}

func (c *CrawlerResourceSubmitV3Request) SetDataId(dataId string) {
	c.DataId = &dataId
}

func (c *CrawlerResourceSubmitV3Request) SetCallback(callback string) {
	c.Callback = &callback
}

func (c *CrawlerResourceSubmitV3Request) SetCheckFlags(checkFlags []int) {
	c.CheckFlags = &checkFlags
}

func (c *CrawlerResourceSubmitV3Request) SetCallbackUrl(callbackUrl string) {
	c.CallbackUrl = &callbackUrl
}

func (c *CrawlerResourceSubmitV3Request) SetSiteName(siteName string) {
	c.SiteName = &siteName
}

func (c *CrawlerResourceSubmitV3Request) SetContent(content string) {
	c.Content = &content
}

func (c *CrawlerResourceSubmitV3Request) SetConfig(config string) {
	c.Config = &config
}

func (c *CrawlerResourceSubmitV3Request) SetCheckStrategyGroupIds(checkStrategyGroupIds []int64) {
	c.CheckStrategyGroupIds = &checkStrategyGroupIds
}

func (c *CrawlerResourceSubmitV3Request) SetExtension(extension string) {
	c.Extension = &extension
}

func (c *CrawlerResourceSubmitV3Request) GetBusinessCustomSignParams() map[string]string {
	result := c.PostFormRequest.GetBusinessCustomSignParams()

	if c.Url != nil {
		result["url"] = *c.Url
	}
	if c.DataId != nil {
		result["dataId"] = *c.DataId
	}
	if c.Callback != nil {
		result["callback"] = *c.Callback
	}
	if c.CheckFlags != nil {
		checkFlagsJson, _ := json.Marshal(c.CheckFlags)
		result["checkFlags"] = strings.Trim(string(checkFlagsJson), "[]")
	}
	if c.CallbackUrl != nil {
		result["callbackUrl"] = *c.CallbackUrl
	}
	if c.SiteName != nil {
		result["siteName"] = *c.SiteName
	}
	if c.Content != nil {
		result["content"] = *c.Content
	}
	if c.Config != nil {
		result["config"] = *c.Config
	}
	if c.CheckStrategyGroupIds != nil {
		checkStrategyGroupIdsJson, _ := json.Marshal(c.CheckStrategyGroupIds)
		result["checkStrategyGroupIds"] = strings.Trim(string(checkStrategyGroupIdsJson), "[]")
	}
	if c.Extension != nil {
		result["extension"] = *c.Extension
	}
	return result
}

func (c *CrawlerResourceSubmitV3Request) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "CrawlerResourceSubmitV3Request"}
	if c.Url != nil && len(*c.Url) > 1024 {
		invalidParams.Add(validation.NewErrParamMaxLen("url最长1024个字符", 1024, strconv.Itoa(len(*c.Url))))
	}
	if c.CallbackUrl != nil && len(*c.CallbackUrl) > 1024 {
		invalidParams.Add(validation.NewErrParamMaxLen("callbackUrl最长1024个字符", 1024, strconv.Itoa(len(*c.CallbackUrl))))
	}
	if c.SiteName != nil && len(*c.SiteName) > 64 {
		invalidParams.Add(validation.NewErrParamMaxLen("siteName最长64个字符", 64, strconv.Itoa(len(*c.SiteName))))
	}
	if c.DataId != nil && len(*c.DataId) > 128 {
		invalidParams.Add(validation.NewErrParamMaxLen("dataId最长128个字符", 128, strconv.Itoa(len(*c.DataId))))
	}
	if c.Callback != nil && len(*c.Callback) > 512 {
		invalidParams.Add(validation.NewErrParamMaxLen("callback最长512个字符", 512, strconv.Itoa(len(*c.Callback))))
	}
	if c.CheckFlags == nil {
		invalidParams.Add(validation.NewErrParamRequired("checkFlags不能为空"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}
