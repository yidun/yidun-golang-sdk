package query

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
)

// CrawlerJobBatchQueryV1Response 网站任务检测批量查询接口响应对象v1.0
type CrawlerJobBatchQueryV1Response struct {
	*types.CommonResponse
	Result *[]CrawlerJobBatchQueryV1Result `json:"result,omitempty"`
}

// CrawlerJobBatchQueryV1Result represents the query result of a crawler job.
type CrawlerJobBatchQueryV1Result struct {
	// 本次查询的任务id
	JobId *int64 `json:"jobId,omitempty"`
	// 任务状态
	Status *int `json:"status,omitempty"`
}
