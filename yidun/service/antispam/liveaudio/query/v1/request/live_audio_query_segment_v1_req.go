package request

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"strconv"
)

type LiveAudioQuerySegmentV1Req struct {
	*types.BizPostFormRequest
	TaskId    *string `json:"taskId,omitempty"`
	StartTime *int64  `json:"startTime,omitempty"`
	EndTime   *int64  `json:"endTime,omitempty"`
}

func NewLiveAudioQuerySegmentV1Req(businessId string) *LiveAudioQuerySegmentV1Req {
	var request = &LiveAudioQuerySegmentV1Req{
		BizPostFormRequest: types.NewBizPostFormRequest(businessId),
	}
	request.SetProductCode("liveAudio")
	request.SetUriPattern("/v1/liveaudio/query/segment")
	request.SetVersion("v1")
	return request
}

// 定义 GetCustomSignParams() 方法的功能
func (req *LiveAudioQuerySegmentV1Req) GetBusinessCustomSignParams() map[string]string {
	params := req.BizPostFormRequest.GetBusinessCustomSignParams()
	if req.TaskId != nil {
		params["taskId"] = *req.TaskId
	}
	if req.StartTime != nil {
		params["startTime"] = strconv.FormatInt(*req.StartTime, 10)
	}
	if req.EndTime != nil {
		params["endTime"] = strconv.FormatInt(*req.EndTime, 10)
	}
	return params
}

// SetTaskId sets the task ID for the request.
// TaskId represents the ID of the task.
func (r *LiveAudioQuerySegmentV1Req) SetTaskId(taskId string) {
	r.TaskId = &taskId
}

// SetStartTime sets the start time for the request.
// StartTime represents the start time of the task.
func (r *LiveAudioQuerySegmentV1Req) SetStartTime(startTime int64) {
	r.StartTime = &startTime
}

// SetEndTime sets the end time for the request.
// EndTime represents the end time of the task.
func (r *LiveAudioQuerySegmentV1Req) SetEndTime(endTime int64) {
	r.EndTime = &endTime
}
