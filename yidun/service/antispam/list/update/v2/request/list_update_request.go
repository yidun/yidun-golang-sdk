package request

import (
	"strconv"

	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/core/validation"
)

// ListUpdateRequest 更新名单请求
type ListUpdateRequest struct {
	*types.BizPostFormRequest
	EntityType  *int    `json:"entityType,omitempty"`  // 名单类型参数
	Uuid        *string `json:"uuid,omitempty"`        // 名单加密id
	Status      *int    `json:"status,omitempty"`      // 生效状态，0 未生效， 1 已生效
	ReleaseTime *int64  `json:"releaseTime,omitempty"` // 释放时间
	SpamType    *int    `json:"spamType,omitempty"`    // 垃圾类型
	ListType    *int    `json:"listType,omitempty"`    // 名单种类
	Entity      *string `json:"entity,omitempty"`      // 名单值
	NotCallback *bool   `json:"notCallback,omitempty"` // 是否不需要回调，默认false，false-需要回调；true：不需要回调
	SceneName   *string `json:"sceneName,omitempty"`   // 策略组名称
}

func (r *ListUpdateRequest) SetEntityType(entityType int) {
	r.EntityType = &entityType
}

func (r *ListUpdateRequest) SetUuid(uuid string) {
	r.Uuid = &uuid
}

func (r *ListUpdateRequest) SetStatus(status int) {
	r.Status = &status
}

func (r *ListUpdateRequest) SetReleaseTime(releaseTime int64) {
	r.ReleaseTime = &releaseTime
}

func (r *ListUpdateRequest) SetSpamType(spamType int) {
	r.SpamType = &spamType
}

func (r *ListUpdateRequest) SetListType(listType int) {
	r.ListType = &listType
}

func (r *ListUpdateRequest) SetEntity(entity string) {
	r.Entity = &entity
}

func (r *ListUpdateRequest) SetNotCallback(notCallback bool) {
	r.NotCallback = &notCallback
}

func (r *ListUpdateRequest) SetSceneName(sceneName string) {
	r.SceneName = &sceneName
}

func NewListUpdateRequest(businessId string) *ListUpdateRequest {
	var request = &ListUpdateRequest{
		BizPostFormRequest: types.NewBizPostFormRequest(businessId),
	}
	request.SetProductCode("list")
	request.SetUriPattern("/v2/list/update")
	request.SetVersion("v2")
	return request
}
func (r *ListUpdateRequest) GetBusinessCustomSignParams() map[string]string {
	parentParams := r.BizPostFormRequest.GetBusinessCustomSignParams()
	if r.EntityType != nil {
		parentParams["entityType"] = strconv.Itoa(*r.EntityType)
	}
	if r.Uuid != nil {
		parentParams["uuid"] = *r.Uuid
	}
	if r.Status != nil {
		parentParams["status"] = strconv.Itoa(*r.Status)
	}
	if r.ReleaseTime != nil {
		parentParams["releaseTime"] = strconv.FormatInt(*r.ReleaseTime, 10)
	}
	if r.SpamType != nil {
		parentParams["spamType"] = strconv.Itoa(*r.SpamType)
	}
	if r.ListType != nil {
		parentParams["listType"] = strconv.Itoa(*r.ListType)
	}
	if r.Entity != nil {
		parentParams["entity"] = *r.Entity
	}
	if r.NotCallback != nil {
		parentParams["notCallback"] = strconv.FormatBool(*r.NotCallback)
	}
	if r.SceneName != nil {
		parentParams["sceneName"] = *r.SceneName
	}

	return parentParams
}

func (r *ListUpdateRequest) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "ListUpdateRequest"}
	if r == nil {
		invalidParams.Add(validation.NewErrParamRequired("ListUpdateRequest"))
	}
	if r.EntityType == nil {
		invalidParams.Add(validation.NewErrParamRequired("EntityType"))
	}
	if r.Status == nil {
		invalidParams.Add(validation.NewErrParamRequired("Status"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}
