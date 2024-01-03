package crawleriface

import (
	crawler "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/crawler"
	jobQuery "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/crawler/v1/query"
	jobSubmit "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/crawler/v1/submit"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/crawler/v3/callback"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/crawler/v3/query"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/crawler/v3/submit"
)

type CrawlerSolutionAPI interface {
	// 网站解决方案-网页单URL检测提交
	ResourceSubmit(req *submit.CrawlerResourceSubmitV3Request) (res *submit.CrawlerResourceSubmitV3Response, err error)

	// 网站解决方案-网页单URL机器检测结果查询
	ResourceQuery(req *query.CrawlerResourceQueryV3Request) (res *callback.CrawlerResourceCallbackV3Response, err error)

	// 网站解决方案-网页单URL结果获取（轮询模式）
	ResourceCallback(req *callback.CrawlerResourceCallbackV3Request) (res *callback.CrawlerResourceCallbackV3Response, err error)

	// 网站解决方案-主站检测任务提交检测
	JobSubmit(req *jobSubmit.CrawlerJobSubmitV1Request) (res *jobSubmit.CrawlerJobSubmitV1Response, err error)

	// 网站解决方案-网站&公众号任务异常检测结果分页查询接口
	JobQuery(req *jobQuery.CrawlerJobCallbackQueryRequest) (res *jobQuery.JobCallbackQueryResponse, err error)
}

var _ CrawlerSolutionAPI = (*crawler.CrawlerClient)(nil)
