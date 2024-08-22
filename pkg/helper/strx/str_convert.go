// 字符串转换
package strx

import (
	"github.com/gertd/go-pluralize"
	"github.com/iancoleman/strcase"
)

// StrPlural 单词单数转复数 eg: user -> users
func StrPlural(s string) string {
	return pluralize.NewClient().Plural(s)
}

// StrSingular 单词复数转单数 eg: users -> user
func StrSingular(s string) string {
	return pluralize.NewClient().Singular(s)
}

// StrSnake 驼峰转蛇形命名 eg: UserAvatar -> user_avatar
func StrSnake(s string) string {
	return strcase.ToSnake(s)
}

// StrCamel 蛇形命名转大驼峰 eg: user_avatar -> UserAvatar
func StrCamel(s string) string {
	return strcase.ToCamel(s)
}

// StrLowerCamel 大驼峰转小驼峰 eg: UserAvatar -> userAvatar
func StrLowerCamel(s string) string {
	return strcase.ToLowerCamel(s)
}

// StrKebab 单词以横杠分割 eg: AnyKind of_string -> any-kind-of-string
func StrKebab(s string) string {
	return strcase.ToKebab(s)
}
