package v5

import (
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/text/query"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/text/v2/feedback"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/text/v5/callback"
    "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/text/v1/similar"
	textfeature "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/text/v1/textfeature"
	asyncBatch "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/text/v5/check/async/batch"
	async "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/text/v5/check/async/single"
	syncBatch "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/text/v5/check/sync/batch"
	sync "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/text/v5/check/sync/single"
)

// 文本同步检测
func (c *TextClient) SyncCheckText(req *sync.TextCheckRequest) (res *sync.TextCheckResponse, err error) {
	res = &sync.TextCheckResponse{}
	err = c.Client.DoExecute(req, res)
	return
}

// 文本批量同步检测
func (c *TextClient) SyncBatchCheckText(req *syncBatch.TextBatchCheckRequest) (res *syncBatch.TextBatchCheckResponse, err error) {
	res = &syncBatch.TextBatchCheckResponse{}
	err = c.Client.DoExecute(req, res)
	return
}

// 文本异步检测
func (c *TextClient) AsyncCheckText(req *async.TextAsyncCheckRequest) (res *async.TextAsyncCheckResponse, err error) {
	res = &async.TextAsyncCheckResponse{}
	err = c.Client.DoExecute(req, res)
	return
}

// 文本批量异步检测
func (c *TextClient) AsyncBatchCheckText(req *asyncBatch.TextAsyncBatchCheckRequest) (res *asyncBatch.TextAsyncBatchCheckResponse, err error) {
	res = &asyncBatch.TextAsyncBatchCheckResponse{}
	err = c.Client.DoExecute(req, res)
	return
}

// 文本检测结果获取（轮询模式）
func (c *TextClient) Callback(req *callback.TextCallBackRequest) (res *callback.TextCallBackResponse, err error) {
	res = &callback.TextCallBackResponse{}
	err = c.Client.DoExecute(req, res)
	return
}

// 文本结果反馈请求
func (c *TextClient) Feedback(req *feedback.TextFeedbackRequest) (res *feedback.TextFeedbackResponse, err error) {
	res = &feedback.TextFeedbackResponse{}
	err = c.Client.DoExecute(req, res)
	return
}

//根据taskIds查询检测结果
func (c *TextClient) QueryTaskIds(req *query.TextTaskIdsQueryRequest) (res *query.TextTaskIdsQueryResponse, err error) {
	res = &query.TextTaskIdsQueryResponse{}
	err = c.Client.DoExecute(req, res)
	return
}

// SimilarTextSubmit 相似文本添加
func (c *TextClient) SimilarTextSubmit(req *similar.SimilarTextSubmitRequest) (res *similar.SimilarTextSubmitResponse, err error) {
    res = &similar.SimilarTextSubmitResponse{}
    err = c.Client.DoExecute(req, res)
    return
}

// TextFeatureEditStatus 文本特征状态编辑
func (c *TextClient) TextFeatureEditStatus(req *textfeature.TextFeatureEditStatusRequest) (res *textfeature.TextFeatureEditStatusResponse, err error) {
    res = &textfeature.TextFeatureEditStatusResponse{}
    err = c.Client.DoExecute(req, res)
    return
}

// TextFeatureAdd 文本特征添加
func (c *TextClient) TextFeatureAdd(req *textfeature.TextFeatureAddRequest) (res *textfeature.TextFeatureAddResponse, err error) {
    res = &textfeature.TextFeatureAddResponse{}
    err = c.Client.DoExecute(req, res)
    return
}

// TextFeatureDelete 文本特征删除
func (c *TextClient) TextFeatureDelete(req *textfeature.TextFeatureDeleteRequest) (res *textfeature.TextFeatureDeleteResponse, err error) {
    res = &textfeature.TextFeatureDeleteResponse{}
    err = c.Client.DoExecute(req, res)
    return
}

// TextFeatureQuery 文本特征查询
func (c *TextClient) TextFeatureQuery(req *textfeature.TextFeatureQueryRequest) (res *textfeature.TextFeatureQueryResponse, err error) {
    res = &textfeature.TextFeatureQueryResponse{}
    err = c.Client.DoExecute(req, res)
    return
}
