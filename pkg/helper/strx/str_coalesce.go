package strx

import (
	"strings"
)

// StrBuilder 高性能字符串拼接
// eg：
// str := strBuilder("11", "+", "22", "=", "33")
// output： "11+22=33"
func StrBuilder(str ...string) string {
	if len(str) == 0 {
		return ""
	}

	var s strings.Builder
	for _, v := range str {
		s.WriteString(v)
	}

	return s.String()
}
