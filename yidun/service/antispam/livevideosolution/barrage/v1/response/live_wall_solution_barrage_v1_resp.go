package response

import "github.com/yidun/yidun-golang-sdk/yidun/core/types"

type LiveWallSolutionBarrageV1Resp struct {
	*types.CommonResponse
	Result *bool `json:"result"` // result
}
