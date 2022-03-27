package arrayx

import (
	"fmt"
	"strings"
)

// Array2Str 将切片转成以指定分隔符分隔的字符串
func Array2Str(slice []string, delimiter string) string {
	return strings.ReplaceAll(strings.Trim(fmt.Sprint(slice), "[]"), " ", delimiter)
}
