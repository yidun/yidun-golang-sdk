package response

import (
	videoOCR "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/video/callback/v4/response"
)

type LiveVideoSolutionCallbackV3Result struct {
	Antispam *LiveVideoSolutionCallbackAntispamV3Result `json:"antispam,omitempty"`
	Ocr      *LiveDataCallbackOcrUnitV4                 `json:"ocr,omitempty"`
	Discern  *LiveDataCallbackDiscernUnitV4             `json:"discern,omitempty"`
	Logo     *LiveDataCallbackLogoUnitV4                `json:"logo,omitempty"`
	Voice    *LiveAudioVoiceCallbackRespV4              `json:"voice,omitempty"` // 人声属性识别结果
	Quality  *LiveCallbackQualityUnitV4                 `json:"quality,omitempty"`
	Face     *LiveDataCallbackFaceUnitV4                `json:"face,omitempty"`
	Language *LiveAudioLanguageCallbackRespV3           `json:"language,omitempty"` // 语种识别结果
	ASR      *LiveAudioAsrContentCallbackRespV3         `json:"asr,omitempty"`      // asr
}

type LiveCallbackQualityUnitV4 struct {
	Video *LiveDataCallbackQualityUnitV4  `json:"video,omitempty"`
	Audio *LiveAudioQualityCallbackRespV4 `json:"audio,omitempty"`
}

// LiveVideoSolutionCallbackAntispamV3Result has been simplified for brevity
type LiveVideoSolutionCallbackAntispamV3Result struct {
	TaskID          *string                           `json:"taskId,omitempty"`
	Callback        *string                           `json:"callback,omitempty"`
	DataID          *string                           `json:"dataId,omitempty"`
	Status          *int                              `json:"status,omitempty"`
	Time            *int64                            `json:"time,omitempty"`
	CensorSource    *int                              `json:"censorSource,omitempty"`
	OperationSource    *int                              `json:"operationSource,omitempty"`
	FailureReason   *int                              `json:"failureReason,omitempty"` // 检测失败原因，当检测失败时返回，1：下载失败，2：直播流不存在，3：解析失败，4：格式错误
	Evidences       *LiveSolutionDataCallbackResult   `json:"evidences,omitempty"`
	ReviewEvidences *LiveSolutionCensorCallbackResult `json:"reviewEvidences,omitempty"`
	RiskLevel       *int                              `json:"riskLevel,omitempty"`
	RiskScore       *int                              `json:"riskScore,omitempty"`
	Duration        *int64                            `json:"duration,omitempty"`
	BillDuration    *int64                            `json:"billDuration,omitempty"`
	CensorExtension *CensorExtensionResult            `json:"censorExtension,omitempty"`
}

// LiveSolutionDataCallbackResult includes Audio and Video information.
// Other nested structs will follow the same pattern.
type LiveSolutionDataCallbackResult struct {
	Audio *LiveAudioCallbackUnitRespV4 `json:"audio,omitempty"`
	Video *MachineEvidences            `json:"video,omitempty"`
}

type SegmentsInfo struct {
	Label     *int             `json:"label,omitempty"`
	Level     *int             `json:"level,omitempty"`
	SubLabels *[]AudioSubLabel `json:"subLabels,omitempty"`
}

type AudioSubLabel struct {
	Details  *HintInfo `json:"details,omitempty"`
	SubLabel *string   `json:"subLabel,omitempty"`
}

type HintInfo struct {
	Evidence *string `json:"evidence,omitempty"`
}

type MachineEvidences struct {
	Evidence *VideoEvidence        `json:"evidence,omitempty"`
	Labels   *[]CallbackImageLabel `json:"labels,omitempty"`
}

type CallbackImageLabel struct {
	Label     *int                   `json:"label,omitempty"`
	Level     *int                   `json:"level,omitempty"`
	Rate      *float32               `json:"rate,omitempty"`
	SubLabels *[]ImageV5SubLabelResp `json:"subLabels,omitempty"`
}

type VideoEvidence struct {
	Suggestion *int              `json:"suggestion,omitempty"`
	Label     *int                `json:"label,omitempty"`
    SecondLabel      *string      `json:"secondLabel,omitempty"`
    ThirdLabel       *string      `json:"thirdLabel,omitempty"`
	Type       *int              `json:"type,omitempty"`
	URL        *string           `json:"url,omitempty"`
	BeginTime  *int64            `json:"beginTime,omitempty"`
	EndTime    *int64            `json:"endTime,omitempty"`
	SpeakerID  *string           `json:"speakerId,omitempty"`
	FrontPics  *[]RelatedPicInfo `json:"frontPics,omitempty"`
	PictureID  *string           `json:"pictureId,omitempty"`
	SpeakerRiskScore *int        `json:"speakerRiskScore,omitempty"`
	PublicOpinionInfo  *string   `json:"publicOpinionInfo,omitempty"`
	RiskDescription    *string   `json:"riskDescription,omitempty"`
}

type RelatedPicInfo struct {
	URL *string `json:"url,omitempty"`
}

// Other nested structs...

type LiveSolutionCensorCallbackResult struct {
	Action           *int               `json:"action,omitempty"`
	ActionTime       *int64             `json:"actionTime,omitempty"`
	Label            *int               `json:"label,omitempty"`
	Detail           *string            `json:"detail,omitempty"`
	WarnCount        *int               `json:"warnCount,omitempty"`
	CensorAccount    *string            `json:"censorAccount,omitempty"`
	Evidence         *[]Evidence        `json:"evidence,omitempty"`
	CensorLabels     *[]CensorLabelInfo `json:"censorLabels,omitempty"`
	Pictures         *[]ReviewPicture   `json:"pictures,omitempty"`
	Segments         *[]ReviewSegment   `json:"segments,omitempty"`
	FrontAudiSegment *FrontAudioSegment `json:"frontAudioSegment,omitempty"`
}

// CensorLabelInfo represents censor label information.
type CensorLabelInfo struct {
	Code       *string `json:"code,omitempty"`       // The code of the label
	Desc       *string `json:"desc,omitempty"`       // The description of the label
	CustomCode *string `json:"customCode,omitempty"` // The custom code of the label
	Name       *string `json:"name,omitempty"`       // The name of the label
}

type Evidence struct {
	Snapshot *string `json:"snapshot,omitempty"`
}

type ReviewPicture struct {
	URL *string `json:"url,omitempty"`
}

type ReviewSegment struct {
	StartTime *int64  `json:"startTime,omitempty"`
	EndTime   *int64  `json:"endTime,omitempty"`
	Content   *string `json:"content,omitempty"`
}

type FrontAudioSegment struct {
	URL     *string `json:"url,omitempty"`
	Content *string `json:"content,omitempty"`
}

type LiveDataCallbackOcrUnitV4 struct {
	TaskID    *string                          `json:"taskId,omitempty"`
	DataID    *string                          `json:"dataId,omitempty"`
	SpeakerID *string                          `json:"speakerId,omitempty"`
	BeginTime *int64                           `json:"beginTime,omitempty"`
	EndTime   *int64                           `json:"endTime,omitempty"`
	Height    *int                             `json:"height,omitempty"`
	Width     *int                             `json:"width,omitempty"`
	URL       *string                          `json:"url,omitempty"`
	Details   *[]videoOCR.VideoDataOcrDetailBo `json:"details,omitempty"`
	PictureID *string                          `json:"pictureId,omitempty"`
}

// LiveDataCallbackDiscernUnitV4 represents discern unit for live data callback.
type LiveDataCallbackDiscernUnitV4 struct {
	TaskID    *string                   `json:"taskId,omitempty"`    // The ID of the task
	DataID    *string                   `json:"dataId,omitempty"`    // The ID of the data
	SpeakerID *string                   `json:"speakerId,omitempty"` // The ID of the speaker
	BeginTime *int64                    `json:"beginTime,omitempty"` // The begin time of the task
	EndTime   *int64                    `json:"endTime,omitempty"`   // The end time of the task
	Details   *[]LiveVideoDiscernItemBo `json:"details,omitempty"`   // The details of the discern task
	PictureID *string                   `json:"pictureId,omitempty"` // The ID of the picture
}

// LiveDataCallbackLogoUnitV4 represents logo unit for live data callback.
type LiveDataCallbackLogoUnitV4 struct {
	TaskID    *string                `json:"taskId,omitempty"`    // The ID of the task
	DataID    *string                `json:"dataId,omitempty"`    // The ID of the data
	SpeakerID *string                `json:"speakerId,omitempty"` // The ID of the speaker
	BeginTime *int64                 `json:"beginTime,omitempty"` // The begin time of the task
	EndTime   *int64                 `json:"endTime,omitempty"`   // The end time of the task
	Details   *[]LiveVideoLogoItemBo `json:"details,omitempty"`   // The details of the logo task
	PictureID *string                `json:"pictureId,omitempty"` // The ID of the picture
}

// LiveVideoLogoItemBo represents logo item in live video.
type LiveVideoLogoItemBo struct {
	LogoName  *string  `json:"logoName,omitempty"`  // The name of the logo
	X1        *float32 `json:"x1,omitempty"`        // The x-coordinate of the top left corner
	Y1        *float32 `json:"y1,omitempty"`        // The y-coordinate of the top left corner
	X2        *float32 `json:"x2,omitempty"`        // The x-coordinate of the bottom right corner
	Y2        *float32 `json:"y2,omitempty"`        // The y-coordinate of the bottom right corner
	Rate      *float32 `json:"rate,omitempty"`      // The sublabel score
	SizeRatio *string  `json:"sizeRatio,omitempty"` // The size ratio of the logo
}

type LiveAudioVoiceCallbackRespV4 struct {
	TaskId           *string  `json:"taskId,omitempty"` // 音频uuid
	StartTime        *int64   `json:"startTime,omitempty"`
	EndTime          *int64   `json:"endTime,omitempty"`
	DataId           *string  `json:"dataId,omitempty"`
	MainAgeGroup     *string  `json:"mainAgeGroup,omitempty"` // 年龄段类型
	MainAgeGroupRate *float64 `json:"mainAgeGroupRate,omitempty"`
	MainGender       *string  `json:"mainGender,omitempty"` // 音频性别建议值，male/female
	Underage         *int     `json:"underage,omitempty"`
	Callback         *string  `json:"callback,omitempty"`
	SegmentId        *string  `json:"segmentId,omitempty"`
	SpeakerId        *string  `json:"speakerId,omitempty"`
	Url              *string  `json:"url,omitempty"` // 音频url
}

// LiveDataCallbackFaceUnitV4 represents face unit for live data callback.
type LiveDataCallbackFaceUnitV4 struct {
	TaskID    *string                 `json:"taskId,omitempty"`    // The ID of the task
	DataID    *string                 `json:"dataId,omitempty"`    // The ID of the data
	SpeakerID *string                 `json:"speakerId,omitempty"` // The ID of the speaker
	BeginTime *int64                  `json:"beginTime,omitempty"` // The begin time of the task
	EndTime   *int64                  `json:"endTime,omitempty"`   // The end time of the task
	URL       *string                 `json:"url,omitempty"`       // The URL of the video
	Details   *[]LiveDataFaceDetailBo `json:"details,omitempty"`   // The details of the face task
	PictureID *string                 `json:"pictureId,omitempty"` // The ID of the picture
}

// LiveDataFaceDetailBo represents face detail in live data.
type LiveDataFaceDetailBo struct {
	FaceNumber   *int               `json:"faceNumber,omitempty"`   // The number of faces
	Deepfake     *int               `json:"deepfake,omitempty"`     // Whether there is a deepfake face
	FaceContents *[]FaceLineContent `json:"faceContents,omitempty"` // The contents of the face line
}

// FaceLineContent represents content of face line.
type FaceLineContent struct {
	Name        *string  `json:"name,omitempty"`        // The name of the character
	X1          *float32 `json:"x1,omitempty"`          // The x-coordinate of the top left corner
	Y1          *float32 `json:"y1,omitempty"`          // The y-coordinate of the top left corner
	X2          *float32 `json:"x2,omitempty"`          // The x-coordinate of the bottom right corner
	Y2          *float32 `json:"y2,omitempty"`          // The y-coordinate of the bottom right corner
	Type        *string  `json:"type,omitempty"`        // The type of the face
	Category    *string  `json:"category,omitempty"`    // The category of the face
	Gender      *string  `json:"gender,omitempty"`      // The gender of the character
	Age         *int     `json:"age,omitempty"`         // The age of the character
	Underage    *string  `json:"underage,omitempty"`    // The underage of the character
	SizeRatio   *string  `json:"sizeRatio,omitempty"`   // The size ratio of the face
	BeautyScore *float32 `json:"beautyScore,omitempty"` // The beauty score of the face
	Expression  *string  `json:"expression,omitempty"`  // The expression of the character
	MaskType    *string  `json:"maskType,omitempty"`    // The type of the mask on the face
	MaskScore   *float32 `json:"maskScore,omitempty"`   // The score of the mask on the face
	Glasses     *string  `json:"glasses,omitempty"`     // Whether the character is wearing glasses
	GrowthStage *string  `json:"growthStage,omitempty"` // The growth stage of the character
}

type LiveAudioLanguageCallbackRespV3 struct {
	TaskId    *string `json:"taskId,omitempty"`
	StartTime *int64  `json:"startTime,omitempty"`
	EndTime   *int64  `json:"endTime,omitempty"`
	Content   *string `json:"content,omitempty"`
	Callback  *string `json:"callback,omitempty"`
	SegmentId *string `json:"segmentId,omitempty"`
	SpeakerId *string `json:"speakerId,omitempty"`
}

// LiveAudioAsrContentCallbackRespV3 represents audio asr content for live data callback.
type LiveAudioAsrContentCallbackRespV3 struct {
	TaskID    *string `json:"taskId,omitempty"`    // The ID of the task
	StartTime *int64  `json:"startTime,omitempty"` // The start time of the task
	EndTime   *int64  `json:"endTime,omitempty"`   // The end time of the task
	Content   *string `json:"content,omitempty"`   // The content of the task
	SpeakerID *string `json:"speakerId,omitempty"` // The ID of the speaker
	DataID    *string `json:"dataId,omitempty"`    // The ID of the data
	Callback  *string `json:"callback,omitempty"`  // The callback of the task
	URL       *string `json:"url,omitempty"`       // The URL of the task
}

type LiveDataCallbackQualityUnitV4 struct {
	TaskId    *string                `json:"taskId,omitempty"`
	DataId    *string                `json:"dataId,omitempty"`
	SpeakerId *string                `json:"speakerId,omitempty"`
	BeginTime *int64                 `json:"beginTime,omitempty"`
	EndTime   *int64                 `json:"endTime,omitempty"`
	Details   []*LiveDataQualityResp `json:"details,omitempty"`
	PictureId *string                `json:"pictureId,omitempty"`
}

type LiveAudioQualityCallbackRespV4 struct {
	TaskId    *string                  `json:"taskId,omitempty"` // 音频uuid
	DataId    *string                  `json:"dataId,omitempty"`
	StartTime *int64                   `json:"startTime,omitempty"`
	EndTime   *int64                   `json:"endTime,omitempty"`
	Callback  *string                  `json:"callback,omitempty"`
	SpeakerId *string                  `json:"speakerId,omitempty"`
	Details   *LiveAudioQualityDetails `json:"details,omitempty"`
}

type LiveAudioQualityDetails struct {
	Silent *bool `json:"silent,omitempty"`
}

// LiveAudioCallbackUnitRespV4 represents audio callback unit.
type LiveAudioCallbackUnitRespV4 struct {
	Suggestion       *int                `json:"suggestion,omitempty"`
	SuggestionLevel  *int                `json:"suggestionLevel,omitempty"`
	Label            *int                `json:"label,omitempty"`
    SecondLabel      *string             `json:"secondLabel,omitempty"`
    ThirdLabel       *string             `json:"thirdLabel,omitempty"`
	StartTime        *int64              `json:"startTime,omitempty"`
	EndTime          *int64              `json:"endTime,omitempty"`
	Content          *string             `json:"content,omitempty"`
	Type             *int                `json:"type,omitempty"`
	Rate             *float64            `json:"rate,omitempty"`
	Labels           *[]SegmentsInfoV4   `json:"labels,omitempty"`
	URL              *string             `json:"url,omitempty"`
	SpeakerID        *string             `json:"speakerId,omitempty"`
	SpeakerRiskScore *int                `json:"speakerRiskScore,omitempty"`
	SegmentID        *string             `json:"segmentId,omitempty"`
	FrontSegment     *RelatedSegmentInfo `json:"frontSegment,omitempty"`
}

// RelatedSegmentInfo represents related segment information.
type RelatedSegmentInfo struct {
	Content *string `json:"content,omitempty"`
	URL     *string `json:"url,omitempty"`
}

// SegmentsInfoV4 represents segments information.
type SegmentsInfoV4 struct {
	Label     *int               `json:"label,omitempty"`
	Level     *int               `json:"level,omitempty"`
	Rate      *float64           `json:"rate,omitempty"`
	SubLabels *[]AudioThirdLabel `json:"subLabels,omitempty"`
}

// AudioThirdLabel represents audio third label.
type AudioThirdLabel struct {
	SecondLabel   *string     `json:"secondLabel,omitempty"`
	ThirdLabel    *string     `json:"thirdLabel,omitempty"`
	SubLabelDepth *int        `json:"subLabelDepth,omitempty"`
	Details       *HintInfoV4 `json:"details,omitempty"`
	SubLabel      *string     `json:"subLabel,omitempty"`
	Rate          *float64    `json:"rate,omitempty"`
}

// HintInfoV4 represents hint information.
type HintInfoV4 struct {
	HitInfos   *[]HintDetailV4 `json:"hitInfos,omitempty"`
	HitLeaders *[]LeaderV4     `json:"hitLeaders,omitempty"`
	Keywords   *[]Keywords     `json:"keywords,omitempty"`
	Intent     *IntentUnit     `json:"intent,omitempty"`
}

// IntentUnit represents intent unit.
type IntentUnit struct {
	Result *int     `json:"result,omitempty"`
	Score  *float64 `json:"score,omitempty"`
}

// HintDetailV4 represents hint detail.
type HintDetailV4 struct {
	Value *string `json:"value,omitempty"`
}

// LeaderV4 represents leader.
type LeaderV4 struct {
	Value *string `json:"value,omitempty"`
}

// Keywords represents keywords.
type Keywords struct {
	Word *string `json:"word,omitempty"`
}
type ImageV5SubLabelResp struct {
	SubLabel      *string                   `json:"subLabel,omitempty"`      // 对外的图片label
	SubLabelDepth *int                      `json:"subLabelDepth,omitempty"` // 命中的最终细分类的层级
	SecondLabel   *string                   `json:"secondLabel,omitempty"`   // 二级分类，必返回
	ThirdLabel    *string                   `json:"thirdLabel,omitempty"`    // 三级分类，可能返回
	HitStrategy   *int                      `json:"hitStrategy,omitempty"`   // 命中标识
	Rate          *float32                  `json:"rate,omitempty"`          // 判断结果，0-正常，1-不确定，2-确定
	Details       *videoOCR.SubLabelDetails `json:"details,omitempty"`       // 二级分类命中详情
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

// LiveVideoDiscernItemBo represents discern item for live video.
type LiveVideoDiscernItemBo struct {
	Type        *int     `json:"type,omitempty"`        // The type of the object being recognized
	DiscernName *string  `json:"discernName,omitempty"` // The name of the object being recognized
	Rate        *float32 `json:"rate,omitempty"`        // The score of recognition
	DiscernKey  *string  `json:"discernKey,omitempty"`  // The key of the object being recognized
}

// LiveDataQualityResp represents live data quality response.
type LiveDataQualityResp struct {
	AestheticsRate *float32        `json:"aestheticsRate,omitempty"` // The aesthetics score
	SharpnessRate  *float32        `json:"sharpnessRate,omitempty"`  // The sharpness score
	GrayValue      *int            `json:"grayValue,omitempty"`      // the grayValue of the image
	MetaInfo       *MetaInfo       `json:"metaInfo,omitempty"`       // The basic information of the image
	BoarderInfo    *BoarderInfo    `json:"boarderInfo,omitempty"`    // The border information of the image
	BackgroundInfo *BackgroundInfo `json:"backgroundInfo,omitempty"` // The background information of the image
}

// BackgroundInfo represents background information.
type BackgroundInfo struct {
	PureBackground *bool `json:"pureBackground,omitempty"` // Whether the background is pure
}

// MetaInfo represents metadata of an image.
type MetaInfo struct {
	ByteSize *int64  `json:"byteSize,omitempty"` // The byte size of the image
	Height   *int    `json:"height,omitempty"`   // The height of the image
	Width    *int    `json:"width,omitempty"`    // The width of the image
	Format   *string `json:"format,omitempty"`   // The format of the image
}

// BoarderInfo represents border information.
type BoarderInfo struct {
	Hit    *bool         `json:"hit,omitempty"`    // Whether the border is hit
	Top    *bool         `json:"top,omitempty"`    // The top of the border
	Right  *bool         `json:"right,omitempty"`  // The right of the border
	Bottom *bool         `json:"bottom,omitempty"` // The bottom of the border
	Left   *bool         `json:"left,omitempty"`   // The left of the border
	Color  *BoarderColor `json:"color,omitempty"`  // The color of the border
}

// BoarderColor represents border color.
type BoarderColor struct {
	Top    *int `json:"top,omitempty"`    // The top color of the border
	Right  *int `json:"right,omitempty"`  // The right color of the border
	Bottom *int `json:"bottom,omitempty"` // The bottom color of the border
	Left   *int `json:"left,omitempty"`   // The left color of the border
}

type CensorExtensionResult struct {
	QualityInspectionTaskId *string `json:"qualityInspectionTaskId,omitempty"`
}