package response

import "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/media/v2/common/response/envidence"

// 反垃圾结果
type MediaAntispamResponse struct {
	// 机审结果
	DataId                 *string                           `json:"dataId,omitempty"`
	TaskId                 *string                           `json:"taskId,omitempty"`
	Suggestion             *int                              `json:"suggestion,omitempty"`      // 0:无结果（检测失败）1:正常 2:异常 3：疑似
	SuggestionLevel        *int                              `json:"suggestionLevel,omitempty"` // 嫌疑级别，针对部分定制策略支持命中嫌疑时返回，1：低嫌疑，2：高嫌疑
	RiskDescription        *string                           `json:"riskDescription,omitempty"` // 风险描述，拼接 label｜secondlabel｜thirdlabel 的中文实时示意
	Label                  *int                              `json:"label,omitempty"`           // 垃圾类别
	SecondLabel            *string                           `json:"secondLabel,omitempty"`     // 二级机审标签
	ThirdLabel             *string                           `json:"thirdLabel,omitempty"`      // 三级机审标签
	ResultType             *int                              `json:"resultType,omitempty"`      // 审核类型，1:机器检测，2:审核
	Callback               *string                           `json:"callback,omitempty"`
	CheckStatus            *int                              `json:"checkStatus,omitempty"`            // 检测状态, 1:检测中 2:检测成功 3:检测失败
	Evidences              *envidence.MediaAntispamEvidence  `json:"evidences,omitempty"`              // 底层业务返回结果
	SolutionEnrichEvidence *envidence.SolutionEnrichEvidence `json:"solutionEnrichEvidence,omitempty"` // 解决方案维度-失败集合
}
