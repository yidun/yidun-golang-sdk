package query

import "github.com/yidun/yidun-golang-sdk/yidun/core/types"

type ImageQueryResponse struct {
	*types.CommonResponse
	Result *[]ImageQueryResp `json:"result"`
}

type ImageQueryResp struct {
	// 图片名称
	Name *string `json:"name,omitempty"`
	// taskId
	TaskId *string
	// 图片状态
	Status *int `json:"status,omitempty"`
	// 审核来源
	CensorSource *int `json:"censorSource,omitempty"`
	// 审核时间
	CensorTime *int64 `json:"censorTime,omitempty"`
	// 检测结果
	Suggestion *int `json:"suggestion,omitempty"`
	// 图片nos地址, 失效时间6小时
	URL *string `json:"url,omitempty"`
	// 图片md5
	ImgMd5 *string `json:"imgMd5,omitempty"`
	// 图片标签
	Labels *[]ImageLabelQueryResp `json:"labels,omitempty"`
}

type ImageLabelQueryResp struct {
	// 对外的图片label
	Label *int `json:"label,omitempty"`
	// 判断结果，0-正常，1-不确定，2-确定
	Level *int `json:"level,omitempty"`
	// 得分，范围为0到1
	Rate *float32 `json:"rate,omitempty"`
	// 二级分类
	SubLabel *string `json:"subLabel,omitempty"`
}

// 获取图片名称
func (r *ImageQueryResp) GetName() *string {
	return r.Name
}

// 设置图片名称
func (r *ImageQueryResp) SetName(name *string) {
	r.Name = name
}

// 获取图片taskid
func (r *ImageQueryResp) GetTaskId() *string {
	return r.TaskId
}

// 设置图片taskid
func (r *ImageQueryResp) SetTaskId(taskId *string) {
	r.TaskId = taskId
}

// 获取status
func (r *ImageQueryResp) GetStatus() *int {
	return r.Status
}

// 设置status
func (r *ImageQueryResp) SetStatus(status *int) {
	r.Status = status
}

// 获取审核来源
func (r *ImageQueryResp) GetCensorSource() *int {
	return r.CensorSource
}

// 设置审核来源
func (r *ImageQueryResp) SetCensorSource(censorSource *int) {
	r.CensorSource = censorSource
}

// 获取审核时间
func (r *ImageQueryResp) GetCensorTime() *int64 {
	return r.CensorTime
}

// 设置审核时间
func (r *ImageQueryResp) SetCensorTime(censorTime *int64) {
	r.CensorTime = censorTime
}

// 获取审核结果
func (r *ImageQueryResp) GetSuggestion() *int {
	return r.Suggestion
}

// 设置审核结果
func (r *ImageQueryResp) SetSuggestion(suggestion *int) {
	r.Suggestion = suggestion
}

// 获取图片nos地址
func (r *ImageQueryResp) GetURL() *string {
	return r.URL
}

// 设置图片nos地址
func (r *ImageQueryResp) SetURL(url *string) {
	r.URL = url
}

// 获取图片md5
func (r *ImageQueryResp) GetImgMd5() *string {
	return r.ImgMd5
}

// 设置图片md5
func (r *ImageQueryResp) SetImgMd5(imgMd5 *string) {
	r.ImgMd5 = imgMd5
}

// 获取图片标签
func (r *ImageQueryResp) GetLabels() *[]ImageLabelQueryResp {
	return r.Labels
}

// 设置图片标签
func (r *ImageQueryResp) SetLabels(labels *[]ImageLabelQueryResp) {
	r.Labels = labels
}

// 获取图片标签的外部label
func (l *ImageLabelQueryResp) GetLabel() *int {
	return l.Label
}

// 设置图片标签的外部label
func (l *ImageLabelQueryResp) SetLabel(label *int) {
	l.Label = label
}

// 获取判断结果
func (l *ImageLabelQueryResp) GetLevel() *int {
	return l.Level
}

// 设置判断结果
func (l *ImageLabelQueryResp) SetLevel(level *int) {
	l.Level = level
}

// 获取得分
func (l *ImageLabelQueryResp) GetRate() *float32 {
	return l.Rate
}

// 设置得分
func (l *ImageLabelQueryResp) SetRate(rate *float32) {
	l.Rate = rate
}

// 获取二级分类
func (l *ImageLabelQueryResp) GetSubLabel() *string {
	return l.SubLabel
}

// 设置二级分类
func (l *ImageLabelQueryResp) SetSubLabel(subLabel *string) {
	l.SubLabel = subLabel
}
