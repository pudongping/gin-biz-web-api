// 取子元素
package arrayx

// ArrayFirstElementString 取切片中的第一个元素
func ArrayFirstElementString(args []string) string {
	if len(args) > 0 {
		return args[0]
	}

	return ""
}

// ArrayLastElementString 取切片中的最后一个元素
func ArrayLastElementString(args []string) string {
	l := len(args)
	if l > 0 {
		return args[l-1]
	}

	return ""
}
