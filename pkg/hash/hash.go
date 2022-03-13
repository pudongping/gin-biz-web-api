// hash 明文密码加密
package hash

import (
	"golang.org/x/crypto/bcrypt"

	"gin-biz-web-api/pkg/logger"
)

// BcryptHash 使用 bcrypt 对密码进行加密
func BcryptHash(password string) string {
	// GenerateFromPassword 的第二个参数是 cost 值。建议大于 12，数值越大耗费时间越长，不过最大值为 31
	bytes, err := bcrypt.GenerateFromPassword([]byte(password), 14)
	logger.LogErrorIf(err)

	return string(bytes)
}

// BcryptCheck 对比明文密码和数据库的哈希值是否一致
func BcryptCheck(password, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(password))
	logger.LogErrorIf(err)
	return err == nil
}

// BcryptIsHashed 判断字符串是否是哈希过的数据
func BcryptIsHashed(str string) bool {
	// bcrypt 加密后的长度等于 60
	return len(str) == 60
}
