package callback

import (
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/callback"
)

type ImageV5ActiveCallbackRequest struct {
	*callback.ActiveCallbackRequest
}

func NewImageV5ActiveCallbackRequest(params map[string]string) *ImageV5ActiveCallbackRequest {

	return &ImageV5ActiveCallbackRequest{callback.NewActiveCallbackRequest(params)}
}
