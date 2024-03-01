package single

import (
	"strconv"

	"github.com/yidun/yidun-golang-sdk/yidun/core/validation"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/text/v5/check/sync/single"
)
// 文本异步检测request
type TextAsyncCheckRequest struct {
    *single.TextCheckSceneRequest
    CheckLabels *string `json:"checkLabels,omitempty" validate:"max=512"` // 业务过检分类,如果没有勾选分类提交返回参数错误，多个垃圾类别以逗号分隔（"100,200"）
	CheckStrategyGroupIds *string                     `json:"checkStrategyGroupIds,omitempty" validate:"max=512"` // 业务指定过检策略组id,多个策略组id以逗号分隔（"xxx,xxx"），最多支持传20个
    Token       *string `json:"token,omitempty" validate:"max=256"`       // 内容安全与反作弊融合版专属字段，来自易盾反作弊SDK返回的token，接入SDK必传,接入流程请参考防刷版说明文档
}
// 构建request
func NewTextAsyncCheckRequest(businessId string) *TextAsyncCheckRequest {
	var checkSceneRequest = single.NewTextCheckSceneRequest(businessId)
	var request = &TextAsyncCheckRequest{
		TextCheckSceneRequest: checkSceneRequest,
	}
	checkSceneRequest.SetProductCode("text-check")
	checkSceneRequest.SetUriPattern("/v5/text/async-check")
	checkSceneRequest.SetVersion("v5.2")
	return request
}


// GetCustomSignParams 获取具体业务中特有的需要参与签名计算的参数
func (r *TextAsyncCheckRequest) GetBusinessCustomSignParams() map[string]string {
	params := r.TextCheckSceneRequest.GetBusinessCustomSignParams()
	if r.CheckLabels != nil {
		params["checkLabels"] = *r.CheckLabels
	}
	if r.Token != nil {
		params["token"] = *r.Token
	}
	return params
}

// check接口 参数校验方法,TODO: 业务参数校验
func (r *TextAsyncCheckRequest) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "TextAsyncCheckRequest"}
	if r == nil {
		invalidParams.Add(validation.NewErrParamRequired("TextAsyncCheckRequest"))
	}
	if (r.CheckLabels != nil && len(*r.CheckLabels) > 512) {
		invalidParams.Add(validation.NewErrParamMaxLen("CheckLabels", 512, strconv.Itoa(len(*r.CheckLabels))))
	}
	if (r.Token != nil && len(*r.Token) > 256) {
		invalidParams.Add(validation.NewErrParamMaxLen("Token", 512, strconv.Itoa(len(*r.Token))))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	err := r.TextCheckSceneRequest.ValidateParam()
	if (err != nil) {
		return err
	}
	return nil
}
//设置CheckLabels参数
func (t *TextAsyncCheckRequest) SetCheckLabels(checkLabels string) *TextAsyncCheckRequest {
    t.CheckLabels = &checkLabels
    return t
}
//设置CheckStrategyGroupIds参数
func (t *TextAsyncCheckRequest) SetCheckStrategyGroupIds(checkStrategyGroupIds string) *TextAsyncCheckRequest {
	t.CheckStrategyGroupIds = &checkStrategyGroupIds
	return t
}
//设置Token参数
func (t *TextAsyncCheckRequest) SetToken(token string) *TextAsyncCheckRequest {
    t.Token = &token
    return t
}