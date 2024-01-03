package callback

import (
)

/**
 * 主动回调的实体对象
 */
type ActiveCallbackRequest struct {
    SecretId        *string `json:"secretId,omitempty"`        // 注释：密钥secretId
    BusinessId      *string `json:"businessId,omitempty"`      // 注释：密钥businessId
    Scene           *string `json:"scene,omitempty"`           // 注释：业务场景
    Signature       *string `json:"signature,omitempty"`       // 注释：数据签名
    SignatureMethod *string `json:"signatureMethod,omitempty"` // 注释：签名方法
    CallbackData    *string `json:"callbackData,omitempty"`    // 注释：回调数据
}
// 构建request
func NewActiveCallbackRequest(params map[string]string) *ActiveCallbackRequest {
	request := ActiveCallbackRequest{}
	if val, ok := params["secretId"]; ok {
		request.SecretId = &val
	}
	if val, ok := params["businessId"]; ok {
		request.BusinessId = &val
	}
	if val, ok := params["scene"]; ok {
		request.Scene = &val
	}
	if val, ok := params["signature"]; ok {
		request.Signature = &val
	}
	if val, ok := params["signatureMethod"]; ok {
		request.SignatureMethod = &val
	}
	if val, ok := params["callbackData"]; ok {
		request.CallbackData = &val
	}
	return &request
}