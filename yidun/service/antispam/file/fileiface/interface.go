package fileiface

import (
	file "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/file/v2"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/file/v2/callback"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/file/v2/submit"
)

type FileSolutionAPI interface {
	// 文档解决方案检测提交
	Submit(req *submit.FileSubmitV2Request) (res *submit.FileSubmitV2Response, err error)
	// 文档解决方案回调（轮询模式）
	Callback(req *callback.FileCallBackV2Request) (res *callback.FileCallbackV2Response, err error)
}

var _ FileSolutionAPI = (*file.FileClient)(nil)
