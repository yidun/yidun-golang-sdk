package request

import (
	"encoding/json"
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/core/validation"
)

// VideoQueryTaskRequest 视频查询任务请求
type VideoQueryTaskRequest struct {
	*types.BizPostFormRequest
	TaskIds *[]string `json:"taskIds,omitempty"` // 任务id列表，逗号分隔，最多100个
}

func (r *VideoQueryTaskRequest) SetTaskIds(taskIds []string) {
	r.TaskIds = &taskIds
}

// NewVideoQueryTaskRequest 初始化VideoQueryTaskRequest对象
func NewVideoQueryTaskRequest(businessId string) *VideoQueryTaskRequest {
	var request = &VideoQueryTaskRequest{
		BizPostFormRequest: types.NewBizPostFormRequest(businessId),
	}
	request.SetProductCode("videoCommon")
	request.SetUriPattern("/v4/video/query/task")
	request.SetVersion("v4")
	return request
}
func (r *VideoQueryTaskRequest) GetBusinessCustomSignParams() map[string]string {
	params := r.BizPostFormRequest.GetBusinessCustomSignParams()
	if r.TaskIds != nil {
		taskIds, _ := json.Marshal(r.TaskIds)
		params["taskIds"] = string(taskIds)
	}

	return params
}

func (r *VideoQueryTaskRequest) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "VideoQueryTaskRequest"}
	if r == nil {
		invalidParams.Add(validation.NewErrParamRequired("VideoQueryTaskRequest"))
	}
	if r.TaskIds == nil || len(*r.TaskIds) == 0 || len(*r.TaskIds) > 100 {
		invalidParams.Add(validation.NewErrParamRequired("TaskIds"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}
