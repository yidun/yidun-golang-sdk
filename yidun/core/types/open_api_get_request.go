package types

import (
	"strconv"
	"strings"

	"github.com/yidun/yidun-golang-sdk/yidun/core/auth"
	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
)

type OpenApiGetRequest struct {
	*PostFormRequest

}

func NewOpenApiGetRequest() *OpenApiGetRequest {
	var request = &OpenApiGetRequest{
		PostFormRequest: NewPostFormRequest(),
	}
	return request
}

func (r *OpenApiGetRequest) GetHeaders() map[string]string {
	headers := make(map[string]string)
	for k, v := range r.PostFormRequest.GetHeaders() {
		headers[k] = v
	}
	headers[http.Nonce] = r.Nonce
	headers[http.Timestamp] = strconv.FormatInt(r.Timestamp, 10)
	return headers
}

func (r *OpenApiGetRequest) GetHeadersWithSign(signer *auth.Signer, credentials *auth.Credentials) map[string]string {
	headers := r.GetHeaders()
	headers[http.SecretID] = credentials.SecretId
	headers[http.Sign] = (*signer).GenSignature(*credentials, r.GetSignParams()).Signature
	return headers
}

// GetSignParams 获取待签名参数列表
func (r *OpenApiGetRequest) GetSignParams() map[string]string {
	// 创建一个新map，将原始map的所有键值对复制到新map中
	params := make(map[string]string)
	for k, v := range r.CustomParams {
		params[k] = v
	}
	// 修改新map的值
	params["timestamp"] = strconv.FormatInt(r.Timestamp, 10)
	params["nonce"] = r.Nonce
	if r.SignatureMethod != "" {
		params["signatureMethod"] = string(r.SignatureMethod)
	}
	return params
}

// ToHttpRequest 构建 http 请求
func (r *OpenApiGetRequest) ToHttpRequest(signer auth.Signer, credentials auth.Credentials) (http.Request, error) {
	req := http.HttpRequest{
		MethodValue:  string(r.Method),
		UrlValue:     r.AssembleUrl(),
		HeadersValue: r.GetHeadersWithSign(&signer, &credentials),
	}
	return &req, nil
}

func (r *OpenApiGetRequest) AssembleUrl() string {
	url := r.PostFormRequest.AssembleUrl()
	mapQueries := make(map[string]string)
	for k, v := range r.CustomParams {
		mapQueries[k] = v
	}
	urlBuilder := strings.Builder{}
	urlBuilder.WriteString(url)
	if len(mapQueries) != 0 {
		if !strings.Contains(urlBuilder.String(), "?") {
			urlBuilder.WriteString("?")
		} else if !strings.HasSuffix(urlBuilder.String(), "?") {
			urlBuilder.WriteString("&")
		}
		urlBuilder.WriteString(UrlEncode(mapQueries))
	}
	url = urlBuilder.String()
	if strings.HasSuffix(url, "?") || strings.HasSuffix(url, "&") {
		url = url[0 : len(url)-1]
	}
	return url
}
