package imageiface

import (
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/image/v5"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/image/v5/callback"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/image/v5/check"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/image/v5/check/async"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/image/v5/check/sync"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/image/v5/feedback"
)

type ImageCheckAPI interface {
	// 图片同步检测
	ImageSyncCheck(req *check.ImageV5CheckRequest) (res *sync.ImageCheckResponse, err error)

	// 图片异步检测
	ImageAsyncCheck(req *async.ImageV5AsyncCheckRequest) (res *async.ImageV5AsyncCheckResp, err error)

	// 图片结果反馈请求
	ImageFeedback(req *feedback.ImageV5FeedbackRequest) (res *feedback.ImageV5FeedbackResponse, err error)

	// 图片检测结果获取（轮询模式）
	ImageCallback(req *callback.ImageV5CallbackRequest) (res *sync.ImageCheckResponse, err error)
}

var _ ImageCheckAPI = (*image.ImageClient)(nil)
