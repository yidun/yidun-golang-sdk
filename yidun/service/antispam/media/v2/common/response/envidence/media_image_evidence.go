package envidence

type MediaImageEvidence struct {
	// 融媒体机审-图片证据信息
	Status        *int              `json:"status,omitempty"` // 检测结果，0 未开始、1检测中、2检测成功、3检测失败
	Suggestion    *int              `json:"suggestion,omitempty"`
	ResultType    *int              `json:"resultType,omitempty"`
	FrameSize     *int              `json:"frameSize,omitempty"`
	CensorType    *int              `json:"censorType,omitempty"`
	FailureReason *int              `json:"failureReason,omitempty"` // 失败原因
	DataId        *string           `json:"dataId,omitempty"`
	Field         *string           `json:"field,omitempty"`
	Name          *string           `json:"name,omitempty"`
	Labels        []*ImageLabelInfo `json:"labels,omitempty"`
}

// 图片标签信息
type ImageLabelInfo struct {
	Label     *int             `json:"label,omitempty"`
	SubLabels []*ImageSubLabel `json:"subLabels,omitempty"`
	Level     *int             `json:"level,omitempty"`
	Rate      *float32         `json:"rate,omitempty"`
}

// 图片子标签信息
type ImageSubLabel struct {
	SubLabel      *string              `json:"subLabel,omitempty"`
	Rate          *float64             `json:"rate,omitempty"`
	Details       *ImageSubLabelDetail `json:"details,omitempty"`
	SubLabelDepth *int                 `json:"subLabelDepth,omitempty"`
	SecondLabel   *string              `json:"secondLabel,omitempty"`
	ThirdLabel    *string              `json:"thirdLabel,omitempty"`
}

// Getter and Setter for MediaImageEvidence
func (m *MediaImageEvidence) GetStatus() *int {
	return m.Status
}

func (m *MediaImageEvidence) SetStatus(status *int) {
	m.Status = status
}

func (m *MediaImageEvidence) GetSuggestion() *int {
	return m.Suggestion
}

func (m *MediaImageEvidence) SetSuggestion(suggestion *int) {
	m.Suggestion = suggestion
}

func (m *MediaImageEvidence) GetResultType() *int {
	return m.ResultType
}

func (m *MediaImageEvidence) SetResultType(resultType *int) {
	m.ResultType = resultType
}

func (m *MediaImageEvidence) GetFrameSize() *int {
	return m.FrameSize
}

func (m *MediaImageEvidence) SetFrameSize(frameSize *int) {
	m.FrameSize = frameSize
}

func (m *MediaImageEvidence) GetCensorType() *int {
	return m.CensorType
}

func (m *MediaImageEvidence) SetCensorType(censorType *int) {
	m.CensorType = censorType
}

func (m *MediaImageEvidence) GetFailureReason() *int {
	return m.FailureReason
}

func (m *MediaImageEvidence) SetFailureReason(failureReason *int) {
	m.FailureReason = failureReason
}

func (m *MediaImageEvidence) GetDataId() *string {
	return m.DataId
}

func (m *MediaImageEvidence) SetDataId(dataId *string) {
	m.DataId = dataId
}

func (m *MediaImageEvidence) GetField() *string {
	return m.Field
}

func (m *MediaImageEvidence) SetField(field *string) {
	m.Field = field
}

func (m *MediaImageEvidence) GetName() *string {
	return m.Name
}

func (m *MediaImageEvidence) SetName(name *string) {
	m.Name = name
}

func (m *MediaImageEvidence) GetLabels() []*ImageLabelInfo {
	return m.Labels
}

func (m *MediaImageEvidence) SetLabels(labels []*ImageLabelInfo) {
	m.Labels = labels
}
