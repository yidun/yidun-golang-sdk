package single

import "github.com/yidun/yidun-golang-sdk/yidun/core/types"

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
	TaskID  *string           `json:"taskId,omitempty"`  // task ID
	DataID  *string           `json:"dataId,omitempty"`  // data ID
	Details []*UserRiskDetail `json:"details,omitempty"` // user risk detail list
}

type UserRiskDetail struct {
	Account   *string             `json:"account,omitempty"`   // account
	AcDetails []*UserRiskAcDetail `json:"acDetails,omitempty"` // user risk account detail list
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
	TaskID           *string            `json:"taskId"`
	DataID           *string            `json:"dataId"`
	Label            *int               `json:"label"`
	SecondLabel      *string            `json:"secondLabel"`
	ThirdLabel       *string            `json:"thirdLabel"`
	Suggestion       *int               `json:"suggestion"`
	SuggestionLevel  *int               `json:"suggestionLevel"`
	CustomAction     *int               `json:"customAction"`
	ResultType       *int               `json:"resultType"`
	CensorType       *int               `json:"censorType"`
	Callback         *string            `json:"callback"`
	CensorLabels     []*CensorLabel     `json:"censorLabels"`
	StrategyVersions []*StrategyVersion `json:"strategyVersions"`
	CensorSource     *int               `json:"censorSource"`
	CensorRound      *int               `json:"censorRound"`
	CensorTime       *int64             `json:"censorTime"`
	IsRelatedHit     *bool              `json:"isRelatedHit"`
	Labels           []*AntispamLabel   `json:"labels"`
	Remark           *string            `json:"remark"`
	Censor           *string            `json:"censor"`
	FilteredContent  *string            `json:"filteredContent"`
	MergeHints       []*string          `json:"mergeHints"`
	RelateContents   *string            `json:"relatedContents"`
	HitSources       *int               `json:"hitSources"`
}

type CensorLabel struct {
	Code *string `json:"code"`
	Desc *string `json:"desc"`
	Name *string `json:"name"`
}

type StrategyVersion struct {
	Label   *int    `json:"label"`
	Version *string `json:"version"`
}

type AntispamLabel struct {
	Label       *int                `json:"label"`
	SecondLabel *string             `json:"secondLabel"`
	ThirdLabel  *string             `json:"thirdLabel"`
	Level       *int                `json:"level"`
	Rate        *float64            `json:"rate"`
	SubLabels   []*AntispamSubLabel `json:"subLabels"`
	HitType     *int                `json:"hitType"`
}

type AntispamSubLabel struct {
	SubLabel      *string                 `json:"subLabel"`
	SubLabelDepth *int                    `json:"subLabelDepth"`
	SecondLabel   *string                 `json:"secondLabel"`
	ThirdLabel    *string                 `json:"thirdLabel"`
	Rate          *float64                `json:"rate"`
	Details       *AntispamSubLabelDetail `json:"details"`
}

type AntispamSubLabelDetail struct {
	Keywords  []*AntispamSubLabelDetailKeyword `json:"keywords"`
	LibInfos  []*AntispamSubLabelDetailLibInfo `json:"libInfos"`
	Anticheat *AntispamSubLabelDetailAnticheat `json:"anticheat"`
	HitInfos  []*AntispamSubLabelDetailHitInfo `json:"hitInfos"`
	Rules     []*AntispamSubLabelDetailRule    `json:"rules"`
}

type AntispamSubLabelDetailKeyword struct {
	Word *string `json:"word"`
}

type AntispamSubLabelDetailRule struct {
	Name *string `json:"name"`
}

type AntispamSubLabelDetailLibInfo struct {
	Type        *int    `json:"type"`
	Entity      *string `json:"entity"`
	ReleaseTime *string `json:"releaseTime"`
}

type AntispamSubLabelDetailAnticheat struct {
	HitType *int
}

type AntispamSubLabelDetailHitInfoPosition struct {
	FieldName *string `json:"fieldName"`
	StartPos  *int    `json:"startPos"`
	EndPos    *int    `json:"endPos"`
}

type AntispamSubLabelDetailHitInfo struct {
	Value     *string                                  `json:"value"`
	Positions []*AntispamSubLabelDetailHitInfoPosition `json:"positions"`
}
