package request

type LiveAudioBarrage struct {
	// 弹幕id
	Id *string `json:"id,omitempty"`
	// 直播UUID
	TaskId *string `json:"taskId,omitempty"`
	// 数据 id
	DataId *string `json:"dataId,omitempty"`
	// 弹幕发送时间
	PublishTime *int64 `json:"publishTime,omitempty"`
	// 弹幕内容
	Content *string `json:"content,omitempty"`
	// 高亮标签
	Hint *[]string `json:"hint,omitempty"`
	// 0-正常 1-嫌疑 2-确定
	Level *int `json:"level,omitempty"`
	// 垃圾类型
	SpamType *int `json:"spamType,omitempty"`
}

// SetId sets the Id field of the LiveAudioBarrage struct.
func (lab *LiveAudioBarrage) SetId(id string) {
	lab.Id = &id
}

// SetTaskId sets the TaskId field of the LiveAudioBarrage struct.
func (lab *LiveAudioBarrage) SetTaskId(taskId string) {
	lab.TaskId = &taskId
}

// SetDataId sets the DataId field of the LiveAudioBarrage struct.
func (lab *LiveAudioBarrage) SetDataId(dataId string) {
	lab.DataId = &dataId
}

// SetPublishTime sets the PublishTime field of the LiveAudioBarrage struct.
func (lab *LiveAudioBarrage) SetPublishTime(publishTime int64) {
	lab.PublishTime = &publishTime
}

// SetContent sets the Content field of the LiveAudioBarrage struct.
func (lab *LiveAudioBarrage) SetContent(content string) {
	lab.Content = &content
}

// SetHint sets the Hint field of the LiveAudioBarrage struct.
func (lab *LiveAudioBarrage) SetHint(hint []string) {
	lab.Hint = &hint
}

// SetLevel sets the Level field of the LiveAudioBarrage struct.
func (lab *LiveAudioBarrage) SetLevel(level int) {
	lab.Level = &level
}

// SetSpamType sets the SpamType field of the LiveAudioBarrage struct.
func (lab *LiveAudioBarrage) SetSpamType(spamType int) {
	lab.SpamType = &spamType
}
