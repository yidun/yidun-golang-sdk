package callback

import (
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/callback"
)

type TextV5ActiveCallbackRequest struct {
    *callback.ActiveCallbackRequest
}

func NewTextV5ActiveCallbackRequest(params map[string]string) *TextV5ActiveCallbackRequest {
	
	return &TextV5ActiveCallbackRequest{callback.NewActiveCallbackRequest(params)}
}