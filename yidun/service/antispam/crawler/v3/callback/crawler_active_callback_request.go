package callback

import (
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/callback"
)

type CrawlerResourceActiveCallbackV3Request struct {
	*callback.ActiveCallbackRequest
}

func NewCrawlerResourceActiveCallbackV3Request(params map[string]string) *CrawlerResourceActiveCallbackV3Request {

	return &CrawlerResourceActiveCallbackV3Request{callback.NewActiveCallbackRequest(params)}
}
