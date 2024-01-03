package fingerprint

const (
	// code:400
	// 说明：请求参数错误。检查参数是否有遗漏或空白
	BAD_REQUEST = 400

	// code:401
	// 说明：businessId无效或过期
	FORBIDDEN_ERROR = 401

	// code:405
	// 说明：请求参数错误。检测businessKey等参数是否有效，签名是否正确
	PARAM_ERROR = 405

	// code:410
	// 说明：签名验证失败。检查 secretId 与 businessId 是否匹配，签名生成逻辑是否正确
	SIGNATURE_FAILURE = 410

	// code:420
	// 说明：请求过期
	REQUEST_EXPIRED = 420

	// code:430
	// 说明：请求重放。检查请求是否被截获并重放
	REPLAY_ATTACK = 430

	// code:503
	// 说明：服务端其它异常。请联系易盾技术支持解决
	SERVICE_UNAVAILABLE = 503
)
