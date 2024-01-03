package v5

const (
	// code:400
	// 说明：请求参数不合法
	BAD_REQUEST = 400
	// code:404
	// 服务不存在
	SERVICE_NOT_FOUND = 404
	// code:411
	// 说明：请求qps超限，请需调整请联系客户经理
	REQUEST_FREQUENCY_UPPER_LIMIT = 411
	// code:500
	// 说明：服务异常
	INTERNAL_SERVER_ERROR = 500
	// code:5501
	// 说明：请求过期，请检查timestamp参数是否正确
	REQUEST_EXPIRED = 5501
	// code:5502
	// 说明：参数错误，请检查参数是否完整
	INVALID_PARAM = 5502
	// code:5504
	// 说明：数据解密失败，请确认数据是否正确
	DECRYPT_FAILURE = 5504
	// code:5505
	// 说明：无效nonce，请注意nonce格式是否正确
	INVALID_NONCE = 5505
	// code:5506
	// 说明：无效业务，请检查businessId是否正确
	INVALID_BUSINESS = 5506
	// code:5507
	// 说明：无上报数据，请检查SDK对接是否正确
	NO_UPLOAD_DATA = 5507
	// code:5508
	// 说明：无效签名，请检查签名算法是否正确
	INVALID_SIGNATURE = 5508
	// code:5509
	// 说明：无效产品，请检查产品开通状态
	INVALID_PRODUCT = 5509
	// code:5511
	// 说明：重复请求，请检查是否存在多次调用
	DUPLICATE_DATA = 5511
	// code:5512
	// 说明：非法token，请检查token数据与产品编号是否一致，或token数据是否跨平台请求
	INVALID_TOKEN = 5512
)
