package single

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
)

/**
 * 文本单次同步检测响应
 */
type TextCheckResponse struct {
	*types.CommonResponse
	Result *TextCheckResult `json:"result"`
}

/**
 * 文本检测结果
 */
type TextCheckResult struct {
	// 文本内容安全检测结果
	Antispam *Antispam `json:"antispam"`
	// 文本情感分析检测结果
	EmotionAnalysis *EmotionAnalysis `json:"emotionAnalysis"`
	// 文本反作弊检测结果
	Anticheat *Anticheat `json:"anticheat"`
	// 文本用户画像检测结果
	UserRisk *UserRisk `json:"userRisk"`
	// 文本语种检测结果
	Language *Language `json:"language"`
	// 风险控制检测结果
	RiskControl *RiskControl `json:"riskControl,omitempty"`
	// 相似文本检测结果
	SimilarText *SimilarText `json:"similarText,omitempty"`
	// 未成年人检测结果
	Underage *Underage `json:"underage,omitempty"`
	// aigc提示分析结果
	AigcPrompt *AigcPrompt `json:"aigcPrompt"`
	// 文本大模型检测结果
	LlmCheckInfo *LlmCheckInfo `json:"llmCheckInfo"`
}

type Language struct {
	TaskID  *string           `json:"taskId,omitempty"`  // task ID
	DataID  *string           `json:"dataId,omitempty"`  // data ID
	Details []*LanguageDetail `json:"details,omitempty"` // language detail list
}

type LanguageDetail struct {
	Type *string `json:"type,omitempty"` // language type
}

type UserRisk struct {
	TaskID        *string              `json:"taskId,omitempty"`        // task ID
	DataID        *string              `json:"dataId,omitempty"`        // data ID
	Details       []*UserRiskDetail    `json:"details,omitempty"`       // user risk detail list
	AccountDetail *RiskPortraitDetail  `json:"accountDetail,omitempty"` // 账号风险画像详情
	PhoneDetail   *RiskPortraitDetail  `json:"phoneDetail,omitempty"`   // 手机号风险画像详情
	IpDetail      *RiskPortraitDetail  `json:"ipDetail,omitempty"`      // IP风险画像详情
}

type UserRiskDetail struct {
	Account      *string             `json:"account,omitempty"`   // account
	AccountLevel *int                `json:"accountLevel"`        // 账号风险等级
	AcDetails    []*UserRiskAcDetail `json:"acDetails,omitempty"` // user risk account detail list
}

type UserRiskAcDetail struct {
	RiskType  *string  `json:"riskType,omitempty"`  // risk type
	RiskLevel *int     `json:"riskLevel,omitempty"` // risk level
	RiskScore *float64 `json:"riskScore,omitempty"` // risk score
}

type Anticheat struct {
	TaskID  *string            `json:"taskId,omitempty"`  // task ID
	DataID  *string            `json:"dataId,omitempty"`  // data ID
	Details []*AnticheatDetail `json:"details,omitempty"` // anticheat detail list
}

type AnticheatDetail struct {
	Suggestion *int                      `json:"suggestion,omitempty"` // suggestion
	HitInfos   []*AnticheatDetailHitInfo `json:"hitInfos,omitempty"`   // hit info list
}

type AnticheatDetailHitInfo struct {
	HitType *int    `json:"hitType,omitempty"` // hit type
	HitMsg  *string `json:"hitMsg,omitempty"`  // hit message
}

type EmotionAnalysis struct {
	TaskID  *string                  `json:"taskId,omitempty"`  // task ID
	DataID  *string                  `json:"dataId,omitempty"`  // data ID
	Details []*EmotionAnalysisDetail `json:"details,omitempty"` // emotion analysis details
}

type EmotionAnalysisDetail struct {
	PositiveProb *float64 `json:"positiveProb,omitempty"` // positive probability
	NegativeProb *float64 `json:"negativeProb,omitempty"` // negative probability
	Sentiment    *string  `json:"sentiment,omitempty"`    // sentiment
}

type Antispam struct {
	TaskID              *string            `json:"taskId"`
	DataID              *string            `json:"dataId"`
	Label               *int               `json:"label"`
	SecondLabel         *string            `json:"secondLabel"`
	ThirdLabel          *string            `json:"thirdLabel"`
	RiskDescription     *string            `json:"riskDescription"`
	Suggestion          *int               `json:"suggestion"`
	SuggestionLevel     *int               `json:"suggestionLevel"`
	SuggestionRiskLevel *int               `json:"suggestionRiskLevel"`
	CustomAction        *int               `json:"customAction"`
	ResultType          *int               `json:"resultType"`
	CensorType          *int               `json:"censorType"`
	Callback            *string            `json:"callback"`
	PublicOpinionInfo   *string            `json:"publicOpinionInfo"`
	CensorLabels        []*CensorLabel     `json:"censorLabels"`
	StrategyVersions    []*StrategyVersion `json:"strategyVersions"`
	CensorSource        *int               `json:"censorSource"`
	CensorRound         *int               `json:"censorRound"`
	CensorTime          *int64             `json:"censorTime"`
	IsRelatedHit        *bool              `json:"isRelatedHit"`
	RelatedHitType      *int               `json:"relatedHitType"`
	Labels              []*AntispamLabel   `json:"labels"`
	Remark              *string            `json:"remark"`
	Censor              *string            `json:"censor"`
	FilteredContent     *string            `json:"filteredContent"`
	MergeHints          []*string          `json:"mergeHints"`
	// 新增：关联内容字段（正确的后端字段名为 relateContent 单数）
	RelateContent       *string            `json:"relateContent,omitempty"`
	// Deprecated: Use RelateContent instead. This field is kept for backward compatibility.
	// Compatibility is handled automatically during JSON unmarshaling.
	RelatedContents     *string            `json:"-"`
	// 新增：命中来源字段（正确的后端字段名为 hitSource 单数）
	HitSource           *int               `json:"hitSource,omitempty"`
	// Deprecated: Use HitSource instead. This field is kept for backward compatibility.
	// Compatibility is handled automatically during JSON unmarshaling.
	HitSources          *int               `json:"-"`
	// 新增：父结构级别的三个字段（修复字段位置错误）
	HitType             *int               `json:"hitType,omitempty"`      // 命中策略类型（整体检测级别）
	StrategyType        *int               `json:"strategyType,omitempty"` // 策略类型（1-公有策略，2-私有策略）
	HitResult           *string            `json:"hitResult,omitempty"`    // 命中结果详情（注意：String类型）
	Status              *int               `json:"status"`
	CustomLabels        []*CustomLabel     `json:"customLabels"`
	CensorExtension     *CensorExtension   `json:"censorExtension"`
}

type CensorLabel struct {
	Code          *string `json:"code"`
	Desc          *string `json:"desc"`
	Name          *string `json:"name"`
	CustomCode    *string `json:"customCode"`
	ParentLabelId *string `json:"parentLabelId,omitempty"`
	Depth         *int    `json:"depth"`
}

type StrategyVersion struct {
	Label   *int    `json:"label"`
	Version *string `json:"version"`
}

type AntispamLabel struct {
	Label     *int                `json:"label"`
	Level     *int                `json:"level"`
	Rate      *float64            `json:"rate"`
	SubLabels []*AntispamSubLabel `json:"subLabels"`
	HitType   *int                `json:"hitType"`
}

type CustomLabel struct {
	Name  *string `json:"name"`
	Code  *string `json:"code"`
	Depth *int    `json:"depth"`
}

type AntispamSubLabel struct {
	SubLabel            *string                 `json:"subLabel"`
	SubLabelDepth       *int                    `json:"subLabelDepth"`
	SecondLabel         *string                 `json:"secondLabel"`
	ThirdLabel          *string                 `json:"thirdLabel"`
	RiskDescription     *string                 `json:"riskDescription"`
	SuggestionRiskLevel *int                    `json:"suggestionRiskLevel"`
	Rate                *float64                `json:"rate"`
	PoliticalSentiment  *int                    `json:"politicalSentiment"`
	IsRelatedLabel      *bool                   `json:"isRelatedLabel"`
	Details             *AntispamSubLabelDetail `json:"details"`
}

type AntispamSubLabelDetail struct {
	Keywords  []*AntispamSubLabelDetailKeyword `json:"keywords"`
	LibInfos  []*AntispamSubLabelDetailLibInfo `json:"libInfos"`
	Anticheat *AntispamSubLabelDetailAnticheat `json:"anticheat"`
	HitInfos  []*AntispamSubLabelDetailHitInfo `json:"hitInfos"`
	Rules     []*AntispamSubLabelDetailRule    `json:"rules"`
}

type AntispamSubLabelDetailKeyword struct {
	Word              *string `json:"word"`
	StrategyGroupName *string `json:"strategyGroupName"`
	StrategyGroupId   *int64  `json:"strategyGroupId"`
}

type AntispamSubLabelDetailRule struct {
	Name *string `json:"name"`
}

type AntispamSubLabelDetailLibInfo struct {
	Type              *int    `json:"type"`
	Entity            *string `json:"entity"`
	ReleaseTime       *int64  `json:"releaseTime"`
	StrategyGroupName *string `json:"strategyGroupName,omitempty"` // 命中的策略组名称（如果命中的词是策略组的，将策略组的名称也返回，若不是策略组词，值为空）
	StrategyGroupId   *int64  `json:"strategyGroupId,omitempty"`   // 命中的策略组ID
}

type AntispamSubLabelDetailAnticheat struct {
	Type *int `json:"type"`
}

type AntispamSubLabelDetailHitInfoPosition struct {
	FieldName *string `json:"fieldName"`
	StartPos  *int    `json:"startPos"`
	EndPos    *int    `json:"endPos"`
}

type AntispamSubLabelDetailHitInfo struct {
	Value     *string                                  `json:"value"`
	Type      *int                                     `json:"type"` // type:1，“通用反垃圾线索信息”，type:2，“增强版线索信息”
	Positions []*AntispamSubLabelDetailHitInfoPosition `json:"positions"`
}

type AigcPrompt struct {
	TaskId          *string              `json:"taskId"`                    // 任务id
	DataId          *string              `json:"dataId"`                    // 数据id
	ProxyAnswerType *int                 `json:"proxyAnswerType,omitempty"` // 代理应答类型（0-知识库，1-大模型，2-自定义知识库）
	Details         *[]*AigcPromptDetail `json:"details"`                   // 详情
}

type AigcPromptDetail struct {
	Type     *int    `json:"type"`     // prompt分类的枚举值
	Answer   *string `json:"answer"`   // 需要回答且能找到回答时返回
	Source   *int    `json:"source"`   // 标记对外输出内容由知识库结果还是大模型生成的结果（0代表知识库,1代表大模型,2代表自定义知识库）
	LibId    *string `json:"libId"`    // 知识库ID
	AnswerId *string `json:"answerId"` // 知识库-答案 ID
}

type LlmCheckInfo struct {
	TaskId  *string               `json:"taskId"`  // 任务id
	DataId  *string               `json:"dataId"`  // 数据id
	Details []*LlmCheckInfoDetail `json:"details"` // 详情
}

type LlmCheckInfoDetail struct {
	ModelIdentifier *string `json:"modelIdentifier"` // 模型标识
	Label           *string `json:"label"` // 大模型识别标签
	Explain         *string `json:"explain"` // llm 对于标签的解释
}

type CensorExtension struct {
	QualityInspectionTaskId *string `json:"qualityInspectionTaskId,omitempty"` // 质量检测任务ID
	InspTaskCreateTime      *int64  `json:"inspTaskCreateTime,omitempty"`      // 质检任务创建时间（Unix时间戳，毫秒）
}

// RiskControl 风险控制检测结果
type RiskControl struct {
	TaskID  *string               `json:"taskId,omitempty"`  // 任务ID
	DataID  *string               `json:"dataId,omitempty"`  // 数据ID
	Details []*RiskControlDetail `json:"details,omitempty"` // 风险控制详情列表
}

// RiskControlDetail 风险控制详情
type RiskControlDetail struct {
	RiskLevel *int                   `json:"riskLevel,omitempty"` // 风险等级 0-正常，10-高风险，20-中风险，30-低风险
	HitInfos  []*RiskControlHitInfo `json:"hitInfos,omitempty"`  // 命中的风险标签信息
}

// RiskControlHitInfo 风险控制命中信息
type RiskControlHitInfo struct {
	Type *string `json:"type,omitempty"` // 命中的标签code
	Name *string `json:"name,omitempty"` // 命中的标签名称
	Desc *string `json:"desc,omitempty"` // 标签描述
}

// SimilarText 相似文本检测结果
type SimilarText struct {
	TaskID  *string              `json:"taskId,omitempty"`  // 任务ID
	DataID  *string              `json:"dataId,omitempty"`  // 数据ID
	Details []*SimilarTextDetail `json:"details,omitempty"` // 相似文本详情列表
}

// SimilarTextDetail 相似文本详情
type SimilarTextDetail struct {
	MatchDataID     *string  `json:"matchDataId,omitempty"` // 匹配到的相似文本 dataId
	SimilarityScore *float64 `json:"rate,omitempty"`        // 相似度评分 (0.0-1.0)，后端字段名为 rate
}

// Underage 未成年人检测结果
type Underage struct {
	TaskID  *string           `json:"taskId,omitempty"`  // 任务ID
	DataID  *string           `json:"dataId,omitempty"`  // 数据ID
	Details []*UnderageDetail `json:"details,omitempty"` // 未成年人详情列表
}

// UnderageDetail 未成年人详情
type UnderageDetail struct {
	Type *int `json:"type,omitempty"` // 1-涉及未成年 2-不涉及未成年
}

// RiskPortraitDetail 风险画像详情
// 用于 UserRisk 的账号、手机号、IP 等维度的风险画像分析
type RiskPortraitDetail struct {
	RiskDetails []*RiskDetail `json:"riskDetails,omitempty"` // 风险详情列表（通常只有一个元素）
	Success     *bool         `json:"success,omitempty"`     // 是否成功获取画像
}

// RiskDetail 风险详情
type RiskDetail struct {
	RiskInfoList []*RiskInfo            `json:"riskInfoList,omitempty"` // 风险信息列表
	PropDetails  map[string]interface{} `json:"propDetails,omitempty"`  // 属性详情（JSON对象）
	Account      *string                `json:"account,omitempty"`      // 账号
	Phone        *string                `json:"phone,omitempty"`        // 手机号
	Ip           *string                `json:"ip,omitempty"`           // IP地址
}

// RiskInfo 风险信息
type RiskInfo struct {
	RiskLevel *int     `json:"riskLevel,omitempty"` // 风险等级
	RiskModel *string  `json:"riskModel,omitempty"` // 风险模型
	RiskType  *int     `json:"riskType,omitempty"`  // 风险类型
	RiskScore *float64 `json:"riskScore,omitempty"` // 风险分数
}
