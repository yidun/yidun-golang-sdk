package submit

// URL过滤条件
type CrawlerUrlFilter struct {
	// 过滤类型
	Type *int    `json:"type,omitempty"`
	Url  *string `json:"url,omitempty"`
}

func (c *CrawlerUrlFilter) SetType(typeCode int) {
	c.Type = &typeCode
}

func (c *CrawlerUrlFilter) SetUrl(url string) {
	c.Url = &url
}
