package callback

import (
	"github.com/yidun/yidun-golang-sdk/yidun/core/types"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/media/v2/common/response"
)

type CrawlerResourceCallbackV3Response struct {
	*types.CommonResponse
	Result []*CrawlerResourceResult `json:"result,omitempty"`
}

// 网站资源检测v3回调结果
type CrawlerResourceResult struct {
	// 反垃圾结果
	Antispam *CrawlerAntispamResponse `json:"antispam,omitempty"`
	// 审核信息
	Censor *CrawlerCensorResponse `json:"censor,omitempty"`
	// 增值服务结果
	ValueAddService *CrawlerValueAddServiceResponse `json:"valueAddService,omitempty"`
}

type CrawlerAntispamResponse struct {
	*response.MediaAntispamResponse

	// 检测失败原因，检测失败时才返回该字段；
	// 1：爬虫失败；2：试用量超限；3：业务已关闭；4：爬取内容为空；
	// 5：连接失败；6：404网页不存在；7：正在爬取中,重复提交；100：其他原因-内容检测失败；
	FailureReason *int `json:"failureReason,omitempty"`
}

type CrawlerCensorResponse struct {
	*response.MediaCensorResponse
	SiteName *string `json:"siteName,omitempty"`
	SiteUrl  *string `json:"siteUrl,omitempty"`
	// 快照URL
	SnapshotUrl *string `json:"snapshotUrl,omitempty"`
}

type CrawlerValueAddServiceResponse struct {
	*response.MediaValueAddServiceResponse
}
