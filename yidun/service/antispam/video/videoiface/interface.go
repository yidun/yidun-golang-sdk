package videoiface

import (
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/video"
	request2 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/video/callback/v4/request"
	response2 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/video/callback/v4/response"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/video/check/v4/request"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/video/check/v4/response"
	request3 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/video/query/v4/request"
	response3 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/video/query/v4/response"
)

type VideoClientAPI interface {
	Check(req *request.VideoCheckRequest) (res *response.VideoCheckResponse, err error)
	Callback(req *request2.VideoCallbackV4Request) (res *response2.VideoCallbackV4Response, err error)
	Query(req *request3.VideoQueryTaskRequest) (res *response3.VideoQueryTaskV4Response, err error)
}

var _ VideoClientAPI = (*video.VideoClient)(nil)
