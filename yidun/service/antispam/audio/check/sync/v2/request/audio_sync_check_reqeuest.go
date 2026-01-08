package request

import (
	"strconv"

	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/audio/check"
)

// DataCheckType 检测类型枚举
type DataCheckType int

const (
	// DataCheckTypeURL URL类型
	DataCheckTypeURL DataCheckType = 0
	// DataCheckTypeBASE64 BASE64类型
	DataCheckTypeBASE64 DataCheckType = 1
)

type AudioSyncCheckRequest struct {
	*check.AudioCommonCheckRequest
	// 检测类型
	DataCheckType *int `json:"dataCheckType,omitempty"`
	// 同步检测接口-语音内容base64
	Data *string `json:"data,omitempty"`
}

// NewAudioSyncCheckRequest 初始化AudioSyncCheckRequest对象
func NewAudioSyncCheckRequest(businessId string) *AudioSyncCheckRequest {
	var request = &AudioSyncCheckRequest{
		AudioCommonCheckRequest: check.NewAudioCommonCheckRequest(businessId),
	}
	request.SetProductCode("audioCheck")
	request.SetUriPattern("/v2/audio/check")
	request.SetVersion("v2.2")
	return request
}

// GetCustomSignParams 获取具体业务中特有的需要参与签名计算的参数
func (r *AudioSyncCheckRequest) GetBusinessCustomSignParams() map[string]string {
	params := r.AudioCommonCheckRequest.GetBusinessCustomSignParams()
	if r.DataCheckType != nil {
		params["dataCheckType"] = strconv.Itoa(*r.DataCheckType)
	}
	if r.Data != nil {
		params["data"] = *r.Data
	}
	return params
}

// SetDataCheckType 设置检测类型
func (r *AudioSyncCheckRequest) SetDataCheckType(dataCheckType int) *AudioSyncCheckRequest {
	r.DataCheckType = &dataCheckType
	return r
}

// SetData 设置同步检测接口-语音内容base64
func (r *AudioSyncCheckRequest) SetData(data string) *AudioSyncCheckRequest {
	r.Data = &data
	return r
}
