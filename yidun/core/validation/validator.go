package validation

import "regexp"

const (
	NAME_PATTERN = `^[\p{Han}・]{2,32}$`
)

// 可以为空，但是不能超过最大值
func EmptyMaxLen(param *string, max int) bool {
	if param != nil && len(*param) > max {
		return false
	}
	return true
}

// 可以为空，不为空时不能小于最小值
func EmptyMinLen(param *string, min int) bool {
	if param != nil && len(*param) < min {
		return false
	}
	return true
}

// 不能为空，不能超过最大长度
func NotEmptyMaxLen(param *string, max int) bool {
	if param == nil || len(*param) > max {
		return false
	}
	return true
}

// 不能为空，且在最小最大值之间
func NotEmptyMinMaxLen(param *string, min int, max int) bool {
	if param == nil || len(*param) > max || len(*param) < min {
		return false
	}
	return true
}

// 不能为空，且不能小于最小值
func NotEmptyMinLen(param *string, min int) bool {
	if param == nil || len(*param) < min {
		return false
	}
	return true
}

// 校验姓名是汉字或者中文的连字符（・）
func CheckName(name *string) bool {
	if name == nil || len(*name) > 32 {
		return false
	}
	reg := regexp.MustCompile(NAME_PATTERN)
	return reg.MatchString(*name)
}

var regex15 = `^[1-9]\d{5}\d{2}(0[1-9]|10|11|12)([0-2][1-9]|10|20|30|31)\d{3}$`
var regex18 = `^[1-9]\d{5}(19|20)\d{2}(0[1-9]|10|11|12)([0-2][1-9]|10|20|30|31)\d{3}[\dXx]$`

// 校验身份证号码是否符合格式
func CheckIdCard(idCard *string) bool {
	// 判断idCard号码是否为空， idCard号码长度是否为15位或者18位
	if idCard == nil {
		return false
	}
	if len(*idCard) != 15 && len(*idCard) != 18 {
		return false
	}
	lenght := len(*idCard)
	if lenght == 15 {
		reg := regexp.MustCompile(regex15)
		return reg.MatchString(*idCard)
	}
	reg := regexp.MustCompile(regex18)
	return reg.MatchString(*idCard)
}
