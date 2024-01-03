package videosolution

import (
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

func (c *VideoSolutionClient) Submit(req *request.VideoSolutionSubmitV2Request) (res *response.VideoSolutionSubmitV2Response, err error) {
	res = &response.VideoSolutionSubmitV2Response{}
	err = c.Client.DoExecute(req, res)
	return
}
func (c *VideoSolutionClient) Callback(req *request2.VideoSolutionCallbackV2Request) (res *response2.VideoSolutionCallbackV2Response, err error) {
	res = &response2.VideoSolutionCallbackV2Response{}
	err = c.Client.DoExecute(req, res)
	return
}
func (c *VideoSolutionClient) Feedback(req *request3.VideoSolutionFeedbackV1Request) (res *response3.VideoSolutionFeedbackV1Response, err error) {
	res = &response3.VideoSolutionFeedbackV1Response{}
	err = c.Client.DoExecute(req, res)
	return
}
func (c *VideoSolutionClient) QueryImageV1(req *request4.VideoSolutionQueryImageV1Request) (res *request6.VideoSolutionQueryImageV1Response, err error) {
	res = &request6.VideoSolutionQueryImageV1Response{}
	err = c.Client.DoExecute(req, res)
	return

}
func (c *VideoSolutionClient) QueryTaskV1(req *request4.VideoSolutionQueryTaskV1Request) (res *request6.VideoSolutionQueryTaskV1Response, err error) {

	res = &request6.VideoSolutionQueryTaskV1Response{}
	err = c.Client.DoExecute(req, res)
	return
}
func (c *VideoSolutionClient) QueryTaskV2(req *request5.VideoSolutionQueryTaskV2Request) (res *request7.VideoSolutionQueryTaskV2Response, err error) {

	res = &request7.VideoSolutionQueryTaskV2Response{}
	err = c.Client.DoExecute(req, res)
	return
}
