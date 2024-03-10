package envidence

type ImageSubLabelDetail struct {
	// 反垃圾自定义敏感词结果
	Keywords []*ImageSubLabelDetailInfo `json:"keywords,omitempty"`
	// 反垃圾自定义图片名单结果
	LibInfos []*ImageSubLabelDetailInfo `json:"libInfos,omitempty"`
	// 反垃圾其他命中信息
	HitInfos []*ImageSubLabelDetailInfo `json:"hitInfos,omitempty"`
	Rules    []*ImageSubLabelDetailInfo `json:"rules,omitempty"`
}

func (d *ImageSubLabelDetail) GetKeywords() []*ImageSubLabelDetailInfo {
	return d.Keywords
}

func (d *ImageSubLabelDetail) SetKeywords(keywords []*ImageSubLabelDetailInfo) {
	d.Keywords = keywords
}

func (d *ImageSubLabelDetail) GetLibInfos() []*ImageSubLabelDetailInfo {
	return d.LibInfos
}

func (d *ImageSubLabelDetail) SetLibInfos(libInfos []*ImageSubLabelDetailInfo) {
	d.LibInfos = libInfos
}

func (d *ImageSubLabelDetail) GetHitInfos() []*ImageSubLabelDetailInfo {
	return d.HitInfos
}

func (d *ImageSubLabelDetail) SetHitInfos(hitInfos []*ImageSubLabelDetailInfo) {
	d.HitInfos = hitInfos
}

type ImageSubLabelDetailInfo struct {
	// 自定义敏感词
	Word *string `json:"word,omitempty"`
	// 自定义图片名单url
	Entity *string `json:"entity,omitempty"`
	// 自定义图片名单命中次数
	HitCount *int    `json:"hitCount,omitempty"`
	Value    *string `json:"value,omitempty"`
	Group    *string `json:"group,omitempty"`
	// 坐标左上一个 右下一个
	X1   *float64 `json:"x1,omitempty"`
	Y1   *float64 `json:"y1,omitempty"`
	X2   *float64 `json:"x2,omitempty"`
	Y2   *float64 `json:"y2,omitempty"`
	Name *string  `json:"name,omitempty"`
}

func (i *ImageSubLabelDetailInfo) GetWord() *string {
	return i.Word
}

func (i *ImageSubLabelDetailInfo) SetWord(word *string) {
	i.Word = word
}

func (i *ImageSubLabelDetailInfo) GetEntity() *string {
	return i.Entity
}

func (i *ImageSubLabelDetailInfo) SetEntity(entity *string) {
	i.Entity = entity
}

func (i *ImageSubLabelDetailInfo) GetHitCount() *int {
	return i.HitCount
}

func (i *ImageSubLabelDetailInfo) SetHitCount(hitCount *int) {
	i.HitCount = hitCount
}

func (i *ImageSubLabelDetailInfo) GetValue() *string {
	return i.Value
}

func (i *ImageSubLabelDetailInfo) SetValue(value *string) {
	i.Value = value
}

func (i *ImageSubLabelDetailInfo) GetGroup() *string {
	return i.Group
}

func (i *ImageSubLabelDetailInfo) SetGroup(group *string) {
	i.Group = group
}

func (i *ImageSubLabelDetailInfo) GetX1() *float64 {
	return i.X1
}

func (i *ImageSubLabelDetailInfo) SetX1(x1 *float64) {
	i.X1 = x1
}

func (i *ImageSubLabelDetailInfo) GetY1() *float64 {
	return i.Y1
}

func (i *ImageSubLabelDetailInfo) SetY1(y1 *float64) {
	i.Y1 = y1
}

func (i *ImageSubLabelDetailInfo) GetX2() *float64 {
	return i.X2
}

func (i *ImageSubLabelDetailInfo) SetX2(x2 *float64) {
	i.X2 = x2
}

func (i *ImageSubLabelDetailInfo) GetY2() *float64 {
	return i.Y2
}

func (i *ImageSubLabelDetailInfo) SetY2(y2 *float64) {
	i.Y2 = y2
}
