package auth

const (

	// code:200
	// 说明：成功
	SUCCESS = 200
	// code:400
	// 说明：请求参数错误
	BAD_REQUEST = 400
	// code:401
	// secretId 或 businessId 错误
	FORBIDDEN = 401
	// code:402
	// 业务未上线
	BUSINESS_OFFLINE = 402
	// code:405
	// 请求参数异常
	PARAMS_ERROR = 405
	// code: 410
	// 签名验证失败，请重新参考demo签名代码
	SIGNATURE_ERROR = 410
	// code:411
	// 请求频率或数量超过限制
	HIGH_FREQUENCY = 411
	// code:420
	// 请求过期
	REQUEST_EXPIRED = 420
	// code:430
	// 重放攻击
	REPLAY_ATTACK = 430
	// code:503
	// 接口异常
	SERVER_ERROR = 503
	// code:504
	//  次数超限
	QUOTA_LIMIT = 504
	// code: 505
	// 服务未开通或已过期
	SERVICE_EXPIRED = 505
	// code: 508
	// 试用条数不足
	QUOTA_NOT_ENOUGH = 508
)
