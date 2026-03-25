package single

import (
	"encoding/json"
	"testing"
)

// TestLoopVariablePointerTrap 测试循环变量指针陷阱的修复
func TestLoopVariablePointerTrap(t *testing.T) {
	// 准备不同的测试JSON数据
	jsonData1 := `{"result": {"antispam": {"taskId": "task1", "dataId": "data1"}}}`
	jsonData2 := `{"result": {"antispam": {"taskId": "task2", "dataId": "data2"}}}`
	jsonData3 := `{"result": {"antispam": {"taskId": "task3", "dataId": "data3"}}}`

	testData := []string{jsonData1, jsonData2, jsonData3}

	// 测试用例1：错误的方式（在循环外定义变量会导致指针复用问题）
	t.Run("IncorrectWayWithOuterVariable", func(t *testing.T) {
		responses := make([]*TextCheckResponse, len(testData))
		var response TextCheckResponse  // 错误：在循环外定义，所有指针都指向同一个变量

		for i, jsonStr := range testData {
			if err := json.Unmarshal([]byte(jsonStr), &response); err != nil {
				t.Fatal(err)
			}
			responses[i] = &response  // 错误：所有指针都指向同一个变量
		}

		// 验证问题：所有响应都会有相同的TaskID（最后一个）
		expectedTaskID := "task3"  // 最后一次解析的值
		for i, resp := range responses {
			if resp.Result == nil || resp.Result.Antispam == nil {
				t.Errorf("响应 %d 的数据为空", i)
				continue
			}

			actualTaskID := *resp.Result.Antispam.TaskID
			if actualTaskID != expectedTaskID {
				t.Errorf("错误的方式应该导致所有响应都有相同的TaskID，期望: %s, 实际: %s (索引: %d)",
					expectedTaskID, actualTaskID, i)
			}
		}

		t.Logf("错误方式验证通过：所有 %d 个响应都指向最后解析的数据 (TaskID: %s)",
			len(responses), expectedTaskID)
	})

	// 测试用例2：正确的方式（修复后的代码）
	t.Run("CorrectWay", func(t *testing.T) {
		responses := make([]*TextCheckResponse, len(testData))

		for i, jsonStr := range testData {
			response := &TextCheckResponse{}  // 正确：每次创建新的变量
			if err := json.Unmarshal([]byte(jsonStr), response); err != nil {
				t.Fatal(err)
			}
			responses[i] = response  // 正确：每个指针指向不同的变量
		}

		// 验证修复：每个响应都应该有对应的TaskID
		expectedTaskIDs := []string{"task1", "task2", "task3"}
		for i, resp := range responses {
			if resp.Result == nil || resp.Result.Antispam == nil {
				t.Errorf("响应 %d 的数据为空", i)
				continue
			}

			actualTaskID := *resp.Result.Antispam.TaskID
			expectedTaskID := expectedTaskIDs[i]

			if actualTaskID != expectedTaskID {
				t.Errorf("响应 %d 的TaskID不正确，期望: %s, 实际: %s",
					i, expectedTaskID, actualTaskID)
			}
		}

		t.Logf("正确方式验证通过：%d 个响应都有正确的独立数据", len(responses))
	})

	// 测试用例3：另一种正确的方式（使用复制）
	t.Run("CorrectWayWithCopy", func(t *testing.T) {
		responses := make([]TextCheckResponse, len(testData))  // 值类型切片

		for i, jsonStr := range testData {
			var response TextCheckResponse
			if err := json.Unmarshal([]byte(jsonStr), &response); err != nil {
				t.Fatal(err)
			}
			responses[i] = response  // 复制值而不是指针
		}

		// 验证：每个响应都应该有对应的TaskID
		expectedTaskIDs := []string{"task1", "task2", "task3"}
		for i, resp := range responses {
			if resp.Result == nil || resp.Result.Antispam == nil {
				t.Errorf("响应 %d 的数据为空", i)
				continue
			}

			actualTaskID := *resp.Result.Antispam.TaskID
			expectedTaskID := expectedTaskIDs[i]

			if actualTaskID != expectedTaskID {
				t.Errorf("响应 %d 的TaskID不正确，期望: %s, 实际: %s",
					i, expectedTaskID, actualTaskID)
			}
		}

		t.Logf("值复制方式验证通过：%d 个响应都有正确的独立数据", len(responses))
	})
}

// 性能测试：基线数据（不含新增的 3 个模块）
const baselineJSON = `{
	"result": {
		"antispam": {
			"taskId": "task123",
			"dataId": "data456",
			"label": 100,
			"suggestion": 0,
			"labels": [{
				"label": 200,
				"level": 1,
				"rate": 0.95,
				"subLabels": [{
					"subLabel": "sublabel1",
					"rate": 0.8,
					"details": {
						"keywords": [{"word": "test"}],
						"libInfos": [{"type": 1, "entity": "entity1"}]
					}
				}]
			}]
		},
		"aigcPrompt": {
			"taskId": "task789",
			"dataId": "data012",
			"details": [{
				"type": 1,
				"answer": "test answer",
				"source": 0
			}]
		}
	}
}`

// 性能测试：包含所有新增模块的完整数据
const fullJSON = `{
	"result": {
		"antispam": {
			"taskId": "task123",
			"dataId": "data456",
			"label": 100,
			"suggestion": 0,
			"relateContent": "相关内容数据",
			"hitSource": 5,
			"labels": [{
				"label": 200,
				"level": 1,
				"rate": 0.95,
				"hitType": 1,
				"strategyType": 2,
				"hitResult": 3,
				"subLabels": [{
					"subLabel": "sublabel1",
					"rate": 0.8,
					"details": {
						"keywords": [{"word": "test"}],
						"libInfos": [{"type": 1, "entity": "entity1"}]
					}
				}]
			}]
		},
		"aigcPrompt": {
			"taskId": "task789",
			"dataId": "data012",
			"details": [{
				"type": 1,
				"answer": "test answer",
				"source": 0,
				"proxyAnswerType": 2
			}]
		},
		"riskControl": {
			"taskId": "task_rc123",
			"dataId": "data_rc456",
			"details": [{
				"riskType": "spam",
				"riskLevel": 2,
				"riskScore": 0.85,
				"hitInfos": [{
					"hitType": 1,
					"hitMsg": "高风险行为"
				}]
			}]
		},
		"similarText": {
			"taskId": "task_st789",
			"dataId": "data_st012",
			"details": [{
				"matchDataId": "match123",
				"rate": 0.92
			}]
		},
		"underage": {
			"taskId": "task_ua345",
			"dataId": "data_ua678",
			"details": [{
				"isUnderage": true
			}]
		}
	}
}`

// BenchmarkUnmarshal_Baseline 基线测试：不含新增模块
func BenchmarkUnmarshal_Baseline(b *testing.B) {
	data := []byte(baselineJSON)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		var response TextCheckResponse
		if err := json.Unmarshal(data, &response); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkUnmarshal_Full 完整测试：包含所有新增模块
func BenchmarkUnmarshal_Full(b *testing.B) {
	data := []byte(fullJSON)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		var response TextCheckResponse
		if err := json.Unmarshal(data, &response); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkUnmarshal_AntispamOnly 测试自定义 UnmarshalJSON 的性能影响
func BenchmarkUnmarshal_AntispamOnly(b *testing.B) {
	jsonData := `{
		"taskId": "task123",
		"dataId": "data456",
		"label": 100,
		"relateContent": "相关内容数据",
		"hitSource": 5
	}`
	data := []byte(jsonData)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		var antispam Antispam
		if err := json.Unmarshal(data, &antispam); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkUnmarshal_DeepNested 测试深度嵌套结构的性能
func BenchmarkUnmarshal_DeepNested(b *testing.B) {
	jsonData := `{
		"taskId": "task123",
		"dataId": "data456",
		"label": 100,
		"labels": [
			{
				"label": 200,
				"level": 1,
				"rate": 0.95,
				"hitType": 1,
				"strategyType": 2,
				"hitResult": 3,
				"subLabels": [
					{
						"subLabel": "sublabel1",
						"rate": 0.8,
						"details": {
							"keywords": [
								{"word": "keyword1"},
								{"word": "keyword2"},
								{"word": "keyword3"}
							],
							"libInfos": [
								{"type": 1, "entity": "entity1"},
								{"type": 2, "entity": "entity2"}
							],
							"hitInfos": [
								{"value": "hit1", "type": 1},
								{"value": "hit2", "type": 2}
							]
						}
					},
					{
						"subLabel": "sublabel2",
						"rate": 0.7,
						"details": {
							"keywords": [
								{"word": "keyword4"},
								{"word": "keyword5"}
							]
						}
					}
				]
			},
			{
				"label": 300,
				"level": 2,
				"rate": 0.85
			}
		]
	}`
	data := []byte(jsonData)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		var antispam Antispam
		if err := json.Unmarshal(data, &antispam); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkMarshal_Full 测试序列化性能
func BenchmarkMarshal_Full(b *testing.B) {
	var response TextCheckResponse
	if err := json.Unmarshal([]byte(fullJSON), &response); err != nil {
		b.Fatal(err)
	}

	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		_, err := json.Marshal(response)
		if err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkUnmarshal_LargeArray 测试大数组场景性能
func BenchmarkUnmarshal_LargeArray(b *testing.B) {
	// 模拟批量检测场景：100 条记录
	jsonData := `{"result": {"antispam": {"taskId": "task", "dataId": "data", "label": 100, "labels": [`
	for i := 0; i < 100; i++ {
		if i > 0 {
			jsonData += ","
		}
		jsonData += `{"label": 200, "level": 1, "rate": 0.95, "hitType": 1}`
	}
	jsonData += `]}}}`

	data := []byte(jsonData)
	b.ResetTimer()
	b.ReportAllocs()

	for i := 0; i < b.N; i++ {
		var response TextCheckResponse
		if err := json.Unmarshal(data, &response); err != nil {
			b.Fatal(err)
		}
	}
}

// BenchmarkMemoryFootprint 测试内存占用
func BenchmarkMemoryFootprint(b *testing.B) {
	data := []byte(fullJSON)

	b.ResetTimer()
	b.ReportAllocs()

	responses := make([]*TextCheckResponse, b.N)
	for i := 0; i < b.N; i++ {
		// 修复：在每次循环中创建新的变量，避免指针复用
		response := &TextCheckResponse{}
		if err := json.Unmarshal(data, response); err != nil {
			b.Fatal(err)
		}
		responses[i] = response
	}

	// 防止编译器优化掉未使用的变量
	_ = responses
}
