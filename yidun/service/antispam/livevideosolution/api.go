package livevideosolution

import (
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

func (c *LiveVideoSolutionClient) Check(req *submitResquest.LiveWallSolutionSubmitV3Req) (res *submitResponse.LiveWallSolutionSubmitV3Resp, err error) {
	res = &submitResponse.LiveWallSolutionSubmitV3Resp{}
	err = c.Client.DoExecute(req, res)
	return
}

func (c *LiveVideoSolutionClient) PushBarrage(req *barrageResquest.LiveWallSolutionBarrageV1Req) (res *barrageResponse.LiveWallSolutionBarrageV1Resp, err error) {
	res = &barrageResponse.LiveWallSolutionBarrageV1Resp{}
	err = c.Client.DoExecute(req, res)
	return
}

func (c *LiveVideoSolutionClient) Callback(req *callbackResquest.LiveWallSolutionCallbackV3Req) (res *callbackResponse.LiveWallSolutionCallbackV3Resp, err error) {
	res = &callbackResponse.LiveWallSolutionCallbackV3Resp{}
	err = c.Client.DoExecute(req, res)
	return
}
func (c *LiveVideoSolutionClient) Feedback(req *feedbackResquest.LiveWallSolutionFeedbackV1Req) (res *feedbackResponse.LiveWallSolutionFeedbackV1Resp, err error) {
	res = &feedbackResponse.LiveWallSolutionFeedbackV1Resp{}
	err = c.Client.DoExecute(req, res)
	return
}

func (c *LiveVideoSolutionClient) QueryAudio(req *queryRequest.LiveWallSolutionQueryAudioV1Req) (res *queryResponse.LiveWallSolutionQueryAudioV1Resp, err error) {
	res = &queryResponse.LiveWallSolutionQueryAudioV1Resp{}
	err = c.Client.DoExecute(req, res)
	return
}

func (c *LiveVideoSolutionClient) QueryImage(req *queryRequest.LiveWallSolutionQueryImageV1Req) (res *queryResponse.LiveWallSolutionQueryImageV1Resp, err error) {
	res = &queryResponse.LiveWallSolutionQueryImageV1Resp{}
	err = c.Client.DoExecute(req, res)
	return
}

func (c *LiveVideoSolutionClient) QueryMonitor(req *queryRequest.LiveWallSolutionQueryMonitorV1Req) (res *queryResponse.LiveWallSolutionQueryMonitorV1Resp, err error) {
	res = &queryResponse.LiveWallSolutionQueryMonitorV1Resp{}
	err = c.Client.DoExecute(req, res)
	return
}

func (c *LiveVideoSolutionClient) Query(req *queryRequest.LiveWallSolutionQueryV1Req) (res *queryResponse.LiveWallSolutionQueryV1Resp, err error) {
	res = &queryResponse.LiveWallSolutionQueryV1Resp{}
	err = c.Client.DoExecute(req, res)
	return
}
