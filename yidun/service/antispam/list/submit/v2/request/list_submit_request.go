package request

import (
	"strconv"

	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/core/validation"
)

// ListSubmitRequest 提交名单请求
type ListSubmitRequest struct {
	*types.BizPostFormRequest
	Description *string `json:"description,omitempty"` // 描述
	ReleaseTime *int64  `json:"releaseTime,omitempty"` // 释放时间
	ListType    *int    `json:"listType,omitempty"`    // 名单种类
	EntityType  *int    `json:"entityType,omitempty"`  // 名单类型参数
	Entities    *string `json:"entities,omitempty"`    // 名单列表参数
	NeedConfirm *bool   `json:"needConfirm,omitempty"` // 是否需要确认
	NotCallback *bool   `json:"notCallback,omitempty"` // 是否不需要回调
	SpamType    *int    `json:"spamType,omitempty"`    // 违规类型
	SceneName   *string `json:"sceneName,omitempty"`   // 策略组名称
}

func NewListSubmitRequest(businessId string) *ListSubmitRequest {
	var request = &ListSubmitRequest{
		BizPostFormRequest: types.NewBizPostFormRequest(businessId),
	}
	request.SetProductCode("list")
	request.SetUriPattern("/v2/list/submit")
	request.SetVersion("v2")
	return request
}
func (r *ListSubmitRequest) GetBusinessCustomSignParams() map[string]string {
	parentParams := r.BizPostFormRequest.GetBusinessCustomSignParams()
	if r.Description != nil {
		parentParams["description"] = *r.Description
	}
	if r.ReleaseTime != nil {
		parentParams["releaseTime"] = strconv.FormatInt(*r.ReleaseTime, 10)
	}
	if r.ListType != nil {
		parentParams["listType"] = strconv.Itoa(*r.ListType)
	}
	if r.EntityType != nil {
		parentParams["entityType"] = strconv.Itoa(*r.EntityType)
	}
	if r.Entities != nil {
		parentParams["entities"] = *r.Entities
	}
	if r.NeedConfirm != nil {
		parentParams["needConfirm"] = strconv.FormatBool(*r.NeedConfirm)
	}
	if r.NotCallback != nil {
		parentParams["notCallback"] = strconv.FormatBool(*r.NotCallback)
	}
	if r.SpamType != nil {
		parentParams["spamType"] = strconv.Itoa(*r.SpamType)
	}
	if r.SceneName != nil {
		parentParams["sceneName"] = *r.SceneName
	}

	return parentParams
}

func (r *ListSubmitRequest) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "ListSubmitRequest"}
	if r == nil {
		invalidParams.Add(validation.NewErrParamRequired("ListSubmitRequest"))
	}
	if r.Entities == nil || len(*r.Entities) == 0 {
		invalidParams.Add(validation.NewErrParamRequired("Entities"))
	}
	if r.EntityType == nil {
		invalidParams.Add(validation.NewErrParamRequired("EntityType"))
	}
	if r.ListType == nil {
		invalidParams.Add(validation.NewErrParamRequired("ListType"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

func (req *ListSubmitRequest) SetDescription(description string) {
	req.Description = &description
}

func (req *ListSubmitRequest) SetReleaseTime(releaseTime int64) {
	req.ReleaseTime = &releaseTime
}

func (req *ListSubmitRequest) SetListType(listType int) {
	req.ListType = &listType
}

func (req *ListSubmitRequest) SetEntityType(entityType int) {
	req.EntityType = &entityType
}

func (req *ListSubmitRequest) SetEntities(entities string) {
	req.Entities = &entities
}

func (req *ListSubmitRequest) SetNeedConfirm(needConfirm bool) {
	req.NeedConfirm = &needConfirm
}

func (req *ListSubmitRequest) SetNotCallback(notCallback bool) {
	req.NotCallback = &notCallback
}

func (req *ListSubmitRequest) SetSpamType(spamType int) {
	req.SpamType = &spamType
}

func (req *ListSubmitRequest) SetSceneName(sceneName string) {
	req.SceneName = &sceneName
}
