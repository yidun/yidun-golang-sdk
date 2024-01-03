package single

import (
	"strconv"

	"github.com/yidun/yidun-golang-sdk/yidun/core/validation"
)
/**
 * 文本单次同步检测请求
 */
 type TextCheckRequest struct {
	*TextCheckSceneRequest
	CheckLabels *string `json:"checkLabels,omitempty"`//业务过检分类,如果没有勾选分类提交返回参数错误，多个垃圾类别以逗号分隔（"100,200"）
	Token *string `json:"token,omitempty"`//内容安全与反作弊融合版专属字段，来自易盾反作弊SDK返回的token，接入SDK必传,接入流程请参考防刷版说明文档
}
// 构建request
func NewTextCheckRequest(businessId string) *TextCheckRequest {
	var checkSceneRequest = NewTextCheckSceneRequest(businessId)
	var request = &TextCheckRequest{
		TextCheckSceneRequest: checkSceneRequest,
	}
	checkSceneRequest.SetProductCode("text-check")
	checkSceneRequest.SetUriPattern("/v5/text/check")
	checkSceneRequest.SetVersion("v5.2")
	return request
}


// GetCustomSignParams 获取具体业务中特有的需要参与签名计算的参数
func (r *TextCheckRequest) GetBusinessCustomSignParams() map[string]string {
	params := r.TextCheckSceneRequest.GetBusinessCustomSignParams()
	if r.CheckLabels != nil {
		params["checkLabels"] = *r.CheckLabels
	}
	if r.Token != nil {
		params["token"] = *r.Token
	}
	return params
}

//参数校验方法
func (r *TextCheckRequest) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "TextCheckRequest"}
	if r == nil {
		invalidParams.Add(validation.NewErrParamRequired("TextCheckRequest"))
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
//设置TextCheckSceneRequest
func (t *TextCheckRequest) SetTextCheckSceneRequest(sceneRequest *TextCheckSceneRequest) {
    t.TextCheckSceneRequest = sceneRequest
}
//设置CheckLabels
func (t *TextCheckRequest) SetCheckLabels(checkLabels string) {
    t.CheckLabels = &checkLabels
}
//设置Token
func (t *TextCheckRequest) SetToken(token string) {
    t.Token = &token
}
