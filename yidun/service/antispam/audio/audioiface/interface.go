package audioiface

import (
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/audio"
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

// 音频检测服务
type AudioClientAPI interface {
	//音频同步检测
	SyncCheckAudio(req *request2.AudioSyncCheckRequest) (res *response2.AudioSyncCheckResponse, err error)

	AsyncCheckAudio(req *request3.AudioAsyncCheckRequest) (res *response3.AudioAsyncCheckV4Response, err error)

	Feedback(req *request4.AudioFeedbackRequest) (res *resopnse.AudioFeedbackResponse, err error)

	Callback(req *request.AudioCallBackRequest) (res *response.AudioCallBackV4Response, err error)

	Query(req *request5.AudioQueryTaskRequest) (res *response4.AudioQueryTaskV4Response, err error)
}

var _ AudioClientAPI = (*audio.AudioClient)(nil)
