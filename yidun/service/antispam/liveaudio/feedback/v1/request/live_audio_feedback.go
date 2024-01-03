package request

type LiveAudioFeedback struct {
	// 直播UUID
	TaskId *string `json:"taskId,omitempty"`
	// 直播状态 100：直播结束
	Status *int `json:"status,omitempty"`
	// 背景图
	BackgroundImage *string `json:"backgroundImage,omitempty"`
	// 观看人数
	ViewCount *int `json:"viewCount,omitempty"`
	// 弹幕数
	BarrageCount *int `json:"barrageCount,omitempty"`
	// 垃圾弹幕数
	GarbageBarrageCount *int `json:"garbageBarrageCount,omitempty"`
	// 举报数
	ReportCount *int `json:"reportCount,omitempty"`
	// 新增异常值
	AddScore *int `json:"addScore,omitempty"`
	// 直播标签, 优化乱序
	LiveTag *LiveAudioTag `json:"liveTag,omitempty"`
	// 分数监控类型
	ScoreMonitorType *int `json:"scoreMonitorType,omitempty"`
}

type LiveAudioTag struct {
	// 直播标签数组
	Tag *[]string `json:"tag,omitempty"`
	// 直播标签更新时间
	UpdateTime *int64 `json:"updateTime,omitempty"`
}

// SetTaskID sets the TaskID field of the LiveAudioFeedback struct
func (f *LiveAudioFeedback) SetTaskID(taskId string) {
	f.TaskId = &taskId
}

// SetStatus sets the Status field of the LiveAudioFeedback struct
func (f *LiveAudioFeedback) SetStatus(status int) {
	f.Status = &status
}

// SetBackgroundImage sets the BackgroundImage field of the LiveAudioFeedback struct
func (f *LiveAudioFeedback) SetBackgroundImage(backgroundImage string) {
	f.BackgroundImage = &backgroundImage
}

// SetViewCount sets the ViewCount field of the LiveAudioFeedback struct
func (f *LiveAudioFeedback) SetViewCount(viewCount int) {
	f.ViewCount = &viewCount
}

// SetBarrageCount sets the BarrageCount field of the LiveAudioFeedback struct
func (f *LiveAudioFeedback) SetBarrageCount(barrageCount int) {
	f.BarrageCount = &barrageCount
}

// SetGarbageBarrageCount sets the GarbageBarrageCount field of the LiveAudioFeedback struct
func (f *LiveAudioFeedback) SetGarbageBarrageCount(garbageBarrageCount int) {
	f.GarbageBarrageCount = &garbageBarrageCount
}

// SetReportCount sets the ReportCount field of the LiveAudioFeedback struct
func (f *LiveAudioFeedback) SetReportCount(reportCount int) {
	f.ReportCount = &reportCount
}

// SetAddScore sets the AddScore field of the LiveAudioFeedback struct
func (f *LiveAudioFeedback) SetAddScore(addScore int) {
	f.AddScore = &addScore
}

// SetLiveTag sets the LiveTag field of the LiveAudioFeedback struct
func (f *LiveAudioFeedback) SetLiveTag(liveTag *LiveAudioTag) {
	f.LiveTag = liveTag
}

// SetScoreMonitorType sets the ScoreMonitorType field of the LiveAudioFeedback struct
func (f *LiveAudioFeedback) SetScoreMonitorType(scoreMonitorType int) {
	f.ScoreMonitorType = &scoreMonitorType
}

// For LiveAudioTag struct

// SetTag sets the Tag field of the LiveAudioTag struct
func (t *LiveAudioTag) SetTag(tag []string) {
	t.Tag = &tag
}

// SetUpdateTime sets the UpdateTime field of the LiveAudioTag struct
func (t *LiveAudioTag) SetUpdateTime(updateTime int64) {
	t.UpdateTime = &updateTime
}
