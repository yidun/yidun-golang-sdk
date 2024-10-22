package query

import (
	"encoding/json"

	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/core/validation"
)

// 图片查询请求
type ImageQueryV2Request struct {
	*types.BizPostFormRequest
	// 图片taskid列表 json字符串
	Taskids *[]string `json:"taskIds,omitempty"`
}

func (r *ImageQueryV2Request) SetTaskIds(taskids []string) {
	r.Taskids = &taskids
}

// 构建request
func NewImageQueryV2Request(businessId string) *ImageQueryV2Request {
	var request = &ImageQueryV2Request{
		BizPostFormRequest: types.NewBizPostFormRequest(businessId),
	}
	request.SetProductCode("imageCommon")
	request.SetUriPattern("/v2/image/query/task")
	request.SetVersion("v1")
	return request
}

// 获取业务自定义参数
func (r *ImageQueryV2Request) GetBusinessCustomSignParams() map[string]string {
	params := r.BizPostFormRequest.GetBusinessCustomSignParams()
	if r.Taskids != nil {
		taskIdsJson, _ := json.Marshal(r.Taskids)
		params["taskIds"] = string(taskIdsJson)
	}

	return params
}

// 参数校验
func (r *ImageQueryV2Request) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "ImageQueryV2Request"}
	if r == nil {
		invalidParams.Add(validation.NewErrParamRequired("ImageQueryV2Request"))
		return invalidParams
	}
	if r.Taskids == nil || len(*r.Taskids) == 0 {
		invalidParams.Add(validation.NewErrParamRequired("Taskids"))
		return invalidParams
	}
	return nil
}
