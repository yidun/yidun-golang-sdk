package test

import (
	"testing"

	"github.com/yidun/yidun-golang-sdk/yidun/core/validation"
)

// 测试中文名
func TestValidaChineseName(t *testing.T) {
	name := "张三"
	if validation.CheckName(&name) {
		print("success")
	}
}

// 测试英文名
func TestIdCardName(t *testing.T) {
	idcard := "50023619911008123x"
	if validation.CheckIdCard(&idcard) {
		print("success")
	}
}
