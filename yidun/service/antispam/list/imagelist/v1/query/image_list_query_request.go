package query

import (
	"strconv"

	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/core/validation"
)

// ImageListQueryRequest 图片名单查询请求
type ImageListQueryRequest struct {
	*types.BizPostFormRequest
	// 名单分类，1: 白名单，2: 黑名单
	ListType *int `json:"listType,omitempty"`
	// 名单分类（加黑名单时传），100：色情，110：性感，200：广告，210：二维码，300：暴恐，400：违禁，500：涉政
	ImageLabel *int `json:"imageLabel,omitempty"`
	// 名单类型，0: 相似名单，1: md5精确名单
	Type *int `json:"type,omitempty"`
	// 页数
	PageNum *int `json:"pageNum,omitempty"`
	// 每页大小
	PageSize *int `json:"pageSize,omitempty"`
	// 查询开始时间
	StartTime *int64 `json:"startTime,omitempty"`
	// 查询结束时间
	EndTime *int64 `json:"endTime,omitempty"`
	// 图片名单url
	URL *string `json:"url,omitempty"`
	// 名单状态
	Status *int `json:"status,omitempty"`
	// 排序方式
	SortType *int `json:"sortType,omitempty"`
	// 业务id
	BusinessId *string `json:"businessId,omitempty"`
}

// NewImageListQueryRequest 初始化ImageListQueryRequest对象
func NewImageListQueryRequest(businessId string) *ImageListQueryRequest {
	var request = &ImageListQueryRequest{
		BizPostFormRequest: types.NewBizPostFormRequest(businessId),
	}
	request.SetProductCode("list")
	request.SetUriPattern("/v1/image/list/pageQuery")
	request.SetMethod(http.HttpMethodPost)
	request.SetVersion("v1")
	return request
}

// 参数校验方法
func (r *ImageListQueryRequest) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "ImageListQueryRequest"}
	if r.Type == nil {
		invalidParams.Add(validation.NewErrParamRequired("Type"))
	}
	if r.PageNum == nil {
		invalidParams.Add(validation.NewErrParamRequired("PageNum"))
	}
	if r.PageSize == nil {
		invalidParams.Add(validation.NewErrParamRequired("PageSize"))
	}
	if r.StartTime == nil {
		invalidParams.Add(validation.NewErrParamRequired("StartTime"))
	}
	if r.EndTime == nil {
		invalidParams.Add(validation.NewErrParamRequired("EndTime"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// SetListType sets the ListType field.
func (req *ImageListQueryRequest) SetListType(listType int) {
	req.ListType = &listType
}

// SetImageLabel sets the ImageLabel field.
func (req *ImageListQueryRequest) SetImageLabel(imageLabel int) {
	req.ImageLabel = &imageLabel
}

// SetType sets the Type field.
func (req *ImageListQueryRequest) SetType(typ int) {
	req.Type = &typ
}

// SetPageNum sets the PageNum field.
func (req *ImageListQueryRequest) SetPageNum(pageNum int) {
	req.PageNum = &pageNum
}

// SetPageSize sets the PageSize field.
func (req *ImageListQueryRequest) SetPageSize(pageSize int) {
	req.PageSize = &pageSize
}

// SetStartTime sets the StartTime field.
func (req *ImageListQueryRequest) SetStartTime(startTime int64) {
	req.StartTime = &startTime
}

// SetEndTime sets the EndTime field.
func (req *ImageListQueryRequest) SetEndTime(endTime int64) {
	req.EndTime = &endTime
}

// SetURL sets the URL field.
func (req *ImageListQueryRequest) SetURL(url string) {
	req.URL = &url
}

// SetStatus sets the Status field.
func (req *ImageListQueryRequest) SetStatus(status int) {
	req.Status = &status
}

// SetSortType sets the SortType field.
func (req *ImageListQueryRequest) SetSortType(sortType int) {
	req.SortType = &sortType
}

// SetBusinessId sets the BusinessId field.
func (req *ImageListQueryRequest) SetBusinessId(businessId string) {
	req.BusinessId = &businessId
}

// GetCustomSignParams returns the custom sign parameters
func (req *ImageListQueryRequest) GetBusinessCustomSignParams() map[string]string {
	params := req.PostFormRequest.GetBusinessCustomSignParams()
	if req.ListType != nil {
		params["listType"] = strconv.Itoa(*req.ListType)
	}
	if req.ImageLabel != nil {
		params["imageLabel"] = strconv.Itoa(*req.ImageLabel)
	}
	if req.Type != nil {
		params["type"] = strconv.Itoa(*req.Type)
	}
	if req.PageNum != nil {
		params["pageNum"] = strconv.Itoa(*req.PageNum)
	}
	if req.PageSize != nil {
		params["pageSize"] = strconv.Itoa(*req.PageSize)
	}
	if req.StartTime != nil {
		params["startTime"] = strconv.FormatInt(*req.StartTime, 10)
	}
	if req.EndTime != nil {
		params["endTime"] = strconv.FormatInt(*req.EndTime, 10)
	}
	if req.URL != nil {
		params["url"] = *req.URL
	}
	if req.Status != nil {
		params["status"] = strconv.Itoa(*req.Status)
	}
	if req.SortType != nil {
		params["sortType"] = strconv.Itoa(*req.SortType)
	}
	if req.BusinessId != nil {
		params["businessId"] = *req.BusinessId
	}
	return params
}
