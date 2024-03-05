package main

import (
	"log"

	"github.com/yidun/yidun-golang-sdk/yidun/service/irisk"

	"fmt"
)

func main() {
	request := irisk.NewIRiskListQueryRequest("BUSINESS_ID")
	request.SetListGroupCode("ayp76c2dmh2k0ktd8jyia2cg22009v02")
	// 可设置分页参数进行分页查询
	request.SetPageNum(1)
	request.SetBeginModifyTime(1656296583895)
	request.SetEndModifyTime(1722752006000)

	iriskClient := irisk.NewIRiskClientWithAccessKey("SECRET_ID", "SECRET_KEY")
	response, err := iriskClient.ListQuery(request)

	if err != nil {
		// 处理错误并打印日志
		log.Fatal(err)
	}
	if response.GetCode() == 200 {
		// 查询成功
		data := response.Data
		// 总数
		fmt.Println("count: ", data.Count)
		// 详情列表数据
		fmt.Println("rows: ", data.Rows)
		// 可以遍历输出详情列表数据
		for _, item := range data.Rows {
			fmt.Println("Content: ", item.Content)
			fmt.Println("Description: ", item.Description)
			// 其它字段 ......
		}
	} else {
		fmt.Println("error code: ", response.GetCode())
		fmt.Println("error msg: ", response.GetMsg())
	}
}
