package request

import (
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/callback"
)

type VideoV4ActiveCallbackRequest struct {
	*callback.ActiveCallbackRequest
}

func NewVideoV4ActiveCallbackRequest(params map[string]string) *VideoV4ActiveCallbackRequest {

	return &VideoV4ActiveCallbackRequest{callback.NewActiveCallbackRequest(params)}
}
