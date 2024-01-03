package mobileverifyiface

import "github.com/yidun/yidun-golang-sdk/yidun/service/mobileverify"

// MobileverifyAPI is the mobileverify interface
type MobileverifyAPI interface {
	// 获取手机号 调用：/mobileverify/getMobileNumber
	GetMobileNumber(input *mobileverify.MobileNumberGetRequest) (output *mobileverify.MobileNumberGetResponse, err error)
	// 验证本机号码是否与指定的号码相同 调用：/mobileverify/verifyMobileNumber
	VerifyMobileNumber(input *mobileverify.MobileNumberVerifyRequest) (output *mobileverify.MobileNumberVerifyResponse, err error)
}

var _ MobileverifyAPI = (*mobileverify.MobileNumberClient)(nil)
