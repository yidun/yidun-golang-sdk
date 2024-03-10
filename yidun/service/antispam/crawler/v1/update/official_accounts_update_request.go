package update

import (
	"strconv"

	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/core/validation"
)

// OfficialAccountsUpdateV1Request 公众号检测更新接口v1.0
type OfficialAccountsUpdateV1Request struct {
	*types.BizPostFormRequest
	// @NotNull
	JobId *int64 `json:"jobId,omitempty"`
	// 微信号
	WechatAccount *string `json:"wechatAccount,omitempty"`
	// 公众号名称
	OfficialAccountName *string `json:"officialAccountName,omitempty"`
	// 循环爬虫时间区间--开始时间
	StartTime *int64 `json:"startTime,omitempty"`
}

// For OfficialAccountsUpdateV1Request
func NewOfficialAccountsUpdateV1Request() *OfficialAccountsUpdateV1Request {
	var request = &OfficialAccountsUpdateV1Request{
		BizPostFormRequest: types.NewBizPostFormRequestWithoutBizId(),
	}
	request.SetProductCode("crawler")
	request.SetUriPattern("/v1/official-accounts/job/update")
	request.SetMethod(http.HttpMethodPost)
	request.SetVersion("v1.0")
	return request
}

// SetJobId sets the jobId field for the OfficialAccountsUpdateV1Request struct.
func (r *OfficialAccountsUpdateV1Request) SetJobId(jobId int64) {
	r.JobId = &jobId
}

// SetWechatAccount sets the wechatAccount field for the OfficialAccountsUpdateV1Request struct.
func (r *OfficialAccountsUpdateV1Request) SetWechatAccount(wechatAccount string) {
	r.WechatAccount = &wechatAccount
}

// SetOfficialAccountName sets the officialAccountName field for the OfficialAccountsUpdateV1Request struct.
func (r *OfficialAccountsUpdateV1Request) SetOfficialAccountName(officialAccountName string) {
	r.OfficialAccountName = &officialAccountName
}

// SetStartTime sets the startTime field for the OfficialAccountsUpdateV1Request struct.
func (r *OfficialAccountsUpdateV1Request) SetStartTime(startTime int64) {
	r.StartTime = &startTime
}

func (c *OfficialAccountsUpdateV1Request) GetBusinessCustomSignParams() map[string]string {
	data := c.PostFormRequest.GetBusinessCustomSignParams()

	if c.JobId != nil {
		data["jobId"] = strconv.FormatInt(*c.JobId, 10)
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
	return data
}

func (c *OfficialAccountsUpdateV1Request) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "OfficialAccountsUpdateV1Request"}
	if c.JobId == nil {
		invalidParams.Add(validation.NewErrParamRequired("JobId不能为空"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}
