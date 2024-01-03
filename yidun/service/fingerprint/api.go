package fingerprint

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
)

// 设备指纹请求入参
type FingerprintQueryRequest struct {
	*types.BizPostJsonRequest
	Token *string `type:"string" required:"true"`
}

func (req *FingerprintQueryRequest) SetToken(token string) {
	req.Token = &token
}

// 设备指纹返回结果
type FingerprintResponse struct {
	*types.CommonResponse
	Data *DeviceQueryResult `json:"data"`
}

type DeviceQueryResult struct {
	TaskId            *string     `json:"taskId"`
	TokenCreationTime *int64      `json:"tokenCreationTime"`
	Device            *DeviceData `json:"device"`
}

type DeviceData struct {
	DeviceId    *string            `json:"deviceId"`
	SdkType     *int               `json:"sdkType"`
	CheckResult *DeviceCheckResult `json:"checkResult"`
}
//设备指纹风险项
type DeviceCheckResult struct {
	IsTampered       *int `json:"isTampered"`
	IsSimulator      *int `json:"isSimulator"`
	IsRooted         *int `json:"isRooted"`
	IsMultiRun       *int `json:"isMultiRun"`
	IsVpn            *int `json:"isVpn"`
	IsProxy          *int `json:"isProxy"`
	IsHooked         *int `json:"isHooked"`
	IsInjected       *int `json:"isInjected"`
	IsDebugged       *int `json:"isDebugged"`
	IsXposed         *int `json:"isXposed"`
	IsCloud          *int `json:"isCloud"`
	IsSuspectCloud   *int `json:"isSuspectCloud"`
	IsRiskRom        *int `json:"isRiskRom"`
	IsVm             *int `json:"isVm"`
	IsModify         *int `json:"isModify"`
	IsModifyApp      *int `json:"isModifyApp"`
	IsFlash          *int `json:"isFlash"`
	IsAutoTouch      *int `json:"isAutoTouch"`
	IsControlApp     *int `json:"isControlApp"`
	IsScript         *int `json:"isScript"`
	SecurityScore    *int `json:"securityScore"`
	IsCydiaSubstrate *int `json:"isCydiaSubstrate"`
	IsM1             *int `json:"isM1"`
	IsSpeedUp        *int `json:"isSpeedUp"`
	IsAntiJailbreak  *int `json:"isAntiJailbreak"`
}

func (c *FingerprintClient) Query(request *FingerprintQueryRequest) (response *FingerprintResponse, err error) {
	response = &FingerprintResponse{}
	err = c.Client.DoExecute(request, response)
	return
}

// 创建设备指纹query请求体
func NewFingerprintQueryRequest(businessId string) *FingerprintQueryRequest {
	request := &FingerprintQueryRequest{
		BizPostJsonRequest: types.NewBizPostJsonRequest((businessId)),
	}
	request.SetProductCode("fingerprint")
	request.SetUriPattern("/v1/device/query")
	request.SetVersion("v1")
	return request
}

// GetCustomSignParams 获取具体业务中特有的需要参与签名计算的参数
func (r *FingerprintQueryRequest) GetBusinessCustomSignParams() map[string]string {
	params := r.BizPostJsonRequest.GetBusinessCustomSignParams()
	if r.Token != nil {
		params["token"] = *r.Token
	}
	return params
}
