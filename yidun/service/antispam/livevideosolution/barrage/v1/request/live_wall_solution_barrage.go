package request

type LiveWallSolutionBarrage struct {
	// 弹幕id
	Id *string `json:"id,omitempty"`
	// 直播UUID
	TaskId *string `json:"taskId,omitempty"`
	// 直播UUID
	DataId *string `json:"dataId,omitempty"`
	// 弹幕发送时间
	PublishTime *int64 `json:"publishTime,omitempty"`
	// 弹幕内容
	Content *string `json:"content,omitempty"`
	// 高亮标签
	Hint *[]string `json:"hint,omitempty"`
	// 0-正常 1-嫌疑 2-确定
	Level *int `json:"level,omitempty"`
	// Spam type
	SpamType *int `json:"spamType,omitempty"`
}

// SetId sets the Id field.
func (l *LiveWallSolutionBarrage) SetId(id string) {
	l.Id = &id
}

// SetTaskId sets the TaskId field.
func (l *LiveWallSolutionBarrage) SetTaskId(taskId string) {
	l.TaskId = &taskId
}

// SetDataId sets the DataId field.
func (l *LiveWallSolutionBarrage) SetDataId(dataId string) {
	l.DataId = &dataId
}

// SetPublishTime sets the PublishTime field.
func (l *LiveWallSolutionBarrage) SetPublishTime(publishTime int64) {
	l.PublishTime = &publishTime
}

// SetContent sets the Content field.
func (l *LiveWallSolutionBarrage) SetContent(content string) {
	l.Content = &content
}

// SetHint sets the Hint field.
func (l *LiveWallSolutionBarrage) SetHint(hint []string) {
	l.Hint = &hint
}

// SetLevel sets the Level field.
func (l *LiveWallSolutionBarrage) SetLevel(level int) {
	l.Level = &level
}

// SetSpamType sets the SpamType field.
func (l *LiveWallSolutionBarrage) SetSpamType(spamType int) {
	l.SpamType = &spamType
}
