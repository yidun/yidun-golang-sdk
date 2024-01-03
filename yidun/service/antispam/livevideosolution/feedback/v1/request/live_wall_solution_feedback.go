package request

// LiveWallSolutionFeedback represents live wall solution feedback.
// LiveWallSolutionFeedback represents live wall solution feedback.
type LiveWallSolutionFeedback struct {
	// 直播UUID
	TaskId *string `json:"taskId,omitempty"`
	// 观看人数,-1或者null代表不更新
	ViewCount *int `json:"viewCount,omitempty"`
	// 礼物数，-1或者null代表不更新
	GiftCount *int `json:"giftCount,omitempty"`
	// 弹幕数,-1或者null代表不更新
	BarrageCount *int `json:"barrageCount,omitempty"`
	// 直播状态 100：直播结束 @see LiveVideoFeedBackStatusEnum
	Status *int `json:"status,omitempty"`
	// 新增分数
	AddScore *int `json:"addScore,omitempty"`
	// 直播标签, 优化乱序
	LiveTag *LiveTag `json:"liveTag,omitempty"`
}

// LiveTag represents live tag.
type LiveTag struct {
	// 直播标签数组
	Labels *[]string `json:"labels,omitempty"`
	// 直播标签更新时间
	UpdateTime *int64 `json:"updateTime,omitempty"`
}

// SetTaskId sets the TaskId.
func (l *LiveWallSolutionFeedback) SetTaskId(taskId string) {
	l.TaskId = &taskId
}

// SetViewCount sets the ViewCount.
func (l *LiveWallSolutionFeedback) SetViewCount(viewCount int) {
	l.ViewCount = &viewCount
}

// SetGiftCount sets the GiftCount.
func (l *LiveWallSolutionFeedback) SetGiftCount(giftCount int) {
	l.GiftCount = &giftCount
}

// SetBarrageCount sets the BarrageCount.
func (l *LiveWallSolutionFeedback) SetBarrageCount(barrageCount int) {
	l.BarrageCount = &barrageCount
}

// SetStatus sets the Status.
func (l *LiveWallSolutionFeedback) SetStatus(status int) {
	l.Status = &status
}

// SetAddScore sets the AddScore.
func (l *LiveWallSolutionFeedback) SetAddScore(addScore int) {
	l.AddScore = &addScore
}

// SetLiveTag sets the LiveTag.
func (l *LiveWallSolutionFeedback) SetLiveTag(liveTag LiveTag) {
	l.LiveTag = &liveTag
}

// SetLabels sets the Labels.
func (l *LiveTag) SetLabels(labels []string) {
	l.Labels = &labels
}

// SetUpdateTime sets the UpdateTime.
func (l *LiveTag) SetUpdateTime(updateTime int64) {
	l.UpdateTime = &updateTime
}
