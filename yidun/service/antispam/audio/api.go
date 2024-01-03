package audio

import (
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/audio/callback/v4/request"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/audio/callback/v4/response"
	request3 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/audio/check/async/v4/request"
	response3 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/audio/check/async/v4/response"
	request2 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/audio/check/sync/v2/request"
	response2 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/audio/check/sync/v2/response"
	request4 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/audio/feedback/v1/request"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/audio/feedback/v1/resopnse"
	request5 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/audio/query/v4/request"
	response4 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/audio/query/v4/response"
)

// 音频同步检测服务
func (c *AudioClient) SyncCheckAudio(req *request2.AudioSyncCheckRequest) (res *response2.AudioSyncCheckResponse, err error) {
	res = &response2.AudioSyncCheckResponse{}
	err = c.Client.DoExecute(req, res)
	return
}

// AsyncCheckAudio 异步检测
func (c *AudioClient) AsyncCheckAudio(req *request3.AudioAsyncCheckRequest) (res *response3.AudioAsyncCheckV4Response, err error) {
	res = &response3.AudioAsyncCheckV4Response{}
	err = c.Client.DoExecute(req, res)
	return
}

// Feedback 反馈
func (c *AudioClient) Feedback(req *request4.AudioFeedbackRequest) (res *resopnse.AudioFeedbackResponse, err error) {
	res = &resopnse.AudioFeedbackResponse{}
	err = c.Client.DoExecute(req, res)
	return
}

// Callback 回调
func (c *AudioClient) Callback(req *request.AudioCallBackRequest) (res *response.AudioCallBackV4Response, err error) {
	res = &response.AudioCallBackV4Response{}
	err = c.Client.DoExecute(req, res)
	return
}

// Query 查询
func (c *AudioClient) Query(req *request5.AudioQueryTaskRequest) (res *response4.AudioQueryTaskV4Response, err error) {
	res = &response4.AudioQueryTaskV4Response{}
	err = c.Client.DoExecute(req, res)
	return
}
