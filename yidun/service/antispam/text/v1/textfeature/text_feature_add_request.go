package textfeature

import (
	"strconv"

	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/core/validation"
)

// TextFeatureAddRequest 文本特征添加请求
type TextFeatureAddRequest struct {
	*types.BizPostFormRequest
	FeatureType *int    `json:"featureType,omitempty"`
	Level       *int    `json:"level,omitempty"`
	MatchType   *int    `json:"matchType,omitempty"`
	SpamType    *int    `json:"spamType,omitempty"`
	SubLabel    *string `json:"subLabel,omitempty"`
	Description *string `json:"description,omitempty"`
	Entities    *string `json:"entities,omitempty"`
	Scope       *int    `json:"scope,omitempty"`
}

func (r *TextFeatureAddRequest) SetFeatureType(featureType int) {
	r.FeatureType = &featureType
}

func (r *TextFeatureAddRequest) SetLevel(level int) {
	r.Level = &level
}

func (r *TextFeatureAddRequest) SetMatchType(matchType int) {
	r.MatchType = &matchType
}

func (r *TextFeatureAddRequest) SetSpamType(spamType int) {
	r.SpamType = &spamType
}

func (r *TextFeatureAddRequest) SetSubLabel(subLabel string) {
	r.SubLabel = &subLabel
}

func (r *TextFeatureAddRequest) SetDescription(description string) {
	r.Description = &description
}

func (r *TextFeatureAddRequest) SetEntities(entities string) {
	r.Entities = &entities
}

func (r *TextFeatureAddRequest) SetScope(scope int) {
	r.Scope = &scope
}

func NewTextFeatureAddRequest(businessId string) *TextFeatureAddRequest {
	request := &TextFeatureAddRequest{
		BizPostFormRequest: types.NewBizPostFormRequest(businessId),
	}
	request.SetProductCode("text-api")
	request.SetUriPattern("/v1/text-feature/add")
	request.SetVersion("v1")
	return request
}

func (r *TextFeatureAddRequest) GetBusinessCustomSignParams() map[string]string {
	params := r.BizPostFormRequest.GetBusinessCustomSignParams()
	if r.FeatureType != nil {
		params["featureType"] = strconv.Itoa(*r.FeatureType)
	}
	if r.Level != nil {
		params["level"] = strconv.Itoa(*r.Level)
	}
	if r.MatchType != nil {
		params["matchType"] = strconv.Itoa(*r.MatchType)
	}
	if r.SpamType != nil {
		params["spamType"] = strconv.Itoa(*r.SpamType)
	}
	if r.SubLabel != nil {
		params["subLabel"] = *r.SubLabel
	}
	if r.Description != nil {
		params["description"] = *r.Description
	}
	if r.Entities != nil {
		params["entities"] = *r.Entities
	}
	if r.Scope != nil {
		params["scope"] = strconv.Itoa(*r.Scope)
	}
	return params
}

func (r *TextFeatureAddRequest) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "TextFeatureAddRequest"}
	if r == nil {
		invalidParams.Add(validation.NewErrParamRequired("TextFeatureAddRequest"))
		return invalidParams
	}
	if r.FeatureType == nil {
		invalidParams.Add(validation.NewErrParamRequired("FeatureType"))
	}
	if r.Level == nil {
		invalidParams.Add(validation.NewErrParamRequired("Level"))
	}
	if r.MatchType == nil {
		invalidParams.Add(validation.NewErrParamRequired("MatchType"))
	}
	if r.Entities == nil || len(*r.Entities) == 0 {
		invalidParams.Add(validation.NewErrParamRequired("Entities"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}
