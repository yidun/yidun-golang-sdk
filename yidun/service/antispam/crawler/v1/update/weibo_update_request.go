package update

import (
	"strconv"

	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/core/validation"
)

// WeiboUpdateV1Request 微博检测更新接口v1.0
type WeiboUpdateV1Request struct {
	*types.BizPostFormRequest
	// @NotNull
	JobId *int64 `json:"jobId,omitempty"` // Although `@NotNull` is provided, it's translated to `omitempty` to follow Go convention.
	// 微博名
	Blogger *string `json:"blogger,omitempty"`
	// 微博链接
	Url *string `json:"url,omitempty"`
	// 循环爬虫时间区间--开始时间
	StartTime *int64 `json:"startTime,omitempty"`
}

// For WeiboUpdateV1Request
func NewWeiboUpdateV1Request() *WeiboUpdateV1Request {
	var request = &WeiboUpdateV1Request{
		BizPostFormRequest: types.NewBizPostFormRequestWithoutBizId(),
	}
	request.SetProductCode("crawler")
	request.SetUriPattern("/v1/crawler/weibo-job/update")
	request.SetMethod(http.HttpMethodPost)
	request.SetVersion("v1.0")
	return request
}

// SetJobId sets the jobId for WeiboUpdateV1Request.
func (r *WeiboUpdateV1Request) SetJobId(jobId int64) {
	r.JobId = &jobId
}

// SetBlogger sets the blogger for WeiboUpdateV1Request.
func (r *WeiboUpdateV1Request) SetBlogger(blogger string) {
	r.Blogger = &blogger
}

// SetUrl sets the url for WeiboUpdateV1Request.
func (r *WeiboUpdateV1Request) SetUrl(url string) {
	r.Url = &url
}

// SetStartTime sets the startTime for WeiboUpdateV1Request.
func (r *WeiboUpdateV1Request) SetStartTime(startTime int64) {
	r.StartTime = &startTime
}

func (c *WeiboUpdateV1Request) GetBusinessCustomSignParams() map[string]string {
	data := c.PostFormRequest.GetBusinessCustomSignParams()

	if c.JobId != nil {
		data["jobId"] = strconv.FormatInt(*c.JobId, 10)
	}
	if c.Blogger != nil {
		data["blogger"] = *c.Blogger
	}
	if c.Url != nil {
		data["url"] = *c.Url
	}
	if c.StartTime != nil {
		data["startTime"] = strconv.FormatInt(*c.StartTime, 10)
	}
	return data
}

func (c *WeiboUpdateV1Request) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "WeiboUpdateV1Request"}
	if c.JobId == nil {
		invalidParams.Add(validation.NewErrParamRequired("JobId不能为空"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}
