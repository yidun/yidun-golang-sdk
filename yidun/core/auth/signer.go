package auth

import (
	"log"
	"sort"
	"strings"
)

type Credentials struct {
	SecretId  string
	SecretKey string
	Debug     bool
}

func NewCredentials(secretId string, secretKey string) *Credentials {
	return &Credentials{secretId, secretKey, false}
}

type SignResult struct {
	SignatureMethod SignatureMethod
	SecretId        string
	Signature       string
}

var (
	TIMESTAMP                        = "timestamp"
	NONCE                            = "nonce"
	DEFAULT_SIGNATURE_METHOD         = MD5
	DEFAULT_OPENAPI_SIGNATURE_METHOD = SHA1
)

type Signer interface {
	GenSignature(credentials Credentials, params map[string]string) SignResult
}

type SignerCommon struct{}

func GetCommonSignerInstance() *SignerCommon {
	return &SignerCommon{}
}
func GetOpenapiSignerInstance() *OpenApiSigner {
	return &OpenApiSigner{}
}

func (s SignerCommon) GenSignature(credentials Credentials, params map[string]string) SignResult {
	if credentials.Debug {
		log.Println("SignerCommon.GenSignature, credentials:", credentials, "params:", params)
	}
	signatureMethod := determineSignatureMethod(params)
	secretId := credentials.SecretId

	target := params
	if _, ok := params["secretId"]; !ok {
		target = make(map[string]string, len(params)+1)
		for k, v := range params {
			target[k] = v
		}
		target["secretId"] = secretId
	}

	signature := genSignature(signatureMethod, credentials.SecretKey, target)

	if credentials.Debug {
		log.Println("SignerCommon.GenSignature, signature:", signature)
	}
	return SignResult{signatureMethod, secretId, signature}
}

func genSignature(signatureMethod SignatureMethod, secretKey string, params map[string]string) string {
	// 1. 参数名按照ASCII码表升序排序
	paramNames := make([]string, 0, len(params))
	for paramName := range params {
		paramNames = append(paramNames, paramName)
	}
	sort.Strings(paramNames)

	// 2. 按照排序拼接参数名与参数值
	var paramBuffer strings.Builder
	for _, paramName := range paramNames {
		paramBuffer.WriteString(paramName)
		paramValue := params[paramName]
		if paramValue != "" {
			paramBuffer.WriteString(paramValue)
		}
	}

	// 3. 将secretKey拼接到最后
	paramBuffer.WriteString(secretKey)

	// 4. 计算签名
	return SignatureMethod.calcSign(signatureMethod, paramBuffer.String())
}

func determineSignatureMethod(params map[string]string) SignatureMethod {
	method := params["signatureMethod"]
	if method == "" {
		return DEFAULT_SIGNATURE_METHOD
	}
	if method, ok := methodMap[strings.ToUpper(method)]; ok {
		return method
	} else {
		panic("method not support!")
	}
}

type OpenApiSigner struct{}

func (s OpenApiSigner) GenSignature(credentials Credentials, params map[string]string) SignResult {
	log.Println("OpenApiSigner.GenSignature, credentials:", credentials, "params:", params)
	// 1. 参数名按照ASCII码表升序排序
	paramNames := make([]string, 0, len(params))
	for paramName := range params {
		paramNames = append(paramNames, paramName)
	}
	sort.Strings(paramNames)

	timestamp := ""
	nonce := ""

	// 2. 按照排序拼接参数名与参数值
	var paramBuffer strings.Builder
	for _, paramName := range paramNames {
		paramValue := params[paramName]
		if paramName == TIMESTAMP {
			timestamp = paramValue
			continue
		}
		if paramName == NONCE {
			nonce = paramValue
			continue
		}

		paramBuffer.WriteString(paramName)
		if paramValue != "" {
			paramBuffer.WriteString(paramValue)
		}
	}

	// 3. 将secretKey，nonce，timestamp拼接到最后
	paramBuffer.WriteString(credentials.SecretKey)
	paramBuffer.WriteString(nonce)
	paramBuffer.WriteString(timestamp)
	signature := DEFAULT_OPENAPI_SIGNATURE_METHOD.calcSign(paramBuffer.String())
	log.Println("OpenApiSigner.GenSignature, signature:", signature)
	return SignResult{DEFAULT_OPENAPI_SIGNATURE_METHOD, credentials.SecretId, signature}
}
