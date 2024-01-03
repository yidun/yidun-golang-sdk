package response

import "github.com/yidun/yidun-golang-sdk/yidun/core/types"

// LiveWallSolutionCallbackV3Resp represents live wall solution callback response.
type LiveWallSolutionCallbackV3Resp struct {
	*types.CommonResponse                                      // Embedded struct, represents common response
	Result                *[]LiveVideoSolutionCallbackV3Result `json:"result,omitempty"` // The result of the callback
}
