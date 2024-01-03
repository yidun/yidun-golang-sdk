package authiface

import "github.com/yidun/yidun-golang-sdk/yidun/service/auth"

type AUTHAPI interface {
	// 银行卡检测接口，调用/v1/bankcard/check
	BankCardCheck(request *auth.BankCardCheckRequest) (response *auth.BankCardResponse, err error)
	// 出入境证件核验, 调用/v1/foreign/check
	EntryExitPermitCheck(request *auth.EntryExitPermitCheckRequest) (response *auth.EntryExitPermitCheckResponse, err error)
	// 身份证核验, 调用/v1/idcard/check
	IdCardCheck(request *auth.IdCardCheckRequest) (response *auth.IdCardCheckResponse, err error)
	//号码状态检测, 调用/v1/idlephone/check
	IdlePhoneCheck(request *auth.IdlePhoneCheckRequest) (response *auth.IdlePhoneCheckResponse, err error)
	// InteractiveLivePersonCheck /v1/liveperson/recheck
	InteractiveLivePersonCheck(request *auth.InteractiveLivePersonCheckRequest) (response *auth.InteractiveLivePersonCheckResponse, err error)
	//交互式活体人脸核身检测 InteractiveLivePersonIdCardCheck  /v1/liveperson/audit
	InteractiveLivePersonIdCardCheck(request *auth.InteractiveLivePersonIdCardCheckRequest) (response *auth.InteractiveLivePersonIdCardCheckResponse, err error)
	//视频活体检测请求 VideoLivePersonCheck /v1/liveperson/h5/check
	VideoLivePersonCheck(request *auth.VideoLivePersonCheckRequest) (response *auth.VideoLivePersonCheckResponse, err error)
	//视频活体人脸核身请求 VideoLivePersonIdCardCheck /v1/liveperson/h5/audit
	VideoLivePersonIdCardCheck(request *auth.VideoLivePersonIdCardCheckRequest) (response *auth.VideoLivePersonIdCardCheckResponse, err error)
	//手机号与所有者身份证号及姓名验证请求 MobileNumberOwnerIdNameCheck /v1/mobile/check
	MobileNumberOwnerIdNameCheck(request *auth.MobileNumberOwnerIdNameCheckRequest) (response *auth.MobileNumberOwnerIdNameCheckResponse, err error)
	//手机号与所有者姓名验证请求 MobileNumberOwnerNameCheck /v1/mobile2Ele/check
	MobileNumberOwnerNameCheck(request *auth.MobileNumberOwnerNameCheckRequest) (response *auth.MobileNumberOwnerNameCheckResponse, err error)
	//身份证OCR请求 IdCardOcrCheck /v1/ocr/check
	IdCardOcrCheck(request *auth.IdCardOcrCheckRequest) (response *auth.IdCardOcrCheckResponse, err error)
	// 实人认证请求 RpCheck /v1/rp/check
	RpCheck(request *auth.RpCheckRequest) (response *auth.RpCheckResponse, err error)
}

var _ AUTHAPI = (*auth.AuthClient)(nil)
