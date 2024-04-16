package query

import (
	"encoding/json"

	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/core/validation"
)

// 图片查询请求
type ImageQueryRequest struct {
	*types.BizPostFormRequest
	// 图片taskid列表 json字符串
	Taskids *[]string `json:"taskIds,omitempty"`
}

func (r *ImageQueryRequest) SetTaskIds(taskids []string) {
	r.Taskids = &taskids
}

// 构建request
func NewImageQueryRequest(businessId string) *ImageQueryRequest {
	var request = &ImageQueryRequest{
		BizPostFormRequest: types.NewBizPostFormRequest(businessId),
	}
	request.SetProductCode("imageCommon")
	request.SetUriPattern("/v1/image/query/task")
	request.SetVersion("v1")
	return request
}

// 获取业务自定义参数
func (r *ImageQueryRequest) GetBusinessCustomSignParams() map[string]string {
	params := r.BizPostFormRequest.GetBusinessCustomSignParams()
	if r.Taskids != nil {
		taskIdsJson, _ := json.Marshal(r.Taskids)
		params["taskIds"] = string(taskIdsJson)
	}

	return params
}

// 参数校验
func (r *ImageQueryRequest) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "ImageQueryRequest"}
	if r == nil {
		invalidParams.Add(validation.NewErrParamRequired("ImageV5FeedBackRequest"))
		return invalidParams
	}
	if r.Taskids == nil || len(*r.Taskids) == 0 {
		invalidParams.Add(validation.NewErrParamRequired("Feedbacks"))
		return invalidParams
	}
	return nil
}
