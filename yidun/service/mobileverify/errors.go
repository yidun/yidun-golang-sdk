package mobileverify

const (
	// code:200
	// 说明：成功
	SUCCESS = 200
	// code:400
	// 请求缺少 secretId 或 businessId
	BAD_REQUEST = 400
	// code:401
	// forbidden
	FORBIDDEN = 401
	// code:405
	// 参数错误
	PARAMS_ERROR = 405
	// code: 410
	// signature failure 签名验证失败，请重新参考demo签名代码
	SIGNATURE_ERROR = 410
	// code:420
	// request expired 请求过期
	REQUEST_EXPIRED = 420
	// code:429
	// too many requests 次数超限
	QUOTA_LIMIT = 429
	// code: 430
	// replay attack 重放攻击
	REPLAY_ATTACK = 430
	// code:440
	// decode error 解密错误
	DECODE_ERROR = 440
	// code: 450
	// wrong token token错误
	TOKEN_ERROR = 450
	// code: 503
	// service unavailable 	接口异常
	SERVER_ERROR = 503
	// code: 507
	// balance not enough 余额不足
	BALANCE_NOT_ENOUGH = 507
	// code: 508
	// rate limit 	QPS超限
	RATE_LIMIT = 508
	// code: 1002
	// other errors 其他错误
	OTHER_ERROR = 1002
)
