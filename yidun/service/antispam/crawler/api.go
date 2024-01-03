package crawler

import (
	jobQuery "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/crawler/v1/query"
	jobSubmit "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/crawler/v1/submit"
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