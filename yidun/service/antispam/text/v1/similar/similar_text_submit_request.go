package similar

import (
    "github.com/yidun/yidun-golang-sdk/yidun/core/types"
    "github.com/yidun/yidun-golang-sdk/yidun/core/validation"
)

// SimilarTextSubmitRequest 相似文本添加请求
type SimilarTextSubmitRequest struct {
    *types.BizPostFormRequest
    SimilarText *string `json:"similarText,omitempty"` // 文本数据，json格式对象数组
}

// SetSimilarText 设置相似文本数据
func (r *SimilarTextSubmitRequest) SetSimilarText(similarText string) {
    r.SimilarText = &similarText
}

// NewSimilarTextSubmitRequest 创建新的相似文本添加请求
func NewSimilarTextSubmitRequest(businessId string) *SimilarTextSubmitRequest {
    request := &SimilarTextSubmitRequest{
        BizPostFormRequest: types.NewBizPostFormRequest(businessId),
    }
    request.SetProductCode("text-api")
    request.SetUriPattern("/v1/similar-text/submit")
    request.SetVersion("v1")
    return request
}

// GetBusinessCustomSignParams 获取业务自定义参数
func (r *SimilarTextSubmitRequest) GetBusinessCustomSignParams() map[string]string {
    params := r.BizPostFormRequest.GetBusinessCustomSignParams()
    if r.SimilarText != nil {
        params["similarText"] = *r.SimilarText
    }
    return params
}

// ValidateParam 参数校验
func (r *SimilarTextSubmitRequest) ValidateParam() error {
    invalidParams := validation.ErrInvalidParams{Context: "SimilarTextSubmitRequest"}
    if r == nil {
        invalidParams.Add(validation.NewErrParamRequired("SimilarTextSubmitRequest"))
        return invalidParams
    }
    if r.SimilarText == nil || len(*r.SimilarText) == 0 {
        invalidParams.Add(validation.NewErrParamRequired("SimilarText"))
    }
    if invalidParams.Len() > 0 {
        return invalidParams
    }
    return nil
}