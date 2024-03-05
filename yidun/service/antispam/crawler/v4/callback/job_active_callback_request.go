package callback

// JobActiveCallbackV4Request represents the active callback request for jobs.
type JobActiveCallbackV4Request struct {
	Antispam *JobAntispamCallbackV4Response `json:"antispam,omitempty"`
	Censor   *JobCensorCallbackV4Response   `json:"censor,omitempty"`
}

// JobAntispamCallbackV4Response represents the machine detection result.
type JobAntispamCallbackV4Response struct {
	JobId           *int64       `json:"jobId,omitempty"`
	TaskId          *string      `json:"taskId,omitempty"`
	ResourceName    *string      `json:"resourceName,omitempty"`
	Resource        *string      `json:"resource,omitempty"`
	ResourceType    *int         `json:"resourceType,omitempty"`
	CheckStatus     *int         `json:"checkStatus,omitempty"`
	Suggestion      *int         `json:"suggestion,omitempty"`
	Labels          *[]LabelInfo `json:"labels,omitempty"`
	CheckTime       *int64       `json:"checkTime,omitempty"`
	ReportUrl       *string      `json:"reportUrl,omitempty"`
	FailureReason   *int         `json:"failureReason,omitempty"`
	TotalUrlCount   *int         `json:"totalUrlCount,omitempty"`
	SuspectUrlCount *int         `json:"suspectUrlCount,omitempty"`
	UnPassUrlCount  *int         `json:"unPassUrlCount,omitempty"`
}

// JobCensorCallbackV4Response represents the manual review result.
type JobCensorCallbackV4Response struct {
	JobId           *int64       `json:"jobId,omitempty"`
	SiteUrl         *string      `json:"siteUrl,omitempty"`
	Suggestion      *int         `json:"suggestion,omitempty"`
	Labels          *[]LabelInfo `json:"labels,omitempty"`
	CensorTime      *int64       `json:"censorTime,omitempty"`
	ReportUrl       *string      `json:"reportUrl,omitempty"`
	TotalUrlCount   *int         `json:"totalUrlCount,omitempty"`
	SuspectUrlCount *int         `json:"suspectUrlCount,omitempty"`
	UnPassUrlCount  *int         `json:"unPassUrlCount,omitempty"`
}

// LabelInfo represents the detection label information.
type LabelInfo struct {
	Label *int `json:"label,omitempty"`
}
