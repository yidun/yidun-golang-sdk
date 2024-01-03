package client

import (
	"encoding/json"
	"fmt"

	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/core/yidunerr"
)

type ClientContext struct {
	// 客户端对象
	*DefaultClient
	// 单次的请求对象
	request types.BaseRequest
	// 单次的响应结果
	response http.CommonResponse
	// 重试次数
	attemptedCount int
	// 是否使用请求对象中预置的域名
	usePreassignedDomain bool
	// 【域名动态获取场景】通过 {@link #endpointResolver} 获取的域名
	allDomains []string
	// 【域名动态获取场景】当前选域名的索引
	currentDomainIndex int
	// 【域名动态获取场景】是否为首次尝试提交请求
	firstAttempt bool
	// 域名动态获取场景】是否已向熔断器申请过所有域名
	allDomainsApplied bool
}

// NewEndpointContext 返回一个新的EndpointContext
func NewClientContext(request types.BaseRequest, client *DefaultClient) *ClientContext {
	return &ClientContext{
		DefaultClient:        client,
		request:              request,
		attemptedCount:       0,
		usePreassignedDomain: request.GetDomain() != "",
		allDomains:           make([]string, 0),
		currentDomainIndex:   -1,
		firstAttempt:         true,
		allDomainsApplied:    false,
	}
}

// 创建请求对象
func (c *ClientContext) CreateRequest() (http.Request, error) {
	err := c.resolveDomain()
	if err != nil {
		return nil, err
	}
	return c.request.ToHttpRequest(c.signer, c.credentials)
}

func (c *ClientContext) resolveDomain() error {
	if c.usePreassignedDomain {
		return nil
	}
	if c.firstAttempt {
		c.firstAttempt = false
		err := c.resolveDomainForFirstAttempt()
		return err
	} else {
		c.resolveDomainForRetry()
	}
	return nil
}

// 首次尝试提交请求时的域名获取逻辑
func (c *ClientContext) resolveDomainForFirstAttempt() error {
	allDomains, err := c.endpointResolver.Resolve(c.request.GetProductCode(), c.request.GetRegionCode())
	if err != nil {
		return err
	}
	c.allDomains = allDomains
	c.useNextAvailableDomainOrFirstAsDefault()
	return nil
}

// 重试提交请求时的域名获取逻辑
func (c *ClientContext) resolveDomainForRetry() {
	if c.allDomainsApplied {
		c.currentDomainIndex = (c.currentDomainIndex + 1) % len(c.allDomains)
		c.request.SetDomain(c.allDomains[c.currentDomainIndex])
		return
	}

	c.useNextAvailableDomainOrFirstAsDefault()
}

// 使用下一个可用的域名，如果没有可用的域名，则使用第一个域名
func (c *ClientContext) useNextAvailableDomainOrFirstAsDefault() {
	for i := c.currentDomainIndex + 1; i < len(c.allDomains); i++ {
		domain := c.allDomains[i]
		if c.breakStrategy.AttemptAccess(c.request.GetProductCode(), c.request.GetRegionCode(), domain) {
			c.request.SetDomain(domain)
			c.currentDomainIndex = i
			return
		}
	}

	c.allDomainsApplied = true
	c.currentDomainIndex = 0
	c.request.SetDomain(c.allDomains[0])
}

// 处理请求失败的情况
func (c *ClientContext) requestFail(response *http.CommonResponse) {
	c.response = *response
	if !c.usePreassignedDomain {
		c.breakStrategy.RequestFail(c.request.GetProductCode(), c.request.GetRegionCode(), c.request.GetDomain())
	}
}

// 请求成功时的执行逻辑
func (c *ClientContext) requestSuccess(response *http.CommonResponse) {
	c.response = *response
	if !c.usePreassignedDomain {
		c.breakStrategy.RequestSuccess(c.request.GetProductCode(), c.request.GetRegionCode(), c.request.GetDomain())
	}
}

// 判断是否需要重试
func (c *ClientContext) canAttempt() bool {
	return c.attemptedCount < int(c.maxAttemptCount)
}

// 当前次数 +1
func (c *ClientContext) incrementAttemptCount() {
	c.attemptedCount++
}

// 将响应结果转换为 BaseResponse 对象
func (c *ClientContext) convertToBusinessResponse(data interface{}) error {
	statusCode := c.response.StatusCode()
	if condition := statusCode >= 200 && statusCode < 300; !condition {
		return yidunerr.NewRequestFailure(
			yidunerr.New(fmt.Sprint(statusCode), http.StatusText(statusCode), nil), statusCode)
	}
	parseErr := json.Unmarshal(c.response.Body(), &data)
	if parseErr != nil {
		return yidunerr.NewUnmarshalError(parseErr, "解析响应结果失败！", c.response.Body())
	}
	return nil
}
