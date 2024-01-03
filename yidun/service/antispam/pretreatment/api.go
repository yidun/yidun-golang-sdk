package pretreatment

import (
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/pretreatment/v1/delete"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/pretreatment/v1/query"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/pretreatment/v1/update"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/pretreatment/v2/add"
)

func (c *PretreatmentClient) Add(req *add.PretreatmentAddRequest) (res *add.PretreatmentAddResponse, err error) {
	res = &add.PretreatmentAddResponse{}
	err = c.Client.DoExecute(req, res)
	return
}

func (c *PretreatmentClient) Delete(req *delete.PretreatmentDeleteRequest) (res *delete.PretreatmentDeleteResponse, err error) {
	res = &delete.PretreatmentDeleteResponse{}
	err = c.Client.DoExecute(req, res)
	return
}

func (c *PretreatmentClient) Query(req *query.PretreatmentQueryRequest) (res *query.PretreatmentQueryResponse, err error) {
	res = &query.PretreatmentQueryResponse{}
	err = c.Client.DoExecute(req, res)
	return
}

func (c *PretreatmentClient) Update(req *update.PretreatmentUpdateRequest) (res *update.PretreatmentUpdateResponse, err error) {
	res = &update.PretreatmentUpdateResponse{}
	err = c.Client.DoExecute(req, res)
	return
}
