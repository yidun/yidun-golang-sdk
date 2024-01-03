package keyword

import (
	request4 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/keyword/delete/v1/request"
	request3 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/keyword/delete/v1/response"
	request6 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/keyword/query/v1/request"
	request7 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/keyword/query/v1/response"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/keyword/submit/v1/request"
	request2 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/keyword/submit/v1/response"
	request5 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/keyword/update/v2/request"
	request8 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/keyword/update/v2/response"
)

func (c *KeywordClient) Submit(req *request.KeywordSubmitRequest) (res *request2.KeywordSubmitResponse, err error) {
	res = &request2.KeywordSubmitResponse{}
	err = c.Client.DoExecute(req, res)
	return
}

func (c *KeywordClient) Delete(req *request4.KeywordDeleteRequest) (res *request3.KeywordDeleteResponse, err error) {
	res = &request3.KeywordDeleteResponse{}
	err = c.Client.DoExecute(req, res)
	return
}

func (c *KeywordClient) Query(req *request6.KeywordQueryRequest) (res *request7.KeywordQueryResponse, err error) {
	res = &request7.KeywordQueryResponse{}
	err = c.Client.DoExecute(req, res)
	return
}

func (c *KeywordClient) Update(req *request5.KeywordUpdateRequest) (res *request8.KeywordUpdateResponse, err error) {
	res = &request8.KeywordUpdateResponse{}
	err = c.Client.DoExecute(req, res)
	return
}
