package request

import (
	"encoding/json"
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/core/validation"
	"strconv"
)

// VideoSolutionQueryImageV1Request 视频解决方案查询截图请求
type VideoSolutionQueryImageV1Request struct {
	*types.PostFormRequest
	TaskId    *string             `json:"taskId,omitempty"` // taskId
	Levels    *[]int              `json:"levels,omitempty"` // 等级数组; json数组格式[1,2,3]
	PageNum   *int                `json:"pageNum,omitempty"`
	PageSize  *int                `json:"pageSize,omitempty"`
	OrderType *VideoDataOrderType `json:"orderType,omitempty"` // 排序类型
}
type VideoDataOrderType int

const (
	VideoDataOrderTypeDefault         VideoDataOrderType = 0
	VideoDataOrderTypeReceiveTimeASC  VideoDataOrderType = 3
	VideoDataOrderTypeReceiveTimeDESC VideoDataOrderType = 4
)

func (r *VideoSolutionQueryImageV1Request) SetTaskId(taskId string) {
	r.TaskId = &taskId
}
func (r *VideoSolutionQueryImageV1Request) SetLevels(levels []int) {
	r.Levels = &levels
}
func (r *VideoSolutionQueryImageV1Request) SetPageNum(pageNum int) {
	r.PageNum = &pageNum
}
func (r *VideoSolutionQueryImageV1Request) SetPageSize(pageSize int) {
	r.PageSize = &pageSize
}
func (r *VideoSolutionQueryImageV1Request) SetOrderType(orderType VideoDataOrderType) {
	r.OrderType = &orderType
}

// NewVideoSolutionQueryImageV1Req 初始化VideoSolutionQueryImageV1Request对象
func NewVideoSolutionQueryImageV1Req() *VideoSolutionQueryImageV1Request {
	var request = &VideoSolutionQueryImageV1Request{
		PostFormRequest: types.NewPostFormRequest(),
	}
	request.SetProductCode("videoSolutionCommon")
	request.SetUriPattern("/v1/videosolution/query/image")
	request.SetVersion("v1.1")
	return request
}
func (r *VideoSolutionQueryImageV1Request) GetBusinessCustomSignParams() map[string]string {
	params := r.PostFormRequest.GetBusinessCustomSignParams()
	if r.TaskId != nil {
		params["taskId"] = *r.TaskId
	}
	if r.Levels != nil {
		var levels, _ = json.Marshal(r.Levels)
		params["levels"] = string(levels)
	}
	if r.PageNum != nil {
		params["pageNum"] = strconv.Itoa(*r.PageNum)
	}
	if r.PageSize != nil {
		params["pageSize"] = strconv.Itoa(*r.PageSize)
	}
	if r.OrderType != nil {
		params["orderType"] = strconv.Itoa(int(*r.OrderType))
	}

	return params
}

func (r *VideoSolutionQueryImageV1Request) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "VideoSolutionQueryImageV1Request"}
	if r == nil {
		invalidParams.Add(validation.NewErrParamRequired("VideoSolutionQueryImageV1Request"))
	}
	if r.PageNum == nil || (*r.PageNum) <= 0 {
		invalidParams.Add(validation.NewErrParamMinValue("PageNum", 1))
	}
	if r.PageSize == nil || (*r.PageSize) < 0 || (*r.PageSize) > 1000 {
		invalidParams.Add(validation.NewErrParamBetweenRange("PageSize", 0, 1000))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}
