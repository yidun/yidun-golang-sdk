package mobileverify

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/core/validation"
)

// 获取手机号
func (c *MobileNumberClient) GetMobileNumber(request *MobileNumberGetRequest) (response *MobileNumberGetResponse, err error) {
	response = &MobileNumberGetResponse{}
	err = c.Client.DoExecute(request, response)
	return
}

// 验证本机号码是否与指定的号码相同
func (c *MobileNumberClient) VerifyMobileNumber(request *MobileNumberVerifyRequest) (response *MobileNumberVerifyResponse, err error) {
	response = &MobileNumberVerifyResponse{}
	err = c.Client.DoExecute(request, response)
	return
}

// 获取手机号请求
type MobileNumberGetRequest struct {
	*types.BizPostFormRequest
	Token       *string
	AccessToken *string
}

// SetToken 设置token
func (r *MobileNumberGetRequest) SetToken(token string) *MobileNumberGetRequest {
	r.Token = &token
	return r
}

// SetAccessToken 设置accessToken
func (r *MobileNumberGetRequest) SetAccessToken(accessToken string) *MobileNumberGetRequest {
	r.AccessToken = &accessToken
	return r
}

// new method
func NewMobileNumberGetRequest(businessId string) *MobileNumberGetRequest {
	request := &MobileNumberGetRequest{
		BizPostFormRequest: types.NewBizPostFormRequest((businessId)),
	}
	request.SetProductCode("mobileverify")
	request.SetUriPattern("/v1/oneclick/check")
	request.SetVersion("v1")
	return request
}

// param check
func (r *MobileNumberGetRequest) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "MobileNumberGetRequest"}
	if r.Token == nil {
		invalidParams.Add(validation.NewErrParamRequired("Token"))
	}
	if r.AccessToken == nil {
		invalidParams.Add(validation.NewErrParamRequired("AccessToken"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// GetNonSignParams 获取具体业务中特有的不参与签名计算的参数
func (r *MobileNumberGetRequest) GetNonSignParams() map[string]string {
	return make(map[string]string)
}

// GetCustomSignParams 获取具体业务中特有的需要参与签名计算的参数
func (r *MobileNumberGetRequest) GetBusinessCustomSignParams() map[string]string {
	params := r.BizPostFormRequest.GetBusinessCustomSignParams()
	if r.Token != nil {
		params["token"] = *r.Token
	}
	if r.AccessToken != nil {
		params["accessToken"] = *r.AccessToken
	}
	return params
}

// 一键登录response
type MobileNumberGetResponse struct {
	*types.CommonResponse
	Data MobileNumberResult `json:"data"`
}

// 一键登录result
type MobileNumberResult struct {
	Phone      *string `json:"phone"`
	ResultCode *string `json:"resultCode"`
}

// 号码认证请求
type MobileNumberVerifyRequest struct {
	*types.BizPostFormRequest
	Token       *string
	AccessToken *string
	Phone       *string
}

// SetToken 设置token
func (r *MobileNumberVerifyRequest) SetToken(token string) *MobileNumberVerifyRequest {
	r.Token = &token
	return r
}

// SetAccessToken 设置accessToken
func (r *MobileNumberVerifyRequest) SetAccessToken(accessToken string) *MobileNumberVerifyRequest {
	r.AccessToken = &accessToken
	return r
}

// SetPhone 设置手机号
func (r *MobileNumberVerifyRequest) SetPhone(phone string) *MobileNumberVerifyRequest {
	r.Phone = &phone
	return r
}

// new method
func NewMobileNumberVerifyRequest(businessId string) *MobileNumberVerifyRequest {
	request := &MobileNumberVerifyRequest{
		BizPostFormRequest: types.NewBizPostFormRequest((businessId)),
	}
	request.SetProductCode("mobileverify")
	request.SetUriPattern("/v1/check")
	request.SetVersion("v1")
	return request
}

// param check
func (r *MobileNumberVerifyRequest) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "MobileNumberVerifyRequest"}
	if r.Token == nil {
		invalidParams.Add(validation.NewErrParamRequired("Token"))
	}
	if r.AccessToken == nil {
		invalidParams.Add(validation.NewErrParamRequired("AccessToken"))
	}
	if r.Phone == nil {
		invalidParams.Add(validation.NewErrParamRequired("Phone"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// GetNonSignParams 获取具体业务中特有的不参与签名计算的参数
func (r *MobileNumberVerifyRequest) GetNonSignParams() map[string]string {
	return make(map[string]string)
}

// GetCustomSignParams 获取具体业务中特有的需要参与签名计算的参数
func (r *MobileNumberVerifyRequest) GetBusinessCustomSignParams() map[string]string {
	params := r.BizPostFormRequest.GetBusinessCustomSignParams()
	if r.Token != nil {
		params["token"] = *r.Token
	}
	if r.AccessToken != nil {
		params["accessToken"] = *r.AccessToken
	}
	if r.Phone != nil {
		params["phone"] = *r.Phone
	}
	return params
}

// 号码认证响应
type MobileNumberVerifyResponse struct {
	*types.CommonResponse
	Data MobileNumberVerifyResult `json:"data"`
}

// 号码认证详细
type MobileNumberVerifyResult struct {
	Result *int `json:"result"`
}
