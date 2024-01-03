package request

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
)

// LiveWallSolutionQueryMonitorV1Req represents a request to LiveWallSolutionQueryMonitorV1 API.
type LiveWallSolutionQueryMonitorV1Req struct {
	*types.PostFormRequest
	// The version, should be "v1"
	TaskId *string `json:"taskId,omitempty"` // taskId
}

// SetTaskId sets the TaskId.
func (l *LiveWallSolutionQueryMonitorV1Req) SetTaskId(taskId string) {
	l.TaskId = &taskId
}

// GetCustomSignParams 获取自定义签名参数
func (l *LiveWallSolutionQueryMonitorV1Req) GetBusinessCustomSignParams() map[string]string {
	result := l.PostFormRequest.GetBusinessCustomSignParams()
	if l.TaskId != nil {
		result["taskId"] = *l.TaskId
	}
	return result
}

// NewLiveAudioCallbackV4Req 创建一个 LiveAudioCallbackV4Req 实例
func NewLiveWallSolutionQueryMonitorV1Req() *LiveWallSolutionQueryMonitorV1Req {
	var request = &LiveWallSolutionQueryMonitorV1Req{
		PostFormRequest: types.NewPostFormRequest(),
	}
	request.SetProductCode("liveVideoSolutionCommon")
	request.SetUriPattern("/v1/livewallsolution/query/monitor")
	request.SetVersion("v1")
	return request
}
