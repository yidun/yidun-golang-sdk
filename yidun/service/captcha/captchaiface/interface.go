package captchaiface

import "github.com/yidun/yidun-golang-sdk/yidun/service/captcha"

type CaptchaAPI interface {
	// 验证码二次校验接口，调用/api/v2/verify
	Verify(request *captcha.CaptchaVerifyRequest) (response *captcha.CaptchaVerifyResponse, err error)
}

var _ CaptchaAPI = (*captcha.CaptchaVerifyClient)(nil)
