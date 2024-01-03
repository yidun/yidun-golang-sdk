package keywordiface

import (
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/keyword"
	request4 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/keyword/delete/v1/request"
	request3 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/keyword/delete/v1/response"
	request6 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/keyword/query/v1/request"
	request7 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/keyword/query/v1/response"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/keyword/submit/v1/request"
	request2 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/keyword/submit/v1/response"
	request5 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/keyword/update/v2/request"
	request8 "github.com/yidun/yidun-golang-sdk/yidun/service/antispam/keyword/update/v2/response"
)

type KeywordClientAPI interface {
	// Submit 提交关键词
	Submit(req *request.KeywordSubmitRequest) (res *request2.KeywordSubmitResponse, err error)
	// Delete 删除关键词
	Delete(req *request4.KeywordDeleteRequest) (res *request3.KeywordDeleteResponse, err error)
	// Query 查询关键词
	Query(req *request6.KeywordQueryRequest) (res *request7.KeywordQueryResponse, err error)
	// Update 更新关键词
	Update(req *request5.KeywordUpdateRequest) (res *request8.KeywordUpdateResponse, err error)
}

var _ KeywordClientAPI = (*keyword.KeywordClient)(nil)
