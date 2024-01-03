package mediaiface

import (
	media "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/media/v2"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/media/v2/callback"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/media/v2/query"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/media/v2/submit"
)

type MediaSolutionAPI interface {
	// 融媒体解决方案检测提交
	Submit(req *submit.MediaSubmitRequestV2) (res *submit.MediaSubmitResponseV2, err error)
	// 融媒体解决方案回调（轮询模式）
	Callback(req *callback.MediaCallBackRequestV2) (res *callback.MediaCallbackResponseV2, err error)

	// 融媒体解决方案查询
	Query(req *query.MediaQueryRequestV2) (res *callback.MediaCallbackResponseV2, err error)
}

var _ MediaSolutionAPI = (*media.MediaClient)(nil)
