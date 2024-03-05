package callback

type CrawlerJobActiveCallbackResponse struct {
	Antispam *JobAntispamCallbackResponseV4 `json:"antispam,omitempty"`
	Censor   *JobCensorCallbackResponse     `json:"censor,omitempty"`
}

// 机器检测结果
type JobAntispamCallbackResponseV4 struct {
	// 任务jobId
	JobId *int64 `json:"jobId,omitempty"`
	// 任务taskId
	TaskId *string `json:"taskId,omitempty"`
	// 资源名称
	ResourceName *string `json:"resourceName,omitempty"`
	// 资源
	Resource *string `json:"resource,omitempty"`
	// 资源类型
	ResourceType *int `json:"resourceType,omitempty"`
	// 检测状态
	CheckStatus *int `json:"checkStatus,omitempty"`
	// 检测结果
	Suggestion *int `json:"suggestion,omitempty"`
	// 命中结果
	Labels *[]LabelInfo `json:"labels,omitempty"`
	// 检测完成时间
	CheckTime *int64 `json:"checkTime,omitempty"`
	// 报告地址
	ReportUrl *string `json:"reportUrl,omitempty"`
	// 检测失败原因code
	FailureReason *int `json:"failureReason,omitempty"`
	// 检测总url数量
	TotalUrlCount *int `json:"totalUrlCount,omitempty"`
	// 嫌疑url数量
	SuspectUrlCount *int `json:"suspectUrlCount,omitempty"`
	// 不通过url数量
	UnPassUrlCount *int `json:"unPassUrlCount,omitempty"`
}

// 人审回调结果
type JobCensorCallbackResponse struct {
	// 任务id
	JobId *int64 `json:"jobId,omitempty"`
	// 主站url
	SiteUrl *string `json:"siteUrl,omitempty"`
	// 检测结果
	Suggestion *int `json:"suggestion,omitempty"`
	// 命中结果
	Labels *[]LabelInfo `json:"labels,omitempty"`
	// 人工审核最后时间
	CensorTime *int `json:"censorTime,omitempty"`
	// 报告地址
	ReportUrl *string `json:"reportUrl,omitempty"`
	// 检测总url数量
	TotalUrlCount *int `json:"totalUrlCount,omitempty"`
	// 嫌疑url数量
	SuspectUrlCount *int `json:"suspectUrlCount,omitempty"`
	// 不通过url数量
	UnPassUrlCount *int `json:"unPassUrlCount,omitempty"`
}

type LabelInfo struct {
	Label *int `json:"label,omitempty"`
}
