package common

import (
	text "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/text/v5/check/sync/single"
)

type AigcStreamCheckResult struct {
	// 当前流式检测活跃会话的主数据 sessionTaskId
	SessionTaskId string `json:"sessionTaskId,omitempty"`
	// 当前流式检测活跃会话的会话ID
	SessionId string `json:"sessionId,omitempty"`
	// 文本内容安全检测结果
	Antispam *text.Antispam `json:"antispam,omitempty"`
	// 文本语种检测结果
	Language *text.Language `json:"language,omitempty"`
	// aigc提示分析结果
	AigcPrompt *text.AigcPrompt `json:"aigcPrompt,omitempty"`
}

// 推送的事件类型

const (
	// 流式输出检测
	StreamOutputCheck int = 1
	// 输入检测
	InputCheck int = 2
	// 会话结束
	SessionEnd int = 3
)
