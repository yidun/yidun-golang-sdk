package callback

import (
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/callback"
)

type FileV2ActiveCallbackRequest struct {
	*callback.ActiveCallbackRequest
}

func NewFileV2ActiveCallbackRequest(params map[string]string) *FileV2ActiveCallbackRequest {

	return &FileV2ActiveCallbackRequest{callback.NewActiveCallbackRequest(params)}
}
