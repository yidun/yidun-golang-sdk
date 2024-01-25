package client

import (
	"errors"
	"math"
	"reflect"

	"github.com/yidun/yidun-golang-sdk/yidun/core/auth"
	"github.com/yidun/yidun-golang-sdk/yidun/core/endpoint"
	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
	myRecover "github.com/yidun/yidun-golang-sdk/yidun/core/recover"
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/core/yidunerr"
)

type DefaultClient struct {
	defaultRegionCode string
	defaultProtocol   http.Protocol
	signer            auth.Signer
	credentials       auth.Credentials
	maxAttemptCount   int64
	customClient      *http.CustomClient
	endpointResolver  endpoint.EndpointResolver
	breakStrategy     endpoint.BreakStrategy
	requestRecover    myRecover.RequestRecover
}

const (
	DEFAULT_REGION_CODE     = "cn-hangzhou"
	DEFAULT_MAX_RETRY_COUNT = 3
	MAX_RETRY_COUNT_LIMIT   = 10
	NO_RETRY_COUNT          = 0
)

type ClientProfile struct {
	PrimaryEndpoints []endpoint.EndpointConfigEntry
	RegionCode       string
	Signer           auth.Signer
	Credentials      auth.Credentials
	HttpClientConfig *http.HttpClientConfig
	BreakerConfig    *endpoint.StrategyConfig
	MaxRetryCount    int
	RequestRecover   myRecover.RequestRecover
}

func NewDefaultClient(profile *ClientProfile) *DefaultClient {
	if reflect.ValueOf(profile).IsZero() {
		panic("profile should not be nil")
	}
	if reflect.ValueOf(profile.Signer).IsZero() {
		panic("signer should not be nil")
	}
	if reflect.ValueOf(profile.Credentials).IsZero() {
		panic("credentials should not be nil")
	}
	signer := profile.Signer
	credentials := profile.Credentials
	defaultRegionCode := http.DEFAULT_REGION_CODE
	if profile.RegionCode != "" {
		defaultRegionCode = profile.RegionCode
	}
	maxAttemptCount := 1 + math.Max(float64(profile.MaxRetryCount), 0)

	httpClientConfig := http.DefaultHttpClientConfig()
	if !reflect.ValueOf(profile.HttpClientConfig).IsZero() {
		httpClientConfig = profile.HttpClientConfig
	}

	defaultProtocol := http.ProtocolEnumHTTPS
	if !reflect.ValueOf(httpClientConfig.Protocol).IsZero() {
		defaultProtocol = httpClientConfig.Protocol
	}

	httpClient := http.CreateByConfig(httpClientConfig)

	primaryEndpoints := profile.PrimaryEndpoints
	var endpointResolver endpoint.EndpointResolver
	if len(primaryEndpoints) == 0 {
		endpointResolver = endpoint.GetEndpointResolver()
	} else {
		endpointResolver = endpoint.NewCompoundEndpointResolverWithConfigEntries(primaryEndpoints)
	}

	breakerConfig := profile.BreakerConfig
	var breakStrategy endpoint.BreakStrategy
	if reflect.ValueOf(breakerConfig).IsZero() {
		breakStrategy = endpoint.CreateDefaultFixWindowBreakStrategy()
	} else {
		breakStrategy = endpoint.NewFixedWindowBreakStrategy(breakerConfig)
	}
	requestRecover := profile.RequestRecover
	return &DefaultClient{
		defaultRegionCode: defaultRegionCode,
		defaultProtocol:   defaultProtocol,
		signer:            signer,
		credentials:       credentials,
		maxAttemptCount:   int64(maxAttemptCount),
		customClient:      httpClient,
		endpointResolver:  endpointResolver,
		breakStrategy:     breakStrategy,
		requestRecover:    requestRecover,
	}
}
// NewOpenapiClient 创建一个 OpenAPI Client
func NewOpenapiClient(profile *ClientProfile) *DefaultClient {
	profile.SetSigner(auth.GetOpenapiSignerInstance())
	return NewDefaultClient(profile)
}

func NewClientProfile(credentials *auth.Credentials) *ClientProfile {
	return &ClientProfile{
		PrimaryEndpoints: nil,
		RegionCode:       DEFAULT_REGION_CODE,
		Signer:           auth.GetCommonSignerInstance(),
		MaxRetryCount:    DEFAULT_MAX_RETRY_COUNT,
		Credentials:      *credentials,
		BreakerConfig:    endpoint.DefaultStrategyConfig,
		HttpClientConfig: http.DefaultHttpClientConfig(),

		// TODO
		// RequestRecover:  recover.DefaultRequestRecover,
	}
}

func (cp *ClientProfile) SetPrimaryEndpoints(primaryEndpoints []endpoint.EndpointConfigEntry) error {
	if len(primaryEndpoints) == 0 {
		return errors.New("primaryEndpoints cannot be empty")
	}
	cp.PrimaryEndpoints = primaryEndpoints
	return nil
}

func (cp *ClientProfile) SetRegionCode(regionCode string) {
	cp.RegionCode = regionCode
}

func (cp *ClientProfile) SetSigner(signer auth.Signer) {
	cp.Signer = signer
}

func (cp *ClientProfile) SetCredentials(credentials auth.Credentials) {
	cp.Credentials = credentials
}

func (cp *ClientProfile) SetHttpClientConfig(httpClientConfig *http.HttpClientConfig) {
	cp.HttpClientConfig = httpClientConfig
}

func (cp *ClientProfile) SetBreakerConfig(breakerConfig *endpoint.StrategyConfig) {
	cp.BreakerConfig = breakerConfig
}

func (cp *ClientProfile) SetMaxRetryCount(maxRetryCount int) error {
	if maxRetryCount < NO_RETRY_COUNT {
		return errors.New("maxRetryCount cannot be negative")
	}
	if maxRetryCount > MAX_RETRY_COUNT_LIMIT {
		return errors.New("maxRetryCount exceeds the limit")
	}
	cp.MaxRetryCount = maxRetryCount
	return nil
}

func (cp *ClientProfile) SetRequestRecover(requestRecover myRecover.RequestRecover) {
	cp.RequestRecover = requestRecover
}

type ExecuteClient interface {
	// 定义了一个方法，接受 BaseRequest 类型的参数，返回 BaseResponse 类型的结果
	DoExecute(request types.BaseRequest, data interface{}) error
}

func (c *DefaultClient) DoExecute(request types.BaseRequest, data interface{}) error {
	// 参数校验
	paramErr := request.ValidateParam()
	if paramErr != nil {
		return paramErr
	}
	if request.GetProtocol() == "" {
		if c.defaultProtocol != "" {
			request.SetProtocol(c.defaultProtocol)
		} else {
			request.SetProtocol(http.DEFAULT_PROTOCOL)
		}
	}
	if request.GetRegionCode() == "" {
		if c.defaultRegionCode != "" {
			request.SetRegionCode(c.defaultRegionCode)
		} else {
			request.SetRegionCode(http.DEFAULT_REGION_CODE)
		}
	}
	// 获取需要签名的参数
	request.SetCustomParams(request.GetBusinessCustomSignParams())
	// 获取不需要签名的参数
	request.SetNonSignParams(request.GetBusinessNonSignParams())
	// 尝试请求（并且按照策略尝试所有的域名）
	err := c.tryDoHttpRequest(NewClientContext(request, c), data)
	return err
}

// 尝试请求下一个地址直到所以可以申请的域名都已经尝试过
func (c *DefaultClient) tryDoHttpRequest(ctx *ClientContext, data interface{}) error {
	var yidunErr yidunerr.Error
	for {
		response := &http.CommonResponse{}
		// 执行请求
		request, createResqErr := ctx.CreateRequest()
		if createResqErr != nil {
			return yidunerr.New("CreateRequestErr", createResqErr.Error(), createResqErr)
		}
		// 发起http请求
		requestErr := c.customClient.Do(request, response)
		if requestErr != nil {
			yidunErr = yidunerr.New("FailedInitiateRequest", "Failed to initiate request！", requestErr)
			ctx.requestFail(response)
		} else {
			if isServerError(*response) {
				ctx.requestFail(response)
			} else {
				ctx.requestSuccess(response)
				yidunErr = nil
				break
			}
		}
		// 已重试次数+1
		ctx.incrementAttemptCount()
		// 判断是否可以继续尝试
		if !ctx.canAttempt() {
			break
		}
	}
	if yidunErr != nil {
		return yidunErr
	} else {
		// 将响应结果转换为 BaseResponse 对象
		convertErr := ctx.convertToBusinessResponse(data)
		if convertErr != nil {
			return convertErr
		}
		return nil
	}
}

// 判断是否是易盾服务的异常
func isServerError(response http.CommonResponse) bool {
	if response.BodyValue == nil || response.StatusCodeValue >= 500 {
		return true
	} else {
		return false
	}
}
