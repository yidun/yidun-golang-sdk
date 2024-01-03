package labeliface

import (
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/label"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/label/request"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/label/response"
)

type LabelApi interface {

	//标签查询
	QueryLabel(req *request.LabelQueryRequest) (res *response.LabelQueryResponse, err error)

}

var _ LabelApi = (*label.LabelClient)(nil)