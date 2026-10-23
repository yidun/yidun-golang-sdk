package single

import (
	"encoding/json"
	"testing"
)

// TestAntispamSubLabelDetailRuleUnmarshal 测试 AntispamSubLabelDetailRule 的 RuleId 字段反序列化
func TestAntispamSubLabelDetailRuleUnmarshal(t *testing.T) {
	// 测试用例1：full 路径下 rules 中的 ruleId 能被正确接收
	t.Run("RuleIdPresent", func(t *testing.T) {
		jsonData := `{
			"result": {
				"antispam": {
					"taskId": "task406",
					"dataId": "data406",
					"label": 100,
					"labels": [{
						"label": 200,
						"level": 1,
						"rate": 0.95,
						"subLabels": [{
							"subLabel": "sublabel1",
							"rate": 0.8,
							"details": {
								"rules": [{
									"name": "x",
									"ruleId": 12345
								}]
							}
						}]
					}]
				}
			}
		}`

		var response TextCheckResponse
		if err := json.Unmarshal([]byte(jsonData), &response); err != nil {
			t.Fatalf("反序列化失败: %v", err)
		}

		if response.Result == nil || response.Result.Antispam == nil {
			t.Fatal("Antispam 为空")
		}

		labels := response.Result.Antispam.Labels
		if len(labels) != 1 {
			t.Fatalf("Labels 长度不匹配，期望: 1, 实际: %d", len(labels))
		}
		if labels[0].SubLabels == nil || len(labels[0].SubLabels) != 1 {
			t.Fatalf("SubLabels 长度不匹配")
		}
		details := labels[0].SubLabels[0].Details
		if details == nil {
			t.Fatal("Details 为空")
		}
		if len(details.Rules) != 1 {
			t.Fatalf("Rules 长度不匹配，期望: 1, 实际: %d", len(details.Rules))
		}

		rule := details.Rules[0]
		if rule.Name == nil || *rule.Name != "x" {
			t.Errorf("Name 不匹配，期望: x, 实际: %v", rule.Name)
		}
		if rule.RuleId == nil || *rule.RuleId != 12345 {
			t.Errorf("RuleId 不匹配，期望: 12345, 实际: %v", rule.RuleId)
		}
	})

	// 测试用例2：JSON 缺失 ruleId 时字段应为 nil（指针类型，兼容字段缺失）
	t.Run("RuleIdMissing", func(t *testing.T) {
		jsonData := `{"name":"x"}`

		var rule AntispamSubLabelDetailRule
		if err := json.Unmarshal([]byte(jsonData), &rule); err != nil {
			t.Fatalf("反序列化失败: %v", err)
		}

		if rule.Name == nil || *rule.Name != "x" {
			t.Errorf("Name 不匹配，期望: x, 实际: %v", rule.Name)
		}
		if rule.RuleId != nil {
			t.Errorf("缺失 ruleId 时 RuleId 应该为 nil，实际: %v", *rule.RuleId)
		}
	})
}
