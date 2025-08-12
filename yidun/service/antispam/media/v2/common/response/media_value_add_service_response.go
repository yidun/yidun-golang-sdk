package response

/**
 * 增值服务结果
 */
type MediaValueAddServiceResponse struct {
	Ocr             *Ocr             `json:"ocr,omitempty"`
	Asr             *Asr             `json:"asr,omitempty"`
	Face            *Face            `json:"face,omitempty"`
	ImageQuality    *ImageQuality    `json:"imageQuality,omitempty"`
	Logo            *Logo            `json:"logo,omitempty"`
	ImageDiscern    *ImageDiscern    `json:"imageDiscern,omitempty"`
	EmotionAnalysis *EmotionAnalysis `json:"emotionAnalysis,omitempty"`
	Language        *Language        `json:"language,omitempty"`
	GrammarFix      *GrammarFix      `json:"grammarFix,omitempty"`
	Aigc            *MediaAigc       `json:"aigc,omitempty"`
}

type Ocr struct {
	Images *[]OcrImageDetail `json:"images,omitempty"`
}

type Asr struct {
	Audios      *[]AsrAudio      `json:"audios,omitempty"`
	Audiovideos *[]AsrAudiovideo `json:"audiovideos,omitempty"`
}

type Face struct {
	Images *[]FaceDetail `json:"images,omitempty"`
}

type ImageQuality struct {
	Images *[]ImageQualityDetail `json:"images,omitempty"`
}

type Logo struct {
	Images *[]LogoDetail `json:"images,omitempty"`
}

type ImageDiscern struct {
	Images *[]ImageDiscernDetail `json:"images,omitempty"`
}

type ImageDiscernDetail struct {
	ValueServiceBaseResponse
	Details *[]ImageDiscernImageDetail `json:"details,omitempty"`
}

type ImageDiscernImageDetail struct {
	DiscernName *string  `json:"discernName,omitempty"`
	Rate        *float32 `json:"rate,omitempty"`
	Type        *int     `json:"type,omitempty"`
}

type EmotionAnalysis struct {
	Texts *[]EmotionAnalysisDetail `json:"texts,omitempty"`
}

type Language struct {
	Texts *[]LanguageDetail `json:"texts,omitempty"`
}

type GrammarFix struct {
	Texts *[]GrammarValueServiceBaseUnit `json:"texts,omitempty"`
}

// ocr识别结果
type OcrImageDetail struct {
	ValueServiceBaseResponse
	Width   *int         `json:"width,omitempty"`
	Height  *int         `json:"height,omitempty"`
	Details *[]OcrDetail `json:"details,omitempty"`
}

type ValueServiceBaseResponse struct {
	DataId *string `json:"dataId,omitempty"`
	Field  *string `json:"field,omitempty"`
	TaskId *string `json:"taskId,omitempty"`
}

type AsrAudio struct {
	TaskId  *string           `json:"taskId,omitempty"`
	DataId  *string           `json:"dataId,omitempty"`
	Field   *string           `json:"field,omitempty"`
	Details *[]AsrAudioDetail `json:"details,omitempty"`
}

type AsrAudiovideo struct {
	TaskId  *string                `json:"taskId,omitempty"`
	DataId  *string                `json:"dataId,omitempty"`
	Field   *string                `json:"field,omitempty"`
	Details *[]AsrAudiovideoDetail `json:"details,omitempty"`
}

type AsrAudioDetail struct {
	StartTime *int64  `json:"startTime,omitempty"`
	EndTime   *int64  `json:"endTime,omitempty"`
	Content   *string `json:"content,omitempty"`
}

type AsrAudiovideoDetail struct {
	StartTime *int64  `json:"startTime,omitempty"`
	EndTime   *int64  `json:"endTime,omitempty"`
	Content   *string `json:"content,omitempty"`
}

type OcrDetail struct {
	Content      *string           `json:"content,omitempty"`
	LineContents *[]OcrLineContent `json:"lineContents,omitempty"`
}

type OcrLineContent struct {
	LineContent *string  `json:"lineContent,omitempty"`
	X1          *float64 `json:"x1,omitempty"`
	Y1          *float64 `json:"y1,omitempty"`
	X2          *float64 `json:"x2,omitempty"`
	Y2          *float64 `json:"y2,omitempty"`
	Lang        *string  `json:"lang,omitempty"`
}

type FaceDetail struct {
	ValueServiceBaseResponse
	Details *[]FaceImageDetail `json:"details,omitempty"`
}

/**
 * 人脸信息
 */
type FaceImageDetail struct {
	FaceNumber   *int               `json:"faceNumber,omitempty"`
	FaceContents *[]FaceLineContent `json:"faceContents,omitempty"`
}

/**
 * ocr文本详情
 */
type FaceLineContent struct {
	Name        *string  `json:"name,omitempty"`
	X1          *float64 `json:"x1,omitempty"`
	Y1          *float64 `json:"y1,omitempty"`
	X2          *float64 `json:"x2,omitempty"`
	Y2          *float64 `json:"y2,omitempty"`
	Type        *string  `json:"type,omitempty"`
	Category    *string  `json:"category,omitempty"`
	Gender      *string  `json:"gender,omitempty"`
	Age         *int     `json:"age,omitempty"`
	SizeRatio   *string  `json:"sizeRatio,omitempty"`
	BeautyScore *float64 `json:"beautyScore,omitempty"`
	Expression  *string  `json:"expression,omitempty"`
	MaskType    *string  `json:"maskType,omitempty"`
}

type BackgroundInfo struct {
	PureBackground *bool `json:"pureBackground,omitempty"`
}

type MetaInfo struct {
	ByteSize *int64  `json:"byteSize,omitempty"`
	Height   *int    `json:"height,omitempty"`
	Width    *int    `json:"width,omitempty"`
	Format   *string `json:"format,omitempty"`
}

type BoarderInfo struct {
	Hit    *bool `json:"hit,omitempty"`
	Top    *bool `json:"top,omitempty"`
	Right  *bool `json:"right,omitempty"`
	Bottom *bool `json:"bottom,omitempty"`
	Left   *bool `json:"left,omitempty"`
}

type GrammarValueServiceBaseUnit struct {
	TaskId  *string                           `json:"taskId,omitempty"`
	Details *[]GrammarfixSubmitResponseDetail `json:"details,omitempty"`
	Level   *int                              `json:"level,omitempty"`
}

type LanguageDetail struct {
	ValueServiceBaseResponse
	Details *[]LanguageTextDetail `json:"details,omitempty"`
}

// 语种信息
type LanguageTextDetail struct {
	Type *string `json:"type,omitempty"`
}

type EmotionAnalysisDetail struct {
	ValueServiceBaseResponse
	Details *[]EmotionAnalysisTextDetail `json:"details,omitempty"`
}

type EmotionAnalysisTextDetail struct {
	PositiveProb *float64 `json:"positiveProb,omitempty"`
	NegativeProb *float64 `json:"negativeProb,omitempty"`
	Sentiment    *string  `json:"sentiment,omitempty"`
}

type ImageQualityDetail struct {
	ValueServiceBaseResponse
	Details *[]ImageQualityImageDetail `json:"details,omitempty"`
}

type ImageQualityImageDetail struct {
	// 美观度分数
	AestheticsRate *float32 `json:"aestheticsRate,omitempty"`
	// 清晰度分数
	SharpnessRate *float32 `json:"sharpnessRate,omitempty"`
	// 图片基本信息
	MetaInfo *MetaInfo `json:"metaInfo,omitempty"`
	// 图片边框信息
	BoarderInfo *BoarderInfo `json:"boarderInfo,omitempty"`
	// 背景信息
	BackgroundInfo *BackgroundInfo `json:"backgroundInfo,omitempty"`
}

type LogoDetail struct {
	ValueServiceBaseResponse
	Details *[]LogoImageDetail `json:"details,omitempty"`
}

// logo信息
type LogoImageDetail struct {
	ValueServiceBaseResponse
	// logo信息
	LogoName *string  `json:"logoName,omitempty"`
	X1       *float32 `json:"x1,omitempty"`
	Y1       *float32 `json:"y1,omitempty"`
	X2       *float32 `json:"x2,omitempty"`
	Y2       *float32 `json:"y2,omitempty"`
}

type GrammarfixSubmitResponseDetail struct {
	CorrectContent *string                             `json:"correctContent,omitempty"`
	Fragments      *[]GrammarfixSubmitResponseFragment `json:"fragments,omitempty"`
}

/**
 * 纠错结果详细信息
 */
type GrammarfixSubmitResponseFragment struct {
	OriWord           *string `json:"oriWord,omitempty"`
	CorrectWord       *string `json:"correctWord,omitempty"`
	HeadWord          *string `json:"headWord,omitempty"`
	Label             *string `json:"label,omitempty"`
	SubLabel          *string `json:"subLabel,omitempty"`
	InspectType       *string `json:"inspectType,omitempty"`
	ModifiedType      *int    `json:"modifiedType,omitempty"`
	StartPos          *int    `json:"startPos,omitempty"`
	EndPos            *int    `json:"endPos,omitempty"`
	HeadStartPos      *int    `json:"headStartPos,omitempty"`
	HeadEndPos        *int    `json:"headEndPos,omitempty"`
	CorrectSuggestion *string `json:"correctSuggestion,omitempty"`
}

// 顶层结构体
type MediaAigc struct {
	Images      []*MediaAigcImageValueServiceUnit `json:"images,omitempty"`
	AudioVideos []*MediaAigcVideoValueServiceUnit `json:"audiovideos,omitempty"`
}

// 图片检测单元 ================================
type MediaAigcImageValueServiceUnit struct {
	ValueServiceBaseResponse                    // 继承父类字段
	Details                  []*AigcImageDetail `json:"details,omitempty"`
}

type AigcImageDetail struct {
	IsAigc    bool                    `json:"isAigc,omitempty"`    // 是否AIGC生成
	AigcLevel int                     `json:"aigcLevel,omitempty"` // 0:正常 1:嫌疑 2:确定
	AigcRate  float32                 `json:"aigcRate,omitempty"`  // AIGC概率分数
	Signage   *ImageAigcV5SignageResp `json:"signage,omitempty"`   // 生成标识信息
}

type ImageAigcV5SignageResp struct {
	OvertSignage         int                       `json:"overtSignage,omitempty"`         // 显式标识 0/1
	CovertSignage        int                       `json:"covertSignage,omitempty"`        // 隐式标识 0/1
	CovertSignageDetails *CovertSignageDetailsResp `json:"covertSignageDetails,omitempty"` // 隐式详情
}

type CovertSignageDetailsResp struct {
	Role     int                               `json:"role,omitempty"`     // 1:生成者 2:传播者
	Platform *CovertSignageDetailsPlatformResp `json:"platform,omitempty"` // 平台/个人信息
}

type CovertSignageDetailsPlatformResp struct {
	Type int    `json:"type,omitempty"` // 1:公司 2:个人
	Info string `json:"info,omitempty"` // 公司名或空
}

// 视频检测单元 ================================
type MediaAigcVideoValueServiceUnit struct {
	TaskID   string                `json:"taskId,omitempty"`
	DataID   string                `json:"dataId,omitempty"`
	Field    string                `json:"field,omitempty"`
	Pictures []*MediaAigcVideoUnit `json:"pictures,omitempty"`
}

type MediaAigcVideoUnit struct {
	StartTime int64                   `json:"startTime,omitempty"` // 起始时间(ms)
	EndTime   int64                   `json:"endTime,omitempty"`   // 结束时间(ms)
	PictureID string                  `json:"pictureId,omitempty"` // 截图ID
	Details   []*MediaAigcVideoItemBo `json:"details,omitempty"`   // 检测详情
}

type MediaAigcVideoItemBo struct {
	IsAigc   bool    `json:"isAigc,omitempty"`   // 是否AIGC生成
	AigcRate float32 `json:"aigcRate,omitempty"` // AIGC概率分数
}
