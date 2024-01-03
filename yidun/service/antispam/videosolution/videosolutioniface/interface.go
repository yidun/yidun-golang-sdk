package videosolutioniface

import (
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/videosolution"
	request2 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/videosolution/callback/v2/request"
	response2 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/videosolution/callback/v2/response"
	request3 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/videosolution/feedback/v1/request"
	response3 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/videosolution/feedback/v1/response"
	request4 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/videosolution/query/v1/request"
	request6 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/videosolution/query/v1/response"
	request5 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/videosolution/query/v2/request"
	request7 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/videosolution/query/v2/response"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/videosolution/submit/v2/request"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/videosolution/submit/v2/response"
)

type VideoSolutionClientAPI interface {
	// Submit 点播音视频解决方案提交接口
	Submit(req *request.VideoSolutionSubmitV2Request) (res *response.VideoSolutionSubmitV2Response, err error)
	// Callback 点播音视频解决方案回调接口
	Callback(req *request2.VideoSolutionCallbackV2Request) (res *response2.VideoSolutionCallbackV2Response, err error)
	// Feedback 点播音视频解决方案反馈接口
	Feedback(req *request3.VideoSolutionFeedbackV1Request) (res *response3.VideoSolutionFeedbackV1Response, err error)
	// QueryImageV1 点播音视频解决方案查询截图接口
	QueryImageV1(req *request4.VideoSolutionQueryImageV1Request) (res *request6.VideoSolutionQueryImageV1Response, err error)
	// QueryTaskV1 点播音视频解决方案查询任务接口 v1
	QueryTaskV1(req *request4.VideoSolutionQueryTaskV1Request) (res *request6.VideoSolutionQueryTaskV1Response, err error)
	// QueryTaskV2  点播音视频解决方案查询任务接口 v2
	QueryTaskV2(req *request5.VideoSolutionQueryTaskV2Request) (res *request7.VideoSolutionQueryTaskV2Response, err error)
}

var _ VideoSolutionClientAPI = (*videosolution.VideoSolutionClient)(nil)
