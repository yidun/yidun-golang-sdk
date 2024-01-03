package callback

import (
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/callback"
)

/**
 * 数字阅读解决方案主动回调请求
 */
type DigitalBookActiveCallbackV2Request struct {
	*callback.ActiveCallbackRequest
}

func NewDigitalBookActiveCallbackV2Request(params map[string]string) *DigitalBookActiveCallbackV2Request {

	return &DigitalBookActiveCallbackV2Request{callback.NewActiveCallbackRequest(params)}
}
