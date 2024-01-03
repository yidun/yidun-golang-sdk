package request

import (
	"strconv"

	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/core/validation"
)

// ListQueryRequest 查询名单请求
type ListQueryRequest struct {
	*types.BizPostFormRequest
	EntityType *int    `json:"entityType,omitempty"` // 名单类型参数
	PageNum    *int    `json:"pageNum,omitempty"`    // 页数
	PageSize   *int    `json:"pageSize,omitempty"`   // 每页大小
	StartTime  *int64  `json:"startTime,omitempty"`  // 查询开始时间
	EndTime    *int64  `json:"endTime,omitempty"`    // 查询结束时间
	ListType   *int    `json:"listType,omitempty"`   // 黑白名单类型
	Source     *int    `json:"source,omitempty"`     // 名单添加来源
	Entity     *string `json:"entity,omitempty"`     // 名单
	Status     *int    `json:"status,omitempty"`     // 名单状态
	OrderType  *int    `json:"orderType,omitempty"`  // 排序方式
	UUID       *string `json:"uuid,omitempty"`       // 名单加密id
	SpamType   *int    `json:"spamType,omitempty"`   // 垃圾类型
}

func (r *ListQueryRequest) SetEntityType(entityType int) {
	r.EntityType = &entityType
}

func (r *ListQueryRequest) SetPageNum(pageNum int) {
	r.PageNum = &pageNum
}

func (r *ListQueryRequest) SetPageSize(pageSize int) {
	r.PageSize = &pageSize
}

func (r *ListQueryRequest) SetStartTime(startTime int64) {
	r.StartTime = &startTime
}

func (r *ListQueryRequest) SetEndTime(endTime int64) {
	r.EndTime = &endTime
}

func (r *ListQueryRequest) SetListType(listType int) {
	r.ListType = &listType
}

func (r *ListQueryRequest) SetSource(source int) {
	r.Source = &source
}

func (r *ListQueryRequest) SetEntity(entity string) {
	r.Entity = &entity
}

func (r *ListQueryRequest) SetStatus(status int) {
	r.Status = &status
}

func (r *ListQueryRequest) SetOrderType(orderType int) {
	r.OrderType = &orderType
}

func (r *ListQueryRequest) SetUUID(uuid string) {
	r.UUID = &uuid
}

func (r *ListQueryRequest) SetSpamType(spamType int) {
	r.SpamType = &spamType
}

func NewListQueryRequest(businessId string) *ListQueryRequest {
	var request = &ListQueryRequest{
		BizPostFormRequest: types.NewBizPostFormRequest(businessId),
	}
	request.SetProductCode("list")
	request.SetUriPattern("/v2/list/pageQuery")
	request.SetVersion("v2")
	return request
}
func (r *ListQueryRequest) GetBusinessCustomSignParams() map[string]string {
	parentParams := r.BizPostFormRequest.GetBusinessCustomSignParams()
	if r.EntityType != nil {
		parentParams["entityType"] = strconv.Itoa(*r.EntityType)
	}
	if r.PageNum != nil {
		parentParams["pageNum"] = strconv.Itoa(*r.PageNum)
	}
	if r.PageSize != nil {
		parentParams["pageSize"] = strconv.Itoa(*r.PageSize)
	}
	if r.StartTime != nil {
		parentParams["startTime"] = strconv.FormatInt(*r.StartTime, 10)
	}
	if r.EndTime != nil {
		parentParams["endTime"] = strconv.FormatInt(*r.EndTime, 10)
	}
	if r.ListType != nil {
		parentParams["listType"] = strconv.Itoa(*r.ListType)
	}
	if r.Source != nil {
		parentParams["source"] = strconv.Itoa(*r.Source)
	}
	if r.Entity != nil {
		parentParams["entity"] = *r.Entity
	}
	if r.Status != nil {
		parentParams["status"] = strconv.Itoa(*r.Status)
	}
	if r.OrderType != nil {
		parentParams["orderType"] = strconv.Itoa(*r.OrderType)
	}
	if r.UUID != nil {
		parentParams["uuid"] = *r.UUID
	}
	if r.SpamType != nil {
		parentParams["spamType"] = strconv.Itoa(*r.SpamType)
	}
	return parentParams
}

func (r *ListQueryRequest) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "ListQueryRequest"}
	if r == nil {
		invalidParams.Add(validation.NewErrParamRequired("ListQueryRequest"))
	}
	if r.EntityType == nil {
		invalidParams.Add(validation.NewErrParamRequired("EntityType"))
	}
	if r.PageNum == nil || *r.PageNum < 1 {
		invalidParams.Add(validation.NewErrParamMinValue("PageNum", 1))
	}
	if r.PageSize == nil || *r.PageSize < 0 || *r.PageSize > 500 {
		invalidParams.Add(validation.NewErrParamBetweenRange("PageSize", 1, 500))
	}
	if r.StartTime == nil || *r.StartTime < 0 {
		invalidParams.Add(validation.NewErrParamMinValue("StartTime", 0))
	}
	if r.EndTime == nil || *r.EndTime < 0 {
		invalidParams.Add(validation.NewErrParamMinValue("EndTime", 0))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}
