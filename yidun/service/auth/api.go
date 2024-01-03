package auth

import (
	"encoding/json"

	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/core/validation"
)

type PicType string
type VideoType string

const (
	//PicType
	URL    = PicType("1")
	BASE64 = PicType("2")

	//VideoType
	VIDEO_URL    = VideoType("1")
	VIDEO_BASE64 = VideoType("2")
)

// 银行卡校验接口，调用/v1/bankcard/check
func (c *AuthClient) BankCardCheck(request *BankCardCheckRequest) (response *BankCardResponse, err error) {
	response = &BankCardResponse{}
	err = c.Client.DoExecute(request, response)
	return
}

// 银行卡校验接口，调用/v1/bankcard/check
type BankCardCheckRequest struct {
	types.BizPostFormRequest
	BankCardNo *string
	Name       *string
	IdCardNo   *string
	Phone      *string
	DataId     *string
}

// 校验参数
func (r *BankCardCheckRequest) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "BankCardCheckRequest"}
	if r.BankCardNo == nil {
		invalidParams.Add(validation.NewErrParamRequired("BankCardNo"))
	}
	//name校验
	if validation.CheckName(r.Name) == false {
		invalidParams.Add(validation.NewErrParamInvalid("Name"))
	}
	//idCardNo校验
	if validation.CheckIdCard(r.IdCardNo) == false {
		invalidParams.Add(validation.NewErrParamInvalid("IdCardNo"))
	}
	//dataId可以为空，但是不能超过64个字符
	if validation.EmptyMaxLen(r.DataId, 64) == false {
		invalidParams.Add(validation.NewErrParamMaxLen("DataId", 64, "DataId长度不合法"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// 提供BankCardNo参数的SET方法
func (r *BankCardCheckRequest) SetBankCardNo(bankCardNo string) *BankCardCheckRequest {
	r.BankCardNo = &bankCardNo
	return r
}

// 提供Name参数的SET方法
func (r *BankCardCheckRequest) SetName(name string) *BankCardCheckRequest {
	r.Name = &name
	return r
}

// 提供IdCardNo参数的SET方法
func (r *BankCardCheckRequest) SetIdCardNo(idCardNo string) *BankCardCheckRequest {
	r.IdCardNo = &idCardNo
	return r
}

// 提供Phone参数的SET方法
func (r *BankCardCheckRequest) SetPhone(phone string) *BankCardCheckRequest {
	r.Phone = &phone
	return r
}

// 提供DataId参数的SET方法
func (r *BankCardCheckRequest) SetDataId(dataId string) *BankCardCheckRequest {
	r.DataId = &dataId
	return r
}

// 提供new方法
func NewBankCardCheckRequest(businessId string) *BankCardCheckRequest {
	request := &BankCardCheckRequest{
		BizPostFormRequest: *types.NewBizPostFormRequest(businessId),
	}
	request.SetProductCode("auth")
	request.SetUriPattern("/v1/bankcard/check")
	request.SetVersion("v1")
	return request
}

// GetNonSignParams 获取具体业务中特有的不参与签名计算的参数
func (r *BankCardCheckRequest) GetNonSignParams() map[string]string {
	return make(map[string]string)
}

// GetCustomSignParams 获取具体业务中特有的需要参与签名计算的参数
func (r *BankCardCheckRequest) GetBusinessCustomSignParams() map[string]string {
	params := r.BizPostFormRequest.GetBusinessCustomSignParams()
	if r.BankCardNo != nil {
		params["bankCardNo"] = *r.BankCardNo
	}
	if r.Name != nil {
		params["name"] = *r.Name
	}
	if r.IdCardNo != nil {
		params["idCardNo"] = *r.IdCardNo
	}
	if r.Phone != nil {
		params["phone"] = *r.Phone
	}
	if r.DataId != nil {
		params["dataId"] = *r.DataId
	}
	return params
}

// 提供response
type BankCardResponse struct {
	*types.CommonResponse
	Result BankCardResult `json:"result"`
}

// 提供result
type BankCardResult struct {
	// 状态
	Status *int `json:"status"`
	//taskId
	TaskId *string `json:"taskId"`
	//reasonType
	ReasonType *int `json:"reasonType"`
	//isPayed
	IsPayed *int `json:"isPayed"`
}

// 出入境校验，调用/v1/foreign/check
func (c *AuthClient) EntryExitPermitCheck(request *EntryExitPermitCheckRequest) (response *EntryExitPermitCheckResponse, err error) {
	response = &EntryExitPermitCheckResponse{}
	err = c.Client.DoExecute(request, response)
	return
}

// 出入境校验
type EntryExitPermitCheckRequest struct {
	types.BizPostFormRequest
	//name
	Name *string
	//cardNo
	CardNo *string
	//verifyType
	VerifyType *string
	//nation
	Nation *string
	//dataId
	DataId *string
}

// 校验参数
func (r *EntryExitPermitCheckRequest) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "EntryExitPermitCheckRequest"}
	//name校验
	if validation.CheckName(r.Name) == false {
		invalidParams.Add(validation.NewErrParamInvalid("Name"))
	}
	//cardNo校验
	if validation.NotEmptyMaxLen(r.CardNo, 18) == false {
		invalidParams.Add(validation.NewErrParamMaxLen("CardNo", 18, "CardNo长度不合法"))
	}
	//verifyType校验
	if validation.NotEmptyMaxLen(r.VerifyType, 11) == false {
		invalidParams.Add(validation.NewErrParamMaxLen("VerifyType", 11, "VerifyType长度不合法"))
	}
	//dataId可以为空，但是不能超过64个字符
	if validation.EmptyMaxLen(r.DataId, 64) == false {
		invalidParams.Add(validation.NewErrParamMaxLen("DataId", 64, "DataId长度不合法"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// 提供Name参数的SET方法
func (r *EntryExitPermitCheckRequest) SetName(name string) *EntryExitPermitCheckRequest {
	r.Name = &name
	return r
}

// 提供CardNo参数的SET方法
func (r *EntryExitPermitCheckRequest) SetCardNo(cardNo string) *EntryExitPermitCheckRequest {
	r.CardNo = &cardNo
	return r
}

// 提供VerifyType参数的SET方法
func (r *EntryExitPermitCheckRequest) SetVerifyType(verifyType string) *EntryExitPermitCheckRequest {
	r.VerifyType = &verifyType
	return r
}

// 提供Nation参数的SET方法
func (r *EntryExitPermitCheckRequest) SetNation(nation string) *EntryExitPermitCheckRequest {
	r.Nation = &nation
	return r
}

// 提供DataId参数的SET方法
func (r *EntryExitPermitCheckRequest) SetDataId(dataId string) *EntryExitPermitCheckRequest {
	r.DataId = &dataId
	return r
}

// 提供new方法
func NewEntryExitPermitCheckRequest(businessId string) *EntryExitPermitCheckRequest {
	request := &EntryExitPermitCheckRequest{
		BizPostFormRequest: *types.NewBizPostFormRequest(businessId),
	}
	request.SetProductCode("auth")
	request.SetUriPattern("/v1/foreign/check")
	request.SetVersion("v1")
	return request
}

// GetNonSignParams 获取具体业务中特有的不参与签名计算的参数
func (r *EntryExitPermitCheckRequest) GetNonSignParams() map[string]string {
	return make(map[string]string)
}

// GetCustomSignParams 获取具体业务中特有的需要参与签名计算的参数
func (r *EntryExitPermitCheckRequest) GetBusinessCustomSignParams() map[string]string {
	params := r.BizPostFormRequest.GetBusinessCustomSignParams()
	if r.Name != nil {
		params["name"] = *r.Name
	}
	if r.CardNo != nil {
		params["cardNo"] = *r.CardNo
	}
	if r.VerifyType != nil {
		params["verifyType"] = *r.VerifyType
	}
	if r.Nation != nil {
		params["nation"] = *r.Nation
	}
	if r.DataId != nil {
		params["dataId"] = *r.DataId
	}
	return params
}

// 提供response
type EntryExitPermitCheckResponse struct {
	*types.CommonResponse
	Result EntryExitPermitResult `json:"result"`
}

// 提供result
type EntryExitPermitResult struct {
	// 状态
	Status *int `json:"status"`
	//taskId
	TaskId *string `json:"taskId"`
	//reasonType
	ReasonType *int `json:"reasonType"`
	//isPayed
	IsPayed *int `json:"isPayed"`
}

// 实证认证/v1/iacard/check
func (c *AuthClient) IdCardCheck(request *IdCardCheckRequest) (response *IdCardCheckResponse, err error) {
	response = &IdCardCheckResponse{}
	err = c.Client.DoExecute(request, response)
	return
}

// 身份证校验
type IdCardCheckRequest struct {
	types.BizPostFormRequest
	//name
	Name *string
	//cardNo
	CardNo *string
	//dataId
	DataId *string
}

// 校验参数
func (r *IdCardCheckRequest) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "IdCardCheckRequest"}
	//name校验
	if validation.CheckName(r.Name) == false {
		invalidParams.Add(validation.NewErrParamInvalid("Name"))
	}
	//cardNo校验
	if validation.CheckIdCard(r.CardNo) == false {
		invalidParams.Add(validation.NewErrParamInvalid("CardNo"))
	}
	//dataId可以为空，但是不能超过64个字符
	if validation.EmptyMaxLen(r.DataId, 64) == false {
		invalidParams.Add(validation.NewErrParamMaxLen("DataId", 64, "DataId长度不合法"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// 提供Name参数的SET方法
func (r *IdCardCheckRequest) SetName(name string) *IdCardCheckRequest {
	r.Name = &name
	return r
}

// 提供CardNo参数的SET方法
func (r *IdCardCheckRequest) SetCardNo(cardNo string) *IdCardCheckRequest {
	r.CardNo = &cardNo
	return r
}

// 提供DataId参数的SET方法
func (r *IdCardCheckRequest) SetDataId(dataId string) *IdCardCheckRequest {
	r.DataId = &dataId
	return r
}

// 提供new方法
func NewIdCardCheckRequest(businessId string) *IdCardCheckRequest {
	request := &IdCardCheckRequest{
		BizPostFormRequest: *types.NewBizPostFormRequest(businessId),
	}
	request.SetProductCode("auth")
	request.SetUriPattern("/v1/idcard/check")
	request.SetVersion("v1")
	return request
}

// GetNonSignParams 获取具体业务中特有的不参与签名计算的参数
func (r *IdCardCheckRequest) GetNonSignParams() map[string]string {
	return make(map[string]string)
}

// GetCustomSignParams 获取具体业务中特有的需要参与签名计算的参数
func (r *IdCardCheckRequest) GetBusinessCustomSignParams() map[string]string {
	params := r.BizPostFormRequest.GetBusinessCustomSignParams()
	if r.Name != nil {
		params["name"] = *r.Name
	}
	if r.CardNo != nil {
		params["cardNo"] = *r.CardNo
	}
	if r.DataId != nil {
		params["dataId"] = *r.DataId
	}
	return params
}

// 提供response
type IdCardCheckResponse struct {
	*types.CommonResponse
	Result IdCardCheckResult `json:"result"`
}

// 提供result
type IdCardCheckResult struct {
	// 状态
	Status *int `json:"status"`
	//taskId
	TaskId *string `json:"taskId"`
	/**
	 * 原因详情
	 * 2-输入姓名和身份证号不一致
	 * 3-查无此身份证
	 * 4-身份证照片信息与输入信息不一致
	 * 7-结果获取失败，请重试
	 */
	ReasonType *int `json:"reasonType"`
	//isPayed
	IsPayed *int `json:"isPayed"`
}

// 实证认证/v1/idlephone/check
func (c *AuthClient) IdlePhoneCheck(request *IdlePhoneCheckRequest) (response *IdlePhoneCheckResponse, err error) {
	response = &IdlePhoneCheckResponse{}
	err = c.Client.DoExecute(request, response)
	return
}

// 号码状态检测
type IdlePhoneCheckRequest struct {
	types.BizPostFormRequest
	//phoneList
	PhoneList *string
	//encryptType
	EncryptType *string
}

// 校验参数
func (r *IdlePhoneCheckRequest) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "IdlePhoneCheckRequest"}
	if validation.NotEmptyMaxLen(r.PhoneList, 50) == false {
		invalidParams.Add(validation.NewErrParamMaxLen("PhoneList", 50, "phoneList长度不合法"))
	}
	if validation.NotEmptyMaxLen(r.EncryptType, 4) == false {
		invalidParams.Add(validation.NewErrParamMaxLen("EncryptType", 4, "encryptType长度不合法"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// 提供PhoneList参数的SET方法
func (r *IdlePhoneCheckRequest) SetPhoneList(phoneList string) *IdlePhoneCheckRequest {
	r.PhoneList = &phoneList
	return r
}

// 提供EncryptType参数的SET方法
func (r *IdlePhoneCheckRequest) SetEncryptType(encryptType string) *IdlePhoneCheckRequest {
	r.EncryptType = &encryptType
	return r
}

// 提供new方法
func NewIdlePhoneCheckRequest(businessId string) *IdlePhoneCheckRequest {
	request := &IdlePhoneCheckRequest{
		BizPostFormRequest: *types.NewBizPostFormRequest(businessId),
	}
	request.SetProductCode("auth")
	request.SetUriPattern("/v1/idlephone/check")
	request.SetVersion("v1")
	return request
}

// GetNonSignParams 获取具体业务中特有的不参与签名计算的参数
func (r *IdlePhoneCheckRequest) GetNonSignParams() map[string]string {
	return make(map[string]string)
}

// GetCustomSignParams 获取具体业务中特有的需要参与签名计算的参数
func (r *IdlePhoneCheckRequest) GetBusinessCustomSignParams() map[string]string {
	params := r.BizPostFormRequest.GetBusinessCustomSignParams()
	if r.PhoneList != nil {
		params["phoneList"] = *r.PhoneList
	}
	if r.EncryptType != nil {
		params["encryptType"] = *r.EncryptType
	}
	return params
}

// 提供response
type IdlePhoneCheckResponse struct {
	*types.CommonResponse
	Result IdlePhoneCheckResult `json:"result"`
}

// result
type IdlePhoneCheckResult struct {
	// 状态
	Status *int `json:"status"`
	//isPayed
	IsPayed *int `json:"isPayed"`
	//requestId
	RequestId *string `json:"requestId"`
	//chargeCount
	ChargeCount *int `json:"chargeCount"`
	//PhoneInfo
	PhoneInfo []PhoneInfo `json:"phoneInfo"`
}

// PhoneInfo
type PhoneInfo struct {
	//mobile
	Mobile *string `json:"mobile"`
	//area-手机号所属区域
	Area *string `json:"area"`
	//numberType-手机号运营商类型。
	NumberType *string `json:"numberType"`
	//status
	Status *string `json:"status"`
}

// 交互式活体检测InteractiveLivePersonCheck  调用/v1/liveperson/recheck
func (c *AuthClient) InteractiveLivePersonCheck(request *InteractiveLivePersonCheckRequest) (response *InteractiveLivePersonCheckResponse, err error) {
	response = &InteractiveLivePersonCheckResponse{}
	err = c.Client.DoExecute(request, response)
	return
}

// checkRequest
type InteractiveLivePersonCheckRequest struct {
	types.BizPostFormRequest
	//token
	Token *string
	//needAvatar
	NeedAvatar *string
	//picType
	PicType *string
	//dataId
	DataId *string
}

// param check
func (r *InteractiveLivePersonCheckRequest) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "InteractiveLivePersonCheckRequest"}
	// token校验
	if r.Token == nil {
		invalidParams.Add(validation.NewErrParamRequired("Token"))
	}
	// dataId可以为空，但是不能超过64个字符
	if validation.EmptyMaxLen(r.DataId, 64) == false {
		invalidParams.Add(validation.NewErrParamMaxLen("DataId", 64, "DataId长度不合法"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// 提供Token参数的SET方法
func (r *InteractiveLivePersonCheckRequest) SetToken(token string) *InteractiveLivePersonCheckRequest {
	r.Token = &token
	return r
}

// 提供NeedAvatar参数的SET方法
func (r *InteractiveLivePersonCheckRequest) SetNeedAvatar(needAvatar string) *InteractiveLivePersonCheckRequest {
	r.NeedAvatar = &needAvatar
	return r
}

// 提供PicType参数的SET方法
func (r *InteractiveLivePersonCheckRequest) SetPicType(picType string) *InteractiveLivePersonCheckRequest {
	r.PicType = &picType
	return r
}

// 提供DataId参数的SET方法
func (r *InteractiveLivePersonCheckRequest) SetDataId(dataId string) *InteractiveLivePersonCheckRequest {
	r.DataId = &dataId
	return r
}

// new method
func NewInteractiveLivePersonCheckRequest(businessId string) *InteractiveLivePersonCheckRequest {
	request := &InteractiveLivePersonCheckRequest{
		BizPostFormRequest: *types.NewBizPostFormRequest(businessId),
	}
	request.SetProductCode("auth")
	request.SetUriPattern("/v1/liveperson/recheck")
	request.SetVersion("v1")
	return request
}

// GetNonSignParams 获取具体业务中特有的不参与签名计算的参数
func (r *InteractiveLivePersonCheckRequest) GetNonSignParams() map[string]string {
	return make(map[string]string)
}

// GetCustomSignParams 获取具体业务中特有的需要参与签名计算的参数
func (r *InteractiveLivePersonCheckRequest) GetBusinessCustomSignParams() map[string]string {
	params := r.BizPostFormRequest.GetBusinessCustomSignParams()
	if r.Token != nil {
		params["token"] = *r.Token
	}
	if r.NeedAvatar != nil {
		params["needAvatar"] = *r.NeedAvatar
	}
	if r.PicType != nil {
		params["picType"] = *r.PicType
	}
	if r.DataId != nil {
		params["dataId"] = *r.DataId
	}
	return params
}

// response
type InteractiveLivePersonCheckResponse struct {
	*types.CommonResponse
	Result InteractiveLivePersonCheckResult `json:"result"`
}

// result
type InteractiveLivePersonCheckResult struct {
	//isPayed
	IsPayed *int `json:"isPayed"`
	//taskId
	TaskId *string `json:"taskId"`
	//picType
	PicType *int `json:"picType"`
	//avatar
	Avatar *string `json:"avatar"`
	//lpCheckStatus
	LpCheckStatus *int `json:"lpCheckStatus"`
	//reasonType
	ReasonType *int `json:"reasonType"`
	//actionImages
	ActionImages []ActionImage `json:"actionImages"`
	//isPayed
}

// actionImage
type ActionImage struct {
	//action
	Action *string `json:"action"`
	//image
	Image *string `json:"image"`
}

// 交互式活体人脸核身检测InteractiveLivePersonIdCardCheck 调用/v1/liveperson/audit
func (c *AuthClient) InteractiveLivePersonIdCardCheck(request *InteractiveLivePersonIdCardCheckRequest) (response *InteractiveLivePersonIdCardCheckResponse, err error) {
	response = &InteractiveLivePersonIdCardCheckResponse{}
	err = c.Client.DoExecute(request, response)
	return
}

// checkRequest
type InteractiveLivePersonIdCardCheckRequest struct {
	types.BizPostFormRequest
	//token
	Token *string
	// name
	Name *string
	//cardNo
	CardNo *string
	//needAvatar
	NeedAvatar *string
	//picType
	PicType *string
	//dataId
	DataId *string
}

// param check
func (r *InteractiveLivePersonIdCardCheckRequest) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "InteractiveLivePersonIdCardCheckRequest"}
	// token校验
	if r.Token == nil {
		invalidParams.Add(validation.NewErrParamRequired("Token"))
	}
	// dataId可以为空，但是不能超过64个字符
	if validation.EmptyMaxLen(r.DataId, 64) == false {
		invalidParams.Add(validation.NewErrParamMaxLen("DataId", 64, "DataId长度不合法"))
	}
	//name校验
	if validation.CheckName(r.Name) == false {
		invalidParams.Add(validation.NewErrParamInvalid("Name"))
	}
	//cardNo校验
	if validation.CheckIdCard(r.CardNo) == false {
		invalidParams.Add(validation.NewErrParamInvalid("CardNo"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// 提供Token参数的SET方法
func (r *InteractiveLivePersonIdCardCheckRequest) SetToken(token string) *InteractiveLivePersonIdCardCheckRequest {
	r.Token = &token
	return r
}

// 提供Name参数的SET方法
func (r *InteractiveLivePersonIdCardCheckRequest) SetName(name string) *InteractiveLivePersonIdCardCheckRequest {
	r.Name = &name
	return r
}

// 提供CardNo参数的SET方法
func (r *InteractiveLivePersonIdCardCheckRequest) SetCardNo(cardNo string) *InteractiveLivePersonIdCardCheckRequest {
	r.CardNo = &cardNo
	return r
}

// 提供NeedAvatar参数的SET方法
func (r *InteractiveLivePersonIdCardCheckRequest) SetNeedAvatar(needAvatar string) *InteractiveLivePersonIdCardCheckRequest {
	r.NeedAvatar = &needAvatar
	return r
}

// 提供PicType参数的SET方法
func (r *InteractiveLivePersonIdCardCheckRequest) SetPicType(picType string) *InteractiveLivePersonIdCardCheckRequest {
	r.PicType = &picType
	return r
}

// 提供DataId参数的SET方法
func (r *InteractiveLivePersonIdCardCheckRequest) SetDataId(dataId string) *InteractiveLivePersonIdCardCheckRequest {
	r.DataId = &dataId
	return r
}

// new method
func NewInteractiveLivePersonIdCardCheckRequest(businessId string) *InteractiveLivePersonIdCardCheckRequest {
	request := &InteractiveLivePersonIdCardCheckRequest{
		BizPostFormRequest: *types.NewBizPostFormRequest(businessId),
	}
	request.SetProductCode("auth")
	request.SetUriPattern("/v1/liveperson/audit")
	request.SetVersion("v1")
	return request
}

// GetNonSignParams 获取具体业务中特有的不参与签名计算的参数
func (r *InteractiveLivePersonIdCardCheckRequest) GetNonSignParams() map[string]string {
	return make(map[string]string)
}

// GetCustomSignParams 获取具体业务中特有的需要参与签名计算的参数
func (r *InteractiveLivePersonIdCardCheckRequest) GetBusinessCustomSignParams() map[string]string {
	params := r.BizPostFormRequest.GetBusinessCustomSignParams()
	if r.Token != nil {
		params["token"] = *r.Token
	}
	if r.Name != nil {
		params["name"] = *r.Name
	}
	if r.CardNo != nil {
		params["cardNo"] = *r.CardNo
	}
	if r.NeedAvatar != nil {
		params["needAvatar"] = *r.NeedAvatar
	}
	if r.PicType != nil {
		params["picType"] = *r.PicType
	}
	if r.DataId != nil {
		params["dataId"] = *r.DataId
	}
	return params
}

// response
type InteractiveLivePersonIdCardCheckResponse struct {
	*types.CommonResponse
	Result InteractiveLivePersonIdCardCheckResult `json:"result"`
}

// result
type InteractiveLivePersonIdCardCheckResult struct {
	//taskId
	TaskId *string `json:"taskId"`
	//picType
	PicType *int `json:"picType"`
	//avatar
	Avatar *string `json:"avatar"`
	//status
	Status *int `json:"status"`
	//faceMatched 公安人像比对是否通过，0-待定 1-通过 2-不通过
	FaceMatched *int `json:"faceMatched"`
	//similarityScore
	SimilarityScore *float64 `json:"similarityScore"`
	//reasonType
	ReasonType *int `json:"reasonType"`
	//actionImages
	ActionImages *[]ActionImage `json:"actionImages"`
	//isPayed
	IsPayed *int `json:"isPayed"`
}

// 视频活体检测请求 VideoLivePersonCheck /v1/liveperson/h5/check
func (c *AuthClient) VideoLivePersonCheck(request *VideoLivePersonCheckRequest) (response *VideoLivePersonCheckResponse, err error) {
	response = &VideoLivePersonCheckResponse{}
	err = c.Client.DoExecute(request, response)
	return
}

// 视频活体检测请求
type VideoLivePersonCheckRequest struct {
	types.BizPostFormRequest
	// 数组 actions
	Actions *[]string
	//数组 actionVideos
	ActionVideos *[]string
	//videoType
	VideoType *string
	//是否需要返回正面照
	NeedAvatar *string
	//dataId
	DataId *string
}

// param check
func (r *VideoLivePersonCheckRequest) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "VideoLivePersonCheckRequest"}
	// token校验
	if r.Actions == nil || len(*r.Actions) == 0 {
		invalidParams.Add(validation.NewErrParamRequired("Actions"))
	}
	if r.ActionVideos == nil || len(*r.ActionVideos) == 0 {
		invalidParams.Add(validation.NewErrParamRequired("ActionVideos"))
	}
	if r.VideoType == nil {
		invalidParams.Add(validation.NewErrParamRequired("VideoType"))
	}
	if validation.EmptyMaxLen(r.DataId, 64) == false {
		invalidParams.Add(validation.NewErrParamMaxLen("DataId", 64, "DataId长度不合法"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// 提供Actions参数的SET方法
func (r *VideoLivePersonCheckRequest) SetActions(actions []string) *VideoLivePersonCheckRequest {
	r.Actions = &actions
	return r
}

// 提供ActionVideos参数的SET方法
func (r *VideoLivePersonCheckRequest) SetActionVideos(actionVideos []string) *VideoLivePersonCheckRequest {
	r.ActionVideos = &actionVideos
	return r
}

// 提供VideoType参数的SET方法
func (r *VideoLivePersonCheckRequest) SetVideoType(videoType string) *VideoLivePersonCheckRequest {
	r.VideoType = &videoType
	return r
}

// 提供NeedAvatar参数的SET方法
func (r *VideoLivePersonCheckRequest) SetNeedAvatar(needAvatar string) *VideoLivePersonCheckRequest {
	r.NeedAvatar = &needAvatar
	return r
}

// 提供DataId参数的SET方法
func (r *VideoLivePersonCheckRequest) SetDataId(dataId string) *VideoLivePersonCheckRequest {
	r.DataId = &dataId
	return r
}

// new method
func NewVideoLivePersonCheckRequest(businessId string) *VideoLivePersonCheckRequest {
	request := &VideoLivePersonCheckRequest{
		BizPostFormRequest: *types.NewBizPostFormRequest(businessId),
	}
	request.SetProductCode("auth")
	request.SetUriPattern("/v1/liveperson/h5/check")
	request.SetVersion("v1")
	return request
}

// GetNonSignParams 获取具体业务中特有的不参与签名计算的参数
func (r *VideoLivePersonCheckRequest) GetNonSignParams() map[string]string {
	actionVideosBites, _ := json.Marshal(*r.ActionVideos)
	//actionVideos不参与签名
	return map[string]string{
		"actionVideos": string(actionVideosBites)}
}

// GetCustomSignParams 获取具体业务中特有的需要参与签名计算的参数
func (r *VideoLivePersonCheckRequest) GetBusinessCustomSignParams() map[string]string {
	params := r.BizPostFormRequest.GetBusinessCustomSignParams()
	if r.Actions != nil {
		actionBytes, _ := json.Marshal(*r.Actions)
		params["actions"] = string(actionBytes)
	}
	if r.VideoType != nil {
		params["videoType"] = *r.VideoType
	}
	if r.NeedAvatar != nil {
		params["needAvatar"] = *r.NeedAvatar
	}
	if r.DataId != nil {
		params["dataId"] = *r.DataId
	}
	return params
}

// response
type VideoLivePersonCheckResponse struct {
	*types.CommonResponse
	Result VideoLivePersonCheckResult `json:"result"`
}

// result
type VideoLivePersonCheckResult struct {
	//taskId
	TaskId *string `json:"taskId"`
	//picType 图片类型：1-URL，2-BASE64
	PicType *int `json:"picType"`
	//avatar
	Avatar *string `json:"avatar"`
	//lpCheckStatus
	LpCheckStatus *int `json:"lpCheckStatus"`
	//reasonType
	ReasonType *int `json:"reasonType"`
}

// 视频活体人脸核身请求 VideoLivePersonIdCardCheck /v1/liveperson/h5/audit
func (c *AuthClient) VideoLivePersonIdCardCheck(request *VideoLivePersonIdCardCheckRequest) (response *VideoLivePersonIdCardCheckResponse, err error) {
	response = &VideoLivePersonIdCardCheckResponse{}
	err = c.Client.DoExecute(request, response)
	return
}

// check request
type VideoLivePersonIdCardCheckRequest struct {
	types.BizPostFormRequest
	// name
	Name *string
	//cardNo
	CardNo *string
	//数组 actions
	Actions *[]string
	//数组 actionVideos
	ActionVideos *[]string
	//videoType
	VideoType *string
	//needAvatar
	NeedAvatar *string
	//dataId
	DataId *string
}

// param check
func (r *VideoLivePersonIdCardCheckRequest) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "VideoLivePersonIdCardCheckRequest"}
	//name校验
	if validation.CheckName(r.Name) == false {
		invalidParams.Add(validation.NewErrParamInvalid("Name"))
	}
	//idCardNo校验
	if validation.CheckIdCard(r.CardNo) == false {
		invalidParams.Add(validation.NewErrParamInvalid("CardNo"))
	}
	// actions校验
	if r.Actions == nil || len(*r.Actions) == 0 {
		invalidParams.Add(validation.NewErrParamRequired("Actions"))
	}
	// actionVideos校验
	if r.ActionVideos == nil || len(*r.ActionVideos) == 0 {
		invalidParams.Add(validation.NewErrParamRequired("ActionVideos"))
	}
	// videoType校验
	if r.VideoType == nil {
		invalidParams.Add(validation.NewErrParamRequired("VideoType"))
	}
	// dataId可以为空，但是不能超过64个字符
	if validation.EmptyMaxLen(r.DataId, 64) == false {
		invalidParams.Add(validation.NewErrParamMaxLen("DataId", 64, "DataId长度不合法"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// 提供Name参数的SET方法
func (r *VideoLivePersonIdCardCheckRequest) SetName(name string) *VideoLivePersonIdCardCheckRequest {
	r.Name = &name
	return r
}

// 提供CardNo参数的SET方法
func (r *VideoLivePersonIdCardCheckRequest) SetCardNo(cardNo string) *VideoLivePersonIdCardCheckRequest {
	r.CardNo = &cardNo
	return r
}

// 提供Actions参数的SET方法
func (r *VideoLivePersonIdCardCheckRequest) SetActions(actions []string) *VideoLivePersonIdCardCheckRequest {
	r.Actions = &actions
	return r
}

// 提供ActionVideos参数的SET方法
func (r *VideoLivePersonIdCardCheckRequest) SetActionVideos(actionVideos []string) *VideoLivePersonIdCardCheckRequest {
	r.ActionVideos = &actionVideos
	return r
}

// 提供VideoType参数的SET方法
func (r *VideoLivePersonIdCardCheckRequest) SetVideoType(videoType string) *VideoLivePersonIdCardCheckRequest {
	r.VideoType = &videoType
	return r
}

// 提供NeedAvatar参数的SET方法
func (r *VideoLivePersonIdCardCheckRequest) SetNeedAvatar(needAvatar string) *VideoLivePersonIdCardCheckRequest {
	r.NeedAvatar = &needAvatar
	return r
}

// 提供DataId参数的SET方法
func (r *VideoLivePersonIdCardCheckRequest) SetDataId(dataId string) *VideoLivePersonIdCardCheckRequest {
	r.DataId = &dataId
	return r
}

// new method
func NewVideoLivePersonIdCardCheckRequest(businessId string) *VideoLivePersonIdCardCheckRequest {
	request := &VideoLivePersonIdCardCheckRequest{
		BizPostFormRequest: *types.NewBizPostFormRequest(businessId),
	}
	request.SetProductCode("auth")
	request.SetUriPattern("/v1/liveperson/h5/audit")
	request.SetVersion("v1")
	return request
}

// GetNonSignParams 获取具体业务中特有的不参与签名计算的参数
func (r *VideoLivePersonIdCardCheckRequest) GetNonSignParams() map[string]string {
	//actionVideos不参与签名
	actionVideosBites, _ := json.Marshal(*r.ActionVideos)
	return map[string]string{
		"actionVideos": string(actionVideosBites)}
}

// GetCustomSignParams 获取具体业务中特有的需要参与签名计算的参数
func (r *VideoLivePersonIdCardCheckRequest) GetBusinessCustomSignParams() map[string]string {
	params := r.BizPostFormRequest.GetBusinessCustomSignParams()
	if r.Name != nil {
		params["name"] = *r.Name
	}
	if r.CardNo != nil {
		params["cardNo"] = *r.CardNo
	}
	if r.Actions != nil {
		actionBytes, _ := json.Marshal(*r.Actions)
		params["actions"] = string(actionBytes)
	}
	if r.VideoType != nil {
		params["videoType"] = *r.VideoType
	}
	if r.NeedAvatar != nil {
		params["needAvatar"] = *r.NeedAvatar
	}
	if r.DataId != nil {
		params["dataId"] = *r.DataId
	}
	return params
}

// response
type VideoLivePersonIdCardCheckResponse struct {
	*types.CommonResponse
	Result VideoLivePersonIdCardCheckResult `json:"result"`
}

// result
type VideoLivePersonIdCardCheckResult struct {
	//taskId
	TaskId *string `json:"taskId"`
	//picType 图片类型：1-URL，2-BASE64
	PicType *int `json:"picType"`
	//avatar
	Avatar *string `json:"avatar"`
	//status
	Status *int `json:"status"`
	//reasonType
	ReasonType *int `json:"reasonType"`
	//similarityScore
	SimilarityScore *float64 `json:"similarityScore"`
	//faceMatched
	FaceMatched *int `json:"faceMatched"`
	//actionImages
	ActionImages *[]ActionImage `json:"actionImages"`
}

// 手机号与所有者身份证号及姓名验证请求 MobileNumberOwnerIdNameCheck /v1/mobile/check
func (c *AuthClient) MobileNumberOwnerIdNameCheck(request *MobileNumberOwnerIdNameCheckRequest) (response *MobileNumberOwnerIdNameCheckResponse, err error) {
	response = &MobileNumberOwnerIdNameCheckResponse{}
	err = c.Client.DoExecute(request, response)
	return
}

// check request
type MobileNumberOwnerIdNameCheckRequest struct {
	types.BizPostFormRequest
	//phone
	Phone *string
	//cardNo
	CardNo *string
	//name
	Name *string
}

// param check
func (r *MobileNumberOwnerIdNameCheckRequest) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "MobileNumberOwnerIdNameCheckRequest"}
	//name校验
	if validation.CheckName(r.Name) == false {
		invalidParams.Add(validation.NewErrParamInvalid("Name"))
	}
	//idCardNo校验
	if validation.CheckIdCard(r.CardNo) == false {
		invalidParams.Add(validation.NewErrParamInvalid("CardNo"))
	}
	//phone校验
	if validation.NotEmptyMinMaxLen(r.Phone, 1, 32) == false {
		invalidParams.Add(validation.NewErrParamMaxLen("Phone", 32, "Phone长度不合法"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// 提供Phone参数的SET方法
func (r *MobileNumberOwnerIdNameCheckRequest) SetPhone(phone string) *MobileNumberOwnerIdNameCheckRequest {
	r.Phone = &phone
	return r
}

// 提供CardNo参数的SET方法
func (r *MobileNumberOwnerIdNameCheckRequest) SetCardNo(cardNo string) *MobileNumberOwnerIdNameCheckRequest {
	r.CardNo = &cardNo
	return r
}

// 提供Name参数的SET方法
func (r *MobileNumberOwnerIdNameCheckRequest) SetName(name string) *MobileNumberOwnerIdNameCheckRequest {
	r.Name = &name
	return r
}

// new method
func NewMobileNumberOwnerIdNameCheckRequest(businessId string) *MobileNumberOwnerIdNameCheckRequest {
	request := &MobileNumberOwnerIdNameCheckRequest{
		BizPostFormRequest: *types.NewBizPostFormRequest(businessId),
	}
	request.SetProductCode("auth")
	request.SetUriPattern("/v1/mobile/check")
	request.SetVersion("v1")
	return request
}

// GetNonSignParams 获取具体业务中特有的不参与签名计算的参数
func (r *MobileNumberOwnerIdNameCheckRequest) GetNonSignParams() map[string]string {
	return make(map[string]string)
}

// GetCustomSignParams 获取具体业务中特有的需要参与签名计算的参数
func (r *MobileNumberOwnerIdNameCheckRequest) GetBusinessCustomSignParams() map[string]string {
	params := r.BizPostFormRequest.GetBusinessCustomSignParams()
	if r.Phone != nil {
		params["phone"] = *r.Phone
	}
	if r.CardNo != nil {
		params["cardNo"] = *r.CardNo
	}
	if r.Name != nil {
		params["name"] = *r.Name
	}
	return params
}

// response
type MobileNumberOwnerIdNameCheckResponse struct {
	*types.CommonResponse
	Result MobileNumberOwnerIdNameCheckResult `json:"result"`
}

// result
type MobileNumberOwnerIdNameCheckResult struct {
	//status
	Status *int `json:"status"`
	//taskId
	TaskId *string `json:"taskId"`
	//reasonType
	ReasonType *int `json:"reasonType"`
	//isPayed
	IsPayed *int `json:"isPayed"`
}

// 手机号与所有者姓名验证请求 MobileNumberOwnerNameCheck /v1/mobile2Ele/check
func (c *AuthClient) MobileNumberOwnerNameCheck(request *MobileNumberOwnerNameCheckRequest) (response *MobileNumberOwnerNameCheckResponse, err error) {
	response = &MobileNumberOwnerNameCheckResponse{}
	err = c.Client.DoExecute(request, response)
	return
}

// check request
type MobileNumberOwnerNameCheckRequest struct {
	types.BizPostFormRequest
	//phone
	Phone *string
	//name
	Name *string
}

// param check
func (r *MobileNumberOwnerNameCheckRequest) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "MobileNumberOwnerNameCheckRequest"}
	//name校验
	if validation.CheckName(r.Name) == false {
		invalidParams.Add(validation.NewErrParamInvalid("Name"))
	}
	//phone校验
	if validation.NotEmptyMinMaxLen(r.Phone, 1, 32) == false {
		invalidParams.Add(validation.NewErrParamMaxLen("Phone", 32, "Phone长度不合法"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// 提供Phone参数的SET方法
func (r *MobileNumberOwnerNameCheckRequest) SetPhone(phone string) *MobileNumberOwnerNameCheckRequest {
	r.Phone = &phone
	return r
}

// 提供Name参数的SET方法
func (r *MobileNumberOwnerNameCheckRequest) SetName(name string) *MobileNumberOwnerNameCheckRequest {
	r.Name = &name
	return r
}

// new method
func NewMobileNumberOwnerNameCheckRequest(businessId string) *MobileNumberOwnerNameCheckRequest {
	request := &MobileNumberOwnerNameCheckRequest{
		BizPostFormRequest: *types.NewBizPostFormRequest(businessId),
	}
	request.SetProductCode("auth")
	request.SetUriPattern("/v1/mobile2Ele/check")
	request.SetVersion("v1")
	return request
}

// GetNonSignParams 获取具体业务中特有的不参与签名计算的参数
func (r *MobileNumberOwnerNameCheckRequest) GetNonSignParams() map[string]string {
	return make(map[string]string)
}

// GetCustomSignParams 获取具体业务中特有的需要参与签名计算的参数
func (r *MobileNumberOwnerNameCheckRequest) GetBusinessCustomSignParams() map[string]string {
	params := r.BizPostFormRequest.GetBusinessCustomSignParams()
	if r.Phone != nil {
		params["phone"] = *r.Phone
	}
	if r.Name != nil {
		params["name"] = *r.Name
	}
	return params
}

// response
type MobileNumberOwnerNameCheckResponse struct {
	*types.CommonResponse
	Result MobileNumberOwnerNameCheckResult `json:"result"`
}

// result
type MobileNumberOwnerNameCheckResult struct {
	//status
	Status *int `json:"status"`
	//taskId
	TaskId *string `json:"taskId"`
	//reasonType
	ReasonType *int `json:"reasonType"`
	//isPayed
	IsPayed *int `json:"isPayed"`
}

// 身份证OCR请求 IdCardOcrCheck /v1/ocr/check
func (c *AuthClient) IdCardOcrCheck(request *IdCardOcrCheckRequest) (response *IdCardOcrCheckResponse, err error) {
	response = &IdCardOcrCheckResponse{}
	err = c.Client.DoExecute(request, response)
	return
}

// check request
type IdCardOcrCheckRequest struct {
	types.BizPostFormRequest
	//frontPicture 身份证正面照
	FrontPicture *string
	//backPicture 身份证反面照
	BackPicture *string
	//picType
	PicType *string
	//dataId
	DataId *string
}

// param check
func (r *IdCardOcrCheckRequest) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "IdCardOcrCheckRequest"}
	//picType校验
	if r.PicType == nil {
		invalidParams.Add(validation.NewErrParamRequired("PicType"))
	}
	//dataId可以为空，但是不能超过64个字符
	if validation.EmptyMaxLen(r.DataId, 64) == false {
		invalidParams.Add(validation.NewErrParamMaxLen("DataId", 64, "DataId长度不合法"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// 提供FrontPicture参数的SET方法
func (r *IdCardOcrCheckRequest) SetFrontPicture(frontPicture string) *IdCardOcrCheckRequest {
	r.FrontPicture = &frontPicture
	return r
}

// 提供BackPicture参数的SET方法
func (r *IdCardOcrCheckRequest) SetBackPicture(backPicture string) *IdCardOcrCheckRequest {
	r.BackPicture = &backPicture
	return r
}

// 提供PicType参数的SET方法
func (r *IdCardOcrCheckRequest) SetPicType(picType string) *IdCardOcrCheckRequest {
	r.PicType = &picType
	return r
}

// 提供DataId参数的SET方法
func (r *IdCardOcrCheckRequest) SetDataId(dataId string) *IdCardOcrCheckRequest {
	r.DataId = &dataId
	return r
}

// new method
func NewIdCardOcrCheckRequest(businessId string) *IdCardOcrCheckRequest {
	request := &IdCardOcrCheckRequest{
		BizPostFormRequest: *types.NewBizPostFormRequest(businessId),
	}
	request.SetProductCode("auth")
	request.SetUriPattern("/v1/ocr/check")
	request.SetVersion("v1")
	return request
}

// GetNonSignParams 获取具体业务中特有的不参与签名计算的参数
func (r *IdCardOcrCheckRequest) GetNonSignParams() map[string]string {
	return make(map[string]string)
}

// GetCustomSignParams 获取具体业务中特有的需要参与签名计算的参数
func (r *IdCardOcrCheckRequest) GetBusinessCustomSignParams() map[string]string {
	params := r.BizPostFormRequest.GetBusinessCustomSignParams()
	if r.FrontPicture != nil {
		params["frontPicture"] = *r.FrontPicture
	}
	if r.BackPicture != nil {
		params["backPicture"] = *r.BackPicture
	}
	if r.PicType != nil {
		params["picType"] = *r.PicType
	}
	if r.DataId != nil {
		params["dataId"] = *r.DataId
	}
	return params
}

// response
type IdCardOcrCheckResponse struct {
	*types.CommonResponse
	Result IdCardOcrCheckResult `json:"result"`
}

// result
type IdCardOcrCheckResult struct {
	//status
	Status *int `json:"status"`
	//taskId
	TaskId *string `json:"taskId"`
	//ocrResponseDetail
	OcrResponseDetail OcrResponseDetail `json:"ocrResponseDetail"`
	//isPayed
	IsPayed *int `json:"isPayed"`
}

// ocrResponseDetail
type OcrResponseDetail struct {
	//ocrName 识别姓名
	OcrName *string `json:"ocrName"`
	//ocrCardNo 识别身份证号
	OcrCardNo *string `json:"ocrCardNo"`
	//expireDate 身份证过期时间
	ExpireDate *string `json:"expireDate"`
	//gender 性别
	Gender *string `json:"Gender"`
	//nation 名族
	Nation *string `json:"nation"`
	// birthday 出身日期
	Birthday *string `json:"birthday"`
	//address 地址
	Address *string `json:"address"`
	//authority 发证机关
	Authority *string `json:"authority"`
}

// 实人认证请求 RpCheck /v1/rp/check
func (c *AuthClient) RpCheck(request *RpCheckRequest) (response *RpCheckResponse, err error) {
	response = &RpCheckResponse{}
	err = c.Client.DoExecute(request, response)
	return
}

// check request
type RpCheckRequest struct {
	types.BizPostFormRequest
	//name
	Name *string
	//cardNo
	CardNo *string
	//picType
	PicType *string
	//avatar
	Avatar *string
	//dataId
	DataId *string
}

// param check
func (r *RpCheckRequest) ValidateParam() error {
	invalidParams := validation.ErrInvalidParams{Context: "RpCheckRequest"}
	//picType校验
	if r.PicType == nil {
		invalidParams.Add(validation.NewErrParamRequired("PicType"))
	}
	//name
	if validation.CheckName(r.Name) == false {
		invalidParams.Add(validation.NewErrParamInvalid("Name"))
	}
	//idCardNo
	if validation.CheckIdCard(r.CardNo) == false {
		invalidParams.Add(validation.NewErrParamInvalid("CardNo"))
	}
	//avatar
	if r.Avatar == nil {
		invalidParams.Add(validation.NewErrParamRequired("Avatar"))
	}
	//dataId可以为空，但是不能超过64个字符
	if validation.EmptyMaxLen(r.DataId, 64) == false {
		invalidParams.Add(validation.NewErrParamMaxLen("DataId", 64, "DataId长度不合法"))
	}
	if invalidParams.Len() > 0 {
		return invalidParams
	}
	return nil
}

// 提供Name参数的SET方法
func (r *RpCheckRequest) SetName(name string) *RpCheckRequest {
	r.Name = &name
	return r
}

// 提供CardNo参数的SET方法
func (r *RpCheckRequest) SetCardNo(cardNo string) *RpCheckRequest {
	r.CardNo = &cardNo
	return r
}

// 提供PicType参数的SET方法
func (r *RpCheckRequest) SetPicType(picType string) *RpCheckRequest {
	r.PicType = &picType
	return r
}

// 提供Avatar参数的SET方法
func (r *RpCheckRequest) SetAvatar(avatar string) *RpCheckRequest {
	r.Avatar = &avatar
	return r
}

// 提供DataId参数的SET方法
func (r *RpCheckRequest) SetDataId(dataId string) *RpCheckRequest {
	r.DataId = &dataId
	return r
}

func NewRpCheckRequest(businessId string) *RpCheckRequest {
	request := &RpCheckRequest{
		BizPostFormRequest: *types.NewBizPostFormRequest(businessId),
	}
	request.SetProductCode("auth")
	request.SetUriPattern("/v1/rp/check")
	request.SetVersion("v1")
	return request
}

// GetNonSignParams 获取具体业务中特有的不参与签名计算的参数
func (r *RpCheckRequest) GetNonSignParams() map[string]string {
	return make(map[string]string)
}

// GetCustomSignParams 获取具体业务中特有的需要参与签名计算的参数
func (r *RpCheckRequest) GetBusinessCustomSignParams() map[string]string {
	params := r.BizPostFormRequest.GetBusinessCustomSignParams()
	if r.Name != nil {
		params["name"] = *r.Name
	}
	if r.CardNo != nil {
		params["cardNo"] = *r.CardNo
	}
	if r.PicType != nil {
		params["picType"] = *r.PicType
	}
	if r.Avatar != nil {
		params["avatar"] = *r.Avatar
	}
	if r.DataId != nil {
		params["dataId"] = *r.DataId
	}
	return params
}

// response
type RpCheckResponse struct {
	*types.CommonResponse
	Result RpCheckResult `json:"result"`
}

// result
type RpCheckResult struct {
	//taskId 本次请求数据标识，可以根据该标识在控制台进行数据查询
	TaskId *string `json:"taskId"`
	//status 姓名身份证号认证结果
	Status *int `json:"status"`
	//reasonType 认证结果对应的原因
	ReasonType *int `json:"reasonType"`
	//similarityScore 用户所传头像与身份证头像的相似度得分
	SimilarityScore *float64 `json:"similarityScore"`
	//faceMatched 人脸比对认证结果
	FaceMatched *int `json:"faceMatched"`
	// isPayed 本次请求是否收费标识，1代表收费，0代表不收费
	IsPayed *int `json:"isPayed"`
}
