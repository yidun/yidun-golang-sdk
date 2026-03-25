package single

import (
	"encoding/json"
	"testing"
)

// TestNilComparisonAndErrorHandlingFixes 验证nil比较和错误处理的修复
func TestNilComparisonAndErrorHandlingFixes(t *testing.T) {
	// 测试用例1：验证指针到切片的正确nil检查
	t.Run("PointerToSliceNilCheck", func(t *testing.T) {
		// 测试Details指针为nil的情况
		var ap AigcPrompt
		if ap.Details != nil {
			t.Error("新建的AigcPrompt.Details应该为nil")
		}

		// 测试Details指针不为nil但切片为nil的情况
		ap.Details = new([]*AigcPromptDetail)
		if ap.Details == nil {
			t.Error("Details指针不应该为nil")
		}
		if *ap.Details != nil {
			t.Error("Details切片应该为nil")
		}

		// 测试Details指针和切片都不为nil的情况
		details := []*AigcPromptDetail{
			{Type: ptrInt(1), Answer: ptrString("test")},
		}
		ap.Details = &details

		if ap.Details == nil {
			t.Error("Details指针不应该为nil")
		}
		if *ap.Details == nil {
			t.Error("Details切片不应该为nil")
		}
		if len(*ap.Details) != 1 {
			t.Errorf("Details切片长度应该为1，实际: %d", len(*ap.Details))
		}
	})

	// 测试用例2：验证JSON反序列化错误处理
	t.Run("JSONUnmarshalErrorHandling", func(t *testing.T) {
		// 测试无效JSON
		invalidJSON := `{"invalid": json}`

		var response TextCheckResponse
		err := json.Unmarshal([]byte(invalidJSON), &response)
		if err == nil {
			t.Error("无效JSON应该返回错误")
		}
		t.Logf("正确捕获JSON错误: %v", err)

		// 测试序列化后重新解析的错误处理
		validData := Antispam{
			TaskID: ptrString("test"),
			Label:  ptrInt(100),
		}

		data, err := json.Marshal(validData)
		if err != nil {
			t.Fatalf("序列化失败: %v", err)
		}

		var result map[string]interface{}
		err = json.Unmarshal(data, &result)
		if err != nil {
			t.Fatalf("反序列化结果失败: %v", err)
		}

		// 验证结果正确
		if result["taskId"] != "test" {
			t.Errorf("taskId不匹配，期望: test, 实际: %v", result["taskId"])
		}
	})

	// 测试用例3：验证修复后的AigcPromptDetailFields测试逻辑
	t.Run("FixedAigcPromptDetailFields", func(t *testing.T) {
		jsonData := `{
			"result": {
				"aigcPrompt": {
					"taskId": "fix_test",
					"dataId": "fix_data",
					"details": [{
						"type": 1,
						"answer": "修复测试答案",
						"source": 1,
						"proxyAnswerType": 3
					}]
				}
			}
		}`

		var response TextCheckResponse
		if err := json.Unmarshal([]byte(jsonData), &response); err != nil {
			t.Fatalf("反序列化失败: %v", err)
		}

		// 验证响应结构不为空
		if response.Result == nil {
			t.Fatal("Result不应该为空")
		}
		if response.Result.AigcPrompt == nil {
			t.Fatal("AigcPrompt不应该为空")
		}

		ap := response.Result.AigcPrompt

		// 使用修复后的正确nil检查逻辑
		if ap.Details == nil {
			t.Fatal("Details指针不应该为空")
		}
		if *ap.Details == nil {
			t.Fatal("Details切片不应该为空")
		}
		if len(*ap.Details) != 1 {
			t.Fatalf("Details长度应该为1，实际: %d", len(*ap.Details))
		}

		detail := (*ap.Details)[0]
		if detail.ProxyAnswerType == nil || *detail.ProxyAnswerType != 3 {
			t.Errorf("ProxyAnswerType不匹配，期望: 3, 实际: %v", detail.ProxyAnswerType)
		}
		if detail.Source == nil || *detail.Source != 1 {
			t.Errorf("Source不匹配，期望: 1, 实际: %v", detail.Source)
		}
	})
}

// TestRiskControlUnmarshal 测试 RiskControl 模块的反序列化
func TestRiskControlUnmarshal(t *testing.T) {
	jsonData := `{
		"result": {
			"riskControl": {
				"taskId": "task123",
				"dataId": "data456",
				"details": [{
					"riskLevel": 2,
					"hitInfos": [{
						"type": "SPAM",
						"name": "垃圾内容",
						"desc": "高风险行为"
					}]
				}]
			}
		}
	}`

	var response TextCheckResponse
	if err := json.Unmarshal([]byte(jsonData), &response); err != nil {
		t.Fatalf("反序列化失败: %v", err)
	}

	if response.Result == nil || response.Result.RiskControl == nil {
		t.Fatal("RiskControl 为空")
	}

	rc := response.Result.RiskControl
	if rc.TaskID == nil || *rc.TaskID != "task123" {
		t.Errorf("TaskID 不匹配，期望: task123, 实际: %v", rc.TaskID)
	}

	if len(rc.Details) != 1 {
		t.Fatalf("Details 长度不匹配，期望: 1, 实际: %d", len(rc.Details))
	}

	detail := rc.Details[0]
	if detail.RiskLevel == nil || *detail.RiskLevel != 2 {
		t.Errorf("RiskLevel 不匹配")
	}
	if len(detail.HitInfos) != 1 {
		t.Fatalf("HitInfos 长度不匹配")
	}
	hitInfo := detail.HitInfos[0]
	if hitInfo.Type == nil || *hitInfo.Type != "SPAM" {
		t.Errorf("HitInfo.Type 不匹配")
	}
	if hitInfo.Name == nil || *hitInfo.Name != "垃圾内容" {
		t.Errorf("HitInfo.Name 不匹配")
	}
	if hitInfo.Desc == nil || *hitInfo.Desc != "高风险行为" {
		t.Errorf("HitInfo.Desc 不匹配")
	}
}

// TestSimilarTextUnmarshal 测试 SimilarText 模块的反序列化
func TestSimilarTextUnmarshal(t *testing.T) {
	jsonData := `{
		"result": {
			"similarText": {
				"taskId": "task789",
				"dataId": "data012",
				"details": [{
					"matchDataId": "match123",
					"rate": 0.92
				}]
			}
		}
	}`

	var response TextCheckResponse
	if err := json.Unmarshal([]byte(jsonData), &response); err != nil {
		t.Fatalf("反序列化失败: %v", err)
	}

	if response.Result == nil || response.Result.SimilarText == nil {
		t.Fatal("SimilarText 为空")
	}

	st := response.Result.SimilarText
	if st.TaskID == nil || *st.TaskID != "task789" {
		t.Errorf("TaskID 不匹配")
	}

	if len(st.Details) != 1 {
		t.Fatalf("Details 长度不匹配")
	}

	detail := st.Details[0]
	if detail.MatchDataID == nil || *detail.MatchDataID != "match123" {
		t.Errorf("MatchDataID 不匹配")
	}
	if detail.SimilarityScore == nil || *detail.SimilarityScore != 0.92 {
		t.Errorf("SimilarityScore 不匹配")
	}
}

// TestUnderageUnmarshal 测试 Underage 模块的反序列化
func TestUnderageUnmarshal(t *testing.T) {
	jsonData := `{
		"result": {
			"underage": {
				"taskId": "task345",
				"dataId": "data678",
				"details": [{
					"type": 1
				}]
			}
		}
	}`

	var response TextCheckResponse
	if err := json.Unmarshal([]byte(jsonData), &response); err != nil {
		t.Fatalf("反序列化失败: %v", err)
	}

	if response.Result == nil || response.Result.Underage == nil {
		t.Fatal("Underage 为空")
	}

	ua := response.Result.Underage
	if ua.TaskID == nil || *ua.TaskID != "task345" {
		t.Errorf("TaskID 不匹配")
	}

	if len(ua.Details) != 1 {
		t.Fatalf("Details 长度不匹配")
	}

	detail := ua.Details[0]
	if detail.Type == nil || *detail.Type != 1 {
		t.Errorf("Type 不匹配，期望: 1, 实际: %v", detail.Type)
	}
}

// TestAntispamBasicFields 测试 Antispam 基本字段反序列化
func TestAntispamBasicFields(t *testing.T) {
	// 测试正确的后端字段名（单数形式）
	jsonData := `{
		"result": {
			"antispam": {
				"taskId": "task111",
				"dataId": "data222",
				"label": 100,
				"relateContent": "相关内容数据",
				"hitSource": 5
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

	as := response.Result.Antispam

	// 验证新字段能正确接收数据
	if as.RelateContent == nil || *as.RelateContent != "相关内容数据" {
		t.Errorf("RelateContent 不匹配，期望: 相关内容数据, 实际: %v", as.RelateContent)
	}

	// 验证 HitSource 新字段
	if as.HitSource == nil || *as.HitSource != 5 {
		t.Errorf("HitSource 不匹配，期望: 5, 实际: %v", as.HitSource)
	}

	// 注意：老字段的同步现在由 HTTP 层处理，这里只验证新字段能正确接收数据
}

// TestAigcPromptDetailFields 测试 AigcPromptDetail 的新增字段
func TestAigcPromptDetailFields(t *testing.T) {
	jsonData := `{
		"result": {
			"aigcPrompt": {
				"taskId": "task555",
				"dataId": "data666",
				"details": [{
					"type": 1,
					"answer": "这是答案",
					"source": 0,
					"libId": "lib123",
					"answerId": "ans456",
					"proxyAnswerType": 2
				}]
			}
		}
	}`

	var response TextCheckResponse
	if err := json.Unmarshal([]byte(jsonData), &response); err != nil {
		t.Fatalf("反序列化失败: %v", err)
	}

	if response.Result == nil || response.Result.AigcPrompt == nil {
		t.Fatal("AigcPrompt 为空")
	}

	ap := response.Result.AigcPrompt
	if ap.Details == nil {
		t.Fatal("Details 指针为空")
	}
	if *ap.Details == nil {
		t.Fatal("Details 切片为空")
	}

	if len(*ap.Details) != 1 {
		t.Fatalf("Details 长度不匹配，期望: 1, 实际: %d", len(*ap.Details))
	}

	detail := (*ap.Details)[0]
	if detail.ProxyAnswerType == nil || *detail.ProxyAnswerType != 2 {
		t.Errorf("ProxyAnswerType 不匹配，期望: 2, 实际: %v", detail.ProxyAnswerType)
	}
	if detail.Source == nil || *detail.Source != 0 {
		t.Errorf("Source 不匹配")
	}
	if detail.LibId == nil || *detail.LibId != "lib123" {
		t.Errorf("LibId 不匹配")
	}
}

// TestAntispamLabelP2Fields 测试 AntispamLabel 的 P2 优先级字段
func TestAntispamLabelP2Fields(t *testing.T) {
	jsonData := `{
		"result": {
			"antispam": {
				"taskId": "task777",
				"dataId": "data888",
				"label": 100,
				"labels": [{
					"label": 200,
					"level": 1,
					"rate": 0.95,
					"hitType": 1,
					"strategyType": 2,
					"hitResult": 3
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
		t.Fatalf("Labels 长度不匹配")
	}

	label := labels[0]
	if label.HitType == nil || *label.HitType != 1 {
		t.Errorf("HitType 不匹配")
	}
	if label.StrategyType == nil || *label.StrategyType != 2 {
		t.Errorf("StrategyType 不匹配，期望: 2, 实际: %v", label.StrategyType)
	}
	if label.HitResult == nil || *label.HitResult != 3 {
		t.Errorf("HitResult 不匹配，期望: 3, 实际: %v", label.HitResult)
	}
}


// TestRelatedHitTypeSerialization 测试 RelatedHitType 字段序列化（修复 P1-003）
func TestRelatedHitTypeSerialization(t *testing.T) {
	jsonData := `{
		"taskId": "task123",
		"dataId": "data456",
		"label": 100,
		"isRelatedHit": true,
		"relatedHitType": 2
	}`

	var as Antispam
	if err := json.Unmarshal([]byte(jsonData), &as); err != nil {
		t.Fatal(err)
	}

	// 验证字段被正确解析
	if as.RelatedHitType == nil {
		t.Fatal("RelatedHitType 未被解析")
	}
	if *as.RelatedHitType != 2 {
		t.Errorf("RelatedHitType 值错误，期望: 2，实际: %d", *as.RelatedHitType)
	}

	// 验证字段能正确序列化
	data, err := json.Marshal(&as)
	if err != nil {
		t.Fatal(err)
	}

	var result map[string]interface{}
	if err := json.Unmarshal(data, &result); err != nil {
		t.Fatalf("序列化结果解析失败: %v", err)
	}

	if result["relatedHitType"] != float64(2) {
		t.Errorf("RelatedHitType 序列化失败，实际: %v", result["relatedHitType"])
	}
}


// 辅助函数
func ptrString(s string) *string {
	return &s
}

func ptrInt(i int) *int {
	return &i
}