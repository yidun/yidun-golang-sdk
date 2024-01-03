package query

import (
	"strconv"

	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/core/validation"
)

// PretreatmentQueryRequest 自定义忽略词查询请求
type PretreatmentQueryRequest struct {
	*types.BizPostFormRequest
	// 忽略词id
	Id *int64 `json:"id,omitempty"`
	// 忽略词内容
	Entity *string `json:"entity,omitempty"`
	// 开始时间
	StartTime *int64 `json:"startTime,omitempty"`
	// 结束时间
	EndTime *int64 `json:"endTime,omitempty"`
	// 分页标号
	PageNum *int `json:"pageNum,omitempty"`
	// 分页大小
	PageSize *int `json:"pageSize,omitempty"`
	// 业务id
	BusinessId *string `json:"businessId,omitempty"`
}

// 参数校验方法
func (r *PretreatmentQueryRequest) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "PretreatmentQueryRequest"}
	if r.Entity != nil && len(*r.Entity) > 128 {
		invalidParams.Add(validation.NewErrParamMaxLen("entity不能超过128", 128, strconv.Itoa(len(*r.Entity))))
	}
	if r.PageNum == nil {
		invalidParams.Add(validation.NewErrParamRequired("PageNum"))
	}
	if r.PageSize == nil {
		invalidParams.Add(validation.NewErrParamRequired("PageSize"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// NewPretreatmentQueryRequest 初始化 PretreatmentDeleteRequest 对象
func NewPretreatmentQueryRequest(businessId string) *PretreatmentQueryRequest {
	var request = &PretreatmentQueryRequest{
		BizPostFormRequest: types.NewBizPostFormRequest(businessId),
	}
	request.SetProductCode("pretreatment")
	request.SetUriPattern("/v1/pretreatment/pageQuery")
	request.SetMethod(http.HttpMethodPost)
	request.SetVersion("v1")
	return request
}

// GetCustomSignParams returns the custom sign parameters for the request
func (req *PretreatmentQueryRequest) GetBusinessCustomSignParams() map[string]string {
	params := req.PostFormRequest.GetBusinessCustomSignParams()
	if req.Id != nil {
		params["id"] = strconv.FormatInt(*req.Id, 10)
	}
	if req.Entity != nil {
		params["entity"] = *req.Entity
	}
	if req.StartTime != nil {
		params["startTime"] = strconv.FormatInt(*req.StartTime, 10)
	}
	if req.EndTime != nil {
		params["endTime"] = strconv.FormatInt(*req.EndTime, 10)
	}
	if req.PageNum != nil {
		params["pageNum"] = strconv.Itoa(*req.PageNum)
	}
	if req.PageSize != nil {
		params["pageSize"] = strconv.Itoa(*req.PageSize)
	}
	if req.BusinessId != nil {
		params["businessId"] = *req.BusinessId
	}
	return params
}

// SetId sets the Id field.
func (req *PretreatmentQueryRequest) SetId(id int64) {
	req.Id = &id
}

// SetEntity sets the Entity field.
func (req *PretreatmentQueryRequest) SetEntity(entity string) {
	req.Entity = &entity
}

// SetStartTime sets the StartTime field.
func (req *PretreatmentQueryRequest) SetStartTime(startTime int64) {
	req.StartTime = &startTime
}

// SetEndTime sets the EndTime field.
func (req *PretreatmentQueryRequest) SetEndTime(endTime int64) {
	req.EndTime = &endTime
}

// SetPageNum sets the PageNum field.
func (req *PretreatmentQueryRequest) SetPageNum(pageNum int) {
	req.PageNum = &pageNum
}

// SetPageSize sets the PageSize field.
func (req *PretreatmentQueryRequest) SetPageSize(pageSize int) {
	req.PageSize = &pageSize
}

// SetBusinessId sets the BusinessId field.
func (req *PretreatmentQueryRequest) SetBusinessId(businessId string) {
	req.BusinessId = &businessId
}
