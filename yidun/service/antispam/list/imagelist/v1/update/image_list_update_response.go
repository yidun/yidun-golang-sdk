package update

import "github.com/yidun/yidun-golang-sdk/yidun/core/types"

type ImageListUpdateResponse struct {
	*types.CommonResponse
	Result *bool `json:"result,omitempty"`
}
