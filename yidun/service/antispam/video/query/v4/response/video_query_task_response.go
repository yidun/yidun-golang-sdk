package response

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/video/callback/v4/response"
)

type VideoQueryTaskV4Response struct {
	*types.CommonResponse
	Result *[]VideoQueryTaskV4Result `json:"result"`
}

// VideoQueryTaskV4Result 视频查询任务结果
type VideoQueryTaskV4Result struct {
	TaskId   *string                                   `json:"taskId,omitempty"`   // 任务id，64位字符串
	Status   *int                                      `json:"status,omitempty"`   // 任务状态
	Antispam *response.VideoCallbackAntispamV4Response `json:"antispam,omitempty"` // 视频内容安全检测结果
}
