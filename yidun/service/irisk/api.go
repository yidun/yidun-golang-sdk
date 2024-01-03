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

// GetMediaCheckResult 数据查询，调用/risk/getCheckResult
func (c *IRiskClient) GetMediaCheckResult(request *IRiskMediaCheckRequest) (response *IRiskMediaCheckResponse, err error) {
	response = &IRiskMediaCheckResponse{}
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
	return params
}

// check接口 参数校验方法
func (r *IRiskCheckRequest) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "IRiskCheckRequest"}
	if r.Ip == nil {
		invalidParams.Add(validation.NewErrParamRequired("Ip"))
	}
	if r.RoleId == nil {
		invalidParams.Add(validation.NewErrParamRequired("RoleId"))
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
}

type HitInfo struct {
	Type      string `json:"type"`
	Name      string `json:"name"`
	RiskLevel int    `json:"riskLevel"`
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
	RoleId         *string
	RiskLevel      *int
	Ip             *string
	SdkIp          *string
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

// SetRoleId 设置角色ID
func (r *IRiskDetailRequest) SetRoleId(roleId string) *IRiskDetailRequest {
	r.RoleId = &roleId
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

// SetSdkIp 设置SDK IP地址
func (r *IRiskDetailRequest) SetSdkIp(sdkIp string) *IRiskDetailRequest {
	r.SdkIp = &sdkIp
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
	if r.RoleId != nil {
		params["roleId"] = *r.RoleId
	}
	if r.RiskLevel != nil {
		params["riskLevel"] = strconv.Itoa(*r.RiskLevel)
	}
	if r.Ip != nil {
		params["ip"] = *r.Ip
	}
	if r.SdkIp != nil {
		params["sdkIp"] = *r.SdkIp
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
