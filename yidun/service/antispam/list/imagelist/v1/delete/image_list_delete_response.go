package delete

import "github.com/yidun/yidun-golang-sdk/yidun/core/types"

type ImageListDeleteResponse struct {
	*types.CommonResponse
	Result *bool `json:"result,omitempty"`
}
