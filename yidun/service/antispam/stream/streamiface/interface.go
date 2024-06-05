package streamiface

import (
	stream "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/stream/v1"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/stream/v1/callback"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/stream/v1/push"
)

type AigcStreamSolutionAPI interface {
	// AIGC流式检测解决方案检测提交
	Push(req *push.AigcStreamPushRequestV1) (res *push.AigcStreamPushResponseV1, err error)
	// AIGC流式检测解决方案回调（轮询模式）
	Callback(req *callback.AigcStreamCallBackRequestV1) (res *callback.AigcStreamCallbackResponseV1, err error)

	// AIGC流式检测解决方案查询
	Query(req *query.AigcStreamQueryRequestV1) (res *callback.AigcStreamCallbackResponseV1, err error)
}

var _ AigcStreamSolutionAPI = (*stream.AigcStreamClient)(nil)
