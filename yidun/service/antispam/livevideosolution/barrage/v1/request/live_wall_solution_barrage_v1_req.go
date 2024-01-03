package request

import (
	"encoding/json"
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
)

type LiveWallSolutionBarrageV1Req struct {
	*types.PostFormRequest
	Barrages []*LiveWallSolutionBarrage `json:"barrages,omitempty"`
}

// NewLiveWallSolutionBarrageV1Req 是一个创建 LiveWallSolutionBarrageV1Req 的函数
func NewLiveWallSolutionBarrageV1Req() *LiveWallSolutionBarrageV1Req {
	var request = &LiveWallSolutionBarrageV1Req{
		PostFormRequest: types.NewPostFormRequest(),
	}
	request.SetProductCode("liveVideoSolutionCommon")
	request.SetUriPattern("/v1/livewallsolution/barrage/push")
	request.SetVersion("v1")
	return request
}

// SetBarrages 方法用于设置 Barrages 属性
func (r *LiveWallSolutionBarrageV1Req) SetBarrages(barrages []*LiveWallSolutionBarrage) {
	r.Barrages = barrages
}

// GetCustomSignParams 方法用于获取自定义签名参数
func (r *LiveWallSolutionBarrageV1Req) GetBusinessCustomSignParams() map[string]string {
	params := r.PostFormRequest.GetBusinessCustomSignParams()
	if r.Barrages != nil {
		barragesBytes, _ := json.Marshal(r.Barrages)
		params["barrages"] = string(barragesBytes)
	}
	return params
}
