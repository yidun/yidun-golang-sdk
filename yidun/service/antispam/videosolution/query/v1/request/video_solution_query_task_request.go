package request

import (
	"encoding/json"
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/core/validation"
)

// VideoSolutionQueryTaskV1Request 视频解决方案查询任务请求
type VideoSolutionQueryTaskV1Request struct {
	*types.PostFormRequest
	TaskIds *[]string `json:"taskIds,omitempty"` // 任务id列表，逗号分隔，最多20个
}

func (r *VideoSolutionQueryTaskV1Request) SetTaskIds(taskIds []string) {
	r.TaskIds = &taskIds
}

func NewVideoSolutionQueryTaskV1Request() *VideoSolutionQueryTaskV1Request {
	var request = &VideoSolutionQueryTaskV1Request{
		PostFormRequest: types.NewPostFormRequest(),
	}
	request.SetProductCode("videoSolutionCommon")
	request.SetUriPattern("/v1/videosolution/query/task")
	request.SetVersion("v1.1")
	return request
}
func (r *VideoSolutionQueryTaskV1Request) GetBusinessCustomSignParams() map[string]string {
	params := r.PostFormRequest.GetBusinessCustomSignParams()
	if r.TaskIds != nil {
		taskIds, _ := json.Marshal(r.TaskIds)
		params["taskIds"] = string(taskIds)
	}

	return params
}

func (r *VideoSolutionQueryTaskV1Request) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "VideoSolutionQueryTaskV1Request"}
	if r == nil {
		invalidParams.Add(validation.NewErrParamRequired("VideoSolutionQueryTaskV1Request"))
	}
	if r.TaskIds == nil || len(*r.TaskIds) == 0 || len(*r.TaskIds) > 20 {
		invalidParams.Add(validation.NewErrParamRequired("TaskIds"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}
