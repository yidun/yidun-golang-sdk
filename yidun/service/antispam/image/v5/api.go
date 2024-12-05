package image

import (
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/image/v5/callback"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/image/v5/check"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/image/v5/check/async"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/image/v5/check/sync"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/image/v5/feedback"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/image/v5/query"
)

// 图片同步检测
func (c *ImageClient) ImageSyncCheck(req *check.ImageV5CheckRequest) (res *sync.ImageCheckResponse, err error) {
	// 图片类型只要有一张base64图片，则使用base64检测接口
	if req.Images != nil && len(*req.Images) > 0 {
		for _, image := range *req.Images {
			if *image.Type == 2 {
				req.UriPattern = "/v5/image/base64Check"
				break
			}
		}
	}

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

// 图片检测结果查询
func (c *ImageClient) ImageQuery(req *query.ImageQueryRequest) (res *query.ImageQueryResponse, err error) {
	res = &query.ImageQueryResponse{}
	err = c.Client.DoExecute(req, res)
	return
}

// 图片检测结果查询
func (c *ImageClient) ImageQueryV2(req *query.ImageQueryV2Request) (res *query.ImageQueryV2Response, err error) {
	res = &query.ImageQueryV2Response{}
	err = c.Client.DoExecute(req, res)
	return
}

// 图片检测结果获取（轮询模式）
func (c *ImageClient) ImageCallback(req *callback.ImageV5CallbackRequest) (res *sync.ImageCheckResponse, err error) {
	res = &sync.ImageCheckResponse{}
	err = c.Client.DoExecute(req, res)
	return
}
