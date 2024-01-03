package liveaudio

import (
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

func (c *LiveAudioClient) Check(req *submitResquest.LiveAudioSubmitV4Req) (res *submitResponse.LiveAudioSubmitV4Resp, err error) {
	res = &submitResponse.LiveAudioSubmitV4Resp{}
	err = c.Client.DoExecute(req, res)
	return
}

func (c *LiveAudioClient) Callback(req *callbackResquest.LiveAudioCallbackV4Req) (res *callbackResponse.LiveAudioCallbackV4Resp, err error) {
	res = &callbackResponse.LiveAudioCallbackV4Resp{}
	err = c.Client.DoExecute(req, res)
	return
}

func (c *LiveAudioClient) PushBarrage(req *barrageResquest.LiveAudioBarrageV1Req) (res *barrageResponse.LiveAudioBarrageV1Resp, err error) {
	res = &barrageResponse.LiveAudioBarrageV1Resp{}
	err = c.Client.DoExecute(req, res)
	return
}

func (c *LiveAudioClient) Feedback(req *feedbackResquest.LiveAudioFeedbackV1Req) (res *feedbackResponse.LiveAudioFeedbackV1Resp, err error) {
	res = &feedbackResponse.LiveAudioFeedbackV1Resp{}
	err = c.Client.DoExecute(req, res)
	return
}

func (c *LiveAudioClient) QueryTaskID(req *queryRequest.LiveAudioQueryTaskIdV1Req) (res *queryResponse.LiveAudioQueryTaskIdV1Resp, err error) {
	res = &queryResponse.LiveAudioQueryTaskIdV1Resp{}
	err = c.Client.DoExecute(req, res)
	return
}

func (c *LiveAudioClient) QueryExtra(req *queryRequest.LiveAudioQueryExtraV1Req) (res *queryResponse.LiveAudioQueryExtraV1Resp, err error) {
	res = &queryResponse.LiveAudioQueryExtraV1Resp{}
	err = c.Client.DoExecute(req, res)
	return
}

func (c *LiveAudioClient) QueryMonitor(req *queryRequest.LiveAudioQueryMonitorV1Req) (res *queryResponse.LiveAudioQueryMonitorV1Resp, err error) {
	res = &queryResponse.LiveAudioQueryMonitorV1Resp{}
	err = c.Client.DoExecute(req, res)
	return
}

func (c *LiveAudioClient) QuerySegment(req *queryRequest.LiveAudioQuerySegmentV1Req) (res *queryResponse.LiveAudioQuerySegmentV1Resp, err error) {
	res = &queryResponse.LiveAudioQuerySegmentV1Resp{}
	err = c.Client.DoExecute(req, res)
	return
}
