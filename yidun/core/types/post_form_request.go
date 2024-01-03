package types

import (
	"strconv"
	"time"

	"github.com/yidun/yidun-golang-sdk/yidun/core/auth"
	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
)

type PostFormRequest struct {
	*BaseRequestConstruct
	Version         string
	Timestamp       int64
	Nonce           string
	SignatureMethod auth.SignatureMethod
	GzipCompress    bool
}

// NewPostFormRequest 创建 PostFormRequest 实例
func NewPostFormRequest() *PostFormRequest {
	result := &PostFormRequest{
		BaseRequestConstruct: NewBaseRequestConstruct(),
		Timestamp:            time.Now().UnixNano() / int64(time.Millisecond),
		Nonce:                RandUUID(),
		GzipCompress:         false,
	}
	result.Method = http.HttpMethodPost
	return result
}

func (r *PostFormRequest) GetHeaders() map[string]string {
	headers := make(map[string]string)
	for k, v := range r.BaseRequestConstruct.GetHeaders() {
		headers[k] = v
	}
	headers[http.ContentType] = "application/x-www-form-urlencoded;charset=utf-8"
	if r.GzipCompress {
		headers[http.ContentEncoding] = "gzip"
	}
	return headers
}

// GetSignParams 获取待签名参数列表
func (r *PostFormRequest) GetSignParams() map[string]string {
	// 创建一个新map，将原始map的所有键值对复制到新map中
	params := make(map[string]string)
	for k, v := range r.CustomParams {
		params[k] = v
	}
	// 修改新map的值
	params["version"] = r.Version
	params["timestamp"] = strconv.FormatInt(r.Timestamp, 10)
	params["nonce"] = r.Nonce
	if r.SignatureMethod != "" {
		params["signatureMethod"] = string(r.SignatureMethod)
	}
	return params
}

// GetBodyWithSign 构建 body，将参数以 form 表单格式组装放入 body 中
func (r *PostFormRequest) GetBodyWithSign(signer auth.Signer, credentials auth.Credentials) ([]byte, error) {
	params := r.GetSignParams()
	signResult := signer.GenSignature(credentials, params)
	params["secretId"] = signResult.SecretId
	params["signature"] = signResult.Signature

	for k, v := range r.NonSignParams {
		params[k] = v
	}

	body := []byte(UrlEncode(params))

	if r.GzipCompress {
		body, err := ToGzipBytes(body)
		if err == nil {
			return body, nil
		} else {
			return nil, err
		}
	}
	return body, nil
}

// ToHttpRequest 构建 http 请求
func (r *PostFormRequest) ToHttpRequest(signer auth.Signer, credentials auth.Credentials) (http.Request, error) {
	body, err := r.GetBodyWithSign(signer, credentials)
	if err != nil {
		return nil, err
	}
	req := http.HttpRequest{
		MethodValue:  string(r.Method),
		UrlValue:     r.AssembleUrl(),
		HeadersValue: r.GetHeaders(),
		BodyValue:    body,
	}
	return &req, nil
}

func (p *PostFormRequest) SetVersion(version string) {
	p.Version = version
}

func (p *PostFormRequest) SetTimestamp(timestamp int64) {
	p.Timestamp = timestamp
}

func (p *PostFormRequest) SetNonce(nonce string) {
	p.Nonce = nonce
}

func (p *PostFormRequest) SetSignatureMethod(signatureMethod auth.SignatureMethod) {
	p.SignatureMethod = signatureMethod
}

func (p *PostFormRequest) SetGzipCompress(gzipCompress bool) {
	p.GzipCompress = gzipCompress
}
