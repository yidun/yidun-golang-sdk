package update

import (
	"strconv"

	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/core/validation"
)

// CrawlerJobUpdateV1Request 网站任务检测更新接口v1.0
type CrawlerJobUpdateV1Request struct {
	*types.BizPostFormRequest
	// jobId
	JobId *int64 `json:"jobId,omitempty"`
	// 主站url
	SiteUrl *string `json:"siteUrl,omitempty"`
	// 网站名称
	SiteName *string `json:"siteName,omitempty"`
	// 循环爬虫时间区间--开始时间
	SliceStartTime *int64 `json:"sliceStartTime,omitempty"`
}

// For NewCrawlerJobUpdateV1Request
func NewCrawlerJobUpdateV1Request() *CrawlerJobUpdateV1Request {
	var request = &CrawlerJobUpdateV1Request{
		BizPostFormRequest: types.NewBizPostFormRequestWithoutBizId(),
	}
	request.SetProductCode("crawler")
	request.SetUriPattern("/v1/crawler/job/update")
	request.SetMethod(http.HttpMethodPost)
	request.SetVersion("v1.0")
	return request
}

// SetJobId sets the jobId field for the CrawlerJobUpdateV1Request struct.
func (r *CrawlerJobUpdateV1Request) SetJobId(jobId int64) {
	r.JobId = &jobId
}

// SetSiteUrl sets the siteUrl field for the CrawlerJobUpdateV1Request struct.
func (r *CrawlerJobUpdateV1Request) SetSiteUrl(siteUrl string) {
	r.SiteUrl = &siteUrl
}

// SetSiteName sets the siteName field for the CrawlerJobUpdateV1Request struct.
func (r *CrawlerJobUpdateV1Request) SetSiteName(siteName string) {
	r.SiteName = &siteName
}

// SetSliceStartTime sets the sliceStartTime field for the CrawlerJobUpdateV1Request struct.
func (r *CrawlerJobUpdateV1Request) SetSliceStartTime(sliceStartTime int64) {
	r.SliceStartTime = &sliceStartTime
}

func (c *CrawlerJobUpdateV1Request) GetBusinessCustomSignParams() map[string]string {
	data := c.PostFormRequest.GetBusinessCustomSignParams()

	if c.JobId != nil {
		data["jobId"] = strconv.FormatInt(*c.JobId, 10)
	}
	if c.SiteUrl != nil {
		data["siteUrl"] = *c.SiteUrl
	}
	if c.SiteName != nil {
		data["siteName"] = *c.SiteName
	}
	if c.SliceStartTime != nil {
		data["sliceStartTime"] = strconv.FormatInt(*c.SliceStartTime, 10)
	}
	return data
}

func (c *CrawlerJobUpdateV1Request) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "CrawlerJobUpdateV1Request"}
	if c.JobId == nil {
		invalidParams.Add(validation.NewErrParamRequired("JobId不能为空"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}
