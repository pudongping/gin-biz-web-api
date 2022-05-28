package arrayx

// ArrayUniqueString 数组去重
// arr := []string{"a1", "a1", "b1", "c1"}
// ArrayUniqueString(arr) ==> output: []string{"a1", "b1", "c1"}
func ArrayUniqueString(arr []string) []string {
	size := len(arr)
	result := make([]string, 0, size)
	temp := map[string]struct{}{}
	for i := 0; i < size; i++ {
		if _, ok := temp[arr[i]]; !ok {
			temp[arr[i]] = struct{}{}
			result = append(result, arr[i])
		}
	}
	return result
}
