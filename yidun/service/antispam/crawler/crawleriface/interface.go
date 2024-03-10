package crawleriface

import (
	crawler "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/crawler"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/crawler/v1/delete"
	jobQuery "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/crawler/v1/query"
	jobSubmit "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/crawler/v1/submit"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/crawler/v1/update"
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

	// 网站解决方案-主站检测任务批量提交检测
	CrawlerJobBatchSubmit(req *jobSubmit.CrawlerJobBatchSubmitV1Request) (res *jobSubmit.CrawlerJobBatchSubmitV1Response, err error)

	// 网站解决方案-公众号检测任务提交检测
	OfficialAccountsSubmit(req *jobSubmit.OfficialAccountsSubmitV1Request) (res *jobSubmit.OfficialAccountsSubmitV1Response, err error)

	// 网站解决方案-公众号检测任务批量提交检测
	OfficialAccountsBatchSubmit(req *jobSubmit.OfficialAccountsBatchSubmitV1Request) (res *jobSubmit.OfficialAccountsBatchSubmitV1Response, err error)

	// 网站解决方案-微博检测任务提交检测
	WeiboSubmit(req *jobSubmit.WeiboSubmitV1Request) (res *jobSubmit.WeiboSubmitV1Response, err error)

	// 网站解决方案-微博检测任务批量提交检测
	WeiboBatchSubmit(req *jobSubmit.WeiboBatchSubmitV1Request) (res *jobSubmit.WeiboBatchSubmitV1Response, err error)

	// 网站解决方案-主站检测任务批量查询
	CrawlerJobBatchQuery(req *jobQuery.CrawlerJobBatchQueryV1Request) (res *jobQuery.CrawlerJobBatchQueryV1Response, err error)

	// 网站解决方案-公众号检测任务批量查询
	OfficialAccountsBatchQuery(req *jobQuery.OfficialAccountsBatchQueryV1Request) (res *jobQuery.OfficialAccountsBatchQueryV1Response, err error)

	// 网站解决方案-微博检测任务批量查询
	WeiboBatchQuery(req *jobQuery.WeiboBatchQueryV1Request) (res *jobQuery.WeiboBatchQueryV1Response, err error)

	// 网站解决方案-主站检测任务批量删除
	CrawlerJobDelete(req *delete.CrawlerJobDeleteV1Request) (res *delete.CrawlerJobDeleteV1Response, err error)

	// 网站解决方案-公众号检测任务批量删除
	OfficialAccountsDelete(req *delete.OfficialAccountsDeleteV1Request) (res *delete.OfficialAccountsDeleteV1Response, err error)

	// 网站解决方案-微博检测任务批量删除
	WeiboDelete(req *delete.WeiboDeleteV1Request) (res *delete.WeiboDeleteV1Response, err error)

	// 网站解决方案-网站检测任务更新
	CrawlerJobUpdate(req *update.CrawlerJobUpdateV1Request) (res *update.CrawlerJobUpdateV1Response, err error)

	// 网站解决方案-网站检测任务更新
	OfficialAccountsUpdate(req *update.OfficialAccountsUpdateV1Request) (res *update.OfficialAccountsUpdateV1Response, err error)

	// 网站解决方案-网站检测任务更新
	WeiboJobUpdate(req *update.WeiboUpdateV1Request) (res *update.WeiboUpdateV1Response, err error)
}

var _ CrawlerSolutionAPI = (*crawler.CrawlerClient)(nil)
