package callback

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
)

/**
 * 文档检测结果获取（轮询模式）
 */
type FileCallBackV2Request struct {
	*types.BizPostFormRequest
	YidunRequestId *string `json:"yidunRequestId,omitempty"` // 请求的唯一ID
}

func (r *FileCallBackV2Request) SetYidunRequestId(yidunRequestId string) {
	r.YidunRequestId = &yidunRequestId
}

func NewFileCallBackRequest() *FileCallBackV2Request {
	var request = &FileCallBackV2Request{
		BizPostFormRequest: types.NewBizPostFormRequestWithoutBizId(),
	}
	request.SetProductCode("file")
	request.SetUriPattern("/v2/file/callback/results")
	request.SetVersion("v2.0")
	return request
}

// GetCustomSignParams 获取具体业务中特有的需要参与签名计算的参数
func (f *FileCallBackV2Request) GetBusinessCustomSignParams() map[string]string {
	result := f.PostFormRequest.GetBusinessCustomSignParams()

	if f.YidunRequestId != nil {
		result["yidunRequestId"] = *f.YidunRequestId
	}

	return result
}
