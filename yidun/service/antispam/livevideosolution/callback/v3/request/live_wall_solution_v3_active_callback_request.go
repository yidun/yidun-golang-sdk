package request

import "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/callback"

type LiveWallSolutionV3ActiveCallbackRequest struct {
	*callback.ActiveCallbackRequest
}

func NewLiveWallSolutionV3ActiveCallbackRequest(params map[string]string) *LiveWallSolutionV3ActiveCallbackRequest {

	return &LiveWallSolutionV3ActiveCallbackRequest{callback.NewActiveCallbackRequest(params)}
}
