package captcha

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/core/validation"
)

// 验证码二次校验接口，调用/api/v2/verify
func (c *CaptchaVerifyClient) Verify(request *CaptchaVerifyRequest) (response *CaptchaVerifyResponse, err error) {
	response = &CaptchaVerifyResponse{}
	err = c.Client.DoExecute(request, response)
	return
}

// CaptchaVerifyRequest 验证码二次校验接口请求参数
type CaptchaVerifyRequest struct {
	*types.PostFormRequest
	CaptchaId *string
	User      *string
	Validate  *string
	Type      *string
	ClientIp  *string
	ClientUa  *string
}

// 提供CaptchaId参数SET方法
func (r *CaptchaVerifyRequest) SetCaptchaId(captchaId string) *CaptchaVerifyRequest {
	r.CaptchaId = &captchaId
	return r
}

// 提供User参数SET方法
func (r *CaptchaVerifyRequest) SetUser(user string) *CaptchaVerifyRequest {
	r.User = &user
	return r
}

// 提供Validate参数SET方法
func (r *CaptchaVerifyRequest) SetValidate(validate string) *CaptchaVerifyRequest {
	r.Validate = &validate
	return r
}

// 提供Type参数SET方法
func (r *CaptchaVerifyRequest) SetType(captchaType string) *CaptchaVerifyRequest {
	r.Type = &captchaType
	return r
}

// 提供ClientIp参数SET方法
func (r *CaptchaVerifyRequest) SetClientIp(clientIp string) *CaptchaVerifyRequest {
	r.ClientIp = &clientIp
	return r
}

// 提供ClientUa参数SET方法
func (r *CaptchaVerifyRequest) SetClientUa(clientUa string) *CaptchaVerifyRequest {
	r.ClientUa = &clientUa
	return r
}

// new method
func NewCaptchaVerifyRequest() *CaptchaVerifyRequest {
	request := &CaptchaVerifyRequest{
		PostFormRequest: types.NewPostFormRequest(),
	}
	request.SetProductCode("captcha")
	request.SetUriPattern("/api/v2/verify")
	request.SetVersion("v2")
	return request
}

// ValidateParam 验证参数有效性
func (r *CaptchaVerifyRequest) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "CaptchaVerifyRequest"}
	if r.CaptchaId == nil {
		invalidParams.Add(validation.NewErrParamRequired("CaptchaId"))
	}
	if r.Validate == nil {
		invalidParams.Add(validation.NewErrParamRequired("Validate"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// GetNonSignParams 获取具体业务中特有的不参与签名计算的参数
func (r *CaptchaVerifyRequest) GetNonSignParams() map[string]string {
	return make(map[string]string)
}

// GetCustomSignParams 获取具体业务中特有的需要参与签名计算的参数
func (r *CaptchaVerifyRequest) GetBusinessCustomSignParams() map[string]string {
	params := r.PostFormRequest.GetBusinessCustomSignParams()
	if r.CaptchaId != nil {
		params["captchaId"] = *r.CaptchaId
	}
	if r.User != nil {
		params["user"] = *r.User
	}
	if r.Validate != nil {
		params["validate"] = *r.Validate
	}
	if r.Type != nil {
		params["type"] = *r.Type
	}
	if r.ClientIp != nil {
		params["clientIp"] = *r.ClientIp
	}
	if r.ClientUa != nil {
		params["clientUa"] = *r.ClientUa
	}
	return params
}

// response
type CaptchaVerifyResponse struct {
	Error       *int    `json:"error"`
	Msg         *string `json:"msg"`
	Result      *bool   `json:"result"`
	ExtraData   *string `json:"extraData"`
	Phone       *string `json:"phone"`
	Token       *string `json:"token"`
	CaptchaType *int    `json:"captchaType"`
	SdkReduce   *string `json:"sdkReduce"`
	ClientIp    *string `json:"clientIp"`
	ClientUa    *string `json:"clientUa"`
}
