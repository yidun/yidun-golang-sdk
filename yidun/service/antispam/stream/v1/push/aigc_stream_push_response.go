package push

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/stream/v1/common"
)

// AIGC流式检测解决方案检测提交返回结果
type AigcStreamPushResponseV1 struct {
	*types.CommonResponse
	Result *common.AigcStreamCheckResult `json:"result"`
}
