package submit

// 微博检测提交接口v1.0
type WeiboBaseSubmitV1Request struct {
	// 循环爬虫时间区间--开始时间
	StartTime *int64 `json:"startTime,omitempty"`
	// 循环爬虫时间区间--结束时间
	EndTime *int64 `json:"endTime,omitempty"`
	// 最大检测主贴数量
	MaxCheckCount *int `json:"maxCheckCount,omitempty"`
	// 最大检测评论数量
	MaxComment *int `json:"maxComment,omitempty"`
	// 检测频率/多久爬取一次，单位毫秒
	Frequency *int64 `json:"frequency,omitempty"`
	// 任务类型
	Type *int `json:"type,omitempty"`
	// 主动回调地址
	CallbackUrl *string `json:"callbackUrl,omitempty"`
	// 爬取策略
	Strategy *int `json:"strategy,omitempty"`
	// 检测标记  检测内容
	CheckFlags *[]int `json:"checkFlags,omitempty"`
}

type WeiboConfig struct {
	// 微博主页链接
	Url *string `json:"url,omitempty"`
	// 微博名
	Blogger *string `json:"blogger,omitempty"`
}

// SetUrl
func (o *WeiboConfig) SetUrl(url string) {
	o.Url = &url
}

// SetBlogger
func (o *WeiboConfig) SetBlogger(blogger string) {
	o.Blogger = &blogger
}

// SetStartTime 设置循环爬虫时间区间--开始时间
func (o *WeiboBaseSubmitV1Request) SetStartTime(startTime int64) {
	o.StartTime = &startTime
}

// SetEndTime 设置循环爬虫时间区间--结束时间
func (o *WeiboBaseSubmitV1Request) SetEndTime(endTime int64) {
	o.EndTime = &endTime
}

// SetMaxCheckCount 设置最大检测数量
func (o *WeiboBaseSubmitV1Request) SetMaxCheckCount(maxCheckCount int) {
	o.MaxCheckCount = &maxCheckCount
}

// SetMaxComment 设置最大评论数
func (o *WeiboBaseSubmitV1Request) SetMaxComment(maxComment int) {
	o.MaxComment = &maxComment
}

// SetFrequency 设置检测频率
func (c *WeiboBaseSubmitV1Request) SetFrequency(frequency int64) {
	c.Frequency = &frequency
}

// SetType 设置任务类型
func (o *WeiboBaseSubmitV1Request) SetType(typ int) {
	o.Type = &typ
}

// SetCallbackUrl 设置主动回调地址
func (o *WeiboBaseSubmitV1Request) SetCallbackUrl(callbackUrl string) {
	o.CallbackUrl = &callbackUrl
}

// SetStrategy 设置爬取策略
func (o *WeiboBaseSubmitV1Request) SetStrategy(strategy int) {
	o.Strategy = &strategy
}

// SetCheckFlags 设置检测标记 检测内容
func (o *WeiboBaseSubmitV1Request) SetCheckFlags(checkFlags []int) {
	o.CheckFlags = &checkFlags
}
