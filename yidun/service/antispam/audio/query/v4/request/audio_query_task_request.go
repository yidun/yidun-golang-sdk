package request

import (
	"encoding/json"
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/core/validation"
)

type AudioQueryTaskRequest struct {
	*types.BizPostFormRequest
	TaskIds *[]string `json:"taskIds,omitempty"` // 任务id列表，逗号分隔，最多100个
}

func (r *AudioQueryTaskRequest) SetTaskIds(taskIds []string) {
	r.TaskIds = &taskIds
}

// NewAudioQueryTaskRequest 初始化AudioQueryTaskRequest对象
func NewAudioQueryTaskRequest(businessId string) *AudioQueryTaskRequest {
	var request = &AudioQueryTaskRequest{
		BizPostFormRequest: types.NewBizPostFormRequest(businessId),
	}
	request.SetProductCode("audioCommon")
	request.SetUriPattern("/v4/audio/query/task")
	request.SetVersion("v4")
	return request
}
func (r *AudioQueryTaskRequest) GetBusinessCustomSignParams() map[string]string {
	params := r.BizPostFormRequest.GetBusinessCustomSignParams()
	if r.TaskIds != nil {
		taskIds, _ := json.Marshal(r.TaskIds)
		params["taskIds"] = string(taskIds)
	}

	return params
}

func (r *AudioQueryTaskRequest) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "AudioQueryTaskRequest"}
	if r == nil {
		invalidParams.Add(validation.NewErrParamRequired("AudioQueryTaskRequest"))
	}
	if r.TaskIds == nil || len(*r.TaskIds) == 0 || len(*r.TaskIds) > 100 {
		invalidParams.Add(validation.NewErrParamRequired("TaskIds"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}
