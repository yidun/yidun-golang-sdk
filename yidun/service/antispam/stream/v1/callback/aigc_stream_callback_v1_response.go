package callback

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/stream/v1/common"
)

type AigcStreamCallbackResponseV1 struct {
	*types.CommonResponse
	Result []*common.AigcStreamCheckResult `json:"result,omitempty"`
}
