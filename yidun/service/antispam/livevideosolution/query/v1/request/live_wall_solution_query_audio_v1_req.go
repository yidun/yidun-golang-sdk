package request

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"strconv"
)

type LiveWallSolutionQueryAudioV1Req struct {
	*types.PostFormRequest
	// taskId
	TaskId *string `json:"taskId,omitempty"`
	// 起始时间
	StartTime *int64 `json:"startTime,omitempty"`
	// 结束时间
	EndTime *int64 `json:"endTime,omitempty"`
}

// NewLiveAudioCallbackV4Req 创建一个 LiveAudioCallbackV4Req 实例
func NewLiveWallSolutionQueryAudioV1Req() *LiveWallSolutionQueryAudioV1Req {
	var request = &LiveWallSolutionQueryAudioV1Req{
		PostFormRequest: types.NewPostFormRequest(),
	}
	request.SetProductCode("liveVideoSolutionCommon")
	request.SetUriPattern("/v1/livewallsolution/query/audio/task")
	request.SetVersion("v1")
	return request
}

// GetCustomSignParams 获取自定义签名参数
func (r *LiveWallSolutionQueryAudioV1Req) GetBusinessCustomSignParams() map[string]string {
	params := r.PostFormRequest.GetBusinessCustomSignParams()
	if r.TaskId != nil {
		params["taskId"] = *r.TaskId
	}
	if r.StartTime != nil {
		params["startTime"] = strconv.FormatInt(*r.StartTime, 10)
	}
	if r.EndTime != nil {
		params["endTime"] = strconv.FormatInt(*r.EndTime, 10)
	}
	return params
}

// SetTaskId sets the TaskId.
func (l *LiveWallSolutionQueryAudioV1Req) SetTaskId(taskId string) {
	l.TaskId = &taskId
}

// SetStartTime sets the StartTime.
func (l *LiveWallSolutionQueryAudioV1Req) SetStartTime(startTime int64) {
	l.StartTime = &startTime
}

// SetEndTime sets the EndTime.
func (l *LiveWallSolutionQueryAudioV1Req) SetEndTime(endTime int64) {
	l.EndTime = &endTime
}
