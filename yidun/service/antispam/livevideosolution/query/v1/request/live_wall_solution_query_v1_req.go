package request

import (
	"encoding/json"
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
)

// LiveWallSolutionQueryV1Req represents a request to LiveWallSolutionQueryV1 API.
type LiveWallSolutionQueryV1Req struct {
	*types.PostFormRequest
	TaskIds *[]string `json:"taskIds,omitempty"` // taskIds
}

// SetTaskIds sets the TaskIds.
func (l *LiveWallSolutionQueryV1Req) SetTaskIds(taskIds []string) {
	l.TaskIds = &taskIds
}

// GetCustomSignParams 获取自定义签名参数
func (l *LiveWallSolutionQueryV1Req) GetBusinessCustomSignParams() map[string]string {
	result := l.PostFormRequest.GetBusinessCustomSignParams()
	if l.TaskIds != nil {
		taskIdsJson, _ := json.Marshal(*l.TaskIds)
		result["taskIds"] = string(taskIdsJson)
	}
	return result
}

// NewLiveAudioCallbackV4Req 创建一个 LiveAudioCallbackV4Req 实例
func NewLiveWallSolutionQueryV1Req() *LiveWallSolutionQueryV1Req {
	var request = &LiveWallSolutionQueryV1Req{
		PostFormRequest: types.NewPostFormRequest(),
	}
	request.SetProductCode("liveVideoSolutionCommon")
	request.SetUriPattern("/v1/livewallsolution/query/task")
	request.SetVersion("v1")
	return request
}
