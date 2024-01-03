package preiface

import (
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/pretreatment"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/pretreatment/v1/delete"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/pretreatment/v1/query"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/pretreatment/v1/update"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/pretreatment/v2/add"
)

type PretreatmentClientAPI interface {
	// ADD 添加忽略词
	Add(req *add.PretreatmentAddRequest) (res *add.PretreatmentAddResponse, err error)
	// Delete 删除忽略词
	Delete(req *delete.PretreatmentDeleteRequest) (res *delete.PretreatmentDeleteResponse, err error)
	// Query 查询名单
	Query(req *query.PretreatmentQueryRequest) (res *query.PretreatmentQueryResponse, err error)
	// Update 更新名单
	Update(req *update.PretreatmentUpdateRequest) (res *update.PretreatmentUpdateResponse, err error)
}

var _ PretreatmentClientAPI = (*pretreatment.PretreatmentClient)(nil)
