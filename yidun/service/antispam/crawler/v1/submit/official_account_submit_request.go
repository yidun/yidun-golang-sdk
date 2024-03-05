package submit

import (
	"encoding/json"
	"strconv"
	"strings"

	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/core/validation"
)

// OfficialAccountsSubmitV1Request 公众号检测提交接口v1.0
type OfficialAccountsSubmitV1Request struct {
	*types.BizPostFormRequest
	// 公众号检测任务公共参数
	OfficialAccountsBaseSubmitV1Request
	// 公众号相关参数
	OfficicalAccountsConfig
}

// For NewOfficialAccountsSubmitV1Request
func NewOfficialAccountsSubmitV1Request() *OfficialAccountsSubmitV1Request {
	var request = &OfficialAccountsSubmitV1Request{
		BizPostFormRequest: types.NewBizPostFormRequestWithoutBizId(),
	}
	request.SetProductCode("crawler")
	request.SetUriPattern("/v1/official-accounts/job/submit")
	request.SetMethod(http.HttpMethodPost)
	request.SetVersion("v1.0")
	return request
}

func (c *OfficialAccountsSubmitV1Request) GetBusinessCustomSignParams() map[string]string {
	data := c.PostFormRequest.GetBusinessCustomSignParams()

	if c.DataId != nil {
		data["dataId"] = *c.DataId
	}
	if c.WechatAccount != nil {
		data["wechatAccount"] = *c.WechatAccount
	}
	if c.OfficialAccountName != nil {
		data["officialAccountName"] = *c.OfficialAccountName
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

func (c *OfficialAccountsSubmitV1Request) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "OfficialAccountsSubmitV1Request"}
	if c.DataId != nil && len(*c.DataId) > 128 {
		invalidParams.Add(validation.NewErrParamMaxLen("dataId最长128个字符", 128, strconv.Itoa(len(*c.DataId))))
	}
	if c.WechatAccount == nil {
		invalidParams.Add(validation.NewErrParamRequired("WechatAccount不能为空"))
	}
	if c.WechatAccount != nil && len(*c.WechatAccount) > 64 {
		invalidParams.Add(validation.NewErrParamMaxLen("WechatAccount最长64个字符", 64, strconv.Itoa(len(*c.WechatAccount))))
	}
	if c.OfficialAccountName != nil && len(*c.OfficialAccountName) > 64 {
		invalidParams.Add(validation.NewErrParamMaxLen("OfficialAccountName最长128个字符", 64, strconv.Itoa(len(*c.OfficialAccountName))))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}
