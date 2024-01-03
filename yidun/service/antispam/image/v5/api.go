package image

import (
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/image/v5/callback"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/image/v5/check"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/image/v5/check/async"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/image/v5/check/sync"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/image/v5/feedback"
)

// 图片同步检测
func (c *ImageClient) ImageSyncCheck(req *check.ImageV5CheckRequest) (res *sync.ImageCheckResponse, err error) {
	res = &sync.ImageCheckResponse{}
	err = c.Client.DoExecute(req, res)
	return
}

// 图片异步检测
func (c *ImageClient) ImageAsyncCheck(req *async.ImageV5AsyncCheckRequest) (res *async.ImageV5AsyncCheckResp, err error) {
	res = &async.ImageV5AsyncCheckResp{}
	err = c.Client.DoExecute(req, res)
	return
}

// 图片结果反馈请求
func (c *ImageClient) ImageFeedback(req *feedback.ImageV5FeedbackRequest) (res *feedback.ImageV5FeedbackResponse, err error) {
	res = &feedback.ImageV5FeedbackResponse{}
	err = c.Client.DoExecute(req, res)
	return
}

// 图片检测结果获取（轮询模式）
func (c *ImageClient) ImageCallback(req *callback.ImageV5CallbackRequest) (res *sync.ImageCheckResponse, err error) {
	res = &sync.ImageCheckResponse{}
	err = c.Client.DoExecute(req, res)
	return
}
