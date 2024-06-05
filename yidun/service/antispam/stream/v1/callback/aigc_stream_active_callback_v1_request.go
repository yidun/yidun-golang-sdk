package callback

import (
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/callback"
)

/**
 * AIGC流式检测解决方案主动回调请求
 */
type AigcStreamActiveCallbackV1Request struct {
	*callback.ActiveCallbackRequest
}

func NewAigcStreamActiveCallbackV1Request(params map[string]string) *AigcStreamActiveCallbackV1Request {

	return &AigcStreamActiveCallbackV1Request{callback.NewActiveCallbackRequest(params)}
}
