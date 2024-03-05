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
type OfficialAccountsBatchSubmitV1Request struct {
	*types.BizPostFormRequest
	// 公众号检测任务公共参数
	OfficialAccountsBaseSubmitV1Request
	// 检测标记 检测内容
	OfficialAccounts *[]OfficicalAccountsConfig `json:"officialAccounts,omitempty"`
}

// For NewOfficialAccountsSubmitV1Request
func NewOfficialAccountsBatchSubmitV1Request() *OfficialAccountsBatchSubmitV1Request {
	var request = &OfficialAccountsBatchSubmitV1Request{
		BizPostFormRequest: types.NewBizPostFormRequestWithoutBizId(),
	}
	request.SetProductCode("crawler")
	request.SetUriPattern("/v1/official-accounts/job/batch-submit")
	request.SetMethod(http.HttpMethodPost)
	request.SetVersion("v1.0")
	return request
}

// SetOfficialAccounts
func (o *OfficialAccountsBatchSubmitV1Request) SetOfficialAccounts(officialAccounts []OfficicalAccountsConfig) {
	o.OfficialAccounts = &officialAccounts
}

func (c *OfficialAccountsBatchSubmitV1Request) GetBusinessCustomSignParams() map[string]string {
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
	if c.OfficialAccounts != nil {
		officialAccountsJson, _ := json.Marshal(c.OfficialAccounts)
		data["officialAccounts"] = string(officialAccountsJson)
	}
	return data
}

func (c *OfficialAccountsBatchSubmitV1Request) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "OfficialAccountsSubmitV1Request"}

	if c.OfficialAccounts != nil {
		for _, accountsConfig := range *c.OfficialAccounts {
			if accountsConfig.DataId != nil && len(*accountsConfig.DataId) > 128 {
				invalidParams.Add(validation.NewErrParamMaxLen("dataId最长128个字符", 128, strconv.Itoa(len(*accountsConfig.DataId))))
			}
			if accountsConfig.WechatAccount == nil {
				invalidParams.Add(validation.NewErrParamRequired("WechatAccount不能为空"))
			}
			if accountsConfig.WechatAccount != nil && len(*accountsConfig.WechatAccount) > 64 {
				invalidParams.Add(validation.NewErrParamMaxLen("WechatAccount最长64个字符", 64, strconv.Itoa(len(*accountsConfig.WechatAccount))))
			}
			if accountsConfig.OfficialAccountName != nil && len(*accountsConfig.OfficialAccountName) > 64 {
				invalidParams.Add(validation.NewErrParamMaxLen("OfficialAccountName最长128个字符", 64, strconv.Itoa(len(*accountsConfig.OfficialAccountName))))
			}
		}
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}
