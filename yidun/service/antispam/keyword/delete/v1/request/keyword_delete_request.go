package request

import (
	"encoding/json"
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/core/validation"
)

// KeywordDeleteRequest 删除关键词请求
type KeywordDeleteRequest struct {
	*types.BizPostFormRequest
	Ids      *string   `json:"ids,omitempty"`      // 关键词id
	Keywords *[]string `json:"keywords,omitempty"` // 关键词
}

func (r *KeywordDeleteRequest) SetIds(ids string) {
	r.Ids = &ids
}
func (r *KeywordDeleteRequest) SetKeywords(keywords []string) {
	r.Keywords = &keywords
}

func NewKeywordDeleteRequest(businessId string) *KeywordDeleteRequest {
	var request = &KeywordDeleteRequest{
		BizPostFormRequest: types.NewBizPostFormRequest(businessId),
	}
	request.SetProductCode("keyword")
	request.SetUriPattern("/v1/keyword/delete")
	request.SetVersion("v1")
	return request
}
func (r *KeywordDeleteRequest) GetBusinessCustomSignParams() map[string]string {
	parentParams := r.BizPostFormRequest.GetBusinessCustomSignParams()
	if r.Ids != nil {
		parentParams["ids"] = *r.Ids
	}
	if r.Keywords != nil {
		// convert to json
		keywords, _ := json.Marshal(r.Keywords)
		parentParams["keywords"] = string(keywords)
	}

	return parentParams
}

func (r *KeywordDeleteRequest) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "KeywordDeleteRequest"}
	if r == nil {
		invalidParams.Add(validation.NewErrParamRequired("KeywordDeleteRequest"))
	}
	if r.Ids == nil && r.Keywords == nil {
		invalidParams.Add(validation.NewErrParamRequired("Ids or Keywords"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}
