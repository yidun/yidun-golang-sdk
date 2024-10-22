package query

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/image/v5/check/sync"
)

type ImageQueryV2Response struct {
	*types.CommonResponse
	Result *[]sync.ImageV5Result `json:"result"`
}

// 获取图片查询结果
func (r *ImageQueryV2Response) GetResult() *[]sync.ImageV5Result {
	return r.Result
}

// 设置图片查询结果
func (r *ImageQueryV2Response) SetResult(results *[]sync.ImageV5Result) {
	r.Result = results
}
