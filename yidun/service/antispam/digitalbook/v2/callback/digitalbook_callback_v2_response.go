package callback

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/media/v2/common/response"
)

type DigitalBookCallbackResponseV2 struct {
	*types.CommonResponse
	Result []*DigitalBookCheckResult `json:"result,omitempty"`
}

type DigitalBookCheckResult struct {
	// 反垃圾结果
	Antispam *response.MediaAntispamResponse `json:"antispam,omitempty"`
	// 增值服务结果
	ValueAddService *response.MediaValueAddServiceResponse `json:"valueAddService,omitempty"`
	// 反作弊信息
	Anticheat *response.MediaAntiCheatResponse `json:"anticheat,omitempty"`
	// 审核信息
	Censor *response.MediaCensorResponse `json:"censor,omitempty"`
}
