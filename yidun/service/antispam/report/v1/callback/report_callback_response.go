package callback

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/media/v2/common/response"
)

/**
 * 投诉举报检测结果获取（轮询模式）
 */
type ReportCallbackResponseV1 struct {
	*types.CommonResponse
	Result []*ReportCheckResult `json:"result,omitempty"`
}

type ReportCheckResult struct {
	// 反垃圾结果
	Antispam *response.MediaAntispamResponse `json:"antispam,omitempty"`
	// 审核信息
	Censor *response.MediaCensorResponse `json:"censor,omitempty"`
}
