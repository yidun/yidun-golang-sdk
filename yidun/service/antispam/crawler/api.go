package crawler

import (
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/crawler/v1/delete"
	jobQuery "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/crawler/v1/query"
	jobSubmit "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/crawler/v1/submit"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/crawler/v1/update"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/crawler/v3/callback"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/crawler/v3/query"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/crawler/v3/submit"
)

// 网站解决方案-网页单URL检测提交
func (c *CrawlerClient) ResourceSubmit(req *submit.CrawlerResourceSubmitV3Request) (res *submit.CrawlerResourceSubmitV3Response, err error) {
	res = &submit.CrawlerResourceSubmitV3Response{}
	err = c.Client.DoExecute(req, res)
	return
}

// 网站解决方案-网页单URL机器检测结果查询
func (c *CrawlerClient) ResourceQuery(req *query.CrawlerResourceQueryV3Request) (res *callback.CrawlerResourceCallbackV3Response, err error) {
	res = &callback.CrawlerResourceCallbackV3Response{}
	err = c.Client.DoExecute(req, res)
	return
}

// 网站解决方案-网页单url检测结果获取（轮询模式）
func (c *CrawlerClient) ResourceCallback(req *callback.CrawlerResourceCallbackV3Request) (res *callback.CrawlerResourceCallbackV3Response, err error) {
	res = &callback.CrawlerResourceCallbackV3Response{}
	err = c.Client.DoExecute(req, res)
	return
}

// 网站解决方案-主站检测任务提交检测
func (c *CrawlerClient) JobSubmit(req *jobSubmit.CrawlerJobSubmitV1Request) (res *jobSubmit.CrawlerJobSubmitV1Response, err error) {
	res = &jobSubmit.CrawlerJobSubmitV1Response{}
	err = c.Client.DoExecute(req, res)
	return
}

// 网站解决方案-网站&公众号任务异常检测结果分页查询接口
func (c *CrawlerClient) JobQuery(req *jobQuery.CrawlerJobCallbackQueryRequest) (res *jobQuery.JobCallbackQueryResponse, err error) {
	res = &jobQuery.JobCallbackQueryResponse{}
	err = c.Client.DoExecute(req, res)
	return
}

// 网站解决方案-主站检测任务批量提交检测
func (c *CrawlerClient) CrawlerJobBatchSubmit(req *jobSubmit.CrawlerJobBatchSubmitV1Request) (res *jobSubmit.CrawlerJobBatchSubmitV1Response, err error) {
	res = &jobSubmit.CrawlerJobBatchSubmitV1Response{}
	err = c.Client.DoExecute(req, res)
	return
}

// 网站解决方案-公众号检测任务提交检测
func (c *CrawlerClient) OfficialAccountsSubmit(req *jobSubmit.OfficialAccountsSubmitV1Request) (res *jobSubmit.OfficialAccountsSubmitV1Response, err error) {
	res = &jobSubmit.OfficialAccountsSubmitV1Response{}
	err = c.Client.DoExecute(req, res)
	return
}

// 网站解决方案-公众号检测任务批量提交检测
func (c *CrawlerClient) OfficialAccountsBatchSubmit(req *jobSubmit.OfficialAccountsBatchSubmitV1Request) (res *jobSubmit.OfficialAccountsBatchSubmitV1Response, err error) {
	res = &jobSubmit.OfficialAccountsBatchSubmitV1Response{}
	err = c.Client.DoExecute(req, res)
	return
}

// 网站解决方案-微博检测任务提交检测
func (c *CrawlerClient) WeiboSubmit(req *jobSubmit.WeiboSubmitV1Request) (res *jobSubmit.WeiboSubmitV1Response, err error) {
	res = &jobSubmit.WeiboSubmitV1Response{}
	err = c.Client.DoExecute(req, res)
	return
}

// 网站解决方案-微博检测任务批量提交检测
func (c *CrawlerClient) WeiboBatchSubmit(req *jobSubmit.WeiboBatchSubmitV1Request) (res *jobSubmit.WeiboBatchSubmitV1Response, err error) {
	res = &jobSubmit.WeiboBatchSubmitV1Response{}
	err = c.Client.DoExecute(req, res)
	return
}

// 网站解决方案-主站检测任务批量查询
func (c *CrawlerClient) CrawlerJobBatchQuery(req *jobQuery.CrawlerJobBatchQueryV1Request) (res *jobQuery.CrawlerJobBatchQueryV1Response, err error) {
	res = &jobQuery.CrawlerJobBatchQueryV1Response{}
	err = c.Client.DoExecute(req, res)
	return
}

// 网站解决方案-公众号检测任务批量查询
func (c *CrawlerClient) OfficialAccountsBatchQuery(req *jobQuery.OfficialAccountsBatchQueryV1Request) (res *jobQuery.OfficialAccountsBatchQueryV1Response, err error) {
	res = &jobQuery.OfficialAccountsBatchQueryV1Response{}
	err = c.Client.DoExecute(req, res)
	return
}

// 网站解决方案-微博检测任务批量查询
func (c *CrawlerClient) WeiboBatchQuery(req *jobQuery.WeiboBatchQueryV1Request) (res *jobQuery.WeiboBatchQueryV1Response, err error) {
	res = &jobQuery.WeiboBatchQueryV1Response{}
	err = c.Client.DoExecute(req, res)
	return
}

// 网站解决方案-主站检测任务批量删除
func (c *CrawlerClient) CrawlerJobDelete(req *delete.CrawlerJobDeleteV1Request) (res *delete.CrawlerJobDeleteV1Response, err error) {
	res = &delete.CrawlerJobDeleteV1Response{}
	err = c.Client.DoExecute(req, res)
	return
}

// 网站解决方案-公众号检测任务批量删除
func (c *CrawlerClient) OfficialAccountsDelete(req *delete.OfficialAccountsDeleteV1Request) (res *delete.OfficialAccountsDeleteV1Response, err error) {
	res = &delete.OfficialAccountsDeleteV1Response{}
	err = c.Client.DoExecute(req, res)
	return
}

// 网站解决方案-微博检测任务批量删除
func (c *CrawlerClient) WeiboDelete(req *delete.WeiboDeleteV1Request) (res *delete.WeiboDeleteV1Response, err error) {
	res = &delete.WeiboDeleteV1Response{}
	err = c.Client.DoExecute(req, res)
	return
}

// 网站解决方案-网站检测任务更新
func (c *CrawlerClient) CrawlerJobUpdate(req *update.CrawlerJobUpdateV1Request) (res *update.CrawlerJobUpdateV1Response, err error) {
	res = &update.CrawlerJobUpdateV1Response{}
	err = c.Client.DoExecute(req, res)
	return
}

// 网站解决方案-网站检测任务更新
func (c *CrawlerClient) OfficialAccountsUpdate(req *update.OfficialAccountsUpdateV1Request) (res *update.OfficialAccountsUpdateV1Response, err error) {
	res = &update.OfficialAccountsUpdateV1Response{}
	err = c.Client.DoExecute(req, res)
	return
}

// 网站解决方案-网站检测任务更新
func (c *CrawlerClient) WeiboJobUpdate(req *update.WeiboUpdateV1Request) (res *update.WeiboUpdateV1Response, err error) {
	res = &update.WeiboUpdateV1Response{}
	err = c.Client.DoExecute(req, res)
	return
}
