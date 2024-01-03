package query

import (
	"strings"

	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/core/validation"
)

type CrawlerResourceQueryV3Request struct {
	*types.BizPostFormRequest
	// 资源URL
	TaskIdList *[]string `json:"taskIdList,omitempty"`
}

// For ImageV5CheckRequest
func NewCrawlerResourceV3QueryRequest() *CrawlerResourceQueryV3Request {
	var request = &CrawlerResourceQueryV3Request{
		BizPostFormRequest: types.NewBizPostFormRequestWithoutBizId(),
	}
	request.SetProductCode("crawler")
	request.SetUriPattern("/v3/crawler/callback/query")
	request.SetMethod(http.HttpMethodPost)
	request.SetVersion("v3.0")
	return request
}

func (c *CrawlerResourceQueryV3Request) SetTaskIdList(taskIdList []string) {
	c.TaskIdList = &taskIdList
}

func (c *CrawlerResourceQueryV3Request) GetBusinessCustomSignParams() map[string]string {
	result := c.PostFormRequest.GetBusinessCustomSignParams()

	if c.TaskIdList != nil {
		result["taskIdList"] = strings.Join(*c.TaskIdList, ",")
	}

	return result
}

func (c *CrawlerResourceQueryV3Request) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "CrawlerResourceQueryV3Request"}
	if c.TaskIdList == nil {
		invalidParams.Add(validation.NewErrParamRequired("TaskIdList不能为空"))
		return invalidParams
	}
	return nil
}
