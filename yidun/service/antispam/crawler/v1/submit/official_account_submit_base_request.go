package submit

// OfficialAccountsBaseSubmitV1Request 公众号检测提交接口v1.0
type OfficialAccountsBaseSubmitV1Request struct {
	// 循环爬虫时间区间--开始时间
	StartTime *int64 `json:"startTime,omitempty"`
	// 循环爬虫时间区间--结束时间
	EndTime *int64 `json:"endTime,omitempty"`
	// 最大检测数量
	MaxCheckCount *int `json:"maxCheckCount,omitempty"`
	// 任务类型
	Type *int `json:"type,omitempty"`
	// 主动回调地址
	CallbackUrl *string `json:"callbackUrl,omitempty"`
	// 爬取策略
	Strategy *int `json:"strategy,omitempty"`
	// 检测标记 检测内容
	CheckFlags *[]int `json:"checkFlags,omitempty"`
}

type OfficicalAccountsConfig struct {
	// 数据唯一标识，能够根据该值定位到该条数据，如对数据结果有异议，可以发送该值给客户经理查询
	DataId *string `json:"dataId,omitempty"`
	// 微信号
	WechatAccount *string `json:"wechatAccount"`
	// 公众号名称
	OfficialAccountName *string `json:"officialAccountName,omitempty"`
}

// SetDataId 设置数据唯一标识
func (req *OfficicalAccountsConfig) SetDataId(dataId string) {
	req.DataId = &dataId
}

// SetWechatAccount 设置微信号
func (req *OfficicalAccountsConfig) SetWechatAccount(wechatAccount string) {
	req.WechatAccount = &wechatAccount
}

// SetOfficialAccountName 设置公众号名称
func (req *OfficicalAccountsConfig) SetOfficialAccountName(officialAccountName string) {
	req.OfficialAccountName = &officialAccountName
}

// SetStartTime 设置循环爬虫时间区间--开始时间
func (o *OfficialAccountsBaseSubmitV1Request) SetStartTime(startTime int64) {
	o.StartTime = &startTime
}

// SetEndTime 设置循环爬虫时间区间--结束时间
func (o *OfficialAccountsBaseSubmitV1Request) SetEndTime(endTime int64) {
	o.EndTime = &endTime
}

// SetMaxCheckCount 设置最大检测数量
func (o *OfficialAccountsBaseSubmitV1Request) SetMaxCheckCount(maxCheckCount int) {
	o.MaxCheckCount = &maxCheckCount
}

// SetType 设置任务类型
func (o *OfficialAccountsBaseSubmitV1Request) SetType(typ int) {
	o.Type = &typ
}

// SetCallbackUrl 设置主动回调地址
func (o *OfficialAccountsBaseSubmitV1Request) SetCallbackUrl(callbackUrl string) {
	o.CallbackUrl = &callbackUrl
}

// SetStrategy 设置爬取策略
func (o *OfficialAccountsBaseSubmitV1Request) SetStrategy(strategy int) {
	o.Strategy = &strategy
}

// SetCheckFlags 设置检测标记 检测内容
func (o *OfficialAccountsBaseSubmitV1Request) SetCheckFlags(checkFlags []int) {
	o.CheckFlags = &checkFlags
}
