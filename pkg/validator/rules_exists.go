// 自定义验证规则
package validator

import (
	"errors"
	"fmt"
	"strings"

	"github.com/thedevsaddam/govalidator"

	"gin-biz-web-api/pkg/database"
)

// 注册自定义表单验证规则
// 数据库中是否存在某些数据验证
func init() {

	// 自定义规则 exists，确保数据库存在某条数据
	// eg: exists:users,id
	govalidator.AddCustomRule("exists", func(field string, rule string, message string, value interface{}) error {
		sl := strings.Split(strings.TrimPrefix(rule, "exists:"), ",")

		// 第一个参数，表名称，如 users
		tableName := sl[0]
		// 第二个参数，字段名称，如 id
		dbFiled := sl[1]

		// 查询数据库
		var count int64
		database.DB.Table(tableName).Where(fmt.Sprintf("`%s` = ?", dbFiled), value).Count(&count)
		// 验证不通过，数据不存在
		if count == 0 {
			// 如果有自定义错误消息的话，使用自定义消息
			if message != "" {
				return errors.New(message)
			}
			return fmt.Errorf("%v 不存在", value)
		}
		return nil
	})

	// 自定义规则 not_exists，验证请求数据必须不存在于数据库中
	// 常用于保证数据库某个字段的值唯一，如用户名、邮箱、手机号等
	// not_exists 参数可以有两种，一种是 2 个参数，一种是 3 个参数：
	// not_exists:users,email 检查数据库表里是否存在同一条信息
	// not_exists:users,email,8 排除用户掉 id 为 8 的用户
	govalidator.AddCustomRule("not_exists", func(field string, rule string, message string, value interface{}) error {
		sl := strings.Split(strings.TrimPrefix(rule, "not_exists:"), ",")

		// 第一个参数，表名称，如 users
		tableName := sl[0]
		// 第二个参数，字段名称，如 email 或者 phone
		dbFiled := sl[1]

		// 第三个参数，排除 ID
		var exceptID string
		if len(sl) > 2 {
			exceptID = sl[2]
		}

		// 拼接 SQL
		query := database.DB.Table(tableName).Where(fmt.Sprintf("`%s` = ?", dbFiled), value)

		// 如果传参第三个参数，加上 SQL Where 过滤
		if exceptID != "" {
			query.Where("id != ?", exceptID)
		}

		// 查询数据库
		var count int64
		query.Count(&count)

		// 验证不通过，数据库能找到对应的数据
		if count != 0 {
			// 如果有自定义错误消息的话
			if message != "" {
				return errors.New(message)
			}
			// 默认的错误消息
			return fmt.Errorf("%v 已被占用", value)
		}
		// 验证通过
		return nil
	})

}
