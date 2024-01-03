package request

import (
	"strconv"
	"strings"

	"github.com/yidun/yidun-golang-sdk/yidun/core/http"
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
)

const (
	MAX_DEPTH = 2
)
// LabelQueryRequest 标签查询请求
type LabelQueryRequest struct {
	*types.OpenApiGetRequest
	ClientID     *string `json:"clientId,omitempty"` // 客户ID
	Language     *string `json:"language,omitempty"` // 语种
    BusinessTypes *[]string `json:"businessTypes,omitempty"` // @see LabelBusinessTypeEnum
    BusinessID   *string `json:"businessId,omitempty"` // 业务ID
    MaxDepth     *int `json:"maxDepth,omitempty"` // 最大标签层级

}
// 构建request
func NewLabelQueryRequest() *LabelQueryRequest {
	var request = &LabelQueryRequest{
		OpenApiGetRequest: types.NewOpenApiGetRequest(),
	}
	request.SetProductCode("openapi")
	request.SetUriPattern("/openapi/v2/antispam/label/query")
	request.SetProtocol(http.ProtocolEnumHTTPS)
	request.SetMethod(http.HttpMethodGet)
	return request
}

// GetCustomSignParams 获取具体业务中特有的需要参与签名计算的参数
func (r *LabelQueryRequest) GetBusinessCustomSignParams() map[string]string {
	params := r.OpenApiGetRequest.GetBusinessCustomSignParams()
	if r.ClientID != nil {
        params["clientId"] = *r.ClientID
    }
    if r.Language != nil {
        params["language"] = *r.Language
    }
    if r.BusinessTypes != nil {
        params["businessTypes"] = strings.Join(*r.BusinessTypes, ",")
    }
    if r.BusinessID != nil {
        params["businessId"] = *r.BusinessID
    }
    if r.MaxDepth != nil {
        params["maxDepth"] = strconv.Itoa(*r.MaxDepth)
    }
	return params
}
// 设置ClientID
func (r *LabelQueryRequest) SetClientID(clientID string) {
    r.ClientID = &clientID
}
// 设置BusinessTypes
func (r *LabelQueryRequest) SetBusinessTypes(businessTypes *[]string) {
    r.BusinessTypes = businessTypes
}
// 设置BusinessID
func (r *LabelQueryRequest) SetBusinessID(businessID string) {
    r.BusinessID = &businessID
}
// 设置MaxDepth
func (r *LabelQueryRequest) SetMaxDepth(maxDepth int) {
    r.MaxDepth = &maxDepth
}
// 设置Language
func (r *LabelQueryRequest) SetLanguage(language string) {
    r.Language = &language
}