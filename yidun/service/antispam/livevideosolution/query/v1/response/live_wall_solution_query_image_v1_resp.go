package response

import "github.com/yidun/yidun-golang-sdk/yidun/core/types"

// LiveWallSolutionQueryImageV1Resp represents live wall solution query image version 1 response.
type LiveWallSolutionQueryImageV1Resp struct {
	*types.CommonResponse                                     // Embedded struct, represents common response
	Result                *LiveWallSolutionQueryImageV1Result `json:"result,omitempty"` // The result of the query
}

// LiveWallSolutionQueryImageV1Result represents query image version 1 result.
type LiveWallSolutionQueryImageV1Result struct {
	// 错误状态，0 代表成功
	Status *int  `json:"status,omitempty"`
	Images *Page `json:"images,omitempty"` // Page of Image units
}

type Page struct {

	// The total number of items.
	Count int64 `json:"count,omitempty"`

	// The items on the current page.
	Rows []*LiveWallSolutionQueryImageV1Unit `json:"rows,omitempty"`
}

// Fill out the fields of this struct based on your requirements

// LiveWallSolutionQueryImageV1Unit represents an image unit in the result.
type LiveWallSolutionQueryImageV1Unit struct {
	Url            *string `json:"url,omitempty"`
	Label          *int    `json:"label,omitempty"`
	LabelLevel     *int    `json:"labelLevel,omitempty"`
	CallbackStatus *int    `json:"callbackStatus,omitempty"`
	BeginTime      *int64  `json:"beginTime,omitempty"`
	EndTime        *int64  `json:"endTime,omitempty"`
	SpeakerId      *string `json:"speakerId,omitempty"`
}

// Constants for task ID status
const (
	OK                 int = 0
	TASK_ID_EXPIRED    int = 20
	TASK_ID_NOT_EXISTS int = 30
)
