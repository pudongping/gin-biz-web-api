package arrayx

import (
	"math"
)

// ArrayChunkString 将一个数组分割成多个
// s := []string{"a1", "a2", "a3", "a4", "a5", "a6", "a7"}
// size := 2
// output: [[a1 a2] [a3 a4] [a5 a6] [a7]]
func ArrayChunkString(s []string, size int) [][]string {
	if size < 1 {
		panic("size: cannot be less than 1")
	}
	length := len(s)
	chunks := int(math.Ceil(float64(length) / float64(size)))
	var n [][]string
	for i, end := 0, 0; chunks > 0; chunks-- {
		end = (i + 1) * size
		if end > length {
			end = length
		}
		n = append(n, s[i*size:end])
		i++
	}
	return n
}
