package helper

import (
	"crypto/md5"
	"encoding/hex"
)

// EncodeMD5 md5 加密字符串
func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))

	return hex.EncodeToString(m.Sum(nil))
}
