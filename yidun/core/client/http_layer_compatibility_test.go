package client

import (
	"net/http"
	"testing"

	"github.com/yidun/yidun-golang-sdk/yidun/core/auth"
	httpCore "github.com/yidun/yidun-golang-sdk/yidun/core/http"
)

// 测试用的Antispam结构体，模拟真实的结构
type TestAntispam struct {
	TaskID          *string `json:"taskId"`
	DataID          *string `json:"dataId"`
	Label           *int    `json:"label"`
	RelateContent   *string `json:"relateContent,omitempty"`
	RelatedContents *string `json:"-"`
	HitSource       *int    `json:"hitSource,omitempty"`
	HitSources      *int    `json:"-"`
}

type TestTextCheckResult struct {
	Antispam *TestAntispam `json:"antispam"`
}

type TestTextCheckResponse struct {
	Code   *int                 `json:"code"`
	Msg    *string              `json:"msg"`
	Result *TestTextCheckResult `json:"result"`
}

// TestCompatibilityProcessorBasic 测试基本的兼容性处理
func TestCompatibilityProcessorBasic(t *testing.T) {
	processor := GetGlobalCompatibilityProcessor()

	t.Run("NewFieldToOldField", func(t *testing.T) {
		// 新字段有值，同步到老字段
		antispam := &TestAntispam{
			TaskID:        stringPtr("test"),
			RelateContent: stringPtr("新字段内容"),
			HitSource:     intPtr(5),
		}

		processor.ProcessForDeserialization(antispam)

		if antispam.RelatedContents == nil || *antispam.RelatedContents != "新字段内容" {
			t.Error("RelateContent -> RelatedContents 同步失败")
		}
		if antispam.HitSources == nil || *antispam.HitSources != 5 {
			t.Error("HitSource -> HitSources 同步失败")
		}
	})

	t.Run("OldFieldToNewField", func(t *testing.T) {
		// 老字段有值，新字段为空，反向同步
		antispam := &TestAntispam{
			TaskID:          stringPtr("test"),
			RelatedContents: stringPtr("老字段内容"),
			HitSources:      intPtr(99),
		}

		processor.ProcessForDeserialization(antispam)

		if antispam.RelateContent == nil || *antispam.RelateContent != "老字段内容" {
			t.Error("RelatedContents -> RelateContent 同步失败")
		}
		if antispam.HitSource == nil || *antispam.HitSource != 99 {
			t.Error("HitSources -> HitSource 同步失败")
		}
	})

	t.Run("NilValues", func(t *testing.T) {
		// 所有字段都为空
		antispam := &TestAntispam{
			TaskID: stringPtr("test"),
		}

		// 不应该panic
		processor.ProcessForDeserialization(antispam)

		if antispam.RelatedContents != nil || antispam.HitSources != nil {
			t.Error("空值不应该被修改")
		}
	})
}

// TestHTTPLayerCompatibility 测试HTTP层的兼容性处理
func TestHTTPLayerCompatibility(t *testing.T) {
	responseJSON := `{
		"code": 200,
		"msg": "OK",
		"result": {
			"antispam": {
				"taskId": "task123",
				"label": 100,
				"relateContent": "测试内容",
				"hitSource": 5
			}
		}
	}`

	httpResponse := &httpCore.CommonResponse{}
	httpResponse.SetStatusCode(200)
	httpResponse.SetHeaders(http.Header{"Content-Type": []string{"application/json"}})
	httpResponse.SetBody([]byte(responseJSON))

	ctx := &ClientContext{
		response: *httpResponse,
	}

	var responseData TestTextCheckResponse
	err := ctx.convertToBusinessResponse(&responseData)
	if err != nil {
		t.Fatalf("处理响应失败: %v", err)
	}

	// 检查链路
	if responseData.Result == nil {
		t.Fatal("Result 为空")
	}
	if responseData.Result.Antispam == nil {
		t.Fatal("Antispam 为空")
	}

	antispam := responseData.Result.Antispam

	// 验证新字段正确解析
	if antispam.RelateContent == nil || *antispam.RelateContent != "测试内容" {
		t.Error("RelateContent 解析失败")
	}
	if antispam.HitSource == nil || *antispam.HitSource != 5 {
		t.Error("HitSource 解析失败")
	}

	// 验证老字段自动同步
	if antispam.RelatedContents == nil || *antispam.RelatedContents != "测试内容" {
		t.Error("RelatedContents 同步失败")
	}
	if antispam.HitSources == nil || *antispam.HitSources != 5 {
		t.Error("HitSources 同步失败")
	}
}

// TestNestedStructCompatibility 测试嵌套结构体中的兼容性处理
func TestNestedStructCompatibility(t *testing.T) {
	processor := GetGlobalCompatibilityProcessor()

	// 模拟响应结构
	response := &TestTextCheckResponse{
		Result: &TestTextCheckResult{
			Antispam: &TestAntispam{
				TaskID:        stringPtr("nested"),
				RelateContent: stringPtr("嵌套内容"),
				HitSource:     intPtr(10),
			},
		},
	}

	processor.ProcessForDeserialization(response)

	if response.Result.Antispam.RelatedContents == nil || *response.Result.Antispam.RelatedContents != "嵌套内容" {
		t.Error("嵌套结构体中的字段同步失败")
	}
}

// TestSliceOfAntispam 测试切片中的Antispam处理
func TestSliceOfAntispam(t *testing.T) {
	processor := GetGlobalCompatibilityProcessor()

	type ResponseWithSlice struct {
		Items []*TestAntispam
	}

	response := &ResponseWithSlice{
		Items: []*TestAntispam{
			{RelateContent: stringPtr("item1"), HitSource: intPtr(1)},
			{RelateContent: stringPtr("item2"), HitSource: intPtr(2)},
		},
	}

	processor.ProcessForDeserialization(response)

	for i, item := range response.Items {
		if item.RelatedContents == nil {
			t.Errorf("Items[%d].RelatedContents 同步失败", i)
		}
		if item.HitSources == nil {
			t.Errorf("Items[%d].HitSources 同步失败", i)
		}
	}
}

// 辅助函数
// 注意：在Go中，返回局部变量的指针是安全的。
// Go编译器会进行逃逸分析，自动将其分配到堆上。

func stringPtr(s string) *string {
	return &s
}

func intPtr(i int) *int {
	return &i
}

// 以下是为了让测试编译通过而保留的接口实现（简化版）

type TestRequestWithCompatibility struct {
	BaseValue string `json:"baseValue"`
}

func (r *TestRequestWithCompatibility) SetProductCode(productCode string)                {}
func (r *TestRequestWithCompatibility) SetRegionCode(regionCode string)                  {}
func (r *TestRequestWithCompatibility) SetDomain(domain string)                          {}
func (r *TestRequestWithCompatibility) SetProtocol(protocol httpCore.Protocol)           {}
func (r *TestRequestWithCompatibility) SetCustomParams(customParams map[string]string)   {}
func (r *TestRequestWithCompatibility) SetNonSignParams(nonSignParams map[string]string) {}
func (r *TestRequestWithCompatibility) GetProductCode() string                           { return "test" }
func (r *TestRequestWithCompatibility) GetRegionCode() string                            { return "test" }
func (r *TestRequestWithCompatibility) GetDomain() string                                { return "test.domain" }
func (r *TestRequestWithCompatibility) GetProtocol() httpCore.Protocol                   { return httpCore.ProtocolEnumHTTPS }
func (r *TestRequestWithCompatibility) GetHeaders() map[string]string                    { return nil }
func (r *TestRequestWithCompatibility) GetPathParameters() map[string]string             { return nil }
func (r *TestRequestWithCompatibility) GetQueryParameters() map[string]string            { return nil }
func (r *TestRequestWithCompatibility) GetBody() []byte                                  { return nil }
func (r *TestRequestWithCompatibility) ToHttpRequest(signer auth.Signer, credentials auth.Credentials) (httpCore.Request, error) {
	return nil, nil
}
func (r *TestRequestWithCompatibility) AssembleUrl() string                        { return "" }
func (r *TestRequestWithCompatibility) GetSignParams() map[string]string           { return nil }
func (r *TestRequestWithCompatibility) GetBusinessCustomSignParams() map[string]string  { return nil }
func (r *TestRequestWithCompatibility) GetBusinessNonSignParams() map[string]string     { return nil }
func (r *TestRequestWithCompatibility) ValidateParam() error                       { return nil }
