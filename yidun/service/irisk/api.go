package irisk

import (
	"strconv"

	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/core/validation"
)

// GetConfig 配置拉取，调用/risk/getConfig
func (c *IRiskClient) GetConfig(request *IRiskConfigRequest) (response *IRiskConfigResponse, err error) {
	response = &IRiskConfigResponse{}
	err = c.Client.DoExecute(request, response)
	return
}

// GetCheckResult 图片外挂识别，调用/risk/getMediaCheckResult
func (c *IRiskClient) GetCheckResult(request *IRiskCheckRequest) (response *IRiskCheckResponse, err error) {
	response = &IRiskCheckResponse{}
	err = c.Client.DoExecute(request, response)
	return
}

// GetDetailResult 在线检测，调用/risk/getDetailResult
func (c *IRiskClient) GetDetailResult(request *IRiskDetailRequest) (response *IRiskDetailResponse, err error) {
	response = &IRiskDetailResponse{}
	err = c.Client.DoExecute(request, response)
	return
}

// GetDetailV6Result 数据查询v6，调用/v6/risk/detail
func (c *IRiskClient) GetDetailV6Result(request *IRiskDetailV6Request) (response *IRiskDetailV6Response, err error) {
	response = &IRiskDetailV6Response{}
	err = c.Client.DoExecute(request, response)
	return
}

// GetMediaCheckResult 数据查询，调用/risk/getCheckResult
func (c *IRiskClient) GetMediaCheckResult(request *IRiskMediaCheckRequest) (response *IRiskMediaCheckResponse, err error) {
	response = &IRiskMediaCheckResponse{}
	err = c.Client.DoExecute(request, response)
	return
}

// GetMediaQueryResult 数据查询，调用/risk/mediaQuery
func (c *IRiskClient) GetMediaQueryResult(request *IRiskMediaQueryRequest) (response *IRiskMediaQueryResponse, err error) {
	response = &IRiskMediaQueryResponse{}
	err = c.Client.DoExecute(request, response)
	return
}

// GetMediaBatchQueryResult 数据查询，调用/risk/mediaBatchQuery
func (c *IRiskClient) GetMediaBatchQueryResult(request *IRiskMediaBatchQueryRequest) (response *IRiskMediaBatchQueryResponse, err error) {
	response = &IRiskMediaBatchQueryResponse{}
	err = c.Client.DoExecute(request, response)
	return
}

// UploadReportData 举报数据上报，调用/risk/reportData
func (c *IRiskClient) UploadReportData(request *IRiskReportDataRequest) (response *IRiskReportDataResponse, err error) {
	response = &IRiskReportDataResponse{}
	err = c.Client.DoExecute(request, response)
	return
}

// UploadDisposeInfo 上报处置数据，调用/disposeInfo/upload
func (c *IRiskClient) UploadDisposeInfo(request *IRiskDisposeUploadRequest) (response *IRiskDisposeUploadResponse, err error) {
	response = &IRiskDisposeUploadResponse{}
	err = c.Client.DoExecute(request, response)
	return
}

func (c *IRiskClient) AntiGoldCheck(request *IRiskAntiGoldCheckRequest) (response *IRiskAntiGoldCheckResponse, err error) {
	response = &IRiskAntiGoldCheckResponse{}
	err = c.Client.DoExecute(request, response)
	return
}

// ListAdd 上报处置数据，调用/list/add
func (c *IRiskClient) ListAdd(request *IRiskListAddRequest) (response *IRiskListAddResponse, err error) {
	response = &IRiskListAddResponse{}
	err = c.Client.DoExecute(request, response)
	return
}

// ListQuery上报处置数据，调用/list/add
func (c *IRiskClient) ListQuery(request *IRiskListQueryRequest) (response *IRiskListQueryResponse, err error) {
	response = &IRiskListQueryResponse{}
	err = c.Client.DoExecute(request, response)
	return
}

type IRiskCheckRequest struct {
	*types.BizPostJsonRequest
	Token        *string
	Ip           *string
	RoleId       *string
	Account      *string
	Nickname     *string
	Server       *string
	Phone        *string
	ActivityId   *string
	Target       *string
	Email        *string
	RegisterTime *int64
	RegisterIp   *string
	Level        *string
	Identity     *string
	PayUser      *bool
	Verified     *bool
	DeviceId     *string
	SceneData    *string
	ExtData      *string
	ClientCode   *int32
}

// SetToken 设置token
func (r *IRiskCheckRequest) SetToken(token string) *IRiskCheckRequest {
	r.Token = &token
	return r
}

// SetIp 设置IP地址
func (r *IRiskCheckRequest) SetIp(ip string) *IRiskCheckRequest {
	r.Ip = &ip
	return r
}

// SetRoleId 设置角色ID
func (r *IRiskCheckRequest) SetRoleId(roleId string) *IRiskCheckRequest {
	r.RoleId = &roleId
	return r
}

// SetAccount 设置账号
func (r *IRiskCheckRequest) SetAccount(account string) *IRiskCheckRequest {
	r.Account = &account
	return r
}

// SetNickname 设置昵称
func (r *IRiskCheckRequest) SetNickname(nickname string) *IRiskCheckRequest {
	r.Nickname = &nickname
	return r
}

// SetServer 设置服务器名称
func (r *IRiskCheckRequest) SetServer(server string) *IRiskCheckRequest {
	r.Server = &server
	return r
}

// SetPhone 设置手机号
func (r *IRiskCheckRequest) SetPhone(phone string) *IRiskCheckRequest {
	r.Phone = &phone
	return r
}

// SetActivityId 设置活动ID
func (r *IRiskCheckRequest) SetActivityId(activityId string) *IRiskCheckRequest {
	r.ActivityId = &activityId
	return r
}

// SetTarget 设置目标
func (r *IRiskCheckRequest) SetTarget(target string) *IRiskCheckRequest {
	r.Target = &target
	return r
}

// SetEmail 设置邮箱
func (r *IRiskCheckRequest) SetEmail(email string) *IRiskCheckRequest {
	r.Email = &email
	return r
}

// SetRegisterTime 设置注册时间
func (r *IRiskCheckRequest) SetRegisterTime(registerTime int64) *IRiskCheckRequest {
	r.RegisterTime = &registerTime
	return r
}

// SetRegisterIp 设置注册IP
func (r *IRiskCheckRequest) SetRegisterIp(registerIp string) *IRiskCheckRequest {
	r.RegisterIp = &registerIp
	return r
}

// SetLevel 设置等级
func (r *IRiskCheckRequest) SetLevel(level string) *IRiskCheckRequest {
	r.Level = &level
	return r
}

// SetIdentity 设置用户身份
func (r *IRiskCheckRequest) SetIdentity(identity string) *IRiskCheckRequest {
	r.Identity = &identity
	return r
}

// SetPayUser 设置是否付费用户
func (r *IRiskCheckRequest) SetPayUser(payUser bool) *IRiskCheckRequest {
	r.PayUser = &payUser
	return r
}

// SetVerified 设置是否认证用户
func (r *IRiskCheckRequest) SetVerified(verified bool) *IRiskCheckRequest {
	r.Verified = &verified
	return r
}

// SetDeviceId 设置设备ID
func (r *IRiskCheckRequest) SetDeviceId(deviceId string) *IRiskCheckRequest {
	r.DeviceId = &deviceId
	return r
}

// SetSceneData 设置场景数据
func (r *IRiskCheckRequest) SetSceneData(sceneData string) *IRiskCheckRequest {
	r.SceneData = &sceneData
	return r
}

// SetExtData 设置扩展数据
func (r *IRiskCheckRequest) SetExtData(extData string) *IRiskCheckRequest {
	r.ExtData = &extData
	return r
}

// SetClientCode 设置客户端接口状态码
func (r *IRiskCheckRequest) SetClientCode(clientCode int32) *IRiskCheckRequest {
	r.ClientCode = &clientCode
	return r
}

// NewIRiskCheckRequest 初始化IRiskCheckRequest对象
func NewIRiskCheckRequest(businessId string) *IRiskCheckRequest {
	request := &IRiskCheckRequest{
		BizPostJsonRequest: types.NewBizPostJsonRequest((businessId)),
	}
	request.SetProductCode("irisk")
	request.SetUriPattern("/v6/risk/check")
	request.SetVersion("603")
	return request
}

// GetNonSignParams 获取具体业务中特有的不参与签名计算的参数
func (r *IRiskCheckRequest) GetNonSignParams() map[string]string {
	return make(map[string]string)
}

// GetCustomSignParams 获取具体业务中特有的需要参与签名计算的参数
func (r *IRiskCheckRequest) GetBusinessCustomSignParams() map[string]string {
	params := r.BizPostJsonRequest.GetBusinessCustomSignParams()
	if r.Token != nil {
		params["token"] = *r.Token
	}
	params["ip"] = *r.Ip
	if r.RoleId != nil {
		params["roleId"] = *r.RoleId
	}
	if r.Account != nil {
		params["account"] = *r.Account
	}
	if r.Nickname != nil {
		params["nickname"] = *r.Nickname
	}
	if r.Server != nil {
		params["server"] = *r.Server
	}
	if r.Phone != nil {
		params["phone"] = *r.Phone
	}
	if r.ActivityId != nil {
		params["activityId"] = *r.ActivityId
	}
	if r.Target != nil {
		params["target"] = *r.Target
	}
	if r.Email != nil {
		params["email"] = *r.Email
	}
	if r.RegisterTime != nil {
		params["registerTime"] = strconv.FormatInt(*r.RegisterTime, 10)
	}
	if r.RegisterIp != nil {
		params["registerIp"] = *r.RegisterIp
	}
	if r.Level != nil {
		params["level"] = *r.Level
	}
	if r.Identity != nil {
		params["identity"] = *r.Identity
	}
	if r.PayUser != nil {
		params["payUser"] = strconv.FormatBool(*r.PayUser)
	}
	if r.Verified != nil {
		params["verified"] = strconv.FormatBool(*r.Verified)
	}
	if r.DeviceId != nil {
		params["deviceId"] = *r.DeviceId
	}
	if r.SceneData != nil {
		params["sceneData"] = *r.SceneData
	}
	if r.ExtData != nil {
		params["extData"] = *r.ExtData
	}
	if r.ClientCode != nil {
		params["clientCode"] = strconv.Itoa(int(*r.ClientCode))
	}
	return params
}

// check接口 参数校验方法
func (r *IRiskCheckRequest) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "IRiskCheckRequest"}
	if r.Ip == nil {
		invalidParams.Add(validation.NewErrParamRequired("Ip"))
	}
	if r.Account == nil {
		invalidParams.Add(validation.NewErrParamRequired("Account"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type IRiskCheckResponse struct {
	*types.CommonResponse
	Data IRiskCheckResult `json:"data"`
	Desc string           `json:"desc"`
}

type IRiskCheckResult struct {
	RiskLevel          int         `json:"riskLevel"`
	HitInfos           []HitInfo   `json:"hitInfos"`
	TaskId             string      `json:"taskId"`
	SdkRespData        string      `json:"sdkRespData"`
	DeviceId           string      `json:"deviceId"`
	PcId               string      `json:"pcId"`
	MatchedRules       []string    `json:"matchedRules"`
	MatchedCustomRules []string    `json:"matchedCustomRules"`
	DeviceInfo         interface{} `json:"deviceInfo"`
	Evidences          []string    `json:"evidences"`
	PhoneInfo          PhoneInfo   `json:"phoneInfo"`
	IpInfo             IpInfo      `json:"ipInfo"`
	Action             int         `json:"action"`
	CustomAction       int         `json:"customAction"`
}

type HitInfo struct {
	Type      string   `json:"type"`
	Name      string   `json:"name"`
	RiskLevel int      `json:"riskLevel"`
	Rules     []string `json:"rules"`
}

type (
	PhoneInfo struct {
		BasicInfo     PhoneBasicInfo `json:"basicInfo"`
		PhoneRiskInfo PhoneRiskInfo  `json:"phoneRiskInfo"`
	}
	PhoneBasicInfo struct {
		Carrier  string `json:"carrier"`
		Province string `json:"province"`
		City     string `json:"city"`
	}
	PhoneRiskInfo struct {
		RiskType  string  `json:"riskType"`
		RiskLevel int     `json:"riskLevel"`
		RiskScore float64 `json:"riskScore"`
	}
)

type (
	IpInfo struct {
		SdkIp      string      `json:"sdkIp"`
		BasicInfo  IpBasicInfo `json:"basicInfo"`
		IpRiskInfo IpRiskInfo  `json:"ipRiskInfo"`
	}
	IpBasicInfo struct {
		Isp       string `json:"isp"`
		Continent string `json:"continent"`
		Country   string `json:"country"`
		Province  string `json:"province"`
		City      string `json:"city"`
		Longitude string `json:"longitude"`
		Latitude  string `json:"latitude"`
	}
	IpRiskInfo struct {
		RiskType  string  `json:"riskType"`
		RiskLevel int     `json:"riskLevel"`
		RiskScore float64 `json:"riskScore"`
	}
)

type IRiskConfigRequest struct {
	*types.BizPostJsonRequest
	Ip      *string `type:"string" required:"true"`
	SdkData *string
}

// SetIp 设置IP地址
func (r *IRiskConfigRequest) SetIp(ip string) *IRiskConfigRequest {
	r.Ip = &ip
	return r
}

// SetSdkData 设置SDK数据
func (r *IRiskConfigRequest) SetSdkData(sdkData string) *IRiskConfigRequest {
	r.SdkData = &sdkData
	return r
}

// NewIRiskConfigRequest 初始化IRiskConfigRequest对象
func NewIRiskConfigRequest(businessId string) *IRiskConfigRequest {
	request := &IRiskConfigRequest{
		BizPostJsonRequest: types.NewBizPostJsonRequest((businessId)),
	}
	request.SetProductCode("irisk")
	request.SetUriPattern("/v5/risk/getConfig")
	request.SetVersion("500")
	return request
}

// GetNonSignParams 获取具体业务中特有的不参与签名计算的参数
func (r *IRiskConfigRequest) GetNonSignParams() map[string]string {
	return make(map[string]string)
}

// GetCustomSignParams 获取具体业务中特有的需要参与签名计算的参数
func (r *IRiskConfigRequest) GetBusinessCustomSignParams() map[string]string {
	params := r.BizPostJsonRequest.GetBusinessCustomSignParams()
	if r.Ip != nil {
		params["ip"] = *r.Ip
	}
	if r.SdkData != nil {
		params["sdkData"] = *r.SdkData
	}
	return params
}

// config接口 参数校验方法
func (r *IRiskConfigRequest) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "IRiskConfigRequest"}
	if r.Ip == nil {
		invalidParams.Add(validation.NewErrParamRequired("Ip"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type IRiskConfigResponse struct {
	*types.CommonResponse
	Data IRiskConfigResult `json:"data"`
}

type IRiskConfigResult struct {
	ConfigData string `json:"configData"`
}

type IRiskDetailRequest struct {
	*types.BizPostJsonRequest
	BeginTimestamp *int64
	EndTimestamp   *int64
	StartFlag      *string
	Account        *string
	Accounts       *string
	RoleId         *string
	RoleIds        *string
	RiskLevel      *int
	Ip             *string
	Ips            *string
	SdkIp          *string
	SdkIps         *string
	PackageName    *string
	AppVersion     *string
}

// SetBeginTimestamp 设置开始时间
func (r *IRiskDetailRequest) SetBeginTimestamp(beginTimestamp int64) *IRiskDetailRequest {
	r.BeginTimestamp = &beginTimestamp
	return r
}

// SetEndTimestamp 设置结束时间
func (r *IRiskDetailRequest) SetEndTimestamp(endTimestamp int64) *IRiskDetailRequest {
	r.EndTimestamp = &endTimestamp
	return r
}

// SetStartFlag 设置开始标识
func (r *IRiskDetailRequest) SetStartFlag(startFlag string) *IRiskDetailRequest {
	r.StartFlag = &startFlag
	return r
}

// SetAccount 设置账号
func (r *IRiskDetailRequest) SetAccount(account string) *IRiskDetailRequest {
	r.Account = &account
	return r
}

// SetAccounts 设置账号集合(JSON字符串)
func (r *IRiskDetailRequest) SetAccounts(accounts string) *IRiskDetailRequest {
	r.Accounts = &accounts
	return r
}

// SetRoleId 设置角色ID
func (r *IRiskDetailRequest) SetRoleId(roleId string) *IRiskDetailRequest {
	r.RoleId = &roleId
	return r
}

// SetRoleId 设置角色ID集合(JSON字符串)
func (r *IRiskDetailRequest) SetRoleIds(roleIds string) *IRiskDetailRequest {
	r.RoleIds = &roleIds
	return r
}

// SetRiskLevel 设置风险等级
func (r *IRiskDetailRequest) SetRiskLevel(riskLevel int) *IRiskDetailRequest {
	r.RiskLevel = &riskLevel
	return r
}

// SetIp 设置IP地址
func (r *IRiskDetailRequest) SetIp(ip string) *IRiskDetailRequest {
	r.Ip = &ip
	return r
}

// SetIp 设置IP地址集合(JSON字符串)
func (r *IRiskDetailRequest) SetIps(ips string) *IRiskDetailRequest {
	r.Ips = &ips
	return r
}

// SetSdkIp 设置SDK IP地址
func (r *IRiskDetailRequest) SetSdkIp(sdkIp string) *IRiskDetailRequest {
	r.SdkIp = &sdkIp
	return r
}

// SetSdkIp 设置SDK IP地址
func (r *IRiskDetailRequest) SetSdkIps(sdkIps string) *IRiskDetailRequest {
	r.SdkIps = &sdkIps
	return r
}

// SetPackageName 设置包名
func (r *IRiskDetailRequest) SetPackageName(packageName string) *IRiskDetailRequest {
	r.PackageName = &packageName
	return r
}

// SetAppVersion 设置APP版本
func (r *IRiskDetailRequest) SetAppVersion(appVersion string) *IRiskDetailRequest {
	r.AppVersion = &appVersion
	return r
}

// NewIRiskDetailRequest 初始化IRiskDetailRequest对象
func NewIRiskDetailRequest(businessId string) *IRiskDetailRequest {
	request := &IRiskDetailRequest{
		BizPostJsonRequest: types.NewBizPostJsonRequest((businessId)),
	}
	request.SetProductCode("irisk")
	request.SetUriPattern("/v5/risk/detail")
	request.SetVersion("500")
	return request
}

// GetNonSignParams 获取具体业务中特有的不参与签名计算的参数
func (r *IRiskDetailRequest) GetNonSignParams() map[string]string {
	return make(map[string]string)
}

// GetCustomSignParams 获取具体业务中特有的需要参与签名计算的参数
func (r *IRiskDetailRequest) GetBusinessCustomSignParams() map[string]string {
	params := r.BizPostJsonRequest.GetBusinessCustomSignParams()
	if r.BeginTimestamp != nil {
		params["beginTimestamp"] = strconv.FormatInt(*r.BeginTimestamp, 10)
	}
	if r.EndTimestamp != nil {
		params["endTimestamp"] = strconv.FormatInt(*r.EndTimestamp, 10)
	}
	if r.StartFlag != nil {
		params["startFlag"] = *r.StartFlag
	}
	if r.Account != nil {
		params["account"] = *r.Account
	}
	if r.Accounts != nil {
		params["accounts"] = *r.Accounts
	}
	if r.RoleId != nil {
		params["roleId"] = *r.RoleId
	}
	if r.RoleIds != nil {
		params["roleIds"] = *r.RoleIds
	}
	if r.RiskLevel != nil {
		params["riskLevel"] = strconv.Itoa(*r.RiskLevel)
	}
	if r.Ip != nil {
		params["ip"] = *r.Ip
	}
	if r.Ips != nil {
		params["ips"] = *r.Ips
	}
	if r.SdkIp != nil {
		params["sdkIp"] = *r.SdkIp
	}
	if r.SdkIps != nil {
		params["sdkIps"] = *r.SdkIps
	}
	if r.PackageName != nil {
		params["packageName"] = *r.PackageName
	}
	if r.AppVersion != nil {
		params["appVersion"] = *r.AppVersion
	}
	return params
}

// detail接口 参数校验方法
func (r *IRiskDetailRequest) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "IRiskDetailRequest"}
	if r.BeginTimestamp == nil {
		invalidParams.Add(validation.NewErrParamRequired("BeginTimestamp"))
	}
	if r.EndTimestamp == nil {
		invalidParams.Add(validation.NewErrParamRequired("EndTimestamp"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type IRiskDetailResponse struct {
	*types.CommonResponse
	Data IRiskDetailResult `json:"data"`
}

type IRiskDetailResult struct {
	Size      int         `json:"size"`
	StartFlag string      `json:"startFlag"`
	Detail    interface{} `json:"detail"`
}

type IRiskDetailV6Request struct {
	*types.BizPostJsonRequest
	BeginTimestamp  *int64
	EndTimestamp    *int64
	StartFlag       *string
	Account         *string
	Accounts        *string
	RoleId          *string
	RoleIds         *string
	RiskLevel       *int
	Ip              *string
	Ips             *string
	SdkIp           *string
	SdkIps          *string
	PackageName     *string
	AppVersion      *string
	MatchedRiskTags *string
	MatchedTypes    *string
}

// SetBeginTimestamp 设置开始时间
func (r *IRiskDetailV6Request) SetBeginTimestamp(beginTimestamp int64) *IRiskDetailV6Request {
	r.BeginTimestamp = &beginTimestamp
	return r
}

// SetEndTimestamp 设置结束时间
func (r *IRiskDetailV6Request) SetEndTimestamp(endTimestamp int64) *IRiskDetailV6Request {
	r.EndTimestamp = &endTimestamp
	return r
}

// SetStartFlag 设置开始标识
func (r *IRiskDetailV6Request) SetStartFlag(startFlag string) *IRiskDetailV6Request {
	r.StartFlag = &startFlag
	return r
}

// SetAccount 设置账号
func (r *IRiskDetailV6Request) SetAccount(account string) *IRiskDetailV6Request {
	r.Account = &account
	return r
}

// SetAccounts 设置账号集合(JSON字符串)
func (r *IRiskDetailV6Request) SetAccounts(accounts string) *IRiskDetailV6Request {
	r.Accounts = &accounts
	return r
}

// SetRoleId 设置角色ID
func (r *IRiskDetailV6Request) SetRoleId(roleId string) *IRiskDetailV6Request {
	r.RoleId = &roleId
	return r
}

// SetRoleIds 设置角色ID集合(JSON字符串)
func (r *IRiskDetailV6Request) SetRoleIds(roleIds string) *IRiskDetailV6Request {
	r.RoleIds = &roleIds
	return r
}

// SetRiskLevel 设置风险等级
func (r *IRiskDetailV6Request) SetRiskLevel(riskLevel int) *IRiskDetailV6Request {
	r.RiskLevel = &riskLevel
	return r
}

// SetIp 设置IP地址
func (r *IRiskDetailV6Request) SetIp(ip string) *IRiskDetailV6Request {
	r.Ip = &ip
	return r
}

// SetIps 设置IP地址集合(JSON字符串)
func (r *IRiskDetailV6Request) SetIps(ips string) *IRiskDetailV6Request {
	r.Ips = &ips
	return r
}

// SetSdkIp 设置SDK IP地址
func (r *IRiskDetailV6Request) SetSdkIp(sdkIp string) *IRiskDetailV6Request {
	r.SdkIp = &sdkIp
	return r
}

// SetSdkIp 设置SDK IP地址集合(JSON字符串)
func (r *IRiskDetailV6Request) SetSdkIps(sdkIps string) *IRiskDetailV6Request {
	r.SdkIps = &sdkIps
	return r
}

// SetPackageName 设置包名
func (r *IRiskDetailV6Request) SetPackageName(packageName string) *IRiskDetailV6Request {
	r.PackageName = &packageName
	return r
}

// SetAppVersion 设置APP版本
func (r *IRiskDetailV6Request) SetAppVersion(appVersion string) *IRiskDetailV6Request {
	r.AppVersion = &appVersion
	return r
}

// SetMatchedRiskTags 设置命中的风险标签
func (r *IRiskDetailV6Request) SetMatchedRiskTags(matchedRiskTags string) *IRiskDetailV6Request {
	r.MatchedRiskTags = &matchedRiskTags
	return r
}

// SetMatchedTypes 设置命中的风险标签Type
func (r *IRiskDetailV6Request) SetMatchedTypes(matchedTypes string) *IRiskDetailV6Request {
	r.MatchedTypes = &matchedTypes
	return r
}

// NewIRiskDetailV6Request 初始化IRiskDetailV6Request对象
func NewIRiskDetailV6Request(businessId string) *IRiskDetailV6Request {
	request := &IRiskDetailV6Request{
		BizPostJsonRequest: types.NewBizPostJsonRequest((businessId)),
	}
	request.SetProductCode("irisk")
	request.SetUriPattern("/v6/risk/detail")
	request.SetVersion("600")
	return request
}

// GetNonSignParams 获取具体业务中特有的不参与签名计算的参数
func (r *IRiskDetailV6Request) GetNonSignParams() map[string]string {
	return make(map[string]string)
}

// GetCustomSignParams 获取具体业务中特有的需要参与签名计算的参数
func (r *IRiskDetailV6Request) GetBusinessCustomSignParams() map[string]string {
	params := r.BizPostJsonRequest.GetBusinessCustomSignParams()
	if r.BeginTimestamp != nil {
		params["beginTimestamp"] = strconv.FormatInt(*r.BeginTimestamp, 10)
	}
	if r.EndTimestamp != nil {
		params["endTimestamp"] = strconv.FormatInt(*r.EndTimestamp, 10)
	}
	if r.StartFlag != nil {
		params["startFlag"] = *r.StartFlag
	}
	if r.Account != nil {
		params["account"] = *r.Account
	}
	if r.Accounts != nil {
		params["accounts"] = *r.Accounts
	}
	if r.RoleId != nil {
		params["roleId"] = *r.RoleId
	}
	if r.RoleIds != nil {
		params["roleIds"] = *r.RoleIds
	}
	if r.RiskLevel != nil {
		params["riskLevel"] = strconv.Itoa(*r.RiskLevel)
	}
	if r.Ip != nil {
		params["ip"] = *r.Ip
	}
	if r.Ips != nil {
		params["ips"] = *r.Ips
	}
	if r.SdkIp != nil {
		params["sdkIp"] = *r.SdkIp
	}
	if r.SdkIps != nil {
		params["sdkIps"] = *r.SdkIps
	}
	if r.PackageName != nil {
		params["packageName"] = *r.PackageName
	}
	if r.AppVersion != nil {
		params["appVersion"] = *r.AppVersion
	}
    if r.MatchedRiskTags != nil {
        params["matchedRiskTags"] = *r.MatchedRiskTags
    }
    if r.MatchedTypes != nil {
        params["matchedTypes"] = *r.MatchedTypes
    }
	return params
}

// detail接口 参数校验方法
func (r *IRiskDetailV6Request) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "IRiskDetailV6Request"}
	if r.BeginTimestamp == nil {
		invalidParams.Add(validation.NewErrParamRequired("BeginTimestamp"))
	}
	if r.EndTimestamp == nil {
		invalidParams.Add(validation.NewErrParamRequired("EndTimestamp"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type IRiskDetailV6Response struct {
	*types.CommonResponse
	Data IRiskDetailV6Result `json:"data"`
}

type IRiskDetailV6Result struct {
	Size      int            `json:"size"`
	StartFlag string         `json:"startFlag"`
	Detail    []DetailDataV6 `json:"detail"`
}

type DetailDataV6 struct {
	ReceiveTime      string `json:"receiveTime"`
	Platform         string `json:"platform"`
	BusinessId       string `json:"businessId"`
	BusinessName     string `json:"businessName"`
	Token            string `json:"token"`
	TaskId           string `json:"taskId"`
	CheckSource      string `json:"checkSource"`
	SceneType        string `json:"sceneType"`
	SessionId        string `json:"sessionId"`
	TimeoutToken     string `json:"timeoutToken"`
	ActivityId       string `json:"activityId"`
	Target           string `json:"target"`
	Account          string `json:"account"`
	RoleId           string `json:"roleId"`
	Nickname         string `json:"nickname"`
	Phone            string `json:"phone"`
	Identity         string `json:"identity"`
	PayUser          string `json:"payUser"`
	Verified         string `json:"verified"`
	Level            string `json:"level"`
	RegisterChannel  string `json:"registerChannel"`
	RegisterDays     string `json:"registerDays"`
	OnlineTimes      string `json:"onlineTimes"`
	Gender           string `json:"gender"`
	AppName          string `json:"appName"`
	PackageName      string `json:"packageName"`
	Author           string `json:"author"`
	SignV1           string `json:"signV1"`
	SignV2           string `json:"signV2"`
	AppVersion       string `json:"appVersion"`
	GameVersion      string `json:"gameVersion"`
	AssetVersion     string `json:"assetVersion"`
	SdkVersion       string `json:"sdkVersion"`
	Server           string `json:"server"`
	ServerId         string `json:"serverId"`
	GameJson         string `json:"gameJson"`
	DeviceId         string `json:"deviceId"`
	DeviceType       string `json:"deviceType"`
	DeviceBrand      string `json:"deviceBrand"`
	DeviceModel      string `json:"deviceModel"`
	OsVersion        string `json:"osVersion"`
	CarrierName      string `json:"carrierName"`
	EmulatorDid      string `json:"emulatorDid"`
	LocalCacheId     string `json:"localCacheId"`
	Mac              string `json:"mac"`
	DeviceGuid       string `json:"deviceGuid"`
	NetworkType      string `json:"networkType"`
	Location         string `json:"location"`
	SdkIp            string `json:"sdkIp"`
	Ip               string `json:"ip"`
	Isp              string `json:"isp"`
	SceneData        string `json:"sceneData"`
	RiskLevel        string `json:"riskLevel"`
	MatchedRiskTags  string `json:"matchedRiskTags"`
	MatchedTypes     string `json:"matchedTypes"`
	ExecAction       string `json:"execAction"`
	CustomAction     string `json:"customAction"`
	ProtectionResult string `json:"protectionResult"`
	Evidence         string `json:"evidence"`
	MatchedRules     string `json:"matchedRules"`
	PhoneCarrierName string `json:"phoneCarrierName"`
	PhoneLocation    string `json:"phoneLocation"`
	PhoneRiskType    string `json:"phoneRiskType"`
	PhoneRiskLevel   string `json:"phoneRiskLevel"`
	PhoneRiskScore   string `json:"phoneRiskScore"`
	IpRiskType       string `json:"ipRiskType"`
	IpRiskLevel      string `json:"ipRiskLevel"`
	IpRiskScore      string `json:"ipRiskScore"`
	ExtData          string `json:"extData"`
}

type IRiskMediaCheckRequest struct {
	*types.BizPostJsonRequest
	MediaData *string
	MediaName *string
	Ip        *string
	RoleId    *string
	Nickname  *string
	Server    *string
}

// SetMediaData 设置media数据
func (r *IRiskMediaCheckRequest) SetMediaData(mediaData string) *IRiskMediaCheckRequest {
	r.MediaData = &mediaData
	return r
}

// SetMediaName 设置media名称
func (r *IRiskMediaCheckRequest) SetMediaName(mediaName string) *IRiskMediaCheckRequest {
	r.MediaName = &mediaName
	return r
}

// SetIp 设置IP地址
func (r *IRiskMediaCheckRequest) SetIp(ip string) *IRiskMediaCheckRequest {
	r.Ip = &ip
	return r
}

// SetRoleId 设置角色ID
func (r *IRiskMediaCheckRequest) SetRoleId(roleId string) *IRiskMediaCheckRequest {
	r.RoleId = &roleId
	return r
}

// SetNickname 设置昵称
func (r *IRiskMediaCheckRequest) SetNickname(nickname string) *IRiskMediaCheckRequest {
	r.Nickname = &nickname
	return r
}

// SetServer 设置服务器名称
func (r *IRiskMediaCheckRequest) SetServer(server string) *IRiskMediaCheckRequest {
	r.Server = &server
	return r
}

// NewIRiskMediaCheckRequest 初始化IRiskMediaCheckRequest对象
func NewIRiskMediaCheckRequest(businessId string) *IRiskMediaCheckRequest {
	request := &IRiskMediaCheckRequest{
		BizPostJsonRequest: types.NewBizPostJsonRequest((businessId)),
	}
	request.SetProductCode("irisk")
	request.SetUriPattern("/v5/risk/mediaCheck")
	request.SetVersion("500")
	return request
}

type IRiskMediaQueryRequest struct {
	*types.BizPostJsonRequest
	TaskId *string
}

// SetTaskId 设置taskId
func (r *IRiskMediaQueryRequest) SetTaskId(taskId string) *IRiskMediaQueryRequest {
	r.TaskId = &taskId
	return r
}

// NewIRiskMediaQueryRequest 初始化NewIRiskMediaQueryRequest对象
func NewIRiskMediaQueryRequest(businessId string) *IRiskMediaQueryRequest {
	request := &IRiskMediaQueryRequest{
		BizPostJsonRequest: types.NewBizPostJsonRequest((businessId)),
	}
	request.SetProductCode("irisk")
	request.SetUriPattern("/v5/risk/mediaQuery")
	request.SetVersion("500")
	return request
}

type IRiskMediaBatchQueryRequest struct {
	*types.BizPostJsonRequest
	TaskIds *string
}

// SetTaskIds 设置taskIds
func (r *IRiskMediaBatchQueryRequest) SetTaskIds(taskIds string) *IRiskMediaBatchQueryRequest {
	r.TaskIds = &taskIds
	return r
}

// NewIRiskMediaBatchQueryRequest 初始化NewIRiskMediaBatchQueryRequest对象
func NewIRiskMediaBatchQueryRequest(businessId string) *IRiskMediaBatchQueryRequest {
	request := &IRiskMediaBatchQueryRequest{
		BizPostJsonRequest: types.NewBizPostJsonRequest((businessId)),
	}
	request.SetProductCode("irisk")
	request.SetUriPattern("/v5/risk/mediaBatchQuery")
	request.SetVersion("500")
	return request
}


// GetNonSignParams 获取具体业务中特有的不参与签名计算的参数
func (r *IRiskMediaCheckRequest) GetNonSignParams() map[string]interface{} {
	return make(map[string]interface{})
}

// GetCustomSignParams 获取具体业务中特有的需要参与签名计算的参数
func (r *IRiskMediaCheckRequest) GetBusinessCustomSignParams() map[string]string {
	params := r.BizPostJsonRequest.GetBusinessCustomSignParams()
	if r.MediaData != nil {
		params["mediaData"] = *r.MediaData
	}
	if r.MediaName != nil {
		params["mediaName"] = *r.MediaName
	}
	if r.Ip != nil {
		params["ip"] = *r.Ip
	}
	if r.RoleId != nil {
		params["roleId"] = *r.RoleId
	}
	if r.Nickname != nil {
		params["nickname"] = *r.Nickname
	}
	if r.Server != nil {
		params["server"] = *r.Server
	}
	return params
}

// mediaCheck接口 参数校验方法
func (r *IRiskMediaCheckRequest) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "IRiskMediaCheckRequest"}
	if r.MediaData == nil {
		invalidParams.Add(validation.NewErrParamRequired("MediaData"))
	}
	if r.MediaName == nil {
		invalidParams.Add(validation.NewErrParamRequired("MediaName"))
	}
	if r.Ip == nil {
		invalidParams.Add(validation.NewErrParamRequired("Ip"))
	}
	if r.RoleId == nil {
		invalidParams.Add(validation.NewErrParamRequired("RoleId"))
	}
	if r.Nickname == nil {
		invalidParams.Add(validation.NewErrParamRequired("Nickname"))
	}
	if r.Server == nil {
		invalidParams.Add(validation.NewErrParamRequired("Server"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type IRiskMediaCheckResponse struct {
	*types.CommonResponse
	Data IRiskMediaCheckResult `json:"data"`
}

type IRiskMediaCheckResult struct {
	TaskId string `json:"taskId"`
}

type IRiskMediaQueryResponse struct {
	*types.CommonResponse
	Data IRiskMediaQueryResult `json:"data"`
}

type IRiskMediaBatchQueryResponse struct {
	*types.CommonResponse
	Data []IRiskMediaQueryResult `json:"data"`
}

type IRiskMediaQueryResult struct {
	*types.BizPostJsonRequest
	ReceiveTime int64    `json:"receiveTime"`
	Ip          string   `json:"ip"`
	RoleId      string   `json:"roleId"`
	Nickname    string   `json:"nickname"`
	Server      string   `json:"server"`
	Status      int      `json:"status"`
	TagNameList []string `json:"tagNameList"`
	Reason      string   `json:"reason"`
}

type IRiskReportDataRequest struct {
	*types.BizPostJsonRequest
	ReportChannel  *string
	ReportTime     *int64
	Whistleblower  *string
	ReportedPerson *string
	ReportType     *string
	ReportScene    *string
	ReportData     *string
}

// SetReportChannel 设置举报渠道
func (r *IRiskReportDataRequest) SetReportChannel(reportChannel string) *IRiskReportDataRequest {
	r.ReportChannel = &reportChannel
	return r
}

// SetReportTime 设置举报时间
func (r *IRiskReportDataRequest) SetReportTime(reportTime int64) *IRiskReportDataRequest {
	r.ReportTime = &reportTime
	return r
}

// SetWhistleblower 设置举报人
func (r *IRiskReportDataRequest) SetWhistleblower(whistleblower string) *IRiskReportDataRequest {
	r.Whistleblower = &whistleblower
	return r
}

// SetReportedPerson 设置被举报人
func (r *IRiskReportDataRequest) SetReportedPerson(reportedPerson string) *IRiskReportDataRequest {
	r.ReportedPerson = &reportedPerson
	return r
}

// SetReportType 设置被举报人
func (r *IRiskReportDataRequest) SetReportType(reportType string) *IRiskReportDataRequest {
	r.ReportType = &reportType
	return r
}

// SetReportScene 设置被举报人
func (r *IRiskReportDataRequest) SetReportScene(reportScene string) *IRiskReportDataRequest {
	r.ReportScene = &reportScene
	return r
}

// SetReportData 设置举报内容
func (r *IRiskReportDataRequest) SetReportData(reportData string) *IRiskReportDataRequest {
	r.ReportData = &reportData
	return r
}

// NewIRiskReportDataRequest 初始化IRiskMediaCheckRequest对象
func NewIRiskReportDataRequest(businessId string) *IRiskReportDataRequest {
	request := &IRiskReportDataRequest{
		BizPostJsonRequest: types.NewBizPostJsonRequest((businessId)),
	}
	request.SetProductCode("irisk")
	request.SetUriPattern("/v5/risk/reportData")
	request.SetVersion("500")
	return request
}

// GetNonSignParams 获取具体业务中特有的不参与签名计算的参数
func (r *IRiskReportDataRequest) GetNonSignParams() map[string]interface{} {
	return make(map[string]interface{})
}

// GetBusinessCustomSignParams 获取具体业务中特有的需要参与签名计算的参数
func (r *IRiskReportDataRequest) GetBusinessCustomSignParams() map[string]string {
	params := r.BizPostJsonRequest.GetBusinessCustomSignParams()
	if r.ReportChannel != nil {
		params["reportChannel"] = *r.ReportChannel
	}
	if r.ReportTime != nil {
		params["reportTime"] = strconv.FormatInt(*r.ReportTime, 10)
	}
	if r.Whistleblower != nil {
		params["whistleblower"] = *r.Whistleblower
	}
	if r.ReportedPerson != nil {
		params["reportedPerson"] = *r.ReportedPerson
	}
	if r.ReportType != nil {
		params["reportType"] = *r.ReportType
	}
	if r.ReportScene != nil {
		params["reportScene"] = *r.ReportScene
	}
	if r.ReportData != nil {
		params["reportData"] = *r.ReportData
	}
	return params
}

// reportData接口 参数校验方法
func (r *IRiskReportDataRequest) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "IRiskReportDataRequest"}
	if r.ReportChannel == nil {
		invalidParams.Add(validation.NewErrParamRequired("ReportChannel"))
	}
	if r.ReportTime == nil {
		invalidParams.Add(validation.NewErrParamRequired("ReportTime"))
	}
	if r.Whistleblower == nil {
		invalidParams.Add(validation.NewErrParamRequired("Whistleblower"))
	}
	if r.ReportedPerson == nil {
		invalidParams.Add(validation.NewErrParamRequired("ReportedPerson"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type IRiskReportDataResponse struct {
	*types.CommonResponse
}

type IRiskDisposeUploadRequest struct {
	*types.BizPostJsonRequest
	HandleObjectType *int
	Action           *int
	ObjectId         *string
	RoleId           *string
	ActionDes        *string
	Reason           *string
}

// SetHandleObjectType 设置处置对象类型
func (r *IRiskDisposeUploadRequest) SetHandleObjectType(handleObjectType int) *IRiskDisposeUploadRequest {
	r.HandleObjectType = &handleObjectType
	return r
}

// SetAction 设置处置动作
func (r *IRiskDisposeUploadRequest) SetAction(action int) *IRiskDisposeUploadRequest {
	r.Action = &action
	return r
}

// SetObjectId 设置处置动作
func (r *IRiskDisposeUploadRequest) SetObjectId(objectId string) *IRiskDisposeUploadRequest {
	r.ObjectId = &objectId
	return r
}

// SetRoleId 设置角色Id
func (r *IRiskDisposeUploadRequest) SetRoleId(roleId string) *IRiskDisposeUploadRequest {
	r.RoleId = &roleId
	return r
}

// SetActionDesc 设置处置动作解释
func (r *IRiskDisposeUploadRequest) SetActionDesc(actionDes string) *IRiskDisposeUploadRequest {
	r.ActionDes = &actionDes
	return r
}

// SetReason 设置处置原因
func (r *IRiskDisposeUploadRequest) SetReason(reason string) *IRiskDisposeUploadRequest {
	r.Reason = &reason
	return r
}

// NewIRiskDisposeUploadRequest 初始化IRiskDisposeUploadRequest对象
func NewIRiskDisposeUploadRequest(businessId string) *IRiskDisposeUploadRequest {
	request := &IRiskDisposeUploadRequest{
		BizPostJsonRequest: types.NewBizPostJsonRequest((businessId)),
	}
	request.SetProductCode("irisk")
	request.SetUriPattern("/v5/disposeInfo/upload")
	request.SetVersion("500")
	return request
}

// GetNonSignParams 获取具体业务中特有的不参与签名计算的参数
func (r *IRiskDisposeUploadRequest) GetNonSignParams() map[string]interface{} {
	return make(map[string]interface{})
}

// GetBusinessCustomSignParams 获取具体业务中特有的需要参与签名计算的参数
func (r *IRiskDisposeUploadRequest) GetBusinessCustomSignParams() map[string]string {
	params := r.BizPostJsonRequest.GetBusinessCustomSignParams()
	if r.HandleObjectType != nil {
		params["handleObjectType"] = strconv.Itoa(*r.HandleObjectType)
	}
	if r.Action != nil {
		params["action"] = strconv.Itoa(*r.Action)
	}
	if r.ObjectId != nil {
		params["objectId"] = *r.ObjectId
	}
	if r.RoleId != nil {
		params["roleId"] = *r.RoleId
	}
	if r.ActionDes != nil {
		params["actionDes"] = *r.ActionDes
	}
	if r.Reason != nil {
		params["reason"] = *r.Reason
	}
	return params
}

// studioQuery接口 参数校验方法
func (r *IRiskDisposeUploadRequest) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "IRiskDisposeUploadRequest"}
	if r.HandleObjectType == nil {
		invalidParams.Add(validation.NewErrParamRequired("handleObjectType"))
	}
	if r.ObjectId == nil {
		invalidParams.Add(validation.NewErrParamRequired("objectId"))
	}
	if r.Action == nil {
		invalidParams.Add(validation.NewErrParamRequired("action"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type IRiskDisposeUploadResponse struct {
	*types.CommonResponse
}

type IRiskAntiGoldCheckRequest struct {
	*types.BizPostJsonRequest
	LogTime  *string
	Account  *string
	RoleId   *string
	Nickname *string
	ServerId *string
	LogType  *string
	LogData  *string
}

// SetLogTime 设置实际记录日志的时间
func (r *IRiskAntiGoldCheckRequest) SetLogTime(logTime string) *IRiskAntiGoldCheckRequest {
	r.LogTime = &logTime
	return r
}

// SetAccount 设置用户账号
func (r *IRiskAntiGoldCheckRequest) SetAccount(account string) *IRiskAntiGoldCheckRequest {
	r.Account = &account
	return r
}

// SetRoleId 设置用户角色ID
func (r *IRiskAntiGoldCheckRequest) SetRoleId(roleId string) *IRiskAntiGoldCheckRequest {
	r.RoleId = &roleId
	return r
}

// SetNickname 设置用户昵称
func (r *IRiskAntiGoldCheckRequest) SetNickname(nickname string) *IRiskAntiGoldCheckRequest {
	r.Nickname = &nickname
	return r
}

// SetServerId 设置用户所在服务器ID
func (r *IRiskAntiGoldCheckRequest) SetServerId(serverId string) *IRiskAntiGoldCheckRequest {
	r.ServerId = &serverId
	return r
}

// SetLogType 设置日志类型
func (r *IRiskAntiGoldCheckRequest) SetLogType(logType string) *IRiskAntiGoldCheckRequest {
	r.LogType = &logType
	return r
}

// SetLogData 设置日志数据
func (r *IRiskAntiGoldCheckRequest) SetLogData(logData string) *IRiskAntiGoldCheckRequest {
	r.LogData = &logData
	return r
}

// NewIRiskAntiGoldCheckRequest 初始化IRiskAntiGoldCheckRequest对象
func NewIRiskAntiGoldCheckRequest(businessId string) *IRiskAntiGoldCheckRequest {
	request := &IRiskAntiGoldCheckRequest{
		BizPostJsonRequest: types.NewBizPostJsonRequest((businessId)),
	}
	request.SetProductCode("irisk")
	request.SetUriPattern("/v5/risk/antiGoldCheck")
	request.SetVersion("500")
	return request
}

// GetNonSignParams 获取具体业务中特有的不参与签名计算的参数
func (r *IRiskAntiGoldCheckRequest) GetNonSignParams() map[string]interface{} {
	return make(map[string]interface{})
}

// GetBusinessCustomSignParams 获取具体业务中特有的需要参与签名计算的参数
func (r *IRiskAntiGoldCheckRequest) GetBusinessCustomSignParams() map[string]string {
	params := r.BizPostJsonRequest.GetBusinessCustomSignParams()
	if r.LogTime != nil {
		params["logTime"] = *r.LogTime
	}
	if r.Account != nil {
		params["account"] = *r.Account
	}
	if r.RoleId != nil {
		params["roleId"] = *r.RoleId
	}
	if r.Nickname != nil {
		params["nickname"] = *r.Nickname
	}
	if r.ServerId != nil {
		params["serverId"] = *r.ServerId
	}
	if r.LogType != nil {
		params["logType"] = *r.LogType
	}
	if r.LogData != nil {
		params["logData"] = *r.LogData
	}
	return params
}

// antiGoldCheck接口 参数校验方法
func (r *IRiskAntiGoldCheckRequest) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "IRiskAntiGoldCheckRequest"}
	if r.LogTime == nil {
		invalidParams.Add(validation.NewErrParamRequired("logTime"))
	}
	if r.Account == nil {
		invalidParams.Add(validation.NewErrParamRequired("account"))
	}
	if r.RoleId == nil {
		invalidParams.Add(validation.NewErrParamRequired("roleId"))
	}
	if r.Nickname == nil {
		invalidParams.Add(validation.NewErrParamRequired("nickname"))
	}
	if r.ServerId == nil {
		invalidParams.Add(validation.NewErrParamRequired("serverId"))
	}
	if r.LogType == nil {
		invalidParams.Add(validation.NewErrParamRequired("logType"))
	}
	if r.LogData == nil {
		invalidParams.Add(validation.NewErrParamRequired("logData"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

type IRiskAntiGoldCheckResponse struct {
	*types.CommonResponse
}

type IRiskListQueryRequest struct {
	*types.BizPostJsonRequest
	ListGroupCode   *string
	PageNum         *int
	BeginModifyTime *int64
	EndModifyTime   *int64
}

type IRiskListQueryResponse struct {
	*types.CommonResponse
	Data IRiskListQueryPageResponse `json:"data"`
}

type IRiskListQueryPageResponse struct {
	Count int                     `json:"count"`
	Rows  []IRiskListItemResponse `json:"rows"`
}

type IRiskListItemResponse struct {
	ListGroupCode string `json:"listGroupCode"`
	Content       string `json:"content"`
	ExpireTime    int64  `json:"expireTime"`
	Description   string `json:"description"`
	HitCount      int64  `json:"hitCount"`
	Status        int    `json:"status"`
	CreateTime    int64  `json:"createTime"`
	ModifyTime    int64  `json:"modifyTime"`
	CreateBy      string `json:"createBy"`
}

type IRiskListAddRequest struct {
	*types.BizPostJsonRequest
	ListGroupCode *string
	Content       *string
	Description   *string
	ExpireTime    *int64
}

type IRiskListAddResponse struct {
	*types.CommonResponse
}

// NewIRiskListAddRequest 初始化IRiskListAddRequest对象
func NewIRiskListAddRequest(businessId string) *IRiskListAddRequest {
	request := &IRiskListAddRequest{
		BizPostJsonRequest: types.NewBizPostJsonRequest((businessId)),
	}
	request.SetProductCode("irisk")
	request.SetUriPattern("/v5/list/add")
	request.SetVersion("500")
	return request
}

// SetListGroupCode 设置名单库code
func (r *IRiskListAddRequest) SetListGroupCode(listGroupCode string) *IRiskListAddRequest {
	r.ListGroupCode = &listGroupCode
	return r
}

// SetContent 设置名单项
func (r *IRiskListAddRequest) SetContent(content string) *IRiskListAddRequest {
	r.Content = &content
	return r
}

// SetDescription 设置添加原因
func (r *IRiskListAddRequest) SetDescription(description string) *IRiskListAddRequest {
	r.Description = &description
	return r
}

// SetExpireTime 设置释放时间
func (r *IRiskListAddRequest) SetExpireTime(expireTime int64) *IRiskListAddRequest {
	r.ExpireTime = &expireTime
	return r
}

// NewIRiskListQueryRequest 初始化IRiskListQueryRequest对象
func NewIRiskListQueryRequest(businessId string) *IRiskListQueryRequest {
	request := &IRiskListQueryRequest{
		BizPostJsonRequest: types.NewBizPostJsonRequest((businessId)),
	}
	request.SetProductCode("irisk")
	request.SetUriPattern("/v5/list/query")
	request.SetVersion("500")
	return request
}

// SetListGroupCode 设置名单库编号
func (r *IRiskListQueryRequest) SetListGroupCode(listGroupCode string) *IRiskListQueryRequest {
	r.ListGroupCode = &listGroupCode
	return r
}

// SetPageNum 设置分页参数页码
func (r *IRiskListQueryRequest) SetPageNum(pageNum int) *IRiskListQueryRequest {
	r.PageNum = &pageNum
	return r
}

// SetBeginModifyTime 设置（起）修改时间
func (r *IRiskListQueryRequest) SetBeginModifyTime(beginModifyTime int64) *IRiskListQueryRequest {
	r.BeginModifyTime = &beginModifyTime
	return r
}

// SetEndModifyTime 设置（止）修改时间
func (r *IRiskListQueryRequest) SetEndModifyTime(endModifyTime int64) *IRiskListQueryRequest {
	r.EndModifyTime = &endModifyTime
	return r
}

// 名单查询接口 参数校验方法
func (r *IRiskListQueryRequest) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "IRiskListQueryRequest"}
	if r.ListGroupCode == nil {
		invalidParams.Add(validation.NewErrParamRequired("ListGroupCode"))
	}
	if r.PageNum == nil {
		invalidParams.Add(validation.NewErrParamRequired("PageNum"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// 名单添加接口 参数校验方法
func (r *IRiskListAddRequest) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "IRiskListAddRequest"}
	if r.ListGroupCode == nil {
		invalidParams.Add(validation.NewErrParamRequired("ListGroupCode"))
	}
	if r.Content == nil {
		invalidParams.Add(validation.NewErrParamRequired("Content"))
	}
	if r.ExpireTime == nil {
		invalidParams.Add(validation.NewErrParamRequired("ExpireTime"))
	}
	if r.Description == nil {
		invalidParams.Add(validation.NewErrParamRequired("Description"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// GetNonSignParams 获取具体业务中特有的不参与签名计算的参数
func (r *IRiskListQueryRequest) GetNonSignParams() map[string]interface{} {
	return make(map[string]interface{})
}

// GetNonSignParams 获取具体业务中特有的不参与签名计算的参数
func (r *IRiskListAddRequest) GetNonSignParams() map[string]interface{} {
	return make(map[string]interface{})
}

// GetBusinessCustomSignParams 获取具体业务中特有的需要参与签名计算的参数
func (r *IRiskListQueryRequest) GetBusinessCustomSignParams() map[string]string {
	params := r.BizPostJsonRequest.GetBusinessCustomSignParams()
	if r.PageNum != nil {
		params["pageNum"] = strconv.Itoa(*r.PageNum)
	}
	if r.ListGroupCode != nil {
		params["listGroupCode"] = *r.ListGroupCode
	}
	if r.BeginModifyTime != nil {
		params["beginModifyTime"] = strconv.FormatInt(*r.BeginModifyTime, 10)
	}
	if r.EndModifyTime != nil {
		params["endModifyTime"] = strconv.FormatInt(*r.EndModifyTime, 10)
	}
	return params
}

// GetBusinessCustomSignParams 获取具体业务中特有的需要参与签名计算的参数
func (r *IRiskListAddRequest) GetBusinessCustomSignParams() map[string]string {
	params := r.BizPostJsonRequest.GetBusinessCustomSignParams()
	if r.ListGroupCode != nil {
		params["listGroupCode"] = *r.ListGroupCode
	}
	if r.Content != nil {
		params["content"] = *r.Content
	}
	if r.Description != nil {
		params["description"] = *r.Description
	}
	if r.ExpireTime != nil {
		params["expireTime"] = strconv.FormatInt(*r.ExpireTime, 10)
	}
	return params
}

// 图片检测接口 参数校验方法
func (r *IRiskMediaQueryRequest) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "IRiskMediaQueryRequest"}
	if r.TaskId == nil {
		invalidParams.Add(validation.NewErrParamRequired("TaskId"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// GetBusinessCustomSignParams 获取具体业务中特有的需要参与签名计算的参数
func (r *IRiskMediaQueryRequest) GetBusinessCustomSignParams() map[string]string {
	params := r.BizPostJsonRequest.GetBusinessCustomSignParams()
	if r.TaskId != nil {
		params["taskId"] = *r.TaskId
	}
	return params
}


// 图片检测批量查询接口 参数校验方法
func (r *IRiskMediaBatchQueryRequest) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "IRiskMediaBatchQueryRequest"}
	if r.TaskIds == nil {
		invalidParams.Add(validation.NewErrParamRequired("TaskIds"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// GetBusinessCustomSignParams 获取具体业务中特有的需要参与签名计算的参数
func (r *IRiskMediaBatchQueryRequest) GetBusinessCustomSignParams() map[string]string {
	params := r.BizPostJsonRequest.GetBusinessCustomSignParams()
	if r.TaskIds != nil {
		params["taskIds"] = *r.TaskIds
	}
	return params
}

// 安卓设备信息
type IRiskAndroidDeviceInfoResponse struct {

	/**
	 * 屏幕信息
	 */
	ScreenInfo string `json:"screenInfo"`

	/**
	 * app版本号
	 */
	AppVersion string `json:"appVersion"`

	/**
	 * 设备启动次数
	 */
	BootCount int `json:"bootCount"`

	/**
	 * 充电状态
	 */
	ChargeState int `json:"chargeState"`

	/**
	 * 是否开启mac随机化
	 */
	MacRandom int `json:"macRandom"`

	/**
	 * 屏幕亮度
	 */
	ScreenBright int `json:"screenBright"`

	/**
	 * adb是否开启
	 */
	CheckAdb int `json:"checkAdb"`

	/**
	 * 手机已启动时间
	 */
	UpTime string `json:"upTime"`

	/**
	 * Android版本
	 */
	OsVersion string `json:"osVersion"`

	/**
	 * 运营商
	 */
	CarrierName string `json:"carrierName"`

	/**
	 * 手机总锁屏时间
	 */
	LockTime string `json:"lockTime"`

	/**
	 * 是否是虚拟摄像头
	 */
	IsFakeCam int `json:"isFakeCam"`

	/**
	 * 手机品牌
	 */
	PhoneBrand string `json:"phoneBrand"`

	/**
	 * 第一次安装时间
	 */
	FirstSetupTime string `json:"firstSetupTime"`

	/**
	 * 是否存在sim卡
	 */
	HasSim int `json:"hasSim"`

	/**
	 * 型号
	 */
	Model string `json:"model"`

	/**
	 * 投屏状态判断
	 */
	IsShareScreen int `json:"isShareScreen"`

	/**
	 * 网络连接类型
	 */
	NetworkType int `json:"networkType"`

	/**
	 * 当前电量信息
	 */
	BatteryLevel int `json:"batteryLevel"`

	/**
	 * 设备广告ID
	 */
	Oaid string `json:"oaid"`

	/**
	 * 设备类型
	 */
	DeviceType int `json:"deviceType"`

	/**
	 * 设备cpu信息
	 */
	CpuInfo string `json:"cpuInfo"`

	/**
	 * 设备sd卡总内存大小
	 */
	SdTotal string `json:"sdTotal"`

	/**
	 * 设备总内存大小
	 */
	MemTotal string `json:"memTotal"`

	/**
	 * 设备语言设置
	 */
	Language string `json:"language"`

	/**
	 * 设备当前地区设置
	 */
	UserRegion string `json:"userRegion"`

	/**
	 * 设备厂商信息
	 */
	ProductManufacturer string `json:"productManufacturer"`

	/**
	 * 设备系统名称
	 */
	ProductOs string `json:"productOs"`

	/**
	 * 应用广告ID
	 */
	Aaid string `json:"aaid"`

	/**
	 * 设备ROM信息
	 */
	Rom string `json:"rom"`

	/**
	 * 运营商设备标识
	 */
	Uaid string `json:"uaid"`
}

// IOS设备信息
type IRiskIosDeviceInfoResponse struct {
	/**
	 * 开机时间
	 */
	BootTime string `json:"bootTime"`

	/**
	 * 硬盘大小，单位字节
	 */
	TotalDiskSpace string `json:"totalDiskSpace"`

	/**
	 * 非完美越狱失效状态
	 */
	Auxo int `json:"auxo"`

	/**
	 * mac序列号
	 */
	MacSerialnumber string `json:"macSerialnumber"`

	/**
	 * 0:未开启电池权限
	 */
	ChargeStatus string `json:"chargeStatus"`

	/**
	 * 客户端本地随机生成的设备ID
	 */
	DeviceId string `json:"deviceId"`

	/**
	 * 是否安装CydiaSubstrate
	 */
	IsCydiaSubstrate int `json:"isCydiaSubstrate"`

	/**
	 * 是否有反越狱插件
	 */
	IsAntiJailbreak int `json:"isAntiJailbreak"`

	/**
	 * 总内存，单位字节
	 */
	TotalMemory string `json:"totalMemory"`

	/**
	 * 手机操作系统版本
	 */
	OsVersion string `json:"osVersion"`

	/**
	 * 运营商名称
	 */
	CarrierName string `json:"carrierName"`

	/**
	 * 手机型号
	 */
	Model string `json:"model"`

	/**
	 * 当前电量信息
	 */
	BatteryLevel string `json:"batteryLevel"`

	/**
	 * 设备名称
	 */
	DeviceName string `json:"deviceName"`

	/**
	 * iOS设备标识符
	 */
	Idfv string `json:"idfv"`

	/**
	 * 运营商设备标识
	 */
	Uaid string `json:"uaid"`

	/**
	 * cpu 使用率
	 */
	CpuUseRate string `json:"cpuUseRate"`

	/**
	 * 网络连接类型
	 */
	NetworkType int `json:"networkType"`

	/**
	 * 动态加载库
	 */
	LoadedDyLibsInfo []string `json:"loadedDyLibsInfo"`

	/**
	 * 设备的屏幕宽度
	 */
	ScreenWidth string `json:"screenWidth"`

	/**
	 * 设备的屏幕高度
	 */
	ScreenHeight string `json:"screenHeight"`
}

// PC设备信息
type IRiskPcDeviceInfoResponse struct {
	/**
	 * 设备物理地址
	 */
	Mac string `json:"mac"`

	/**
	 * 设备标识符
	 */
	Guid string `json:"guid"`

	/**
	 * 设备系统版本
	 */
	OsVersion string `json:"osVersion"`

	/**
	 * 屏幕分辨率
	 */
	ScreenResolution string `json:"screenResolution"`

	/**
	 * 操作系统型号信息
	 */
	OsType string `json:"osType"`

	/**
	 * 计算机名
	 */
	ComputerName string `json:"computerName"`

	/**
	 * 设备制造商
	 */
	Manufacturer string `json:"manufacturer"`

	/**
	 * 操作系统名称
	 */
	OsName string `json:"osName"`
}

// web设备信息
type IRiskWebDeviceInfoResponse struct {

	/**
	 * web类型
	 */
	RunEnvType int `json:"runEnvType"`

	/**
	 * 浏览器标识信息
	 */
	UserAgent string `json:"userAgent"`

	/**
	 * 浏览器的语言设置
	 */
	Language string `json:"language"`

	/**
	 * 浏览器所在设备的系统类型
	 */
	Platform string `json:"platform"`

	/**
	 * 地址栏当前url
	 */
	Url string `json:"url"`

	/**
	 * 屏幕尺寸
	 */
	ScreenResolution string `json:"screenResolution"`

	/**
	 * 运营商设备标识
	 */
	Uaid string `json:"uaid"`
}

// 小程序设备信息
type IRiskAppletDeviceInfoResponse struct {
	/**
	 * 当前小程序所在设备的联网类型
	 */
	NetworkType string `json:"networkType"`

	/**
	 * 运行小程序的设备的品牌
	 */
	Brand string `json:"brand"`

	/**
	 * 运行小程序的系统的平台类型
	 */
	Platform string `json:"platform"`

	/**
	 * 小程序版本号
	 */
	Version string `json:"version"`

	/**
	 * 小程序类型
	 */
	Type int `json:"type"`

	/**
	 * 运营商设备标识
	 */
	Uaid string `json:"uaid"`
}

// Harmony设备信息
type IRiskHarmonyDeviceInfoResponse struct {
	/**
	 * 设备广告ID
	 */
	Oaid string `json:"oaid"`

	/**
	 * 应用广告ID
	 */
	Aaid string `json:"aaid"`

	/**
	 * 屏幕信息
	 */
	SceneInfo string `json:"sceneInfo"`

	/**
	 * 设备cpu信息
	 */
	CpuInfo string `json:"cpuInfo"`

	/**
	 * 内部存储空间大小
	 */
	MemTotal string `json:"memTotal"`

	/**
	 * 内部可用存储空间大小
	 */
	MemFree string `json:"memFree"`

	/**
	 * 设备语言设置
	 */
	Language string `json:"language"`

	/**
	 * 硬件设备类型，是平板、电视还是手机
	 */
	DeviceType string `json:"deviceType"`

	/**
	 * 设备系统版本
	 */
	OsVersion string `json:"osVersion"`

	/**
	 * 运营商设备标识
	 */
	Uaid string `json:"uaid"`
}
