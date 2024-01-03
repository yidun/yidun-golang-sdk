package audioiface

import (
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/liveaudio"
	barrageResquest "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/liveaudio/barrage/v1/request"
	barrageResponse "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/liveaudio/barrage/v1/response"
	callbackResquest "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/liveaudio/callback/v4/request"
	callbackResponse "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/liveaudio/callback/v4/response"
	feedbackResquest "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/liveaudio/feedback/v1/request"
	feedbackResponse "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/liveaudio/feedback/v1/response"
	queryRequest "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/liveaudio/query/v1/request"
	queryResponse "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/liveaudio/query/v1/response"
	submitResquest "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/liveaudio/submit/v4/request"
	submitResponse "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/liveaudio/submit/v4/response"
)

type LiveAudioClientAPI interface {
	Check(req *submitResquest.LiveAudioSubmitV4Req) (res *submitResponse.LiveAudioSubmitV4Resp, err error)
	Callback(req *callbackResquest.LiveAudioCallbackV4Req) (res *callbackResponse.LiveAudioCallbackV4Resp, err error)
	PushBarrage(req *barrageResquest.LiveAudioBarrageV1Req) (res *barrageResponse.LiveAudioBarrageV1Resp, err error)
	Feedback(req *feedbackResquest.LiveAudioFeedbackV1Req) (res *feedbackResponse.LiveAudioFeedbackV1Resp, err error)
	QueryTaskID(req *queryRequest.LiveAudioQueryTaskIdV1Req) (res *queryResponse.LiveAudioQueryTaskIdV1Resp, err error)
	QueryExtra(req *queryRequest.LiveAudioQueryExtraV1Req) (res *queryResponse.LiveAudioQueryExtraV1Resp, err error)
	QueryMonitor(req *queryRequest.LiveAudioQueryMonitorV1Req) (res *queryResponse.LiveAudioQueryMonitorV1Resp, err error)
	QuerySegment(req *queryRequest.LiveAudioQuerySegmentV1Req) (res *queryResponse.LiveAudioQuerySegmentV1Resp, err error)
}

var _ LiveAudioClientAPI = (*liveaudio.LiveAudioClient)(nil)
