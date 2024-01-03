package request

import "github.com/yidun/yidun-golang-sdk/yidun/core/types"

// VideoSolutionQueryTaskV1Response 视频解决方案查询任务响应
type VideoSolutionQueryTaskV1Response struct {
	*types.CommonResponse
	Result *[]VideoSolutionQueryTaskV1Result `json:"result,omitempty"`
}

// VideoSolutionQueryTaskV1Result 视频解决方案查询任务结果
type VideoSolutionQueryTaskV1Result struct {
	// 数据状态 0-检测完成或检测失败 20-非7天内数据 30-数据不存在 40-检测中
	Status int `json:"status,omitempty"`
	// 调用音视频检测时传递的dataId字段
	DataId *string `json:"dataId,omitempty"`
	// 音视频数据请求标识，可以根据该标识查询音视频数据最新结果
	TaskId *string `json:"taskId,omitempty"`
	// 调用音视频检测时传递的callback字段
	Callback *string `json:"callback,omitempty"`
	// 检测结果, 1:正常 2:异常 3:疑似
	Result *int `json:"result,omitempty"`
	// 检测状态 0-检测中 1-检测成功 2-检测失败
	CheckStatus *int `json:"checkStatus,omitempty"`
	// 审核来源，0：易盾人审，1：客户人审，2：易盾机审
	CensorSource *int `json:"censorSource,omitempty"`
	// 机器检测结束时间，毫秒单位的时间戳形式 <br>（2021.1.1 01:00转换时间戳1609434000000）
	CheckTime *int64 `json:"checkTime,omitempty"`
	// 机审证据信息
	Evidences *Evidence `json:"evidences,omitempty"`
	// 人审证据信息
	ReviewEvidences *ReviewEvidence `json:"reviewEvidences,omitempty"`
}

type Evidence struct {
	Text   *TextCallbackUnit    `json:"text,omitempty"`   // 标题证据信息
	Images *[]ImageCallbackUnit `json:"images,omitempty"` // 图片证据信息
	Audio  *AudioCallbackUnit   `json:"audio,omitempty"`  // 语音证据信息
	Video  *VideoCallbackUnit   `json:"video,omitempty"`  // 视频证据信息
}
type TextCallbackUnit struct {
	Action        *int         `json:"action,omitempty"`        // 检测结果，0：通过，1：嫌疑，2：不通过
	TaskId        *string      `json:"taskId,omitempty"`        // 唯一标识
	SuggestReason *int         `json:"suggestReason,omitempty"` // Not commented in original code
	Labels        []*LabelInfo `json:"labels,omitempty"`        // 分类信息 详细数据与文本在线检测 labels 数据结构一致
}

type LabelInfo struct {
	Label     *int            `json:"label,omitempty"`     // Not commented in original code
	Level     *int            `json:"level,omitempty"`     // Not commented in original code
	Details   *TextHintInfo   `json:"details,omitempty"`   // Not commented in original code
	SubLabels []*TextSubLabel `json:"subLabels,omitempty"` // 二级分类标签
}

type TextHintInfo struct {
	Hint *[]string `json:"hint,omitempty"` // Not commented in original code
}

type TextSubLabel struct {
	SubLabel *string `json:"subLabel,omitempty"` // 值定义见 SpamType
}

type ImageCallbackUnit struct {
	Name          *string               `json:"name,omitempty"`
	TaskId        *string               `json:"taskId,omitempty"`
	SuggestReason *int                  `json:"suggestReason,omitempty"`
	Labels        []*CallbackImageLabel `json:"labels,omitempty"`
}

type CallbackImageLabel struct {
	Label     *int             `json:"label,omitempty"`
	Level     *int             `json:"level,omitempty"`
	Rate      *float32         `json:"rate,omitempty"`
	SubLabels []*ImageSubLabel `json:"subLabels,omitempty"`
}

type ImageSubLabel struct {
	SubLabel *int             `json:"subLabel,omitempty"` // 对外的图片label
	Rate     *float32         `json:"rate,omitempty"`     // 判断结果，0-正常，1-不确定，2-确定
	Details  *SubLabelDetails `json:"details,omitempty"`  // 得分，范围为0到1
}

type SubLabelDetails struct {
	HitInfos         *[]string          `json:"hitInfos,omitempty"`         // 命中信息，包含规则，敏感词，图片模型结果
	AnticheatInfo    *AnticheatInfo     `json:"anticheatInfo,omitempty"`    // 反作弊结果
	ImageTagInfos    []*ImageTagInfo    `json:"imageTagInfos,omitempty"`    // 图片命中标签信息
	ImageListInfos   []*ImageListInfo   `json:"imageListInfos,omitempty"`   // 命中用户添加图片名单信息
	HitLocationInfos []*HitLocationInfo `json:"hitLocationInfos,omitempty"` // 命中内容证据信息
}

type AnticheatInfo struct {
	HitType *int `json:"hitType,omitempty"`
}

type ImageTagInfo struct {
	TagName  *string `json:"tagName,omitempty"`
	TagGroup *string `json:"tagGroup,omitempty"`
}

type ImageListInfo struct {
	Type     *int    `json:"type,omitempty"` // 1为自定义图片名单，2为自定义敏感词
	URL      *string `json:"url,omitempty"`
	HitCount *int    `json:"hitCount,omitempty"`
	Word     *string `json:"word,omitempty"`
}

type HitLocationInfo struct {
	HitInfo *string  `json:"hitInfo,omitempty"` // 命中信息
	X1      *float32 `json:"x1,omitempty"`      // 坐标左上一个 右下一个
	Y1      *float32 `json:"y1,omitempty"`
	X2      *float32 `json:"x2,omitempty"`
	Y2      *float32 `json:"y2,omitempty"`
}
type AudioCallbackUnit struct {
	Action        *int                 `json:"action,omitempty"`
	AsrStatus     *int                 `json:"asrStatus,omitempty"` // 语音翻译状态
	AsrResult     *int                 `json:"asrResult,omitempty"` // 语音翻译失败错误码 @see VideoSolutionAudioAsrResult
	TaskId        *string              `json:"taskId,omitempty"`
	SuggestReason *int                 `json:"suggestReason,omitempty"`
	Labels        []*LabelInfoResponse `json:"labels,omitempty"`
	Segments      []*AudioDataInfo     `json:"segments,omitempty"`
}

type LabelInfoResponse struct {
	Label     *int      `json:"label,omitempty"`
	Level     *int      `json:"level,omitempty"`
	SubLabels *[]string `json:"subLabels,omitempty"`
	Details   *HintInfo `json:"details,omitempty"`
}

type HintInfo struct {
	Hint    []*HintDetail `json:"hint,omitempty"`
	HitType *int          `json:"hitType,omitempty"`
}

type HintDetail struct {
	Value    *string         `json:"value,omitempty"`
	Segments []*SegmentsInfo `json:"segments,omitempty"`
}

type SegmentsInfo struct {
	StartTime *int `json:"startTime,omitempty"`
	EndTime   *int `json:"endTime,omitempty"`
}

type AudioDataInfo struct {
	StartTime *int      `json:"startTime,omitempty"`
	EndTime   *int      `json:"endTime,omitempty"`
	Content   *string   `json:"content,omitempty"`
	Label     *int      `json:"label,omitempty"`
	Level     *int      `json:"level,omitempty"`
	HintList  *[]string `json:"hintList,omitempty"`
}
type AudioSubLabel struct {
	SubLabel *string   `json:"subLabel,omitempty"`
	Details  *HintInfo `json:"details,omitempty"`
}

type VideoCallbackUnit struct {
	TaskId        *string          `json:"taskId,omitempty"`
	Status        *int             `json:"status,omitempty"`
	Level         *int             `json:"level,omitempty"`
	SuggestReason *int             `json:"suggestReason,omitempty"`
	Evidences     []*VideoEvidence `json:"evidences,omitempty"`
}

type VideoEvidence struct {
	Type      *int                  `json:"type,omitempty"`
	URL       *string               `json:"url,omitempty"`
	BeginTime *int64                `json:"beginTime,omitempty"`
	EndTime   *int64                `json:"endTime,omitempty"`
	Labels    []*CallbackImageLabel `json:"labels,omitempty"`
}

type ReviewEvidence struct {
	Reason *string               `json:"reason,omitempty"` // 人工判定原因
	Detail *ReviewEvidenceDetail `json:"detail,omitempty"` // 详细人审证据信息
}

type ReviewEvidenceDetail struct {
	Text  []*TextCensorEvidenceCallback  `json:"text,omitempty"`  // 文本证据信息
	Image []*ImageCensorEvidenceCallback `json:"image,omitempty"` // 图片证据信息
	Audio []*AudioCensorEvidenceCallback `json:"audio,omitempty"` // 人审音频片段证据信息
	Video []*VideoCensorEvidenceCallback `json:"video,omitempty"` // 人审视频截图证据信息
}

type TextCensorEvidenceCallback struct {
	Word        *string `json:"word,omitempty"`        // 证据信息文本
	Description *string `json:"description,omitempty"` // 加入原因
}

type ImageCensorEvidenceCallback struct {
	URL         *string `json:"url,omitempty"`         // 证据截图 url
	Description *string `json:"description,omitempty"` // 加入原因
}

type AudioCensorEvidenceCallback struct {
	StartTime   *int64  `json:"startTime,omitempty"`   // 音频断句证据开始时间戳，单位秒
	EndTime     *int64  `json:"endTime,omitempty"`     // 音频断句证据结束时间戳，单位秒
	Description *string `json:"description,omitempty"` // 加入原因
}

type VideoCensorEvidenceCallback struct {
	URL         *string `json:"url,omitempty"`         // 证据截图url
	StartTime   *int64  `json:"startTime,omitempty"`   // 证据开始时间戳，单位毫秒
	EndTime     *int64  `json:"endTime,omitempty"`     // 证据结束时间戳，单位毫秒
	Description *string `json:"description,omitempty"` // 加入原因
}
