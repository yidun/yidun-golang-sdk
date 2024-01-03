package request

import (
	"encoding/json"
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
)

// LiveAudioBarrageV1Req 结构体定义
type LiveAudioBarrageV1Req struct {
	*types.BizPostFormRequest
	Barrages []*LiveAudioBarrage `json:"barrages,omitempty"`
}

// NewLiveAudioBarrageV1Req 是一个创建 LiveAudioBarrageV1Req 的函数
func NewLiveAudioBarrageV1Req(businessId string) *LiveAudioBarrageV1Req {
	var request = &LiveAudioBarrageV1Req{
		BizPostFormRequest: types.NewBizPostFormRequest(businessId),
	}
	request.SetProductCode("liveAudio")
	request.SetUriPattern("/v1/liveaudio/barrage/push")
	request.SetVersion("v1")
	return request
}

// SetBarrages 方法用于设置 Barrages 属性
func (r *LiveAudioBarrageV1Req) SetBarrages(barrages []*LiveAudioBarrage) {
	r.Barrages = barrages
}

// GetCustomSignParams 方法用于获取自定义签名参数
func (r *LiveAudioBarrageV1Req) GetBusinessCustomSignParams() map[string]string {
	params := r.BizPostFormRequest.GetBusinessCustomSignParams()
	if r.Barrages != nil {
		barragesBytes, _ := json.Marshal(r.Barrages)
		params["barrages"] = string(barragesBytes)
	}
	return params
}
