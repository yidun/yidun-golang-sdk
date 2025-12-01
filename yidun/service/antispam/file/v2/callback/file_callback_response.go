package callback

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/image/v5/check/sync"
)

type FileCallbackV2Response struct {
	*types.CommonResponse
	Result []*FileCallbackV2Resp `json:"result,omitempty"`
}

type FileCallbackV2Resp struct {
	Antispam        *AntispamCheckResult   `json:"antispam,omitempty"`
	ValueAddService *ValueAddServiceResult `json:"valueAddService,omitempty"`
}

// 文档反垃圾结果
type AntispamCheckResult struct {
	DataId          *string               `json:"dataId,omitempty"`
	TaskId          *string               `json:"taskId,omitempty"`
	Callback        *string               `json:"callback,omitempty"`
	Label           *int                  `json:"label,omitempty"`
	SecondLabel     *string               `json:"secondLabel,omitempty"`
	ThirdLabel      *string               `json:"thirdLabel,omitempty"`
	Suggestion      *int                  `json:"suggestion,omitempty"`
	Message         *string               `json:"message,omitempty"`
	ResultType      *int                  `json:"resultType,omitempty"`
	Evidences       *EvidenceResult       `json:"evidences,omitempty"`
	ReviewEvidences *ReviewEvidenceResult `json:"reviewEvidences,omitempty"`
	FailureReason   *int                  `json:"failureReason,omitempty"`
	Details         []*Detail             `json:"details,omitempty"`
	CensorLabels    []*CensorLabel        `json:"censorLabels,omitempty"`
	CensorTime      *int64                `json:"censorTime,omitempty"` //审核完成时间，结果类型是机器结果，代表机器审核完成时间；审核结果是人审结果，代表人审完成时间
	Status          *int                  `json:"status,omitempty"`     //检测状态 3：失败 2：成功
	Md5             *string               `json:"md5,omitempty"`
}

// 压缩文件反垃圾结果，每个detail对应一个压缩包内的子文件
type Detail struct {
	TaskId          *string               `json:"taskId,omitempty"`
	Name            *string               `json:"name,omitempty"`
	FileType        *int                  `json:"fileType,omitempty"`
	Suggestion      *int                  `json:"suggestion,omitempty"`
	Evidences       *EvidenceResult       `json:"evidences,omitempty"`
	ReviewEvidences *ReviewEvidenceResult `json:"reviewEvidences,omitempty"`
	FailureReason   *int                  `json:"failureReason,omitempty"`
	Message         *string               `json:"message,omitempty"`
	Label           *int                  `json:"label,omitempty"`
	SecondLabel     *string               `json:"secondLabel,omitempty"`
	ThirdLabel      *string               `json:"thirdLabel,omitempty"`
	CensorTime      *int64                `json:"censorTime,omitempty"` //压缩文件子文件审核完成时间，结果类型是机器结果，代表机器审核完成时间；审核结果是人审结果，代表人审完成时间
	Status          *int                  `json:"status,omitempty"`     //检测状态 3：失败 2：成功
	Md5             *string               `json:"md5,omitempty"`
}

// 文档反垃圾结果标签
type CensorLabel struct {
	Code       			*string `json:"code,omitempty"`
	Name       			*string `json:"name,omitempty"`
	Desc       			*string `json:"desc,omitempty"`
	CustomCode 			*string `json:"customCode,omitempty"`
	ParentLabelId       *string `json:"parentLabelId,omitempty"`
    Depth       		*int    `json:"depth,omitempty"`
}

// 文档反垃圾人审结果
type ReviewEvidenceResult struct {
	Reason *string          `json:"reason,omitempty"`
	Remark *string          `json:"remark,omitempty"`
	Detail *ReviewEvidences `json:"detail,omitempty"`
}

// 文档反垃圾人审证据
type ReviewEvidences struct {
	Text  []*TextDetail  `json:"text,omitempty"`
	Image []*ImageDetail `json:"image,omitempty"`
}

// 文本反垃圾机器结果详情
type TextDetail struct {
	Text   *string `json:"text,omitempty"`
	Reason *string `json:"reason,omitempty"`
}

// 图片反垃圾机器结果详情
type ImageDetail struct {
	Url    *string `json:"url,omitempty"`
	Reason *string `json:"reason,omitempty"`
}

// 文档反垃圾机器审核结果证据
type EvidenceResult struct {
	Texts  []*TextEvidence  `json:"texts,omitempty"`
	Images []*ImageEvidence `json:"images,omitempty"`
	Audios []*AudioEvidence `json:"audios,omitempty"`
	Videos []*VideoEvidence `json:"videos,omitempty"`
}

// 文本反垃圾机审结果证据
type TextEvidence struct {
	TaskId              *string              `json:"taskId,omitempty"`
	Sequence            *int                 `json:"sequence,omitempty"`
	StartText           *string              `json:"startText,omitempty"`
	EndText             *string              `json:"endText,omitempty"`
	Suggestion          *int                 `json:"suggestion,omitempty"`
	SuggestionRiskLevel *int                 `json:"suggestionRiskLevel,omitempty"`
	PublicOpinionInfo   *string              `json:"publicOpinionInfo,omitempty"`
	Labels              []*TextEvidenceLabel `json:"labels,omitempty"`
	Page                *int                 `json:"page,omitempty"`
}

// 文本反垃圾机审结果证据标签
type TextEvidenceLabel struct {
	Label     *int            `json:"label,omitempty"`
	Level     *int            `json:"level,omitempty"`
	Rate      *float64        `json:"rate,omitempty"`
	SubLabels []*TextSubLabel `json:"subLabels,omitempty"`
}

// 文本反垃圾机审结果二级分类
type TextSubLabel struct {
	SubLabel      *string             `json:"subLabel,omitempty"`
	SubLabelDepth *int                `json:"subLabelDepth,omitempty"`
	SecondLabel   *string             `json:"secondLabel,omitempty"`
	ThirdLabel    *string             `json:"thirdLabel,omitempty"`
	Details       *TextSubLabelDetail `json:"details,omitempty"`
}

// TextSubLabelDetail 文本标签详细信息
type TextSubLabelDetail struct {
	Keywords *[]Keyword  `json:"keywords,omitempty"`
	LibInfos *[]LibInfo  `json:"libInfos,omitempty"`
	HitInfos *[]HintInfo `json:"hitInfos,omitempty"`
	Rules    *[]RuleInfo `json:"rules,omitempty"`
}

// Keyword
type Keyword struct {
	Word *string `json:"word,omitempty"`
}

// LibInfo
type LibInfo struct {
	Type   *int    `json:"type,omitempty"`
	Entity *string `json:"entity,omitempty"`
}

// HintInfo
type HintInfo struct {
	Type      *int                `json:"type,omitempty"` // type
	Value     *string             `json:"value,omitempty"`
	Positions *[]HintInfoPosition `json:"positions,omitempty"`
}

// HintInfoPosition
type HintInfoPosition struct {
	FieldName *string `json:"fieldName,omitempty"`
	StartPos  *int    `json:"startPos,omitempty"`
	EndPos    *int    `json:"endPos,omitempty"`
}

type RuleInfo struct {
	Name *string `json:"name,omitempty"`
}

// ImageEvidence 图片证据
type ImageEvidence struct {
	TaskId              *string               `json:"taskId,omitempty"`
	Sequence            *int                  `json:"sequence,omitempty"`
	ImageUrl            *string               `json:"imageUrl,omitempty"`
	Suggestion          *int                  `json:"suggestion,omitempty"`
	SuggestionRiskLevel *int                  `json:"suggestionRiskLevel,omitempty"`
	PublicOpinionInfo   *string               `json:"publicOpinionInfo,omitempty"`
	Labels              []*ImageEvidenceLabel `json:"labels,omitempty"`
	Page                *int                  `json:"page,omitempty"`
}

// ImageEvidenceLabel 图片证据标签
type ImageEvidenceLabel struct {
	Label     *int             `json:"label,omitempty"`
	Level     *int             `json:"level,omitempty"`
	Rate      *float64         `json:"rate,omitempty"`
	SubLabels []*ImageSubLabel `json:"subLabels,omitempty"`
}

// ImageSubLabel 图片证据子标签
type ImageSubLabel struct {
	SubLabel      *string              `json:"subLabel,omitempty"`
	Rate          *float64             `json:"rate,omitempty"`
	SubLabelDepth *int                 `json:"subLabelDepth,omitempty"`
	SecondLabel   *string              `json:"secondLabel,omitempty"`
	ThirdLabel    *string              `json:"thirdLabel,omitempty"`
	Details       *ImageSubLabelDetail `json:"details,omitempty"`
}

// ImageSubLabelDetail 图片标签详细信息
type ImageSubLabelDetail struct {
	Keywords *[]ImageSubLabelDetailInfo `json:"keywords,omitempty"` // 反垃圾自定义敏感词结果
	LibInfos *[]ImageSubLabelDetailInfo `json:"libInfos,omitempty"` // 反垃圾自定义图片名单结果
	HitInfos *[]ImageSubLabelDetailInfo `json:"hitInfos,omitempty"` // 反垃圾其他命中信息
	Rules    *[]ImageSubLabelDetailInfo `json:"rules,omitempty"`
}

// ImageSubLabelDetailInfo
type ImageSubLabelDetailInfo struct {
	Word     *string  `json:"word,omitempty"`     // 自定义敏感词
	Entity   *string  `json:"entity,omitempty"`   // 自定义图片名单url
	HitCount *int     `json:"hitCount,omitempty"` // 自定义图片名单命中次数
	Value    *string  `json:"value,omitempty"`
	Group    *string  `json:"group,omitempty"`
	X1       *float64 `json:"x1,omitempty"` // 坐标左上
	Y1       *float64 `json:"y1,omitempty"` // 坐标左上
	X2       *float64 `json:"x2,omitempty"` // 坐标右下
	Y2       *float64 `json:"y2,omitempty"` // 坐标右下
	Name     *string  `json:"name,omitempty"`
	Type     *int     `json:"type,omitempty"` // type
}

// AudioEvidence 音频证据
type AudioEvidence struct {
	TaskId                   *string                   `json:"taskId,omitempty"`
	Sequence                 *int                      `json:"sequence,omitempty"`
	AudioUrl                 *string                   `json:"audioUrl,omitempty"`
	Suggestion               *int                      `json:"suggestion,omitempty"`
	Status                   *int                      `json:"status,omitempty"`
	VideoSolutionCheckResult *VideoSolutionCheckResult `json:"videoSolutionCheckResult,omitempty"`
}

type VideoSolutionCheckResult struct {
	Segments *[]AudioEvidenceSegment  `json:"segments,omitempty"`
	Pictures *[]VideoEvidencePictures `json:"pictures,omitempty"`
}

// AudioEvidenceSegment 音频证据分段
type AudioEvidenceSegment struct {
	StartTime  *int64                `json:"startTime,omitempty"`
	EndTime    *int64                `json:"endTime,omitempty"`
	Type       *int                  `json:"type,omitempty"`
	LeaderName *string               `json:"leaderName,omitempty"`
	Content    *string               `json:"content,omitempty"`
	Labels     *[]AudioEvidenceLabel `json:"labels,omitempty"`
}

// AudioEvidenceLabel 音频证据标签
type AudioEvidenceLabel struct {
	Label     *int                     `json:"label,omitempty"`
	Level     *int                     `json:"level,omitempty"`
	SubLabels *[]AudioEvidenceSubLabel `json:"subLabels,omitempty"`
}

// AudioEvidenceSubLabel 音频证据子标签
type AudioEvidenceSubLabel struct {
	SubLabel      *int64               `json:"subLabel,omitempty"`
	SubLabelDepth *int                 `json:"subLabelDepth,omitempty"`
	SecondLabel   *string              `json:"secondLabel,omitempty"`
	ThirdLabel    *string              `json:"thirdLabel,omitempty"`
	Details       *AudioSubLabelDetail `json:"details,omitempty"`
}

// AudioSubLabelDetail 音频标签详细信息
type AudioSubLabelDetail struct {
	Keywords *[]AudioSubLabelKeyword  `json:"keywords,omitempty"` // 反垃圾自定义敏感词结果
	LibInfos *[]AudioSubLabelLibInfo  `json:"libInfos,omitempty"` // 反垃圾自定义图片名单结果
	HitInfos *[]AudioSubLabelHitInfo  `json:"hitInfos,omitempty"` // 反垃圾其他命中信息
	Rules    *[]AudioSubLabelRuleInfo `json:"rules,omitempty"`
}

// AudioSubLabelKeyword 自定义添加敏感词
type AudioSubLabelKeyword struct {
	Word *string `json:"word,omitempty"`
}

// AudioSubLabelLibInfo
type AudioSubLabelLibInfo struct {
	ListType *int    `json:"listType,omitempty"` // 名单类型
	Entity   *string `json:"entity,omitempty"`   // 名单内容
}

// AudioSubLabelHitInfo
type AudioSubLabelHitInfo struct {
	Type     *int    `json:"type,omitempty"`     // type
	Value    *string `json:"value,omitempty"`    // 命中的敏感词或者声纹检测的分值
	SongName *string `json:"songName,omitempty"` // 命中的涉政歌曲名称
}

type AudioSubLabelRuleInfo struct {
	Name *string `json:"name,omitempty"`
}

// VideoEvidence 视频证据
type VideoEvidence struct {
	TaskID                   *string                   `json:"taskId,omitempty"`
	Sequence                 *int                      `json:"sequence,omitempty"`
	VideoURL                 *string                   `json:"videoUrl,omitempty"`
	Suggestion               *int                      `json:"suggestion,omitempty"`
	Status                   *int                      `json:"status,omitempty"`
	VideoSolutionCheckResult *VideoSolutionCheckResult `json:"videoSolutionCheckResult,omitempty"`
}

// VideoEvidencePictures 视频证据图片
type VideoEvidencePictures struct {
	Type                *int                  `json:"type,omitempty"`
	URL                 *string               `json:"url,omitempty"`
	StartTime           *int64                `json:"startTime,omitempty"`
	EndTime             *int64                `json:"endTime,omitempty"`
	SuggestionRiskLevel *int                  `json:"suggestionRiskLevel,omitempty"`
	PublicOpinionInfo   *string               `json:"publicOpinionInfo,omitempty"`
	FrontPics           *[]VideoEvidencePics  `json:"frontPics,omitempty"`
	BackPics            *[]VideoEvidencePics  `json:"backPics,omitempty"`
	Labels              *[]VideoEvidenceLabel `json:"labels,omitempty"`
}

// VideoEvidenceLabel 视频证据图片nos地址
type VideoEvidencePics struct {
	URL *string `json:"url,omitempty"`
}

// VideoEvidenceLabel 视频证据标签
type VideoEvidenceLabel struct {
	Label     *int                     `json:"label,omitempty"`
	Level     *int                     `json:"level,omitempty"`
	Rate      *float64                 `json:"rate,omitempty"`
	SubLabels []*VideoEvidenceSubLabel `json:"subLabels,omitempty"`
}

// VideoEvidenceSubLabel 视频证据二级分类
type VideoEvidenceSubLabel struct {
	SubLabel      *int64               `json:"subLabel,omitempty"`
	SubLabelDepth *int                 `json:"subLabelDepth,omitempty"`
	SecondLabel   *string              `json:"secondLabel,omitempty"`
	ThirdLabel    *string              `json:"thirdLabel,omitempty"`
	Details       *VideoSubLabelDetail `json:"details,omitempty"`
}

// VideoSubLabelDetail 音频标签详细信息
type VideoSubLabelDetail struct {
	Keywords *[]VideoSubLabelKeyword  `json:"keywords,omitempty"` // 反垃圾自定义敏感词结果
	LibInfos *[]VideoSubLabelLibInfo  `json:"libInfos,omitempty"` // 反垃圾自定义图片名单结果
	HitInfos *[]VideoSubLabelHitInfo  `json:"hitInfos,omitempty"` // 反垃圾其他命中信息
	Rules    *[]VideoSubLabelRuleInfo `json:"rules,omitempty"`    // 自定义规则
}

// VideoSubLabelKeyword 自定义添加敏感词
type VideoSubLabelKeyword struct {
	Word *string  `json:"word,omitempty"`
	X1   *float64 `json:"x1,omitempty"` // 坐标左上
	Y1   *float64 `json:"y1,omitempty"` // 坐标左上
	X2   *float64 `json:"x2,omitempty"` // 坐标右下
	Y2   *float64 `json:"y2,omitempty"` // 坐标右下
}

// VideoSubLabelLibInfo
type VideoSubLabelLibInfo struct {
	Entity   *string `json:"entity,omitempty"`   // 该图片命中自定义图片名单对应原始添加的根源图片url
	HitCount *int    `json:"hitCount,omitempty"` // 历史针对该数据源图片命中所有次数
	Value    *string `json:"value,omitempty"`
	Group    *string `json:"group,omitempty"`
}

// VideoSubLabelHitInfo
type VideoSubLabelHitInfo struct {
	Value *string  `json:"value,omitempty"`
	Group *string  `json:"group,omitempty"`
	X1    *float64 `json:"x1,omitempty"`   // 坐标左上
	Y1    *float64 `json:"y1,omitempty"`   // 坐标左上
	X2    *float64 `json:"x2,omitempty"`   // 坐标右下
	Y2    *float64 `json:"y2,omitempty"`   // 坐标右下
	Type  *int     `json:"type,omitempty"` // type
}

type VideoSubLabelRuleInfo struct {
	Name *string `json:"name,omitempty"`
}

type ValueAddServiceResult struct {
	Ocr *Ocr `json:"ocr,omitempty"`
}

type Ocr struct {
	Images *[]sync.ImageV5OcrResp `json:"images,omitempty"`
}
