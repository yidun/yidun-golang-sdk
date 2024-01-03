package query

import "github.com/yidun/yidun-golang-sdk/yidun/core/types"

// PretreatmentQueryResponse 自定义忽略词添加响应
type PretreatmentQueryResponse struct {
	*types.CommonResponse
	Result *PretreatmentQueryResult `json:"result,omitempty"`
}

// PretreatmentQueryResult 自定义忽略词查询结果
type PretreatmentQueryResult struct {
	Count *int64 `json:"count,omitempty"`
	Rows  []*Row `json:"rows,omitempty"`
}

// Row row struct
type Row struct {
	Id *int64 `json:"id,omitempty"`
	// 忽略词
	Entity *string `json:"entity,omitempty"`
	// 忽略词类型
	EntityType *int `json:"entityType,omitempty"`
	// 状态
	Source *int `json:"source,omitempty"`
	// 创建人
	Creator *string `json:"creator,omitempty"`
	// 产品id
	ProductId *int64 `json:"productId,omitempty"`
	// 业务id
	TargetId *int64 `json:"targetId,omitempty"`
	// 创建时间
	InsertTime *int64 `json:"insertTime,omitempty"`
	// 修改时间
	LastModifyTime *int64 `json:"lastModifyTime,omitempty"`
}
