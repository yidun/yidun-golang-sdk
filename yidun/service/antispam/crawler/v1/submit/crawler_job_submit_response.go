package submit

import "github.com/yidun/yidun-golang-sdk/yidun/core/types"

// 网站任务检测提交接口响应对象v1.0
type CrawlerJobSubmitV1Response struct {
	*types.CommonResponse
	Result *CrawlerJobSubmitResult `json:"result,omitempty"`
}

type CrawlerJobSubmitResult struct {
	JobId  *int64  `json:"jobId,omitempty"`
	DataId *string `json:"dataId,omitempty"`
}
