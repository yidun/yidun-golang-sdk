package submit

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/core/validation"
)

// WeiboBatchSubmitV1Request 微博检测提交接口v1.0
type WeiboBatchSubmitV1Request struct {
	*types.BizPostFormRequest
	// 公众号检测任务公共参数
	WeiboBaseSubmitV1Request
	// 公众号相关参数
	WeiboBloggers *[]WeiboConfig `json:"weiboBloggers,omitempty"`
}

// For NewOfficialAccountsSubmitV1Request
func NewWeiboBatchSubmitV1Request() *WeiboBatchSubmitV1Request {
	var request = &WeiboBatchSubmitV1Request{
		BizPostFormRequest: types.NewBizPostFormRequestWithoutBizId(),
	}
	request.SetProductCode("crawler")
	request.SetUriPattern("/v1/crawler/weibo-job/batch-submit")
	request.SetMethod(http.HttpMethodPost)
	request.SetVersion("v1.0")
	return request
}

// SetWeiboBloggers 设置检测标记 检测内容
func (o *WeiboBatchSubmitV1Request) SetWeiboBloggers(blogger []WeiboConfig) {
	o.WeiboBloggers = &blogger
}

func (c *WeiboBatchSubmitV1Request) GetBusinessCustomSignParams() map[string]string {
	data := c.PostFormRequest.GetBusinessCustomSignParams()
	if c.StartTime != nil {
		data["startTime"] = strconv.FormatInt(*c.StartTime, 10)
	}
	if c.EndTime != nil {
		data["endTime"] = strconv.FormatInt(*c.EndTime, 10)
	}
	if c.MaxCheckCount != nil {
		data["maxCheckCount"] = strconv.Itoa(*c.MaxCheckCount)
	}
	if c.MaxComment != nil {
		data["maxComment"] = strconv.Itoa(*c.MaxComment)
	}
	if c.Frequency != nil {
		data["frequency"] = strconv.FormatInt(*c.Frequency, 10)
	}
	if c.Type != nil {
		data["type"] = strconv.Itoa(*c.Type)
	}
	if c.CallbackUrl != nil {
		data["callbackUrl"] = *c.CallbackUrl
	}
	if c.Strategy != nil {
		data["strategy"] = strconv.Itoa(*c.Strategy)
	}
	if c.CheckFlags != nil {
		checkFlagsJson, _ := json.Marshal(c.CheckFlags)
		data["checkFlags"] = strings.Trim(string(checkFlagsJson), "[]")
	}
	if c.WeiboBloggers != nil {
		weiboBloggersJson, _ := json.Marshal(c.WeiboBloggers)
		data["weiboBloggers"] = string(weiboBloggersJson)
	}
	return data
}

func (c *WeiboBatchSubmitV1Request) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "WeiboBatchSubmitV1Request"}
	for _, weiboConfig := range *c.WeiboBloggers {
		if weiboConfig.Url == nil {
			invalidParams.Add(validation.NewErrParamRequired("Url不能为空"))
		}
		if weiboConfig.Url != nil && len(*weiboConfig.Url) > 1024 {
			invalidParams.Add(validation.NewErrParamMaxLen("Url最长1024个字符", 1024, strconv.Itoa(len(*weiboConfig.Url))))
		}
		if weiboConfig.Blogger != nil && len(*weiboConfig.Blogger) > 64 {
			invalidParams.Add(validation.NewErrParamMaxLen("Blogger最长128个字符", 64, strconv.Itoa(len(*weiboConfig.Blogger))))
		}
	}

	if c.MaxCheckCount != nil && *c.MaxCheckCount > 2000 {
		invalidParams.Add(validation.NewErrParamInvalid("MaxCheckCount"))
	}
	if c.MaxComment != nil && *c.MaxComment > 2000 {
		invalidParams.Add(validation.NewErrParamInvalid("MaxComment"))
	}
	if c.CallbackUrl != nil && len(*c.CallbackUrl) > 1024 {
		invalidParams.Add(validation.NewErrParamMaxLen("CallbackUrl最长1024个字符", 1024, strconv.Itoa(len(*c.CallbackUrl))))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}
