package batch

import (
	"encoding/json"
	"strconv"

	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/core/validation"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/text/v5/check/sync/single"
)
// 文本批量异步检测request
type TextBatchCheckRequest struct {
	*types.BizPostFormRequest
    CheckLabels *string                     `json:"checkLabels,omitempty"` // 业务过检分类,如果没有勾选分类提交返回参数错误，多个垃圾类别以逗号分隔（"100,200"）
	CheckStrategyGroupIds *string                     `json:"checkStrategyGroupIds,omitempty"` // 业务指定过检策略组id,多个策略组id以逗号分隔（"xxx,xxx"），最多支持传20个
    Texts       []*single.TextSceneRequest    `json:"texts,omitempty"`       // 1-100条文本数据
    Token       *string                     `json:"token,omitempty"`       // 内容安全与反作弊融合版专属字段，来自易盾反作弊SDK返回的token，接入SDK必传,接入流程请参考防刷版说明文档
}

// 构建request
func NewTextBatchCheckRequest(businessId string) *TextBatchCheckRequest {
	var request = &TextBatchCheckRequest{
		BizPostFormRequest: types.NewBizPostFormRequest(businessId),
	}
	request.SetProductCode("text-check")
	request.SetUriPattern("/v5/text/batch-check")
	request.SetVersion("v5.2")
	return request
}


// GetCustomSignParams 获取具体业务中特有的需要参与签名计算的参数
func (r *TextBatchCheckRequest) GetBusinessCustomSignParams() map[string]string {
	params := r.BizPostFormRequest.GetBusinessCustomSignParams()
	if r.CheckLabels != nil {
		params["checkLabels"] = *r.CheckLabels
	}
	if r.Token != nil {
		params["token"] = *r.Token
	}
	if (r.Texts != nil) {
		texts, _ := json.Marshal(r.Texts)
		params["texts"] = string(texts)
	}
	return params
}

// check接口 参数校验方法
func (r *TextBatchCheckRequest) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "TextBatchCheckRequest"}
	if r == nil {
		invalidParams.Add(validation.NewErrParamRequired("TextCheckRequest"))
	}
	if (r.CheckLabels != nil && len(*r.CheckLabels) > 512) {
		invalidParams.Add(validation.NewErrParamMaxLen("CheckLabels", 512, strconv.Itoa(len(*r.CheckLabels))))
	}
	if (r.Token != nil && len(*r.Token) > 256) {
		invalidParams.Add(validation.NewErrParamMaxLen("Token", 512, strconv.Itoa(len(*r.Token))))
	}
	if (r.Texts == nil || len(r.Texts) == 0) {
		invalidParams.Add(validation.NewErrParamMinLen("Texts", 1))
	}
	if (r.Texts != nil && len(r.Texts) > 100) {
		invalidParams.Add(validation.NewErrParamMaxLen("Texts", 100, strconv.Itoa(len(r.Texts))))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetCheckLabels sets the CheckLabels field value of the TextBatchCheckRequest struct.
func (t *TextBatchCheckRequest) SetCheckLabels(value string) *TextBatchCheckRequest {
    t.CheckLabels = &value
    return t
}

// SetCheckStrategyGroupIds sets the CheckStrategyGroupIds field value of the TextBatchCheckRequest struct.
func (t *TextBatchCheckRequest) SetCheckStrategyGroupIds(value string) *TextBatchCheckRequest {
	t.CheckStrategyGroupIds = &value
	return t
}

// SetTexts sets the Texts field value of the TextBatchCheckRequest struct.
func (t *TextBatchCheckRequest) SetTexts(value []*single.TextSceneRequest) *TextBatchCheckRequest {
    t.Texts = value
    return t
}

// SetToken sets the Token field value of the TextBatchCheckRequest struct.
func (t *TextBatchCheckRequest) SetToken(value string) *TextBatchCheckRequest {
    t.Token = &value
    return t
}

