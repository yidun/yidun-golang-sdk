package envidence

type MediaTextEvidence struct {
	// 融媒体机审-文本证据信息
	DataId           *string            `json:"dataId,omitempty"`
	Field            *string            `json:"field,omitempty"`
	Suggestion       *int               `json:"suggestion,omitempty"` // 检测结果，0：通过，1：嫌疑，2：不通过
	ResultType       *int               `json:"resultType,omitempty"`
	CensorType       *int               `json:"censorType,omitempty"`
	StrategyVersions []*StrategyVersion `json:"strategyVersions,omitempty"`
	Labels           []*TextLabelInfo   `json:"labels,omitempty"`
	IsRelatedHit     *bool              `json:"isRelatedHit,omitempty"`
}

// 文本标签信息
type TextLabelInfo struct {
	Label     *int            `json:"label,omitempty"`
	SubLabels []*TextSubLabel `json:"subLabels,omitempty"`
	Level     *int            `json:"level,omitempty"`
}

// 文本子标签信息
type TextSubLabel struct {
	SubLabel      *string             `json:"subLabel,omitempty"`
	Details       *TextSubLabelDetail `json:"details,omitempty"`
	SubLabelDepth *int                `json:"subLabelDepth,omitempty"`
	SecondLabel   *string             `json:"secondLabel,omitempty"`
	ThirdLabel    *string             `json:"thirdLabel,omitempty"`
}

type StrategyVersion struct {
	Label   *int    `json:"label,omitempty"`
	Version *string `json:"version,omitempty"`
}

// Getter and Setter for MediaTextEvidence
func (m *MediaTextEvidence) GetDataId() *string {
	return m.DataId
}

func (m *MediaTextEvidence) SetDataId(dataId *string) {
	m.DataId = dataId
}

func (m *MediaTextEvidence) GetField() *string {
	return m.Field
}

func (m *MediaTextEvidence) SetField(field *string) {
	m.Field = field
}

func (m *MediaTextEvidence) GetSuggestion() *int {
	return m.Suggestion
}

func (m *MediaTextEvidence) SetSuggestion(suggestion *int) {
	m.Suggestion = suggestion
}

func (m *MediaTextEvidence) GetResultType() *int {
	return m.ResultType
}

func (m *MediaTextEvidence) SetResultType(resultType *int) {
	m.ResultType = resultType
}

func (m *MediaTextEvidence) GetCensorType() *int {
	return m.CensorType
}

func (m *MediaTextEvidence) SetCensorType(censorType *int) {
	m.CensorType = censorType
}

func (m *MediaTextEvidence) GetStrategyVersions() []*StrategyVersion {
	return m.StrategyVersions
}

func (m *MediaTextEvidence) SetStrategyVersions(strategyVersions []*StrategyVersion) {
	m.StrategyVersions = strategyVersions
}

func (m *MediaTextEvidence) GetLabels() []*TextLabelInfo {
	return m.Labels
}

func (m *MediaTextEvidence) SetLabels(labels []*TextLabelInfo) {
	m.Labels = labels
}

func (m *MediaTextEvidence) GetIsRelatedHit() *bool {
	return m.IsRelatedHit
}

func (m *MediaTextEvidence) SetIsRelatedHit(isRelatedHit *bool) {
	m.IsRelatedHit = isRelatedHit
}
