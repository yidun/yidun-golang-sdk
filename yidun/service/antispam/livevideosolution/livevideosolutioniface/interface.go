package livevideosolutioniface

import (
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/livevideosolution"
	barrageResquest "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/livevideosolution/barrage/v1/request"
	barrageResponse "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/livevideosolution/barrage/v1/response"
	callbackResquest "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/livevideosolution/callback/v3/request"
	callbackResponse "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/livevideosolution/callback/v3/response"
	feedbackResquest "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/livevideosolution/feedback/v1/request"
	feedbackResponse "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/livevideosolution/feedback/v1/response"
	queryRequest "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/livevideosolution/query/v1/request"
	queryResponse "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/livevideosolution/query/v1/response"
	submitResquest "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/livevideosolution/submit/v3/request"
	submitResponse "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/livevideosolution/submit/v3/response"
)

type LiveVideoSolutionClientAPI interface {
	Check(req *submitResquest.LiveWallSolutionSubmitV3Req) (res *submitResponse.LiveWallSolutionSubmitV3Resp, err error)
	PushBarrage(req *barrageResquest.LiveWallSolutionBarrageV1Req) (res *barrageResponse.LiveWallSolutionBarrageV1Resp, err error)
	Callback(req *callbackResquest.LiveWallSolutionCallbackV3Req) (res *callbackResponse.LiveWallSolutionCallbackV3Resp, err error)
	Feedback(req *feedbackResquest.LiveWallSolutionFeedbackV1Req) (res *feedbackResponse.LiveWallSolutionFeedbackV1Resp, err error)
	QueryAudio(req *queryRequest.LiveWallSolutionQueryAudioV1Req) (res *queryResponse.LiveWallSolutionQueryAudioV1Resp, err error)
	QueryImage(req *queryRequest.LiveWallSolutionQueryImageV1Req) (res *queryResponse.LiveWallSolutionQueryImageV1Resp, err error)
	QueryMonitor(req *queryRequest.LiveWallSolutionQueryMonitorV1Req) (res *queryResponse.LiveWallSolutionQueryMonitorV1Resp, err error)
	Query(req *queryRequest.LiveWallSolutionQueryV1Req) (res *queryResponse.LiveWallSolutionQueryV1Resp, err error)
}

var _ LiveVideoSolutionClientAPI = (*livevideosolution.LiveVideoSolutionClient)(nil)
