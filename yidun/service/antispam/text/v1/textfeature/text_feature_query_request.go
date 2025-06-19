package textfeature

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/core/validation"
	"strconv"
)

// TextFeatureQueryRequest 文本特征查询请求
type TextFeatureQueryRequest struct {
	*types.BizPostFormRequest
	StartTime   *int64 `json:"startTime,omitempty"`
	EndTime     *int64 `json:"endTime,omitempty"`
	FeatureType *int   `json:"featureType,omitempty"`
	Level       *int   `json:"level,omitempty"`
	MatchType   *int   `json:"matchType,omitempty"`
	SpamType    *int   `json:"spamType,omitempty"`
	Status      *int   `json:"status,omitempty"`
	PageNum     *int   `json:"pageNum,omitempty"`
	PageSize    *int   `json:"pageSize,omitempty"`
	Scope       *int   `json:"scope,omitempty"`
}

func (r *TextFeatureQueryRequest) SetStartTime(startTime int64) {
	r.StartTime = &startTime
}

func (r *TextFeatureQueryRequest) SetEndTime(endTime int64) {
	r.EndTime = &endTime
}

func (r *TextFeatureQueryRequest) SetFeatureType(featureType int) {
	r.FeatureType = &featureType
}

func (r *TextFeatureQueryRequest) SetLevel(level int) {
	r.Level = &level
}

func (r *TextFeatureQueryRequest) SetMatchType(matchType int) {
	r.MatchType = &matchType
}

func (r *TextFeatureQueryRequest) SetSpamType(spamType int) {
	r.SpamType = &spamType
}

func (r *TextFeatureQueryRequest) SetStatus(status int) {
	r.Status = &status
}

func (r *TextFeatureQueryRequest) SetPageNum(pageNum int) {
	r.PageNum = &pageNum
}

func (r *TextFeatureQueryRequest) SetPageSize(pageSize int) {
	r.PageSize = &pageSize
}

func (r *TextFeatureQueryRequest) SetScope(scope int) {
	r.Scope = &scope
}

func NewTextFeatureQueryRequest(businessId string) *TextFeatureQueryRequest {
	request := &TextFeatureQueryRequest{
		BizPostFormRequest: types.NewBizPostFormRequest(businessId),
	}
	request.SetProductCode("text-api")
	request.SetUriPattern("/v1/text-feature/query")
	request.SetVersion("v1")
	return request
}

func (r *TextFeatureQueryRequest) GetBusinessCustomSignParams() map[string]string {
	params := r.BizPostFormRequest.GetBusinessCustomSignParams()
	if r.StartTime != nil {
		params["startTime"] = strconv.Itoa(int(*r.StartTime))
	}
	if r.EndTime != nil {
		params["endTime"] = strconv.Itoa(int(*r.EndTime))
	}
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
	if r.Status != nil {
		params["status"] = strconv.Itoa(*r.Status)
	}
	if r.PageNum != nil {
		params["pageNum"] = strconv.Itoa(*r.PageNum)
	}
	if r.PageSize != nil {
		params["pageSize"] = strconv.Itoa(*r.PageSize)
	}
	if r.Scope != nil {
		params["scope"] = strconv.Itoa(*r.Scope)
	}
	return params
}

func (r *TextFeatureQueryRequest) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "TextFeatureQueryRequest"}
	if r == nil {
		invalidParams.Add(validation.NewErrParamRequired("TextFeatureQueryRequest"))
		return invalidParams
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}
