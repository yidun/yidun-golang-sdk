package iriskiface

import (
	"github.com/yidun/yidun-golang-sdk/yidun/service/irisk"
)

type IRISKAPI interface {
	// GetConfig 配置拉取，调用/v5/risk/getConfig
	GetConfig(request *irisk.IRiskConfigRequest) (response *irisk.IRiskConfigResponse, err error)
	// GetMediaCheckResult 图片外挂识别，调用/v5/risk/mediaCheck
	GetMediaCheckResult(request *irisk.IRiskMediaCheckRequest) (response *irisk.IRiskMediaCheckResponse, err error)	
	// GetMediaQueryResult 图片外挂识别结果查询，调用/v5/risk/mediaQuery
	GetMediaQueryResult(request *irisk.IRiskMediaQueryRequest) (response *irisk.IRiskMediaQueryResponse, err error)
	// GetDetailResult 数据查询，调用/v5/risk/detail
	GetDetailResult(request *irisk.IRiskDetailRequest) (response *irisk.IRiskDetailResponse, err error)
	// GetDetailV6Result 数据查询，调用/v6/risk/detail
	GetDetailV6Result(request *irisk.IRiskDetailV6Request) (response *irisk.IRiskDetailV6Response, err error)
	// GetCheckResult 在线检测，调用/v6/risk/check
	GetCheckResult(request *irisk.IRiskCheckRequest) (response *irisk.IRiskCheckResponse, err error)
	// UploadReportData 举报数据上报，调用/v5/risk/reportData
	UploadReportData(request *irisk.IRiskReportDataRequest) (response *irisk.IRiskReportDataResponse, err error)
	// UploadDisposeInfo 上报处置数据，调用/v5/disposeInfo/upload
	UploadDisposeInfo(request *irisk.IRiskDisposeUploadRequest) (response *irisk.IRiskDisposeUploadResponse, err error)
	// AntiGoldCheck 反打金检测，调用/v5/risk/antigold
	AntiGoldCheck(request *irisk.IRiskAntiGoldCheckRequest) (response *irisk.IRiskAntiGoldCheckResponse, err error)
	// ListQuery 名单查询，调用/v5/list/query
	ListQuery(request *irisk.IRiskListQueryRequest) (response *irisk.IRiskListQueryResponse, err error)
	// ListAdd 名单添加，调用/v5/list/add
	ListAdd(request *irisk.IRiskListAddRequest) (response *irisk.IRiskListAddResponse, err error)

}

var _ IRISKAPI = (*irisk.IRiskClient)(nil)
