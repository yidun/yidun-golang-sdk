package request

import (
	"encoding/json"
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/core/validation"
)

// VideoSolutionQueryTaskV2Request 视频解决方案查询任务请求
type VideoSolutionQueryTaskV2Request struct {
	*types.PostFormRequest
	TaskIds *[]string `json:"taskIds,omitempty"` // 任务id列表，最多100个
	DataIds *[]string `json:"dataIds,omitempty"` // 数据id列表，
}

func (r *VideoSolutionQueryTaskV2Request) SetTaskIds(taskIds []string) {
	r.TaskIds = &taskIds
}
func (r *VideoSolutionQueryTaskV2Request) SetDataIds(dataIds []string) {
	r.DataIds = &dataIds
}

func NewVideoSolutionQueryTaskV2Request() *VideoSolutionQueryTaskV2Request {
	var request = &VideoSolutionQueryTaskV2Request{
		PostFormRequest: types.NewPostFormRequest(),
	}
	request.SetProductCode("videoSolutionCommon")
	request.SetUriPattern("/v2/videosolution/query/task")
	request.SetVersion("v2")
	return request
}
func (r *VideoSolutionQueryTaskV2Request) GetBusinessCustomSignParams() map[string]string {
	params := r.PostFormRequest.GetBusinessCustomSignParams()
	if r.TaskIds != nil {
		taskIds, _ := json.Marshal(r.TaskIds)
		params["taskIds"] = string(taskIds)
	}
	if r.DataIds != nil {
		dataIds, _ := json.Marshal(r.DataIds)
		params["dataIds"] = string(dataIds)
	}
	return params
}

func (r *VideoSolutionQueryTaskV2Request) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "VideoSolutionQueryTaskV2Request"}
	if r == nil {
		invalidParams.Add(validation.NewErrParamRequired("VideoSolutionQueryTaskV2Request"))
	}
	if r.TaskIds == nil || len(*r.TaskIds) == 0 || len(*r.TaskIds) > 100 {
		invalidParams.Add(validation.NewErrParamRequired("TaskIds"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}
