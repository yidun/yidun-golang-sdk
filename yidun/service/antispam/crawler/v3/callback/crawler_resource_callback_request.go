package callback

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/core/validation"
)

type CrawlerResourceCallbackV3Request struct {
	*types.BizPostFormRequest
	YidunRequestId *string `json:"yidunRequestId,omitempty"`
}

func (r *CrawlerResourceCallbackV3Request) SetYidunRequestId(yidunRequestId string) {
	r.YidunRequestId = &yidunRequestId
}

// 构建request
func NewCrawlerResourceCallbackRequest() *CrawlerResourceCallbackV3Request {
	var request = &CrawlerResourceCallbackV3Request{
		BizPostFormRequest: types.NewBizPostFormRequestWithoutBizId(),
	}
	request.SetProductCode("crawler")
	request.SetUriPattern("/v3/crawler/callback/results")
	request.SetVersion("v3.0")
	return request
}

// 获取业务自定义参数
func (r *CrawlerResourceCallbackV3Request) GetBusinessCustomSignParams() map[string]string {
	params := r.BizPostFormRequest.GetBusinessCustomSignParams()
	if r.YidunRequestId != nil {
		params["yidunRequestId"] = *r.YidunRequestId
	}

	return params
}

// 参数校验
func (r *CrawlerResourceCallbackV3Request) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "CrawlerResourceCallbackV3Request"}
	if r == nil {
		invalidParams.Add(validation.NewErrParamRequired("CrawlerResourceCallbackV3Request"))
		return invalidParams
	}

	return nil
}
