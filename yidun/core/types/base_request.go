package types

import (
	"bytes"
	"compress/gzip"
	"crypto/rand"
	"encoding/hex"
	"net/url"
	"reflect"
	"strings"

	"github.com/yidun/yidun-golang-sdk/yidun/core/auth"
	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
)

// 基础的请求对象构造体
type BaseRequestConstruct struct {
	RegionCode    string
	Protocol      http.Protocol
	Domain        string
	ProductCode   string
	UriPattern    string
	Method        http.HttpMethod
	EnableRecover bool
	CustomParams  map[string]string
	NonSignParams map[string]string
}

// 返回一个新的基础的请求对象构造体
func NewBaseRequestConstruct() *BaseRequestConstruct {
	return &BaseRequestConstruct{
		// 默认启用revover
		EnableRecover: true,
	}
}

// NewBaseRequestConstruct 基础的请求对象构造体set方法 start
func (b *BaseRequestConstruct) SetProtocol(protocol http.Protocol) {
	b.Protocol = protocol
}

func (b *BaseRequestConstruct) SetDomain(domain string) {
	b.Domain = domain
}

func (b *BaseRequestConstruct) SetUriPattern(uriPattern string) {
	b.UriPattern = uriPattern
}

func (b *BaseRequestConstruct) SetMethod(method http.HttpMethod) {
	b.Method = method
}

func (b *BaseRequestConstruct) SetEnableRecover(enableRecover bool) {
	b.EnableRecover = enableRecover
}

// NewBaseRequestConstruct 基础的请求对象构造体set方法 end

// BaseRequest 基类接口
type BaseRequest interface {
	// 基类的set方法
	SetProductCode(productCode string)
	SetRegionCode(regionCode string)
	SetDomain(domain string)
	SetProtocol(protocol http.Protocol)
	SetCustomParams(customParams map[string]string)
	SetNonSignParams(nonSignParams map[string]string)
	// 基类的get方法
	GetProductCode() string
	GetRegionCode() string
	GetDomain() string
	GetProtocol() http.Protocol
	// 需要实现的处理方法
	GetHeaders() map[string]string
	GetPathParameters() map[string]string
	GetQueryParameters() map[string]string
	GetBody() []byte
	ToHttpRequest(signer auth.Signer, credentials auth.Credentials) (http.Request, error)
	AssembleUrl() string
	GetSignParams() map[string]string
	GetBusinessCustomSignParams() map[string]string
	GetBusinessNonSignParams() map[string]string
	ValidateParam() error
}

// ParamsFilled returns if the request's parameters have been populated
// and the parameters are valid. False is returned if no parameters are
// provided or invalid.
func (r *BaseRequestConstruct) ParamsFilled() bool {
	return r != nil && reflect.ValueOf(r).Elem().IsValid()
}

// ValidateParam 验证参数，如果参数不合法，则返回错误，子类可以自己实现自己的校验方法
func (r *BaseRequestConstruct) ValidateParam() error {
	return nil
}

// GetHeaders 获取请求头
func (r *BaseRequestConstruct) GetHeaders() map[string]string {
	baseHeaders := make(map[string]string)
	baseHeaders[http.SDKVersion] = http.DefaultSDKVersion
	return baseHeaders
}

// GetPathParameters 获取路径参数
func (r *BaseRequestConstruct) GetPathParameters() map[string]string {
	return make(map[string]string)
}

// GetQueryParameters 获取查询参数
func (r *BaseRequestConstruct) GetQueryParameters() map[string]string {
	return make(map[string]string)
}

// GetBody 获取请求体
func (r *BaseRequestConstruct) GetBody() []byte {
	return nil
}

// GetCustomSignParams 获取具体业务中特有的需要参与签名计算的参数
func (r *BaseRequestConstruct) GetBusinessCustomSignParams() map[string]string {
	return make(map[string]string)
}

// GetNonSignParams 获取具体业务中特有的不参与签名计算的参数
func (r *BaseRequestConstruct) GetBusinessNonSignParams() map[string]string {
	return make(map[string]string)
}

// SetCustomParams 设置自定义参数
func (r *BaseRequestConstruct) SetCustomParams(customParams map[string]string) {
	r.CustomParams = customParams
}

// SetNonSignParams 设置不参与签名的参数
func (r *BaseRequestConstruct) SetNonSignParams(nonSignParams map[string]string) {
	r.NonSignParams = nonSignParams
}

// 获取产品编码
func (r *BaseRequestConstruct) GetProductCode() string {
	return r.ProductCode
}

// 获取区域编码
func (r *BaseRequestConstruct) GetRegionCode() string {
	return r.RegionCode
}

// 获取协议
func (r *BaseRequestConstruct) GetProtocol() http.Protocol {
	return r.Protocol
}

// SetProductCode 设置产品编码
func (r *BaseRequestConstruct) SetProductCode(productCode string) {
	r.ProductCode = productCode
}

// SetRegionCode 设置区域编码
func (r *BaseRequestConstruct) SetRegionCode(regionCode string) {
	r.RegionCode = regionCode
}

// ToHttpRequest 构建 http 请求
func (r *BaseRequestConstruct) ToHttpRequest(signer auth.Signer, credentials auth.Credentials) (http.Request, error) {
	req := http.HttpRequest{
		MethodValue:  string(r.Method),
		UrlValue:     r.AssembleUrl(),
		HeadersValue: r.GetHeaders(),
		BodyValue:    r.GetBody(),
	}
	return &req, nil
}

// 获取域名
func (r *BaseRequestConstruct) GetDomain() string {
	return r.Domain
}

func (r *BaseRequestConstruct) AssembleUrl() string {
	mapQueries := make(map[string]string)
	if len(r.GetQueryParameters()) != 0 {
		mapQueries = r.GetQueryParameters()
	}
	urlBuilder := strings.Builder{}
	urlBuilder.WriteString(string(r.Protocol))
	urlBuilder.WriteString("://")
	urlBuilder.WriteString(r.Domain)
	if r.UriPattern != "" {
		urlBuilder.WriteString(populatePathParams(r.UriPattern, r.GetPathParameters()))
	}
	if !strings.Contains(urlBuilder.String(), "?") {
		urlBuilder.WriteString("?")
	} else if !strings.HasSuffix(urlBuilder.String(), "?") {
		urlBuilder.WriteString("&")
	}
	if len(mapQueries) != 0 {
		urlBuilder.WriteString(UrlEncode(mapQueries))
	}
	url := urlBuilder.String()

	if strings.HasSuffix(url, "?") || strings.HasSuffix(url, "&") {
		url = url[0 : len(url)-1]
	}
	return url
}

func populatePathParams(uriPattern string, pathParams map[string]string) string {
	if len(pathParams) == 0 {
		return uriPattern
	}
	result := uriPattern
	for key, value := range pathParams {
		target := "[" + key + "]"
		result = strings.Replace(result, target, url.QueryEscape(value), -1)
	}
	return result
}

// ToGzipBytes 使用 gzip 压缩数据
func ToGzipBytes(rawBytes []byte) ([]byte, error) {
	var buf bytes.Buffer

	gzipWriter, err := gzip.NewWriterLevel(&buf, gzip.BestCompression)
	if err != nil {
		return nil, err
	}

	_, err = gzipWriter.Write(rawBytes)
	if err != nil {
		return nil, err
	}

	err = gzipWriter.Close()
	if err != nil {
		return nil, err
	}

	return buf.Bytes(), nil
}

// UrlEncode 将 map[string]string 类型的数据转换为 url encode 格式的字符串
func UrlEncode(params map[string]string) string {
	parts := url.Values{}
	for key, value := range params {
		parts.Add(key, value)
	}
	return parts.Encode()
}

// RandUUID 生成随机 UUID
func RandUUID() string {
	buf := make([]byte, 16)
	rand.Read(buf)
	return hex.EncodeToString(buf)
}
