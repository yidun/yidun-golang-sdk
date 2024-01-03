package util

import (
	"net/url"

	"github.com/yidun/yidun-golang-sdk/yidun/core/auth"
)

const (
    ParamSign      = "signature"
    ParamSecretId  = "secretId"
)

func GenSignature(secretKey string, params map[string]string) string {
    secretId := params[ParamSecretId]
    if secretId == "" {
        // 签名参数为空，直接返回空
        return ""
    }
    
	sigher := &auth.SignerCommon{}
    credentials := *auth.NewCredentials(secretId, secretKey)
	signature := sigher.GenSignature(credentials, params)
    
    return signature.Signature
}

func VerifySignature(requestParams url.Values, secretKey string) bool {
    if secretKey == "" || requestParams == nil {
        return false
    }

    // 获取签名参数和其他参数
    params := make(map[string]string)
    var signature string
    for k, v := range requestParams {
        if len(v) > 0 {
            if k == ParamSign {
                signature = v[0]
            } else {
                params[k] = v[0]
            }
        }
    }
    // 生成签名并比较
    generatedSignature := GenSignature(secretKey, params)
    return signature != "" && generatedSignature != "" && signature == generatedSignature
}