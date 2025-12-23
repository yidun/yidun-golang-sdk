package response

type LiveAudioCallbackUnitRespV4 struct {
	Suggestion          *int                `json:"suggestion,omitempty"`
	PublicOpinionInfo   *string             `json:"publicOpinionInfo,omitempty"`
	Label               *int                `json:"label,omitempty"`
	SecondLabel         *string             `json:"secondLabel,omitempty"`
	ThirdLabel          *string             `json:"thirdLabel,omitempty"`
	SuggestionRiskLevel *int                `json:"suggestionRiskLevel,omitempty"`
	RiskDescription     *string             `json:"riskDescription,omitempty"`
	SuggestionLevel     *int                `json:"suggestionLevel,omitempty"`
	StartTime           *int64              `json:"startTime,omitempty"`
	EndTime             *int64              `json:"endTime,omitempty"`
	Content             *string             `json:"content,omitempty"`
	Type                *int                `json:"type,omitempty"`
	Rate                *float64            `json:"rate,omitempty"`
	Labels              []*SegmentsInfoV4   `json:"labels,omitempty"`
	URL                 *string             `json:"url,omitempty"`
	SpeakerID           *string             `json:"speakerId,omitempty"`
	SpeakerRiskScore    *int                `json:"speakerRiskScore,omitempty"`
	SegmentID           *string             `json:"segmentId,omitempty"`
	FrontSegment        *RelatedSegmentInfo `json:"frontSegment,omitempty"`
}

type RelatedSegmentInfo struct {
	Content *string `json:"content,omitempty"`
	URL     *string `json:"url,omitempty"`
}

type SegmentsInfoV4 struct {
	Label     *int               `json:"label,omitempty"`
	Level     *int               `json:"level,omitempty"`
	Rate      *float64           `json:"rate,omitempty"`
	SubLabels []*AudioThirdLabel `json:"subLabels,omitempty"`
}

type AudioThirdLabel struct {
	Details             *HintInfoV4 `json:"details,omitempty"`
	SubLabel            *string     `json:"subLabel,omitempty"`
	SecondLabel         *string     `json:"secondLabel,omitempty"`
	ThirdLabel          *string     `json:"thirdLabel,omitempty"`
	SuggestionRiskLevel *int        `json:"suggestionRiskLevel,omitempty"`
	RiskDescription     *string     `json:"riskDescription,omitempty"`
	SubLabelDepth       *int        `json:"subLabelDepth,omitempty"`
	Rate                *float64    `json:"rate,omitempty"`
}

type HintInfoV4 struct {
	HitInfos   []*HintDetailV4 `json:"hitInfos,omitempty"`
	HitLeaders []*LeaderV4     `json:"hitLeaders,omitempty"`
	Keywords   []*Keyword      `json:"keywords,omitempty"`
	Intent     *IntentUnit     `json:"intent,omitempty"`
}

type IntentUnit struct {
	Result *int     `json:"result,omitempty"`
	Score  *float64 `json:"score,omitempty"`
}

type HintDetailV4 struct {
	Value *string `json:"value,omitempty"`
}

type LeaderV4 struct {
	Value *string `json:"value,omitempty"`
}

type Keyword struct {
	Word *string `json:"word,omitempty"`
	StrategyGroupName *string `json:"strategyGroupName,omitempty"`
	StrategyGroupId *int64 `json:"strategyGroupId,omitempty"`
}
