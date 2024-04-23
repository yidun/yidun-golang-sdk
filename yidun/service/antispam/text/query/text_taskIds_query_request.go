package query

import (
	"encoding/json"

	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/core/validation"
)

type TextTaskIdsQueryRequest struct {
	*types.BizPostFormRequest
	TaskIds *[]string `json:"taskIds,omitempty"` // 任务ID json数组
}

func (r *TextTaskIdsQueryRequest) SetTaskIds(taskIds []string) {
	r.TaskIds = &taskIds

}

// 构建request
func NewTextTaskIdsQueryRequest(businessId string) *TextTaskIdsQueryRequest {
	var request = &TextTaskIdsQueryRequest{
		BizPostFormRequest: types.NewBizPostFormRequest(businessId),
	}
	request.SetProductCode("text-api")
	request.SetUriPattern("/v1/text/query/task")
	request.SetVersion("v1")
	return request
}
//获取业务自定义参数
func (r *TextTaskIdsQueryRequest) GetBusinessCustomSignParams() map[string]string {
	params := r.BizPostFormRequest.GetBusinessCustomSignParams()
	if r.TaskIds != nil {
		taskIds, err := json.Marshal(*r.TaskIds)
		if (err == nil) { 
			params["taskIds"] = string(taskIds)
		}
	}

	return params
}
//参数校验
func (r *TextTaskIdsQueryRequest) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "TextTaskIdsQueryRequest"}
	if r == nil {
		invalidParams.Add(validation.NewErrParamRequired("TextTaskIdsQueryRequest"))
		return invalidParams
	}
	if (r.TaskIds == nil) {
		invalidParams.Add(validation.NewErrParamRequired("TaskIds"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}