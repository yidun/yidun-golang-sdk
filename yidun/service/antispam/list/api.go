package list

import (
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

// Imagelist 删除名单
func (c *ListClient) ImagelistDelete(req *delete.ImageListDeleteRequest) (res *delete.ImageListDeleteResponse, err error) {
	res = &delete.ImageListDeleteResponse{}
	err = c.Client.DoExecute(req, res)
	return
}

// ImagelistQuery 查询名单
func (c *ListClient) ImagelistQuery(req *query.ImageListQueryRequest) (res *query.ImageListQueryResponse, err error) {
	res = &query.ImageListQueryResponse{}
	err = c.Client.DoExecute(req, res)
	return
}

// ImagelistSubmit 提交名单
func (c *ListClient) ImagelistSubmit(req *submit.ImageListSubmitRequest) (res *submit.ImageListSubmitResponse, err error) {
	res = &submit.ImageListSubmitResponse{}
	err = c.Client.DoExecute(req, res)
	return
}

// ImagelistSubmit 提交名单
func (c *ListClient) ImagelistUpdate(req *update.ImageListUpdateRequest) (res *update.ImageListUpdateResponse, err error) {
	res = &update.ImageListUpdateResponse{}
	err = c.Client.DoExecute(req, res)
	return
}

// ImagelistUpdate 更新名单
func (c *ListClient) Submit(req *request.ListSubmitRequest) (res *request2.ListSubmitResponse, err error) {
	res = &request2.ListSubmitResponse{}
	err = c.Client.DoExecute(req, res)
	return
}

func (c *ListClient) Delete(req *request4.ListDeleteRequest) (res *request3.ListDeleteResponse, err error) {
	res = &request3.ListDeleteResponse{}
	err = c.Client.DoExecute(req, res)
	return
}

func (c *ListClient) Query(req *request6.ListQueryRequest) (res *request7.ListQueryResponse, err error) {
	res = &request7.ListQueryResponse{}
	err = c.Client.DoExecute(req, res)
	return
}

func (c *ListClient) Update(req *request5.ListUpdateRequest) (res *request8.ListUpdateResponse, err error) {
	res = &request8.ListUpdateResponse{}
	err = c.Client.DoExecute(req, res)
	return
}
