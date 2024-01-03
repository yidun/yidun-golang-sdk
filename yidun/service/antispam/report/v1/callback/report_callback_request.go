package callback

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
)

/**
 * 投诉举报检测结果获取（轮询模式）
 */
type ReportCallBackRequestV1 struct {
	*types.BizPostFormRequest
	YidunRequestId *string `json:"yidunRequestId,omitempty"` // 请求的唯一ID
}

func (r *ReportCallBackRequestV1) SetYidunRequestId(yidunRequestId string) {
	r.YidunRequestId = &yidunRequestId
}

func NewReportCallBackRequest() *ReportCallBackRequestV1 {
	var request = &ReportCallBackRequestV1{
		BizPostFormRequest: types.NewBizPostFormRequestWithoutBizId(),
	}
	request.SetProductCode("report")
	request.SetUriPattern("/v1/report/callback/results")
	request.SetVersion("v1")
	return request
}

// GetBusinessCustomSignParams 获取具体业务中特有的需要参与签名计算的参数
func (f *ReportCallBackRequestV1) GetBusinessCustomSignParams() map[string]string {
	result := f.PostFormRequest.GetBusinessCustomSignParams()

	if f.YidunRequestId != nil {
		result["yidunRequestId"] = *f.YidunRequestId
	}

	return result
}
