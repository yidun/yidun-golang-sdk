package http

var (
	DEFAULT_REGION_CODE string   = "cn-hangzhou"
	DEFAULT_PROTOCOL    Protocol = ProtocolEnumHTTPS
)

type Protocol string

const (
	ProtocolEnumHTTPS Protocol = "https"
	ProtocolEnumHTTP  Protocol = "http"
)

type HttpMethod string

const (
	HttpMethodGet  HttpMethod = "GET"
	HttpMethodPost HttpMethod = "POST"
)

type HttpClientConfig struct {
	Protocol                   Protocol
	SocketTimeoutMillis        int64
	MaxIdleTimeMillis          int64
	ConnectionKeepAliveMillis  int64
	ConnectionRequestTimeoutMs int64
	ConnectionTimeoutMs        int64
	ResponseTimeoutMs          int64
	MaxConnectionCount         int
	MaxConnectionCountPerRoute int
}

func DefaultHttpClientConfig() *HttpClientConfig {
	return &HttpClientConfig{
		Protocol:                   ProtocolEnumHTTPS,
		SocketTimeoutMillis:        5000,
		MaxIdleTimeMillis:          60 * 1000,
		ConnectionKeepAliveMillis:  2 * 60 * 1000,
		ConnectionRequestTimeoutMs: 1000,
		ConnectionTimeoutMs:        15000,
		ResponseTimeoutMs:          20000,
		MaxConnectionCount:         200,
		MaxConnectionCountPerRoute: 20,
	}
}

// 请求头
const (
	ContentType     = "Content-Type"
	ContentEncoding = "Content-Encoding"
	SecretID        = "X-YD-SECRETID"
	Timestamp       = "X-YD-TIMESTAMP"
	Nonce           = "X-YD-NONCE"
	Sign            = "X-YD-SIGN"
	SDKVersion      = "SdkVer"
)

// DefaultSDKVersion 默认 SDK 版本号
const (
	DefaultSDKVersion = "1.0"
	HttpCLientVersion = "Go-http-client/1.1"
)
