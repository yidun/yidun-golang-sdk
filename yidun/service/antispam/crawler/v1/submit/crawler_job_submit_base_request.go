package submit

// CrawlerJobBaseSubmitV1Request 网站任务检测提交接口v1.0 公共提交参数
type CrawlerJobBaseSubmitV1Request struct {
	// 循环爬虫时间区间--开始时间
	SliceStartTime *int64 `json:"sliceStartTime,omitempty"`
	// 循环爬虫时间区间--结束时间
	SliceEndTime *int64 `json:"sliceEndTime,omitempty"`
	// 爬虫深度/网站层级
	Level *int `json:"level,omitempty"`
	// 检测频率/多久爬取一次，单位毫秒
	Frequency *int64 `json:"frequency,omitempty"`
	// 单次任务周期内爬取页面的最大数量
	MaxResourceAmount *int `json:"maxResourceAmount,omitempty"`
	// 任务类型；0：循环任务；1：单次任务
	Type *int `json:"type,omitempty"`
	// 主动回调地址
	CallbackUrl *string `json:"callbackUrl,omitempty"`
	// 检测策略；当循环任务时，此配置生效。1：全量页面；2：首次全量，后续增量
	CheckStrategy *int `json:"checkStrategy,omitempty"`
	// 网站爬取配置
	Config *string `json:"config,omitempty"`
	// user agent 匹配规则
	// @see com.netease.yidun.sdk.antispam.enums.UserAgentMatchType
	UserAgentMatchType *int `json:"userAgentMatchType,omitempty"`
	// 浏览器user agent
	UserAgent *string `json:"userAgent,omitempty"`
	// 检测内容 默认为1和2。1-检测文本，2-检测图片，4-检测点播音频，5-检测文档附件，6-检测点播音视频
	CheckFlags *[]int `json:"checkFlag,omitempty"`
}

func (c *CrawlerJobBaseSubmitV1Request) SetSliceStartTime(sliceStartTime int64) {
	c.SliceStartTime = &sliceStartTime
}

func (c *CrawlerJobBaseSubmitV1Request) SetSliceEndTime(sliceEndTime int64) {
	c.SliceEndTime = &sliceEndTime
}

func (c *CrawlerJobBaseSubmitV1Request) SetLevel(level int) {
	c.Level = &level
}

func (c *CrawlerJobBaseSubmitV1Request) SetFrequency(frequency int64) {
	c.Frequency = &frequency
}

func (c *CrawlerJobBaseSubmitV1Request) SetMaxResourceAmount(maxResourceAmount int) {
	c.MaxResourceAmount = &maxResourceAmount
}

func (c *CrawlerJobBaseSubmitV1Request) SetType(typeCode int) {
	c.Type = &typeCode
}

func (c *CrawlerJobBaseSubmitV1Request) SetCallbackUrl(callbackUrl string) {
	c.CallbackUrl = &callbackUrl
}

func (c *CrawlerJobBaseSubmitV1Request) SetCheckStrategy(checkStrategy int) {
	c.CheckStrategy = &checkStrategy
}

func (c *CrawlerJobBaseSubmitV1Request) SetConfig(config string) {
	c.Config = &config
}

func (c *CrawlerJobBaseSubmitV1Request) SetUserAgentMatchType(userAgentMatchType int) {
	c.UserAgentMatchType = &userAgentMatchType
}

func (c *CrawlerJobBaseSubmitV1Request) SetUserAgent(userAgent string) {
	c.UserAgent = &userAgent
}

func (c *CrawlerJobBaseSubmitV1Request) SetCheckFlags(checkFlag []int) {
	c.CheckFlags = &checkFlag
}
