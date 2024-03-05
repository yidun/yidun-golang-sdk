package submit

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/core/validation"
)

// WeiboSubmitV1Request 微博检测提交接口v1.0
type WeiboSubmitV1Request struct {
	*types.BizPostFormRequest
	// 公众号检测任务公共参数
	WeiboBaseSubmitV1Request
	// 公众号相关参数
	WeiboConfig
}

// For NewOfficialAccountsSubmitV1Request
func NewWeiboSubmitV1Request() *WeiboSubmitV1Request {
	var request = &WeiboSubmitV1Request{
		BizPostFormRequest: types.NewBizPostFormRequestWithoutBizId(),
	}
	request.SetProductCode("crawler")
	request.SetUriPattern("/v1/crawler/weibo-job/submit")
	request.SetMethod(http.HttpMethodPost)
	request.SetVersion("v1.0")
	return request
}

func (c *WeiboSubmitV1Request) GetBusinessCustomSignParams() map[string]string {
	data := c.PostFormRequest.GetBusinessCustomSignParams()

	if c.Url != nil {
		data["url"] = *c.Url
	}
	if c.Blogger != nil {
		data["blogger"] = *c.Blogger
	}
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
	return data
}

func (c *WeiboSubmitV1Request) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "WeiboSubmitV1Request"}
	if c.Url == nil {
		invalidParams.Add(validation.NewErrParamRequired("Url不能为空"))
	}
	if c.Url != nil && len(*c.Url) > 1024 {
		invalidParams.Add(validation.NewErrParamMaxLen("Url最长1024个字符", 1024, strconv.Itoa(len(*c.Url))))
	}
	if c.Blogger != nil && len(*c.Blogger) > 64 {
		invalidParams.Add(validation.NewErrParamMaxLen("Blogger最长128个字符", 64, strconv.Itoa(len(*c.Blogger))))
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
