package response

type MediaCensorResponse struct {
	// 融媒体人工审核响应结果
	DataId              *string         `json:"dataId,omitempty"`
	TaskId              *string         `json:"taskId,omitempty"`
	Callback            *string         `json:"callback,omitempty"`
	Suggestion          *int            `json:"suggestion,omitempty"`
	ResultType          *int            `json:"resultType,omitempty"` // 审核类型，1:机器检测，2:审核
	CensorSource        *int            `json:"censorSource,omitempty"`
	CensorOperateSource *int            `json:"censorOperateSource,omitempty"`
	CensorRound         *int            `json:"censorRound,omitempty"`
	CensorTime          *int64          `json:"censorTime,omitempty"` // 最近一次人工审核或质检的时间
	SkipCensor          *int            `json:"skipCensor,omitempty"`
	CensorLabels        []*CensorLabel  `json:"censorLabels,omitempty"` // 标签
	ReviewEvidences     *ReviewEvidence `json:"reviewEvidences,omitempty"`
}

// 人审证据信息
type ReviewEvidence struct {
	Reason       *string        `json:"reason,omitempty"`
	Remark       *string        `json:"remark,omitempty"`
	CensorLabels []*CensorLabel `json:"censorLabels,omitempty"` // 人审证据详细信息
	Detail       *Detail        `json:"detail,omitempty"`
}

// 人审证据详细信息
type Detail struct {
	Texts       []*TextDetail       `json:"texts,omitempty"`
	Images      []*ImageDetail      `json:"images,omitempty"`
	Audios      []*AudioDetail      `json:"audios,omitempty"`
	Videos      []*VideoDetail      `json:"videos,omitempty"`
	Audiovideos []*AudioVideoDetail `json:"audiovideos,omitempty"`
	Files       []*FileDetail       `json:"files,omitempty"`
}

// 文本证据信息
type TextDetail struct {
	Field         *string       `json:"field,omitempty"`
	DataId        *string       `json:"dataId,omitempty"`
	CensorResult  *int          `json:"censorResult,omitempty"`
	ReviseContent *string       `json:"reviseContent,omitempty"`
	Reasons       []*TextReason `json:"reasons,omitempty"`
}

// 图片证据信息
type ImageDetail struct {
	Field        *string        `json:"field,omitempty"`
	DataId       *string        `json:"dataId,omitempty"`
	CensorResult *int           `json:"censorResult,omitempty"`
	Reasons      []*ImageReason `json:"reasons,omitempty"`
}

// 视频证据信息
type VideoDetail struct {
	Field        *string             `json:"field,omitempty"`
	DataId       *string             `json:"dataId,omitempty"`
	CensorResult *int                `json:"censorResult,omitempty"`
	Reasons      []*AudioVideoReason `json:"reasons,omitempty"`
}

// 音频证据信息
type AudioDetail struct {
	Field        *string             `json:"field,omitempty"`
	DataId       *string             `json:"dataId,omitempty"`
	CensorResult *int                `json:"censorResult,omitempty"`
	Reasons      []*AudioVideoReason `json:"reasons,omitempty"`
}

// 音视频证据信息
type AudioVideoDetail struct {
	Field        *string             `json:"field,omitempty"`
	DataId       *string             `json:"dataId,omitempty"`
	CensorResult *int                `json:"censorResult,omitempty"`
	Videos       []*AudioVideoReason `json:"videos,omitempty"`
	Audios       []*AudioVideoReason `json:"audios,omitempty"`
}

// 文件证据信息
type FileDetail struct {
	Field        *string        `json:"field,omitempty"`
	DataId       *string        `json:"dataId,omitempty"`
	CensorResult *int           `json:"censorResult,omitempty"`
	Texts        []*TextReason  `json:"texts,omitempty"`
	Images       []*ImageReason `json:"images,omitempty"`
}

// 文本原因信息
type TextReason struct {
	Text         *string        `json:"text,omitempty"`
	Reason       *string        `json:"reason,omitempty"`
	CensorLabels []*CensorLabel `json:"censorLabels,omitempty"`
}

// 图片原因信息
type ImageReason struct {
	Url          *string        `json:"url,omitempty"`
	Reason       *string        `json:"reason,omitempty"`
	CensorLabels []*CensorLabel `json:"censorLabels,omitempty"`
	// 图片标注信息
	DetailMarks []*DetailMark `json:"detailMarks,omitempty"`
}

type DetailMark struct {
	// 标注描述
	Desc *string `json:"desc,omitempty"`
	// 标注标签
	CensorLabels []*CensorLabel `json:"censorLabels,omitempty"`
	// 标注位置
	Position []*Position `json:"position,omitempty"`
}

// 位置信息
type Position struct {
	// x轴坐标
	X *float32 `json:"x,omitempty"`
	Y *float32 `json:"y,omitempty"`
}

// 音视频原因信息
type AudioVideoReason struct {
	URL          *string        `json:"url,omitempty"`
	StartTime    *int64         `json:"startTime,omitempty"`
	EndTime      *int64         `json:"endTime,omitempty"`
	Reason       *string        `json:"reason,omitempty"`
	CensorLabels []*CensorLabel `json:"censorLabels,omitempty"`
}

// 审核标签信息
type CensorLabel struct {
	Code       			*string `json:"code,omitempty"`
	Name       			*string `json:"name,omitempty"`
	CustomCode 			*string `json:"customCode,omitempty"`
	Desc       			*string `json:"desc,omitempty"`
	ParentLabelId       *string `json:"parentLabelId,omitempty"`
    Depth       		*int    `json:"depth,omitempty"`
}
