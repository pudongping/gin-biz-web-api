package validator

import (
	"github.com/pkg/errors"

	"strconv"
	"strings"
	"unicode/utf8"

	"github.com/thedevsaddam/govalidator"
)

// 中文字符长度验证
func init() {

	// min_cn:2
	govalidator.AddCustomRule("min_cn", func(field string, rule string, message string, value interface{}) error {
		relVal, is := value.(string)
		if !is {
			// 如果当前值不为字符串，那么则不需要去做校验
			return nil
		}

		mustLen := strings.TrimPrefix(rule, "min_cn:")
		l, err := strconv.Atoi(mustLen)
		if err != nil {
			return errors.New("字符串转整数失败")
		}

		valLength := utf8.RuneCountInString(relVal)
		if valLength < l {
			if message != "" {
				return errors.New(message)
			}
			return errors.Errorf("长度需大于 %d 个字符", l)
		}

		return nil
	})

	// max_cn:8
	govalidator.AddCustomRule("max_cn", func(field string, rule string, message string, value interface{}) error {
		relVal, is := value.(string)
		if !is {
			// 如果当前值不为字符串，那么则不需要去做校验
			return nil
		}

		l, err := strconv.Atoi(strings.TrimPrefix(rule, "max_cn:"))
		if err != nil {
			return errors.New("字符串转整数失败")
		}

		valLength := utf8.RuneCountInString(relVal)
		if valLength > l {
			if message != "" {
				return errors.New(message)
			}
			return errors.Errorf("长度不能超过 %d 个字符", l)
		}

		return nil
	})

}
