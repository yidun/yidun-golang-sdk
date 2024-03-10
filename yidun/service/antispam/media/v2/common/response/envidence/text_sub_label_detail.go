package envidence

type TextSubLabelDetail struct {
	// 文本标签详细信息
	Keywords []*Keyword  `json:"keywords,omitempty"`
	LibInfos []*LibInfo  `json:"libInfos,omitempty"`
	HitInfos []*HintInfo `json:"hitInfos,omitempty"`
	Rules    []*RuleInfo `json:"rules,omitempty"`
}

func (d *TextSubLabelDetail) GetKeywords() []*Keyword {
	return d.Keywords
}

func (d *TextSubLabelDetail) SetKeywords(keywords []*Keyword) {
	d.Keywords = keywords
}

func (d *TextSubLabelDetail) GetLibInfos() []*LibInfo {
	return d.LibInfos
}

func (d *TextSubLabelDetail) SetLibInfos(libInfos []*LibInfo) {
	d.LibInfos = libInfos
}

func (d *TextSubLabelDetail) GetHitInfos() []*HintInfo {
	return d.HitInfos
}

func (d *TextSubLabelDetail) SetHitInfos(hitInfos []*HintInfo) {
	d.HitInfos = hitInfos
}

// 文本提示信息
type HintInfo struct {
	Value     *string             `json:"value,omitempty"`
	Positions []*HintInfoPosition `json:"positions,omitempty"`
}

func (h *HintInfo) GetValue() *string {
	return h.Value
}

func (h *HintInfo) SetValue(value *string) {
	h.Value = value
}

func (h *HintInfo) GetPositions() []*HintInfoPosition {
	return h.Positions
}

func (h *HintInfo) SetPositions(positions []*HintInfoPosition) {
	h.Positions = positions
}

// 文本提示信息位置
type HintInfoPosition struct {
	FieldName *string  `json:"fieldName,omitempty"`
	StartPos  *float64 `json:"startPos,omitempty"`
	EndPos    *float64 `json:"endPos,omitempty"`
}

func (h *HintInfoPosition) GetFieldName() *string {
	return h.FieldName
}

func (h *HintInfoPosition) SetFieldName(fieldName *string) {
	h.FieldName = fieldName
}

func (h *HintInfoPosition) GetStartPos() *float64 {
	return h.StartPos
}

func (h *HintInfoPosition) SetStartPos(startPos *float64) {
	h.StartPos = startPos
}

func (h *HintInfoPosition) GetEndPos() *float64 {
	return h.EndPos
}

func (h *HintInfoPosition) SetEndPos(endPos *float64) {
	h.EndPos = endPos
}

type LibInfo struct {
	Type   *int    `json:"type,omitempty"`
	Entity *string `json:"entity,omitempty"`
}

func (l *LibInfo) GetType() *int {
	return l.Type
}

func (l *LibInfo) SetType(_type *int) {
	l.Type = _type
}

func (l *LibInfo) GetEntity() *string {
	return l.Entity
}

func (l *LibInfo) SetEntity(entity *string) {
	l.Entity = entity
}

type Keyword struct {
	Word *string `json:"word,omitempty"`
}

func (k *Keyword) GetWord() *string {
	return k.Word
}

func (k *Keyword) SetWord(word *string) {
	k.Word = word
}

type RuleInfo struct {
	Name *string `json:"name,omitempty"`
}

func (r *RuleInfo) GetName() *string {
	return r.Name
}

func (r *RuleInfo) SetName(name *string) {
	r.Name = name
}
