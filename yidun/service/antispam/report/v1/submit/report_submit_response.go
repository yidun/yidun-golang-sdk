package submit

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/media/v2/common/response"
)

// 投诉举报v1版本提交接口返回结果
type ReportSubmitResponseV1 struct {
	*types.CommonResponse
	Result *ReportSubmitResult `json:"result"`
}

type ReportSubmitResult struct {
	// 反垃圾结果
	Antispam *response.MediaAntispamResponse `json:"antispam,omitempty"`
}
