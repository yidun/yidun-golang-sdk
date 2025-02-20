package response

import (
	"encoding/json"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/audio/callback/v4/response"
	response2 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/video/callback/v4/response"
)

// VideoSolutionAntispamCallbackV2Response 视频解决方案回调反垃圾响应
type VideoSolutionAntispamCallbackV2Response struct {
	TaskID          *string                                        `json:"taskId,omitempty"`   // 音视频数据请求标识，可以根据该标识查询音视频数据最新结果
	DataID          *string                                        `json:"dataId,omitempty"`   // 调用音视频检测时传递的dataId字段
	Callback        *string                                        `json:"callback,omitempty"` // 调用音视频检测时传递的callback字段
	Suggestion      *int                                           `json:"suggestion,omitempty"`
	Status          *int                                           `json:"status,omitempty"`
	ResultType      *int                                           `json:"resultType,omitempty"`
	CensorRound     *int                                           `json:"censorRound,omitempty"`     // 人审轮次
	Censor          *string                                        `json:"censor,omitempty"`          // 人审操作人
	CensorSource    *int                                           `json:"censorSource,omitempty"`    // 审核来源，0：易盾人审，1：客户人审，2：易盾机审
	CheckTime       *int64                                         `json:"checkTime,omitempty"`       // 机器检测结束时间，毫秒单位的时间戳形式（2021.1.1 01:00转换时间戳1609434000000）
	CensorTime      *int64                                         `json:"censorTime,omitempty"`      // 人工审核完成时间，毫秒单位时间戳(13位)
	Duration        *int64                                         `json:"duration,omitempty"`        // 音视频时长字段，单位毫秒
	DurationMs      *int64                                         `json:"durationMs,omitempty"`      // 音频时长字段，单位毫秒
	Label           *int                                           `json:"label,omitempty"`           // 一级垃圾类型
	SecondLabel     *string                                        `json:"secondLabel,omitempty"`     // 二级垃圾类型
	ThirdLabel      *string                                        `json:"thirdLabel,omitempty"`      // 三级垃圾类型
	PicCount        *int64                                         `json:"picCount,omitempty"`        // 截图数量
	Evidences       *VideoSolutionCallbackEvidenceV2Response       `json:"evidences,omitempty"`       // 机器检测证据信息，接入机器检测时参考证据信息
	SolutionExtra   *VideoSolutionCallbackExtraV2Response          `json:"solutionExtra,omitempty"`   // 音视频解决方案额外信息
	ReviewEvidences *VideoSolutionCallbackReviewEvidenceV2Response `json:"reviewEvidences,omitempty"` // 人审证据信息，接入人工审核后，参考人审证据信息，人审证据信息与机器检测证据信息不共存
	CensorLabels    *[]CensorLabelInfoV2Response                   `json:"censorLabels,omitempty"`
}

type VideoSolutionCallbackExtraV2Response struct {
	FailUnit *VideoSolutionFailUnitResponse `json:"failUnit,omitempty"`
}

type VideoSolutionFailUnitResponse struct {
	Images *[]VideoSolutionImageFailUnitResponse `json:"images,omitempty"`
	Audio  *VideoSolutionTargetFailUnitResponse  `json:"audio,omitempty"`
	Video  *VideoSolutionTargetFailUnitResponse  `json:"video,omitempty"`
}

type VideoSolutionTargetFailUnitResponse struct {
	FailureReason *int `json:"failureReason,omitempty"`
}

type VideoSolutionImageFailUnitResponse struct {
	FailureReason *int    `json:"failureReason,omitempty"`
	Name          *string `json:"name,omitempty"`
}

type VideoSolutionCallbackEvidenceV2Response struct {
	Text   *TextCallbackUnitV2                        `json:"text,omitempty"`
	Images *[]ImageCallbackUnitV2                     `json:"images,omitempty"`
	Audio  *response.AudioAntispamCallbackV4Response  `json:"audio,omitempty"`
	Video  *response2.VideoCallbackAntispamV4Response `json:"video,omitempty"`
}

type VideoSolutionCallbackReviewEvidenceV2Response struct {
	Description *string                                               `json:"description,omitempty"`
	Detail      *string                                               `json:"detail,omitempty"`
	Texts       *[]VideoSolutionCallbackReviewEvidenceTextV2Response  `json:"texts,omitempty"`
	Images      *[]VideoSolutionCallbackReviewEvidenceImageV2Response `json:"images,omitempty"`
	Audios      *[]VideoSolutionCallbackReviewEvidenceAudioV2Response `json:"audios,omitempty"`
	Videos      *[]VideoSolutionCallbackReviewEvidenceVideoV2Response `json:"videos,omitempty"`
}
type VideoSolutionCallbackReviewEvidenceTextV2Response struct {
	Snippet     *string `json:"snippet,omitempty"`
	Description *string `json:"description,omitempty"`
}

type VideoSolutionCallbackReviewEvidenceImageV2Response struct {
	URL         *string `json:"url,omitempty"`
	Description *string `json:"description,omitempty"`
}

type VideoSolutionCallbackReviewEvidenceAudioV2Response struct {
	StartTime   *int64  `json:"startTime,omitempty"`
	EndTime     *int64  `json:"endTime,omitempty"`
	Description *string `json:"description,omitempty"`
}

type VideoSolutionCallbackReviewEvidenceVideoV2Response struct {
	StartTime   *int64  `json:"startTime,omitempty"`
	EndTime     *int64  `json:"endTime,omitempty"`
	URL         *string `json:"url,omitempty"`
	Description *string `json:"description,omitempty"`
}
type CensorLabelInfoV2Response struct {
	Code *string `json:"code,omitempty"`
	Desc *string `json:"desc,omitempty"`
}

type ImageCallbackUnitV2 struct {
	Name       *string               `json:"name,omitempty"`
	DataID     *string               `json:"dataId,omitempty"`
	Suggestion *int                  `json:"suggestion,omitempty"`
	ResultType *int                  `json:"resultType,omitempty"`
	Labels     *[]CallbackImageLabel `json:"labels,omitempty"`
	Status     *int                  `json:"status,omitempty"`
}

type CallbackImageLabel struct {
	Label     *int                 `json:"label,omitempty"`
	Level     *int                 `json:"level,omitempty"`
	Rate      *float32             `json:"rate,omitempty"`
	SubLabels *[]ImageSubLabelResp `json:"subLabels,omitempty"`
}
type ImageSubLabelResp struct {
	SubLabel      *json.Number     `json:"subLabel,omitempty"`      // 对外的图片label
	SubLabelDepth *int             `json:"subLabelDepth,omitempty"` // 命中的最终细分类的层级
	SecondLabel   *string          `json:"secondLabel,omitempty"`   // 二级分类，必返回
	ThirdLabel    *string          `json:"thirdLabel,omitempty"`    // 三级分类，可能返回
	HitStrategy   *int             `json:"hitStrategy,omitempty"`   // 命中标识
	Rate          *float32         `json:"rate,omitempty"`          // 判断结果，0-正常，1-不确定，2-确定
	Details       *SubLabelDetails `json:"details,omitempty"`       // 得分，范围为0到1
}

type SubLabelDetails struct {
	Keywords  []AntispamInfo `json:"keywords,omitempty"`  // 反垃圾自定义敏感词结果
	LibInfos  []AntispamInfo `json:"libInfos,omitempty"`  // 反垃圾自定义图片名单结果
	HitInfos  []AntispamInfo `json:"hitInfos,omitempty"`  // 反垃圾其他命中信息
	Anticheat *AnticheatInfo `json:"anticheat,omitempty"` // 反作弊结果
}

type AnticheatInfo struct {
	HitType *int `json:"hitType,omitempty"`
}

type AntispamInfo struct {
	Word     *string  `json:"word,omitempty"`     // 自定义敏感词
	Entity   *string  `json:"entity,omitempty"`   // 自定义图片名单url
	HitCount *int     `json:"hitCount,omitempty"` // 自定义图片名单命中次数
	Value    *string  `json:"value,omitempty"`
	Group    *string  `json:"group,omitempty"`
	X1       *float32 `json:"x1,omitempty"` // 坐标左上一个 右下一个
	Y1       *float32 `json:"y1,omitempty"`
	X2       *float32 `json:"x2,omitempty"`
	Y2       *float32 `json:"y2,omitempty"`
}
type TextCallbackUnitV2 struct {
	TaskID       *string      `json:"taskId,omitempty"`
	DataID       *string      `json:"dataId,omitempty"`
	Suggestion   *int         `json:"suggestion,omitempty"`
	ResultType   *int         `json:"resultType,omitempty"`
	CensorType   *int         `json:"censorType,omitempty"`
	IsRelatedHit *bool        `json:"isRelatedHit,omitempty"`
	Labels       *[]LabelInfo `json:"labels,omitempty"`
}

type LabelInfo struct {
	Label     *int            `json:"label,omitempty"`
	Level     *int            `json:"level,omitempty"`
	SubLabels *[]TextSubLabel `json:"subLabels,omitempty"`
}

type TextSubLabel struct {
	SubLabel      *string                                `json:"subLabel,omitempty"`
	SubLabelDepth *int                                   `json:"subLabelDepth,omitempty"`
	SecondLabel   *string                                `json:"secondLabel,omitempty"`
	ThirdLabel    *string                                `json:"thirdLabel,omitempty"`
	Details       *TextCheckResultAntispamSubLabelDetail `json:"details,omitempty"`
}

type TextCheckResultAntispamSubLabelDetail struct {
	Keywords  *[]AntispamSubLabelDetailKeyword `json:"keywords,omitempty"`
	LibInfos  *[]AntispamSubLabelDetailLibInfo `json:"libInfos,omitempty"`
	Anticheat *AntispamSubLabelDetailAnticheat `json:"anticheat,omitempty"`
	HitInfos  *[]AntispamSubLabelDetailHitInfo `json:"hitInfos,omitempty"`
}
type AntispamSubLabelDetailKeyword struct {
	Word *string `json:"word,omitempty"`
}

type AntispamSubLabelDetailLibInfo struct {
	Type        *int    `json:"type,omitempty"`
	Entity      *string `json:"entity,omitempty"`
	ReleaseTime *int64  `json:"releaseTime,omitempty"`
}

type AntispamSubLabelDetailAnticheat struct {
	HitType *int `json:"hitType,omitempty"`
}

type AntispamSubLabelDetailHitInfo struct {
	Value     *string                                  `json:"value,omitempty"`
	Positions *[]AntispamSubLabelDetailHitInfoPosition `json:"positions,omitempty"`
}

type AntispamSubLabelDetailHitInfoPosition struct {
	FieldName *string `json:"fieldName,omitempty"`
	StartPos  *int    `json:"startPos,omitempty"`
	EndPos    *int    `json:"endPos,omitempty"`
}
