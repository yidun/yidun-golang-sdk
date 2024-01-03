package submit

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/media/v2/common/response"
)

// 融媒体解决方案检测提交返回结果
type MediaSubmitResponseV2 struct {
	*types.CommonResponse
	Result *MediaSubmitResult `json:"result"`
}

type MediaSubmitResult struct {
	// 反垃圾结果
	Antispam *response.MediaAntispamResponse `json:"antispam,omitempty"`
	// 增值服务结果
	ValueAddService *response.MediaValueAddServiceResponse `json:"valueAddService,omitempty"`
	// 反作弊信息
	Anticheat *response.MediaAntiCheatResponse `json:"anticheat,omitempty"`
}
