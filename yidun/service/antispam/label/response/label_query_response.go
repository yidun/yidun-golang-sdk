package response

import "github.com/yidun/yidun-golang-sdk/yidun/core/types"

type LabelQueryResponse struct {
    *types.CommonResponse
    Data []*LabelQueryInfo `json:"data,omitempty"` // data
}
