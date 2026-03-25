package single

import (
	"encoding/json"
	"runtime"
	"testing"

	"github.com/yidun/yidun-golang-sdk/yidun/core/client"
)

// TestErrorHandlingInPerformanceTests 验证性能测试中的错误处理修复
func TestErrorHandlingInPerformanceTests(t *testing.T) {
	// 测试用例1：测试无效JSON会被正确处理
	t.Run("InvalidJSONHandling", func(t *testing.T) {
		invalidJSON := `{"invalid": json}`

		// 测试Marshal基准测试的准备阶段会正确处理错误
		var response TextCheckResponse
		err := json.Unmarshal([]byte(invalidJSON), &response)
		if err == nil {
			t.Error("期望JSON解析错误，但没有发生错误")
		}
		t.Logf("正确捕获到JSON错误: %v", err)
	})

	// 测试用例2：验证正确的JSON能被正常处理
	t.Run("ValidJSONHandling", func(t *testing.T) {
		validJSON := `{"result": {"antispam": {"taskId": "test123"}}}`

		var response TextCheckResponse
		err := json.Unmarshal([]byte(validJSON), &response)
		if err != nil {
			t.Errorf("有效JSON解析失败: %v", err)
		}

		// 验证数据解析正确
		if response.Result == nil {
			t.Error("Result应该不为空")
		}
		if response.Result != nil && response.Result.Antispam == nil {
			t.Error("Antispam应该不为空")
		}
		if response.Result != nil && response.Result.Antispam != nil &&
		   (response.Result.Antispam.TaskID == nil || *response.Result.Antispam.TaskID != "test123") {
			t.Error("TaskID解析错误")
		}
	})
}

// 测试不同指针深度的内存影响
type ValueBasedStruct struct {
	TaskID  string
	DataID  string
	Details []ValueBasedDetail
}

type ValueBasedDetail struct {
	RiskType  string
	RiskLevel int
	RiskScore float64
}

type PointerBasedStruct struct {
	TaskID  *string
	DataID  *string
	Details []*PointerBasedDetail
}

type PointerBasedDetail struct {
	RiskType  *string
	RiskLevel *int
	RiskScore *float64
}

// BenchmarkPointerVsValue_Unmarshal 对比指针类型和值类型的反序列化性能
func BenchmarkPointerVsValue_Unmarshal(b *testing.B) {
	jsonData := `{
		"taskId": "task123",
		"dataId": "data456",
		"details": [
			{"riskType": "spam", "riskLevel": 2, "riskScore": 0.85},
			{"riskType": "abuse", "riskLevel": 1, "riskScore": 0.65}
		]
	}`

	b.Run("ValueBased", func(b *testing.B) {
		data := []byte(jsonData)
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			var result ValueBasedStruct
			if err := json.Unmarshal(data, &result); err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("PointerBased", func(b *testing.B) {
		data := []byte(jsonData)
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			var result PointerBasedStruct
			if err := json.Unmarshal(data, &result); err != nil {
				b.Fatal(err)
			}
		}
	})
}

// BenchmarkCompatibilityProcessing_Impact 测试兼容性处理的性能开销
func BenchmarkCompatibilityProcessing_Impact(b *testing.B) {
	jsonData := `{
		"taskId": "task123",
		"dataId": "data456",
		"label": 100,
		"relateContent": "相关内容",
		"hitSource": 5,
		"suggestion": 0
	}`

	processor := client.GetGlobalCompatibilityProcessor()

	b.Run("WithCompatibilityProcessing", func(b *testing.B) {
		data := []byte(jsonData)
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			var antispam Antispam
			if err := json.Unmarshal(data, &antispam); err != nil {
				b.Fatal(err)
			}
			// 应用兼容性处理
			if err := processor.ProcessForDeserialization(&antispam); err != nil {
				b.Fatal(err)
			}
			// 验证字段同步逻辑执行
			if antispam.RelateContent != nil && antispam.RelatedContents == nil {
				b.Error("字段同步失败")
			}
		}
	})

	b.Run("WithoutCompatibilityProcessing", func(b *testing.B) {
		data := []byte(jsonData)
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			var antispam Antispam
			if err := json.Unmarshal(data, &antispam); err != nil {
				b.Fatal(err)
			}
			// 不应用兼容性处理，只验证基本 JSON 反序列化性能
		}
	})
}

// BenchmarkScalability_10x 测试 10 倍数据量的性能
func BenchmarkScalability_10x(b *testing.B) {
	// 生成 10 倍数据量的 JSON
	jsonData := `{"result": {"riskControl": {"taskId": "task", "dataId": "data", "details": [`
	for i := 0; i < 100; i++ {
		if i > 0 {
			jsonData += ","
		}
		jsonData += `{"riskType": "spam", "riskLevel": 2, "riskScore": 0.85, "hitInfos": [{"hitType": 1, "hitMsg": "msg"}]}`
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

// BenchmarkScalability_100x 测试 100 倍数据量的性能
func BenchmarkScalability_100x(b *testing.B) {
	// 生成 100 倍数据量的 JSON
	jsonData := `{"result": {"riskControl": {"taskId": "task", "dataId": "data", "details": [`
	for i := 0; i < 1000; i++ {
		if i > 0 {
			jsonData += ","
		}
		jsonData += `{"riskType": "spam", "riskLevel": 2, "riskScore": 0.85}`
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

// BenchmarkConcurrency 测试并发场景性能
func BenchmarkConcurrency(b *testing.B) {
	data := []byte(fullJSON)

	b.RunParallel(func(pb *testing.PB) {
		for pb.Next() {
			var response TextCheckResponse
			if err := json.Unmarshal(data, &response); err != nil {
				b.Fatal(err)
			}
		}
	})
}

// TestMemoryLeakCheck 内存泄漏检测测试
func TestMemoryLeakCheck(t *testing.T) {
	data := []byte(fullJSON)

	// 记录初始内存状态
	runtime.GC()
	var m1 runtime.MemStats
	runtime.ReadMemStats(&m1)

	// 执行 1000 次反序列化
	for i := 0; i < 1000; i++ {
		var response TextCheckResponse
		if err := json.Unmarshal(data, &response); err != nil {
			t.Fatal(err)
		}
	}

	// 强制 GC 并记录最终内存状态
	runtime.GC()
	var m2 runtime.MemStats
	runtime.ReadMemStats(&m2)

	// 计算内存增长（使用 int64 避免下溢）
	allocDiff := int64(m2.Alloc) - int64(m1.Alloc)
	totalAllocDiff := int64(m2.TotalAlloc) - int64(m1.TotalAlloc)

	t.Logf("初始内存分配: %d bytes", m1.Alloc)
	t.Logf("最终内存分配: %d bytes", m2.Alloc)
	t.Logf("内存分配差异: %d bytes", allocDiff)
	t.Logf("总分配增长: %d bytes", totalAllocDiff)
	t.Logf("GC 次数: %d", m2.NumGC-m1.NumGC)
	t.Logf("平均每次操作分配: %d bytes", totalAllocDiff/1000)

	// 如果持续分配超过 10MB 且最终内存增长明显，可能存在内存泄漏
	if allocDiff > 10*1024*1024 {
		t.Errorf("可能存在内存泄漏，内存增长: %d bytes", allocDiff)
	} else if allocDiff < 0 {
		t.Logf("内存使用减少（正常现象，GC 生效）: %d bytes", -allocDiff)
	}
}

// BenchmarkOmitempty_Impact 测试 omitempty 对性能的影响
func BenchmarkOmitempty_Impact(b *testing.B) {
	// 大部分字段为空的场景
	sparseJSON := `{"result": {"riskControl": {"taskId": "task123"}}}`

	// 所有字段都有值的场景
	denseJSON := `{"result": {"riskControl": {"taskId": "task123", "dataId": "data456", "details": [{"riskType": "spam", "riskLevel": 2, "riskScore": 0.85}]}}}`

	b.Run("Sparse_Unmarshal", func(b *testing.B) {
		data := []byte(sparseJSON)
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			var response TextCheckResponse
			if err := json.Unmarshal(data, &response); err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("Dense_Unmarshal", func(b *testing.B) {
		data := []byte(denseJSON)
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			var response TextCheckResponse
			if err := json.Unmarshal(data, &response); err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("Sparse_Marshal", func(b *testing.B) {
		var response TextCheckResponse
		if err := json.Unmarshal([]byte(sparseJSON), &response); err != nil {
			b.Fatal("准备Sparse测试数据失败:", err)
		}
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_, err := json.Marshal(response)
			if err != nil {
				b.Fatal(err)
			}
		}
	})

	b.Run("Dense_Marshal", func(b *testing.B) {
		var response TextCheckResponse
		if err := json.Unmarshal([]byte(denseJSON), &response); err != nil {
			b.Fatal("准备Dense测试数据失败:", err)
		}
		b.ResetTimer()
		b.ReportAllocs()

		for i := 0; i < b.N; i++ {
			_, err := json.Marshal(response)
			if err != nil {
				b.Fatal(err)
			}
		}
	})
}
