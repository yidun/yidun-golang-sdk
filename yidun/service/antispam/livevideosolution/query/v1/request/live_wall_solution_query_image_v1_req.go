package request

import (
	"encoding/json"
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"strconv"
)

// LiveWallSolutionQueryImageV1Req represents a request to LiveWallSolutionQueryImageV1 API.
type LiveWallSolutionQueryImageV1Req struct {
	*types.PostFormRequest
	// taskId
	TaskId *string `json:"taskId,omitempty"`
	// 等级数组，@See LabelLevel; json数组格式[1,2,3]
	Levels   *[]int `json:"levels,omitempty"`
	PageNum  *int   `json:"pageNum,omitempty"`
	PageSize *int   `json:"pageSize,omitempty"`
	// 回调状态
	CallbackStatus *int `json:"callbackStatus,omitempty"`
	// 接收开始时间
	StartTime *int64 `json:"startTime,omitempty"`
	// 接收结束时间
	EndTime *int64 `json:"endTime,omitempty"`
	// 排序类型 @See LiveVideoOrderType;
	OrderType *int `json:"orderType,omitempty"`
}

// NewLiveAudioCallbackV4Req 创建一个 LiveAudioCallbackV4Req 实例
func NewLiveWallSolutionQueryImageV1Req() *LiveWallSolutionQueryImageV1Req {
	var request = &LiveWallSolutionQueryImageV1Req{
		PostFormRequest: types.NewPostFormRequest(),
	}
	request.SetProductCode("liveVideoSolutionCommon")
	request.SetUriPattern("/v1/livewallsolution/query/image")
	request.SetVersion("v1")
	return request
}

// GetCustomSignParams 获取自定义签名参数
func (l *LiveWallSolutionQueryImageV1Req) GetBusinessCustomSignParams() map[string]string {
	result := l.PostFormRequest.GetBusinessCustomSignParams()
	if l.TaskId != nil {
		result["taskId"] = *l.TaskId
	}
	if l.Levels != nil {
		levelsJson, _ := json.Marshal(*l.Levels)
		result["levels"] = string(levelsJson)
	}
	if l.PageNum != nil {
		result["pageNum"] = strconv.Itoa(*l.PageNum)
	}
	if l.PageSize != nil {
		result["pageSize"] = strconv.Itoa(*l.PageSize)
	}
	if l.CallbackStatus != nil {
		result["callbackStatus"] = strconv.Itoa(*l.CallbackStatus)
	}
	if l.StartTime != nil {
		result["startTime"] = strconv.FormatInt(*l.StartTime, 10)
	}
	if l.EndTime != nil {
		result["endTime"] = strconv.FormatInt(*l.EndTime, 10)
	}
	if l.OrderType != nil {
		result["orderType"] = strconv.Itoa(*l.OrderType)
	}
	return result
}

// SetTaskId sets the TaskId.
func (l *LiveWallSolutionQueryImageV1Req) SetTaskId(taskId string) {
	l.TaskId = &taskId
}

// SetLevels sets the Levels.
func (l *LiveWallSolutionQueryImageV1Req) SetLevels(levels []int) {
	l.Levels = &levels
}

// SetPageNum sets the PageNum.
func (l *LiveWallSolutionQueryImageV1Req) SetPageNum(pageNum int) {
	l.PageNum = &pageNum
}

// SetPageSize sets the PageSize.
func (l *LiveWallSolutionQueryImageV1Req) SetPageSize(pageSize int) {
	l.PageSize = &pageSize
}

// SetCallbackStatus sets the CallbackStatus.
func (l *LiveWallSolutionQueryImageV1Req) SetCallbackStatus(callbackStatus int) {
	l.CallbackStatus = &callbackStatus
}

// SetStartTime sets the StartTime.
func (l *LiveWallSolutionQueryImageV1Req) SetStartTime(startTime int64) {
	l.StartTime = &startTime
}

// SetEndTime sets the EndTime.
func (l *LiveWallSolutionQueryImageV1Req) SetEndTime(endTime int64) {
	l.EndTime = &endTime
}

// SetOrderType sets the OrderType.
func (l *LiveWallSolutionQueryImageV1Req) SetOrderType(orderType int) {
	l.OrderType = &orderType
}
