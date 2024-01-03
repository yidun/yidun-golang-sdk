package query

import (
	"strconv"

	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/core/validation"
)

type CrawlerJobCallbackQueryRequest struct {
	*types.BizPostFormRequest
	// 任务id
	JobId *int64 `json:"jobId,omitempty"`
	// size
	PageSize *int `json:"pageSize,omitempty"`
	// page
	PageNum *int `json:"pageNum,omitempty"`
}

// For ImageV5CheckRequest  构造请求
func NewCrawlerJobCallbackQueryRequest() *CrawlerJobCallbackQueryRequest {
	var request = &CrawlerJobCallbackQueryRequest{
		BizPostFormRequest: types.NewBizPostFormRequestWithoutBizId(),
	}
	request.SetProductCode("crawler")
	request.SetUriPattern("/v1/crawler/callback-result/query")
	request.SetMethod(http.HttpMethodPost)
	request.SetVersion("v1.0")
	return request
}

func (j *CrawlerJobCallbackQueryRequest) SetJobId(jobId int64) {
	j.JobId = &jobId
}

func (j *CrawlerJobCallbackQueryRequest) SetPageSize(pageSize int) {
	j.PageSize = &pageSize
}

func (j *CrawlerJobCallbackQueryRequest) SetPageNum(pageNum int) {
	j.PageNum = &pageNum
}

// 获取参与签名的参数
func (c *CrawlerJobCallbackQueryRequest) GetBusinessCustomSignParams() map[string]string {
	result := c.PostFormRequest.GetBusinessCustomSignParams()

	if c.JobId != nil {
		result["jobId"] = strconv.FormatInt(*c.JobId, 10)
	}
	if c.PageNum != nil {
		result["pageNum"] = strconv.Itoa(*c.PageNum)
	}
	if c.PageSize != nil {
		result["pageSize"] = strconv.Itoa(*c.PageSize)
	}

	return result
}

// 参数校验
func (c *CrawlerJobCallbackQueryRequest) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "CrawlerJobCallbackQueryRequest"}
	if c.JobId == nil {
		invalidParams.Add(validation.NewErrParamRequired("任务Id不能为空"))
	}

	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}
