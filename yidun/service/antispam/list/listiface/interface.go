package listiface

import (
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/list"
	request4 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/list/delete/v2/request"
	request3 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/list/delete/v2/response"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/list/imagelist/v1/delete"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/list/imagelist/v1/query"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/list/imagelist/v1/submit"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/list/imagelist/v1/update"
	request6 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/list/query/v2/request"
	request7 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/list/query/v2/response"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/list/submit/v2/request"
	request2 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/list/submit/v2/response"
	request5 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/list/update/v2/request"
	request8 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/list/update/v2/response"
)

type ListClientAPI interface {
	// Submit 提交名单
	Submit(req *request.ListSubmitRequest) (res *request2.ListSubmitResponse, err error)
	// Delete 删除名单
	Delete(req *request4.ListDeleteRequest) (res *request3.ListDeleteResponse, err error)
	// Query 查询名单
	Query(req *request6.ListQueryRequest) (res *request7.ListQueryResponse, err error)
	// Update 更新名单
	Update(req *request5.ListUpdateRequest) (res *request8.ListUpdateResponse, err error)
	// Imagelist 删除名单
	ImagelistDelete(req *delete.ImageListDeleteRequest) (res *delete.ImageListDeleteResponse, err error)
	// ImagelistQuery 查询名单
	ImagelistQuery(req *query.ImageListQueryRequest) (res *query.ImageListQueryResponse, err error)
	// ImagelistSubmit 提交名单
	ImagelistSubmit(req *submit.ImageListSubmitRequest) (res *submit.ImageListSubmitResponse, err error)
	// ImagelistUpdate 更新名单
	ImagelistUpdate(req *update.ImageListUpdateRequest) (res *update.ImageListUpdateResponse, err error)
}

var _ ListClientAPI = (*list.ListClient)(nil)
