package request

import "github.com/yidun/yidun-golang-sdk/yidun/core/types"

// KeywordQueryResponse 查询关键词响应
type KeywordQueryResponse struct {
	*types.CommonResponse
	Result *KeywordQueryResult `json:"result,omitempty"`
}
type KeywordQueryResult struct {
	Words *KeywordQueryPageResultWords `json:"words,omitempty"`
}
type KeywordQueryPageResultWords struct {
	Count *int64  `json:"count,omitempty"`
	Rows  *[]Word `json:"rows,omitempty"`
}

type Word struct {
	Id         *int64  `json:"id,omitempty"`
	Word       *string `json:"word,omitempty"`
	Category   *int    `json:"category,omitempty"`
	SubLabel   *string `json:"subLabel,omitempty"`
	Type       *int    `json:"type,omitempty"`
	Status     *int    `json:"status,omitempty"`
	Level      *int    `json:"level,omitempty"`
	ProductId  *int    `json:"productId,omitempty"`
	TargetId   *int64  `json:"targetId,omitempty"`
	UpdateTime *int64  `json:"updateTime,omitempty"`
}
