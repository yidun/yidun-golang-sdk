package v1

import (
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/stream/v1/callback"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/stream/v1/push"
)

// AIGC流式检测解决方案检测提交
func (c *AigcStreamClient) Push(req *push.AigcStreamPushRequestV1) (res *push.AigcStreamPushResponseV1, err error) {
	res = &push.AigcStreamPushResponseV1{}
	err = c.Client.DoExecute(req, res)
	return
}

// AIGC流式检测解决方案回调（轮询模式）
func (c *AigcStreamClient) Callback(req *callback.AigcStreamCallBackRequestV1) (res *callback.AigcStreamCallbackResponseV1, err error) {
	res = &callback.AigcStreamCallbackResponseV1{}
	err = c.Client.DoExecute(req, res)
	return
}
