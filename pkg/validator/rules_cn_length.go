package validator

import (
	"github.com/pkg/errors"

	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/thedevsaddam/govalidator"
)

// 注册自定义表单验证规则
// 中文字符长度验证
func init() {

	// min_cn:2
	govalidator.AddCustomRule("min_cn", func(field string, rule string, message string, value interface{}) error {
		valLength := utf8.RuneCountInString(value.(string))
		l, _ := strconv.Atoi(strings.TrimPrefix(rule, "min_cn:"))
		if valLength < l {
			if message != "" {
				return errors.New(message)
			}
			return errors.Errorf("长度需大于 %d 个字", l)
		}

		return nil
	})

	// max_cn:8
	govalidator.AddCustomRule("max_cn", func(field string, rule string, message string, value interface{}) error {
		valLength := utf8.RuneCountInString(value.(string))
		l, _ := strconv.Atoi(strings.TrimPrefix(rule, "max_cn:"))
		if valLength > l {
			if message != "" {
				return errors.New(message)
			}
			return errors.Errorf("长度不能超过 %d 个字", l)
		}

		return nil
	})

}
