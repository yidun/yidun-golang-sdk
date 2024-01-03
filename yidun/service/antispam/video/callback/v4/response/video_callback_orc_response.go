package response

// VideoCallbackOcrV4Response 点播视频回调ocr识别结构
type VideoCallbackOcrV4Response struct {
	TaskId   *string                   `json:"taskId,omitempty"`
	DataId   *string                   `json:"dataId,omitempty"`
	Pictures []*VideoDataOcrV4Response `json:"pictures,omitempty"`
}

type VideoDataOcrV4Response struct {
	URL       *string                 `json:"url,omitempty"`
	StartTime *int64                  `json:"startTime,omitempty"`
	EndTime   *int64                  `json:"endTime,omitempty"`
	Height    *int                    `json:"height,omitempty"`
	Width     *int                    `json:"width,omitempty"`
	Details   []*VideoDataOcrDetailBo `json:"details,omitempty"`
	PictureID *string                 `json:"pictureId,omitempty"`
}

type VideoDataOcrDetailBo struct {
	// ocr文本
	Content *string `json:"content,omitempty"`
	// ocr文本详情
	LineContents []*OcrLineContent `json:"lineContents,omitempty"`
}

type OcrLineContent struct {
	// 行ocr文本
	LineContent *string `json:"lineContent,omitempty"`
	// 相对坐标
	X1 *float64 `json:"x1,omitempty"`
	Y1 *float64 `json:"y1,omitempty"`
	X2 *float64 `json:"x2,omitempty"`
	Y2 *float64 `json:"y2,omitempty"`
	// 语种信息
	Lang *string `json:"lang,omitempty"`
}
