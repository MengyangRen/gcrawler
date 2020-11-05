package validator

import (
	"strconv"
	"strings"
	"time"
	"unicode"
	"unicode/utf8"

	"github.com/asaskevich/govalidator"
	//"fmt"
)

// 判断字符是否为数字
func isDigit(r rune) bool {
	return '0' <= r && r <= '9'
}

// 判断字符是否为英文字符
func isAlpha(r rune) bool {

	if r >= 'A' && r <= 'Z' {
		return true
	} else if r >= 'a' && r <= 'z' {
		return true
	}
	return false
}

func isPriv(s string) bool {

	if s == "" {
		return false
	}

	for _, r := range s {
		if (r < 'A' || r > 'Z') && (r < 'a' || r > 'z') && r != '_' {
			return false
		}
	}

	return true
}

// 匹配值是否为空
func checkStr(str string) bool {
	n := len(str)
	if n <= 0 {
		return false
	}
	return true
}

// 判断是否为bool
func checkBool(str string) bool {

	_, err := strconv.ParseBool(str)
	if err != nil {
		return false
	}
	return true
}

// 判断是否为float
func checkFloat(str string) bool {

	_, err := strconv.ParseFloat(str, 64)
	if err != nil {
		return false
	}
	return true
}

// 判断长度
func checkLength(str string, min, max int) bool {

	if min == 0 && max == 0 {
		return true
	}

	n := len(str)
	if n < min || n > max {
		return false
	}

	return true
}

// 判断字符串长度
func CheckStringLength(val string, _min, _max int) bool {

	if _min == 0 && _max == 0 {
		return true
	}

	count := utf8.RuneCountInString(val)
	if count < _min || count > _max {

		return false
	}
	return true
}

// 判断数字范围
func checkIntScope(s string, min, max int64) bool {

	val, err := strconv.ParseInt(s, 10, 64)
	if err != nil {
		return false
	}

	if val < min || val > max {
		return false
	}

	return true
}

// 判断是否全为数字
func CheckStringDigit(s string) bool {

	if s == "" {
		return false
	}
	for _, r := range s {
		if (r < '0' || r > '9') && r != '-' {
			return false
		}
	}
	return true
}

// 判断是否全为数字+逗号
func CheckStringCommaDigit(s string) bool {

	if s == "" {
		return false
	}
	for _, r := range s {
		if (r < '0' || r > '9') && r != ',' {
			return false
		}
	}
	return true
}

// 判断是不是中文
func CheckStringCHN(str string) bool {

	for _, r := range str {
		if !unicode.Is(unicode.Han, r) &&
			!isAlpha(r) && (r < '0' || r > '9') && r != '_' &&
			r != ' ' && r != '-' && r != '!' && r != '@' && r != ':' &&
			r != '?' && r != '+' && r != '.' && r != '/' && r != '\'' &&
			r != '(' && r != ')' && r != '·' && r != '&' {
			return false
		}
	}
	return true
}

// 判断是否module格式
func CheckStringModule(s string) bool {

	if s == "" {
		return false
	}

	for _, r := range s {
		if (r < 'A' || r > 'Z') && (r < 'a' || r > 'z') && r != '/' {
			return false
		}
	}

	return true
}

// 判断是否全英文字母
func CheckStringAlpha(s string) bool {

	if s == "" {
		return false
	}

	for _, r := range s {
		if (r < 'A' || r > 'Z') && (r < 'a' || r > 'z') && r != ' ' {
			return false
		}
	}

	return true
}

// 判断是否全英文字母+逗号
func CheckStringCommaAlpha(s string) bool {

	if s == "" {
		return false
	}

	for _, r := range s {
		if (r < 'A' || r > 'Z') && (r < 'a' || r > 'z') && r != ',' {
			return false
		}
	}

	return true
}

// 判断是否全为英文字母和数字组合
func CheckStringAlnum(s string) bool {

	if s == "" {
		return false
	}
	for _, r := range s {
		if !isDigit(r) && !isAlpha(r) &&
			r != ' ' && r != '-' && r != '!' && r != '_' &&
			r != '@' && r != '?' && r != '+' && r != ':' &&
			r != '.' && r != '/' && r != '(' && r != '\'' &&
			r != ')' && r != '·' && r != '&' {
			return false
		}
	}
	return true
}

// 检查日期格式"YYYY-MM-DD"
func CheckDate(str string) bool {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	_, err := time.ParseInLocation("2006-01-02", str, loc)
	if err != nil {
		return false
	}
	return true
}

// 匹配时间 "HH:ii" or "HH:ii:ss"
func checkTime(str string) bool {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	_, err := time.ParseInLocation("15:04:05", str, loc)
	if err != nil {
		return false
	}
	return true
}

// 检查日期时间格式"YYYY-MM-DD HH:ii:ss"
func CheckDateTime(str string) bool {
	loc, _ := time.LoadLocation("Asia/Shanghai")
	_, err := time.ParseInLocation("2006-01-02 15:04:05", str, loc)
	if err != nil {
		return false
	}
	return true
}

func CheckMoney(money string) bool {

	// 金额小数验证
	_, errm := strconv.Atoi(money)
	if errm != nil {
		return false
	}
	_, errm = strconv.ParseFloat(money, 64)
	if errm != nil {
		return false
	}
	return true
}

// 判断字符串是不是数字
func CtypeDigit(s string) bool {

	if s == "" {
		return false
	}
	for _, r := range s {
		if !isDigit(r) {
			return false
		}
	}
	return true
}

// 判断字符串是不是字母+数字
func CtypeAlnum(s string) bool {

	if s == "" {
		return false
	}
	for _, r := range s {
		if !isDigit(r) && !isAlpha(r) {
			return false
		}
	}
	return true
}

// 检查url
func CheckUrl(s string) bool {
	return govalidator.IsURL(s)
}

func checkOdds(str string) bool {
	odds := strings.Split(str, ",")
	for _, r := range odds {
		if !checkFloat(r) {
			return false
		}
	}
	return true
}
