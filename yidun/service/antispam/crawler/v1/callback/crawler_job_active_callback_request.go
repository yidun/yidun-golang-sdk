package callback

import "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/callback"

type CrawlerJobActiveCallbackRequest struct {
	*callback.ActiveCallbackRequest
}

func NewCrawlerJobActiveCallbackRequest(params map[string]string) *CrawlerJobActiveCallbackRequest {

	return &CrawlerJobActiveCallbackRequest{callback.NewActiveCallbackRequest(params)}
}
