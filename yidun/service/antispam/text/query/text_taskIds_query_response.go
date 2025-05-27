package query

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
)

type TextTaskIdsQueryResponse struct {
	*types.CommonResponse
    Result []*TextTaskIdsQueryResult `json:"result,omitempty"`
}

type TextTaskIdsQueryResult struct {
    Callback     *string     `json:"callback,omitempty"`
    Action       *int        `json:"action,omitempty"`
    TaskID       *string     `json:"taskId,omitempty"`
    Status       *int        `json:"status,omitempty"`
    Labels       []*LabelInfo `json:"labels,omitempty"`
    CensorSource *int        `json:"censorSource,omitempty"`
    CensorTime   *int64      `json:"censorTime,omitempty"`
    Censor       *string     `json:"censor,omitempty"`
}

type LabelInfo struct {
    Label    *int      `json:"label,omitempty"`
    Level    *int      `json:"level,omitempty"`
    SubLabel *string   `json:"subLabel,omitempty"`
    RiskDescription *string   `json:"riskDescription,omitempty"`
    Details  *HintInfo `json:"details,omitempty"`
}

type HintInfo struct {
    Hint *[]string `json:"hint,omitempty"`
}
