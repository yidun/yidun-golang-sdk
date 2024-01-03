package label

import (
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/label/request"
	"github.com/yidun/yidun-golang-sdk/yidun/service/antispam/label/response"
)

// 标签查询
func (c *LabelClient) QueryLabel(req *request.LabelQueryRequest) (res *response.LabelQueryResponse, err error) {
	res = &response.LabelQueryResponse{}
	err = c.Client.DoExecute(req, res)
	return
}