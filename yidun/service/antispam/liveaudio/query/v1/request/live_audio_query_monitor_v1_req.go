package request

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
)

type LiveAudioQueryMonitorV1Req struct {
	*types.BizPostFormRequest
	TaskId *string `json:"taskId,omitempty"`
}

func NewLiveAudioQueryMonitorV1Req(businessId string) *LiveAudioQueryMonitorV1Req {
	var request = &LiveAudioQueryMonitorV1Req{
		BizPostFormRequest: types.NewBizPostFormRequest(businessId),
	}
	request.SetProductCode("liveAudio")
	request.SetUriPattern("/v1/liveaudio/query/monitor")
	request.SetVersion("v1")
	return request
}

// 定义GetCustomSignParams() 方法的功能
func (req *LiveAudioQueryMonitorV1Req) GetBusinessCustomSignParams() map[string]string {
	params := req.BizPostFormRequest.GetBusinessCustomSignParams()
	if req.TaskId != nil {
		params["taskId"] = *req.TaskId
	}
	return params
}

func (req *LiveAudioQueryMonitorV1Req) SetTaskId(taskId string) {
	req.TaskId = &taskId
}
