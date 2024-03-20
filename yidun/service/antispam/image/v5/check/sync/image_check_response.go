package sync

import "github.com/yidun/yidun-golang-sdk/yidun/core/types"

type ImageCheckResponse struct {
	*types.CommonResponse
	Result *[]ImageV5Result `json:"result,omitempty"`
}

type ImageV5Result struct {
	// 反垃圾结果
	Antispam *ImageV5AntispamResp `json:"antispam,omitempty"`
	// ocr结果
	Ocr *ImageV5OcrResp `json:"ocr,omitempty"`
	// 人脸结果
	Face *ImageV5FaceResp `json:"face,omitempty"`
	// 图片质量结果
	Quality *ImageV5QualityResp `json:"quality,omitempty"`
	// 图片logo结果
	Logo *ImageV5LogoResp `json:"logo,omitempty"`
	// 图片识别结果
	Discern *ImageV5DiscernResp `json:"discern,omitempty"`
	// 用户画像结果
	UserRisk *ImageV5UserRiskResp `json:"userRisk,omitempty"`
	// 反作弊结果
	Anticheat *ImageAnticheatV5Resp `json:"anticheat,omitempty"`
	// 智能风控结果
	RiskControl *ImageRiskControlV5Resp `json:"riskControl,omitempty"`
}

type ImageV5AntispamResp struct {
	// 即易盾生成的uuid，唯一标识一张图片
	TaskId *string `json:"taskId,omitempty"`
	// 检测状态 0 未开始、1检测中、2检测成功、3检测失败
	Status *int `json:"status,omitempty"`
	// 检测失败原因
	FailureReason *int `json:"failureReason,omitempty"`
	// 检测建议结果
	Suggestion *int `json:"suggestion,omitempty"`
	// 一级分类
	Label       *int   `json:"label,omitempty"`
	// 二级分类
    SecondLabel *string `json:"secondLabel,omitempty"`
    // 三级分类
    ThirdLabel  *string `json:"thirdLabel,omitempty"`
	// 图片人审状态
	CensorType *int `json:"censorType,omitempty"`
	// 策略版本字段
	StrategyVersions *[]ImageV5VersionDetail `json:"strategyVersions,omitempty"`
	// 产品方传的图片标识，原样返回
	Name *string `json:"name,omitempty"`
	// 客户图片唯一标识
	DataId *string `json:"dataId,omitempty"`
	// 各个类别检测结果
	Labels *[]ImageV5LabelDetail `json:"labels,omitempty"`
	// 机器或人审结果
	ResultType *int `json:"resultType,omitempty"`
	// 审核时间，回调返回
	CensorTime *int64 `json:"censorTime,omitempty"`
	// 审核来源，回调返回
	CensorSource *int `json:"censorSource,omitempty"`
	// 审核轮数，回调返回
	CensorRound *int `json:"censorRound,omitempty"`
	// 审核标签，回调返回
	CensorLabels *[]CensorLabelInfo `json:"censorLabels,omitempty"`
	// 云信融合业务转换结果
	CustomAction *int `json:"customAction,omitempty"`
	// 分帧数
	FrameSize *int `json:"frameSize,omitempty"`
	// 是否有隐藏
	Hidden *bool `json:"hidden,omitempty"`
	// 隐藏文件的格式
	HiddenFormat *string `json:"hiddenFormat,omitempty"`
	// 图片md5
	ImgMd5 *string `json:"imgMd5,omitempty"`
}

type ImageV5VersionDetail struct {
	// 垃圾类别
	Label *int `json:"label,omitempty"`
	// 版本
	Version *string `json:"version,omitempty"`
}

type ImageV5LabelDetail struct {
	// 对外的图片label
	Label *int `json:"label,omitempty"`
	// 判断结果
	Level *int `json:"level,omitempty"`
	// 得分，范围为0到1
	Rate *float32 `json:"rate,omitempty"`
	// 各个类别检测结果
	SubLabels *[]ImageV5SubLabelDetail `json:"subLabels,omitempty"`
}

type ImageV5SubLabelDetail struct {
	// 对外的图片label
	SubLabel *string `json:"subLabel,omitempty"`
	// 命中的最终细分类的层级
	SubLabelDepth *int `json:"subLabelDepth,omitempty"`
	// 二级分类，必返回
	SecondLabel *string `json:"secondLabel,omitempty"`
	// 三级分类，可能返回
	ThirdLabel *string `json:"thirdLabel,omitempty"`
	// 命中标识
	HitStrategy *int `json:"hitStrategy,omitempty"`
	// 判断结果，0-正常，1-不确定，2-确定
	Rate *float32 `json:"rate,omitempty"`
	// 二级分类命中详情
	Details *SubLabelDetails `json:"details,omitempty"`
}

type SubLabelDetails struct {
	// 反垃圾自定义敏感词结果
	Keywords *[]AntispamInfo `json:"keywords,omitempty"`
	// 反垃圾自定义图片名单结果
	LibInfos *[]AntispamInfo `json:"libInfos,omitempty"`
	// 反垃圾其他命中信息
	HitInfos *[]AntispamInfo `json:"hitInfos,omitempty"`
	// 反作弊结果
	Anticheat *AnticheatInfo `json:"anticheat,omitempty"`
}

type AnticheatInfo struct {
	HitType *int `json:"hitType,omitempty"`
}

type AntispamInfo struct {
	// 自定义敏感词
	Word *string `json:"word,omitempty"`
	// 自定义图片名单url
	Entity *string `json:"entity,omitempty"`
	// 自定义图片名单命中次数
	HitCount *int    `json:"hitCount,omitempty"`
	Value    *string `json:"value,omitempty"`
	Group    *string `json:"group,omitempty"`
	// 坐标左上一个 右下一个
	X1          *float32 `json:"x1,omitempty"`
	Y1          *float32 `json:"y1,omitempty"`
	X2          *float32 `json:"x2,omitempty"`
	Y2          *float32 `json:"y2,omitempty"`
	ReleaseTime *int64   `json:"releaseTime,omitempty"`
}

type CensorLabelInfo struct {
	Code       *string `json:"code,omitempty"`
	Name       *string `json:"name,omitempty"`
	Desc       *string `json:"desc,omitempty"`
	CustomCode *string `json:"customCode,omitempty"`
}

type ImageV5OcrResp struct {
	// 即易盾生成的uuid，唯一标识一张图片
	TaskId *string `json:"taskId,omitempty"`
	// 产品方传的图片标识，原样返回
	Name *string `json:"name,omitempty"`
	// 客户图片唯一标识
	DataId *string `json:"dataId,omitempty"`
	// 长
	Height *int `json:"height,omitempty"`
	// 宽
	Width *int `json:"width,omitempty"`
	// ocr检测结果
	Details *[]ImageV5OcrDetail `json:"details,omitempty"`
}

type ImageV5OcrDetail struct {
	// ocr文本
	Content *string `json:"content,omitempty"`
	// ocr文本详情
	LineContents *[]OcrLineContent `json:"lineContents,omitempty"`
}

type OcrLineContent struct {
	// 行ocr文本
	LineContent *string `json:"lineContent,omitempty"`
	// 相对坐标
	X1 *float32 `json:"x1,omitempty"`
	Y1 *float32 `json:"y1,omitempty"`
	X2 *float32 `json:"x2,omitempty"`
	Y2 *float32 `json:"y2,omitempty"`
	// 语种信息
	Lang *string `json:"lang,omitempty"`
}

type ImageV5FaceResp struct {
	// 即易盾生成的uuid，唯一标识一张图片
	TaskId *string `json:"taskId,omitempty"`
	// 产品方传的图片标识，原样返回
	Name *string `json:"name,omitempty"`
	// 客户图片唯一标识
	DataId *string `json:"dataId,omitempty"`
	// face检测结果
	Details *[]ImageV5FaceDetail `json:"details,omitempty"`
}

type ImageV5FaceDetail struct {
	// 人脸数量
	FaceNumber *int `json:"faceNumber,omitempty"`
	// 是否有伪造人脸，0：无，1：有
	Deepfake *int `json:"deepfake,omitempty"`
	// 人脸详情
	FaceContents *[]FaceLineContent `json:"faceContents,omitempty"`
}

type FaceLineContent struct {
	// 人名
	Name *string `json:"name,omitempty"`
	// 坐标左上一个 右下一个
	X1 *float32 `json:"x1,omitempty"`
	Y1 *float32 `json:"y1,omitempty"`
	X2 *float32 `json:"x2,omitempty"`
	Y2 *float32 `json:"y2,omitempty"`
	// cartoon（动漫脸）、normal（真人脸）
	Type *string `json:"type,omitempty"`
	// star（明星）、normal（非明星）
	Category *string `json:"category,omitempty"`
	// 性别
	Gender *string `json:"gender,omitempty"`
	// 年龄
	Age *int `json:"age,omitempty"`
	// 人脸大小占比
	SizeRatio *string `json:"sizeRatio,omitempty"`
	// 人脸颜值分类
	BeautyScore *float32 `json:"beautyScore,omitempty"`
	// 人脸情绪
	Expression *string `json:"expression,omitempty"`
	// 人脸遮挡
	MaskType *string `json:"maskType,omitempty"`
	// 遮挡分数
	MaskScore *float32 `json:"maskScore,omitempty"`
	// 成长阶段
	GrowthStage *string `json:"growthStage,omitempty"`
	// 人脸角度list
	PoseInfoList *[]FacePoseInfo `json:"poseInfoList,omitempty"`
}

type FacePoseInfo struct {
	// 人脸角度标签
	Label *string `json:"label,omitempty"`
	// 人脸角度值
	Angle *float64 `json:"angle,omitempty"`
}

type ImageV5QualityResp struct {
	// 即易盾生成的uuid，唯一标识一张图片
	TaskId *string `json:"taskId,omitempty"`
	// 产品方传的图片标识，原样返回
	Name *string `json:"name,omitempty"`
	// 客户图片唯一标识
	DataId *string `json:"dataId,omitempty"`
	// 质量检测结果
	Details *[]ImageV5QualityDetail `json:"details,omitempty"`
}

type ImageV5QualityDetail struct {
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

type BackgroundInfo struct {
	PureBackground *bool `json:"pureBackground,omitempty"`
}

type BoarderColor struct {
	Top    *int `json:"top,omitempty"`
	Right  *int `json:"right,omitempty"`
	Bottom *int `json:"bottom,omitempty"`
	Left   *int `json:"left,omitempty"`
}

type MetaInfo struct {
	ByteSize *int64  `json:"byteSize,omitempty"`
	Height   *int    `json:"height,omitempty"`
	Width    *int    `json:"width,omitempty"`
	Format   *string `json:"format,omitempty"`
	// 是否有隐藏
	Hidden *bool `json:"hidden,omitempty"`
	// 隐藏文件的格式
	HiddenFormat *string `json:"hiddenFormat,omitempty"`
}

type BoarderInfo struct {
	Hit    *bool `json:"hit,omitempty"`
	Top    *bool `json:"top,omitempty"`
	Right  *bool `json:"right,omitempty"`
	Bottom *bool `json:"bottom,omitempty"`
	Left   *bool `json:"left,omitempty"`
	// 图片边框颜色
	Color *BoarderColor `json:"color,omitempty"`
}

type ImageV5LogoResp struct {
	// 即易盾生成的uuid，唯一标识一张图片
	TaskId *string `json:"taskId,omitempty"`
	// 产品方传的图片标识，原样返回
	Name *string `json:"name,omitempty"`
	// 客户图片唯一标识
	DataId *string `json:"dataId,omitempty"`
	// 质量检测结果
	Details *[]ImageV5LogoDetail `json:"details,omitempty"`
}

type ImageV5LogoDetail struct {
	// logo信息
	LogoName *string  `json:"logoName,omitempty"`
	X1       *float32 `json:"x1,omitempty"`
	Y1       *float32 `json:"y1,omitempty"`
	X2       *float32 `json:"x2,omitempty"`
	Y2       *float32 `json:"y2,omitempty"`
	// 子标签分数
	Rate *float32 `json:"rate,omitempty"`
	// logo大小占比
	SizeRatio *string `json:"sizeRatio,omitempty"`
}

type ImageV5DiscernResp struct {
	TaskId  *string                 `json:"taskId,omitempty"`  // 即易盾生成的uuid，唯一标识一张图片
	Name    *string                 `json:"name,omitempty"`    // 产品方传的图片标识，原样返回
	DataId  *string                 `json:"dataId,omitempty"`  // 客户图片唯一标识
	Details []*ImageV5DiscernDetail `json:"details,omitempty"` // 质量检测结果
}

type ImageV5DiscernDetail struct {
	Type        *int     `json:"type,omitempty"`        // 识别物体类型 1 场景
	DiscernName *string  `json:"discernName,omitempty"` // 识别名称
	DiscernKey  *string  `json:"discernKey,omitempty"`  // 识别标识
	Rate        *float32 `json:"rate,omitempty"`        // 分数
}

type ImageV5UserRiskResp struct {
	TaskId  *string                  `json:"taskId,omitempty"`  // 即易盾生成的uuid，唯一标识一张图片
	Name    *string                  `json:"name,omitempty"`    // 产品方传的图片标识，原样返回
	DataId  *string                  `json:"dataId,omitempty"`  // 客户图片唯一标识
	Details []*ImageV5UserRiskDetail `json:"details,omitempty"` // 命中详细结果
}

type ImageV5UserRiskDetail struct {
	Account      *string `json:"account,omitempty"`      // 账号
	AccountLevel *int    `json:"accountLevel,omitempty"` // 用户账号风险级别
}

// ImageAnticheatV5Resp v5用户画像结果
type ImageAnticheatV5Resp struct {
	TaskId  *string                          `json:"taskId,omitempty"`  // 任务id
	DataId  *string                          `json:"dataId,omitempty"`  // 数据id
	Details []*ImageAnticheatDetailHitInfoV5 `json:"details,omitempty"` // 反作弊命中详情
}

// ImageAnticheatDetailHitInfoV5
type ImageAnticheatDetailHitInfoV5 struct {
	Suggestion *int                     `json:"suggestion,omitempty"` // 反作弊结果
	HitInfos   []*ImageAnticheatHitInfo `json:"hitInfos,omitempty"`
}

// ImageAnticheatHitInfo
type ImageAnticheatHitInfo struct {
	HitType *int    `json:"hitType,omitempty"`
	HitMsg  *string `json:"hitMsg,omitempty"`
}

// ImageRiskControlV5Resp v5用户画像结果
type ImageRiskControlV5Resp struct {
	TaskId  *string                            `json:"taskId,omitempty"`  // 任务id
	DataId  *string                            `json:"dataId,omitempty"`  // 数据id
	Details []*ImageRiskControlDetailHitInfoV5 `json:"details,omitempty"` // 智能风控命中详情
}

// ImageRiskControlDetailHitInfoV5
type ImageRiskControlDetailHitInfoV5 struct {
	RiskLevel *int                       `json:"riskLevel,omitempty"` // 风控结果
	HitInfos  []*ImageRiskControlHitInfo `json:"hitInfos,omitempty"`  // 命中详情
}

// ImageRiskControlHitInfo
type ImageRiskControlHitInfo struct {
	Type *string `json:"type,omitempty"` // 风险标签type
	Name *string `json:"name,omitempty"` // 风险标签名称
	Desc *string `json:"desc,omitempty"` // 风险标签描述信息
}
