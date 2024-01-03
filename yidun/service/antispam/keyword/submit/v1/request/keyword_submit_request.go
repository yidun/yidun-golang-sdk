package request

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/core/validation"
	"strconv"
	"strings"
)

// KeywordSubmitRequest 提交关键词请求
type KeywordSubmitRequest struct {
	*types.BizPostFormRequest
	Category *int      `json:"category"` // 关键词返回的垃圾分类
	SubLabel *string   `json:"subLabel"` // 细分类
	Keywords *[]string `json:"keywords"` // 关键词检测列表
	Level    *int      `json:"level"`    // 关键词等级，1: 嫌疑，2: 不通过
	Type     *int      `json:"type"`     // 关键词类型，1: 普通词，2: 组合词，3: 拼音词

}

func NewKeywordSubmitRequest(businessId string) *KeywordSubmitRequest {
	var request = &KeywordSubmitRequest{
		BizPostFormRequest: types.NewBizPostFormRequest(businessId),
	}
	request.SetProductCode("keyword")
	request.SetUriPattern("/v1/keyword/submit")
	request.SetVersion("v1.1")
	return request
}
func (r *KeywordSubmitRequest) GetBusinessCustomSignParams() map[string]string {
	parentParams := r.BizPostFormRequest.GetBusinessCustomSignParams()
	if r.Category != nil {
		parentParams["category"] = strconv.Itoa(*r.Category)
	}
	if r.SubLabel != nil {
		parentParams["subLabel"] = *r.SubLabel
	}
	if r.Keywords != nil {
		// string join by ,
		parentParams["keywords"] = strings.Join(*r.Keywords, ",")
	}
	if r.Level != nil {
		parentParams["level"] = strconv.Itoa(*r.Level)
	}
	if r.Type != nil {
		parentParams["type"] = strconv.Itoa(*r.Type)
	}

	return parentParams
}

func (r *KeywordSubmitRequest) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "KeywordSubmitRequest"}
	if r == nil {
		invalidParams.Add(validation.NewErrParamRequired("KeywordSubmitRequest"))
	}
	if r.Keywords == nil || len(*r.Keywords) == 0 {
		invalidParams.Add(validation.NewErrParamRequired("Keywords"))
	}
	if len(*r.Keywords) > 200 || len(*r.Keywords) < 1 {
		invalidParams.Add(validation.NewErrParamBetweenRange("Keywords", 1, 200))
	}
	if r.Category == nil {
		invalidParams.Add(validation.NewErrParamRequired("Category"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (req *KeywordSubmitRequest) SetCategory(category int) {
	req.Category = &category
}

func (req *KeywordSubmitRequest) SetSubLabel(subLabel string) {
	req.SubLabel = &subLabel
}
func (req *KeywordSubmitRequest) SetKeywords(keywords []string) {
	req.Keywords = &keywords
}
func (req *KeywordSubmitRequest) SetLevel(level int) {
	req.Level = &level
}
func (req *KeywordSubmitRequest) SetType(type_ int) {
	req.Type = &type_
}
