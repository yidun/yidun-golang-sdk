package request

import (
	"strconv"

	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/core/validation"
)

// ListDeleteRequest 删除名单请求
type ListDeleteRequest struct {
	*types.BizPostFormRequest
	EntityType *int    `json:"entityType,omitempty"` //名单类型参数
	Uuids      *string `json:"uuids,omitempty"`      //名单加密id列表参数
	ListType   *int    `json:"listType,omitempty"`   //名单加密id列表参数
	Entities   *string `json:"entities,omitempty"`   //名单加密id列表参数
	SceneName  *string `json:"sceneName,omitempty"`  //策略组名称
}

func (s *ListDeleteRequest) SetEntityType(entityType int) {
	s.EntityType = &entityType
}

func (s *ListDeleteRequest) SetUuids(uuids string) {
	s.Uuids = &uuids
}

func (s *ListDeleteRequest) SetListType(listType int) {
	s.ListType = &listType
}

func (s *ListDeleteRequest) SetEntities(entities string) {
	s.Entities = &entities
}

func (s *ListDeleteRequest) SetSceneName(sceneName string) {
	s.SceneName = &sceneName
}

func NewListDeleteRequest(businessId string) *ListDeleteRequest {
	var request = &ListDeleteRequest{
		BizPostFormRequest: types.NewBizPostFormRequest(businessId),
	}
	request.SetProductCode("list")
	request.SetUriPattern("/v2/list/batchDelete")
	request.SetVersion("v2")
	return request
}
func (r *ListDeleteRequest) GetBusinessCustomSignParams() map[string]string {
	parentParams := r.BizPostFormRequest.GetBusinessCustomSignParams()

	if r.EntityType != nil {
		parentParams["entityType"] = strconv.Itoa(*r.EntityType)
	}

	if r.Uuids != nil {
		parentParams["uuids"] = *r.Uuids
	}

	if r.ListType != nil {
		parentParams["listType"] = strconv.Itoa(*r.ListType)
	}

	if r.Entities != nil {
		parentParams["entities"] = *r.Entities
	}

	if r.SceneName != nil {
		parentParams["sceneName"] = *r.SceneName
	}

	return parentParams
}

func (r *ListDeleteRequest) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "ListDeleteRequest"}
	if r == nil {
		invalidParams.Add(validation.NewErrParamRequired("ListDeleteRequest"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}
