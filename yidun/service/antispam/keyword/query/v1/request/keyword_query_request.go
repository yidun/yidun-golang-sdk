package request

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/core/validation"
	"strconv"
)

// KeywordQueryRequest 查询关键词请求
type KeywordQueryRequest struct {
	*types.BizPostFormRequest
	Id        *int64  `json:"id,omitempty"`      // 关键词id
	Keyword   *string `json:"keyword,omitempty"` // 关键词
	Category  *int    `json:"category"`          // 关键词返回的垃圾分类
	OrderType *int    `json:"orderType"`         // 排序类型
	PageNum   *int    `json:"pageNum"`           // 当前页码
	PageSize  *int    `json:"pageSize"`          // 每页条数
}

func (r *KeywordQueryRequest) SetId(id int64) {
	r.Id = &id
}
func (r *KeywordQueryRequest) SetKeyword(keyword string) {
	r.Keyword = &keyword
}
func (r *KeywordQueryRequest) SetCategory(category int) {
	r.Category = &category
}
func (r *KeywordQueryRequest) SetOrderType(orderType int) {
	r.OrderType = &orderType
}
func (r *KeywordQueryRequest) SetPageNum(pageNum int) {
	r.PageNum = &pageNum
}
func (r *KeywordQueryRequest) SetPageSize(pageSize int) {
	r.PageSize = &pageSize
}

func NewKeywordQueryRequest(businessId string) *KeywordQueryRequest {
	var request = &KeywordQueryRequest{
		BizPostFormRequest: types.NewBizPostFormRequest(businessId),
	}
	request.SetProductCode("keyword")
	request.SetUriPattern("/v1/keyword/query")
	request.SetVersion("v1")
	return request
}
func (r *KeywordQueryRequest) GetBusinessCustomSignParams() map[string]string {
	parentParams := r.BizPostFormRequest.GetBusinessCustomSignParams()
	if r.Id != nil {
		parentParams["id"] = strconv.FormatInt(*r.Id, 10)
	}
	if r.Category != nil {
		parentParams["category"] = strconv.Itoa(*r.Category)
	}
	if r.Keyword != nil {
		parentParams["keyword"] = *r.Keyword
	}
	if r.OrderType != nil {
		parentParams["orderType"] = strconv.Itoa(*r.OrderType)
	}
	if r.PageNum != nil {
		parentParams["pageNum"] = strconv.Itoa(*r.PageNum)
	}
	if r.PageSize != nil {
		parentParams["pageSize"] = strconv.Itoa(*r.PageSize)
	}
	return parentParams
}

func (r *KeywordQueryRequest) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "KeywordQueryRequest"}
	if r == nil {
		invalidParams.Add(validation.NewErrParamRequired("KeywordQueryRequest"))
	}
	if r.Keyword != nil && len(*r.Keyword) > 128 {
		invalidParams.Add(validation.NewErrParamMaxLen("Keyword", 128, *r.Keyword))
	}
	if r.PageNum == nil || *r.PageNum < 1 {
		invalidParams.Add(validation.NewErrParamMinValue("PageNum", 1))
	}
	if r.PageSize == nil || *r.PageSize < 0 || *r.PageSize > 1000 {
		invalidParams.Add(validation.NewErrParamBetweenRange("PageSize", 1, 1000))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}
