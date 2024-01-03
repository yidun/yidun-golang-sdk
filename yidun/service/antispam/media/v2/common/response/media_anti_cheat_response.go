package response

type MediaAntiCheatResponse struct {
	// 反作弊结果
	// 0:无结果（检测失败）1:正常 2:异常 3：疑似
	Action  *int     `json:"action,omitempty"`
	TaskId  *string  `json:"taskId,omitempty"`
	HitInfo *HitInfo `json:"hitInfo,omitempty"`
	Code    *int     `json:"code,omitempty"`
}

type HitInfo struct {
	// 见官网文档
	HitType *int `json:"hitType,omitempty"`
	// 见官网文档
	HitMsg *string `json:"hitMsg,omitempty"`
}
