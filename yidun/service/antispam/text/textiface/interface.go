package textiface

import (
	feedback "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/text/v2/feedback"
	callback "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/text/v5/callback"
	asyncBatch "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/text/v5/check/async/batch"
	async "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/text/v5/check/async/single"
	syncBatch "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/text/v5/check/sync/batch"
	sync "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/text/v5/check/sync/single"
	v5 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/text"
)




type TextCheckAPI interface {

	//文本同步检测
	SyncCheckText(req *sync.TextCheckRequest) (res *sync.TextCheckResponse, err error)

	//文本批量同步检测
	SyncBatchCheckText(req *syncBatch.TextBatchCheckRequest) (res *syncBatch.TextBatchCheckResponse, err error)

	//文本异步检测
	AsyncCheckText(req *async.TextAsyncCheckRequest) (res *async.TextAsyncCheckResponse, err error)

	//文本批量异步检测
	AsyncBatchCheckText(req *asyncBatch.TextAsyncBatchCheckRequest) (res *asyncBatch.TextAsyncBatchCheckResponse, err error)

	//文本检测结果获取（轮询模式）
	Callback(req *callback.TextCallBackRequest) (res *callback.TextCallBackResponse, err error)

	//文本结果反馈请求
	Feedback(req *feedback.TextFeedbackRequest) (res *feedback.TextFeedbackResponse, err error)
}
var _ TextCheckAPI = (*v5.TextClient)(nil)