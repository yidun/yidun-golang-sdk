package video

import (
	request2 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/video/callback/v4/request"
	response2 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/video/callback/v4/response"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/video/check/v4/request"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/video/check/v4/response"
	request3 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/video/query/v4/request"
	response3 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/video/query/v4/response"
)

func (c *VideoClient) Check(req *request.VideoCheckRequest) (res *response.VideoCheckResponse, err error) {
	res = &response.VideoCheckResponse{}
	err = c.Client.DoExecute(req, res)
	return
}

func (c *VideoClient) Callback(req *request2.VideoCallbackV4Request) (res *response2.VideoCallbackV4Response, err error) {
	res = &response2.VideoCallbackV4Response{}
	err = c.Client.DoExecute(req, res)
	return
}
func (c *VideoClient) Query(req *request3.VideoQueryTaskRequest) (res *response3.VideoQueryTaskV4Response, err error) {
	res = &response3.VideoQueryTaskV4Response{}
	err = c.Client.DoExecute(req, res)
	return
}
