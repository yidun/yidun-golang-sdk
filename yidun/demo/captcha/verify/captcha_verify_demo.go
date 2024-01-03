package main

import (
	"fmt"
	"log"

	captcha "github.com/yidun/yidun-golang-sdk/yidun/service/captcha"
)

// 验证码二次校验请求
func main() {
	request := captcha.NewCaptchaVerifyRequest()
	var captchaId string = "a05f036b70ab447b87cc788af9a60974"
	var validate string = "VnQ-P4nETS4ToEP6w0ZQIhAbjyxPod1fceKXkxScvMSkkH4ider5prw8BVafKBP79z4MBNRacPo.qPjeE.VJu6-MpOOIueY2zoA4aEHf5EM9Eh7RN6UKRb-BGh09H8CD-9HFX7rgwOmviPquD4ioVnsYwMTaP8LioFXZx26nLCzl6adYCW0rIQdMA17p4rSF71or5R66i4rSLRNN5Tqy8BzO70XtxA2LpOeK0TcfPgavHrYwocvxCPRDQN4XDjdAfTJMwbLGqYQavO26j01gIeTEQ_8WAqCgkw6JdwofxYg6Ni.YixIqlmYDu92FcfPZP-0NFmKU8KBJIFHTLYnzxQPiFMOQ0FvhnZtJPR6IkM4Z6deBkcbw.eo5TaxGQ2x7DAjhYZh-riegT2iT8EjLxIUpwY6.h989LrnoB.QCc50QbN2EgLbuikh.tqVIZHQUOC7pIQ9XKIVjpp5HObZPSc-.QiU1-8VYIi-P8Jy-MQ0hvHQP-GjhwR8yXZa3"
	var user string = ""
	request.SetCaptchaId(captchaId).SetValidate(validate).SetUser(user)
	captchaClient := captcha.NewCaptchaVerifyClientWithAccessKey("SECRET_ID", "SECRET_KEY")
	response, err := captchaClient.Verify(request)

	if err != nil {
		// 处理错误并打印日志
		log.Fatal(err)
	}
	log.Println(response)
	if *response.Error == 0 {
		fmt.Println(*response.Error, *response.Msg, *response.Result)
	}
}
