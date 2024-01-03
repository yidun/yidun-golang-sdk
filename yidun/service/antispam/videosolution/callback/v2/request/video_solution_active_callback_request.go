package request

import (
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/callback"
)

type VideoSolutionV2ActiveCallbackRequest struct {
	*callback.ActiveCallbackRequest
}

func NewVideoSolutionV2ActiveCallbackRequest(params map[string]string) *VideoSolutionV2ActiveCallbackRequest {

	return &VideoSolutionV2ActiveCallbackRequest{callback.NewActiveCallbackRequest(params)}
}
