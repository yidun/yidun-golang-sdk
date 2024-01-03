package callback

import (
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/callback"
)

/**
 * 融媒体解决方案主动回调请求
 */
type MediaActiveCallbackV2Request struct {
	*callback.ActiveCallbackRequest
}

func NewMediaActiveCallbackV2Request(params map[string]string) *MediaActiveCallbackV2Request {

	return &MediaActiveCallbackV2Request{callback.NewActiveCallbackRequest(params)}
}
