package captcha

const (

	// code:415
	// 说明：签名校验错误
	SIGNATURE_ERROR = 415
	// code:419
	// 说明：参数校验错误，例如参数类型错误、参数值错误、必填项为空等
	PARAM_ERROR = 419
	// code:430
	// 说明：qps超限
	REQUEST_FREQUENCY_UPPER_LIMIT = 430
	// code:421
	// 说明：验证码版本不匹配
	VERSION_NOT_MATCH = 421
)
