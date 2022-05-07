package validator

import (
	"github.com/pkg/errors"

	"fmt"
	"reflect"
	"strings"

	"github.com/thedevsaddam/govalidator"

	"gin-biz-web-api/pkg/database"
	"gin-biz-web-api/pkg/helper"
)

func init() {

	// 自定义验证规则 required_all_in_db ，确保传入的数组值均存在数据库中
	// eg：required_all_in_db:users,user_id
	// 表示传入的值必须都是 users 表中的 user_id
	govalidator.AddCustomRule("required_all_in_db", func(field string, rule string, message string, value interface{}) error {

		// 如果为【零值】则不做校验
		if helper.Empty(value) {
			return nil
		}

		var valueLen int
		// 需要验证的值必须为数组类型，否则不做校验
		v := reflect.ValueOf(value)
		if v.Kind() == reflect.Array || v.Kind() == reflect.Slice {
			valueLen = v.Len()
		} else {
			return nil
		}

		sl := strings.Split(strings.TrimPrefix(rule, "required_all_in_db:"), ",")
		if len(sl) < 1 {
			return errors.New("required_all_in_db 规则至少存在一个参数")
		}

		// 第一个参数，表名称，如：users
		tableName := sl[0]

		f := field
		// 如果超过一个参数时，那么第二个参数则为字段名
		if len(sl) >= 2 {
			f = sl[1]
		}

		var count int64
		database.DB.Table(tableName).Where(fmt.Sprintf("`%s` in ?", f), value).Count(&count)
		if int64(valueLen) != count {
			// 如果有自定义错误消息的话，使用自定义消息
			if message != "" {
				return errors.New(message)
			}
			return errors.Errorf("字段 %s 的值不合法", field)
		}

		return nil
	})

}
