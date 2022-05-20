package arrayx

// InArrayString 判断某个值是否在切片中存在
// 类似于 PHP in_array 函数
// needle := "a"
// haystack := []string{"a", "b", "c"}
// output: true
func InArrayString(needle string, haystack []string) bool {
	if len(haystack) == 0 {
		return false
	}

	for i := 0; i < len(haystack); i++ {
		if haystack[i] == needle {
			return true
		}
	}

	return false
}
