package request

import "github.com/yidun/yidun-golang-sdk/yidun/core/types"

// ListQueryResponse 查询名单响应
type ListQueryResponse struct {
	*types.CommonResponse
	Result *ListQueryResult `json:"result,omitempty"`
}
type ListQueryResult struct {
	Count *int64      `json:"count,omitempty"`
	Rows  *[]ListResp `json:"rows,omitempty"`
}

type ListResp struct {
	UUID        string `json:"uuid,omitempty"`
	ProductId   *int64 `json:"product_id,omitempty"`
	TargetId    *int64 `json:"target_id,omitempty"`
	EntityType  *int   `json:"entity_type,omitempty"`
	Entity      string `json:"entity,omitempty"`
	ListType    *int   `json:"list_type,omitempty"`
	ReleaseTime *int64 `json:"release_time,omitempty"`
	InsertTime  *int64 `json:"insert_time,omitempty"`
	UpdateTime  *int64 `json:"update_time,omitempty"`
	Operator    string `json:"operator,omitempty"`
	Description string `json:"description,omitempty"`
	Ext         string `json:"ext,omitempty"`
	SpamType    *int   `json:"spam_type,omitempty"`
	Source      *int   `json:"source,omitempty"`
	HitCount    *int64 `json:"hit_count,omitempty"`
	Status      *int   `json:"status,omitempty"`
}
