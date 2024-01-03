package request

import (
	"encoding/json"
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/core/validation"
	"strconv"
)

// KeywordUpdateRequest 更新关键词请求
type KeywordUpdateRequest struct {
	*types.BizPostFormRequest
	Category *int      `json:"category,omitempty"` // 关键词返回的垃圾分类
	SubLabel *string   `json:"subLabel,omitempty"` // 细分类
	Ids      *string   `json:"ids,omitempty"`      // 关键词唯一标识id列表，逗号分隔
	Keywords *[]string `json:"keywords,omitempty"` // 关键词检测列表
	Status   *int      `json:"status,omitempty"`   // 关键词状态，0: 未启用，1: 已启用
	Level    *int      `json:"level,omitempty"`    // 关键词等级，1: 嫌疑，2: 不通过
}

func (r *KeywordUpdateRequest) SetCategory(category int) {
	r.Category = &category
}
func (r *KeywordUpdateRequest) SetSubLabel(subLabel string) {
	r.SubLabel = &subLabel
}
func (r *KeywordUpdateRequest) SetIds(ids string) {
	r.Ids = &ids
}
func (r *KeywordUpdateRequest) SetKeywords(keywords []string) {
	r.Keywords = &keywords
}
func (r *KeywordUpdateRequest) SetStatus(status int) {
	r.Status = &status
}
func (r *KeywordUpdateRequest) SetLevel(level int) {
	r.Level = &level
}

func NewKeywordUpdateRequest(businessId string) *KeywordUpdateRequest {
	var request = &KeywordUpdateRequest{
		BizPostFormRequest: types.NewBizPostFormRequest(businessId),
	}
	request.SetProductCode("keyword")
	request.SetUriPattern("/v2/keyword/update")
	request.SetVersion("v2")
	return request
}
func (r *KeywordUpdateRequest) GetBusinessCustomSignParams() map[string]string {
	parentParams := r.BizPostFormRequest.GetBusinessCustomSignParams()
	if r.Category != nil {
		parentParams["category"] = strconv.Itoa(*r.Category)
	}
	if r.Status != nil {
		parentParams["status"] = strconv.Itoa(*r.Status)
	}
	if r.Ids != nil {
		parentParams["ids"] = *r.Ids
	}
	if r.SubLabel != nil {
		parentParams["subLabel"] = *r.SubLabel
	}
	if r.Keywords != nil {
		// convert to json
		keywords, _ := json.Marshal(r.Keywords)
		parentParams["keywords"] = string(keywords)
	}
	if r.Level != nil {
		parentParams["level"] = strconv.Itoa(*r.Level)
	}

	return parentParams
}

func (r *KeywordUpdateRequest) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "KeywordUpdateRequest"}
	if r == nil {
		invalidParams.Add(validation.NewErrParamRequired("KeywordUpdateRequest"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}
