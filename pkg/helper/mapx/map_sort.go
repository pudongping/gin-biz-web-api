// 通过 key 进行排序
package mapx

import (
	"sort"

	"github.com/spf13/cast"
)

// SortAscKeyString map 通过 key 进行正序排列
func SortAscKeyString(i interface{}) (s []string) {
	switch v := i.(type) {
	case map[string]int:
		for k := range v {
			s = append(s, k)
		}
	case map[string]int64:
		for k := range v {
			s = append(s, k)
		}
	case map[string]int32:
		for k := range v {
			s = append(s, k)
		}
	case map[string]int16:
		for k := range v {
			s = append(s, k)
		}
	case map[string]int8:
		for k := range v {
			s = append(s, k)
		}
	case map[string]uint:
		for k := range v {
			s = append(s, k)
		}
	case map[string]uint64:
		for k := range v {
			s = append(s, k)
		}
	case map[string]uint32:
		for k := range v {
			s = append(s, k)
		}
	case map[string]uint16:
		for k := range v {
			s = append(s, k)
		}
	case map[string]uint8:
		for k := range v {
			s = append(s, k)
		}
	case map[string]float64:
		for k := range v {
			s = append(s, k)
		}
	case map[string]float32:
		for k := range v {
			s = append(s, k)
		}
	case map[string]bool:
		for k := range v {
			s = append(s, k)
		}
	case map[string]string:
		for k := range v {
			s = append(s, k)
		}
	case map[string][]string:
		for k := range v {
			s = append(s, k)
		}
	case map[string]interface{}:
		for k := range v {
			s = append(s, k)
		}
	case map[interface{}]interface{}:
		for k := range v {
			s = append(s, cast.ToString(k))
		}
	}
	sort.Strings(s)
	return
}

// SortDescKeyString map 通过 key 进行倒序排列
func SortDescKeyString(i interface{}) (s []string) {
	switch v := i.(type) {
	case map[string]int:
		for k := range v {
			s = append(s, k)
		}
	case map[string]int64:
		for k := range v {
			s = append(s, k)
		}
	case map[string]int32:
		for k := range v {
			s = append(s, k)
		}
	case map[string]int16:
		for k := range v {
			s = append(s, k)
		}
	case map[string]int8:
		for k := range v {
			s = append(s, k)
		}
	case map[string]uint:
		for k := range v {
			s = append(s, k)
		}
	case map[string]uint64:
		for k := range v {
			s = append(s, k)
		}
	case map[string]uint32:
		for k := range v {
			s = append(s, k)
		}
	case map[string]uint16:
		for k := range v {
			s = append(s, k)
		}
	case map[string]uint8:
		for k := range v {
			s = append(s, k)
		}
	case map[string]float64:
		for k := range v {
			s = append(s, k)
		}
	case map[string]float32:
		for k := range v {
			s = append(s, k)
		}
	case map[string]bool:
		for k := range v {
			s = append(s, k)
		}
	case map[string]string:
		for k := range v {
			s = append(s, k)
		}
	case map[string][]string:
		for k := range v {
			s = append(s, k)
		}
	case map[string]interface{}:
		for k := range v {
			s = append(s, k)
		}
	case map[interface{}]interface{}:
		for k := range v {
			s = append(s, cast.ToString(k))
		}
	}
	sort.Sort(sort.Reverse(sort.StringSlice(s)))
	return
}

// SortAscKeyInt map 通过 key 进行正序排列
func SortAscKeyInt(i interface{}) (s []int) {
	switch v := i.(type) {
	case map[int]int:
		for k := range v {
			s = append(s, k)
		}
	case map[int]int64:
		for k := range v {
			s = append(s, k)
		}
	case map[int]int32:
		for k := range v {
			s = append(s, k)
		}
	case map[int]int16:
		for k := range v {
			s = append(s, k)
		}
	case map[int]int8:
		for k := range v {
			s = append(s, k)
		}
	case map[int]uint:
		for k := range v {
			s = append(s, k)
		}
	case map[int]uint64:
		for k := range v {
			s = append(s, k)
		}
	case map[int]uint32:
		for k := range v {
			s = append(s, k)
		}
	case map[int]uint16:
		for k := range v {
			s = append(s, k)
		}
	case map[int]uint8:
		for k := range v {
			s = append(s, k)
		}
	case map[int]float64:
		for k := range v {
			s = append(s, k)
		}
	case map[int]float32:
		for k := range v {
			s = append(s, k)
		}
	case map[int]bool:
		for k := range v {
			s = append(s, k)
		}
	case map[int]string:
		for k := range v {
			s = append(s, k)
		}
	case map[int][]string:
		for k := range v {
			s = append(s, k)
		}
	case map[int]interface{}:
		for k := range v {
			s = append(s, k)
		}
	case map[interface{}]interface{}:
		for k := range v {
			s = append(s, cast.ToInt(k))
		}
	}
	sort.Ints(s)
	return
}

// SortDescKeyInt map 通过 key 进行倒序排列
func SortDescKeyInt(i interface{}) (s []int) {
	switch v := i.(type) {
	case map[int]int:
		for k := range v {
			s = append(s, k)
		}
	case map[int]int64:
		for k := range v {
			s = append(s, k)
		}
	case map[int]int32:
		for k := range v {
			s = append(s, k)
		}
	case map[int]int16:
		for k := range v {
			s = append(s, k)
		}
	case map[int]int8:
		for k := range v {
			s = append(s, k)
		}
	case map[int]uint:
		for k := range v {
			s = append(s, k)
		}
	case map[int]uint64:
		for k := range v {
			s = append(s, k)
		}
	case map[int]uint32:
		for k := range v {
			s = append(s, k)
		}
	case map[int]uint16:
		for k := range v {
			s = append(s, k)
		}
	case map[int]uint8:
		for k := range v {
			s = append(s, k)
		}
	case map[int]float64:
		for k := range v {
			s = append(s, k)
		}
	case map[int]float32:
		for k := range v {
			s = append(s, k)
		}
	case map[int]bool:
		for k := range v {
			s = append(s, k)
		}
	case map[int]string:
		for k := range v {
			s = append(s, k)
		}
	case map[int][]string:
		for k := range v {
			s = append(s, k)
		}
	case map[int]interface{}:
		for k := range v {
			s = append(s, k)
		}
	case map[interface{}]interface{}:
		for k := range v {
			s = append(s, cast.ToInt(k))
		}
	}
	sort.Sort(sort.Reverse(sort.IntSlice(s)))
	return
}

// SortAscKeyFloat64 map 通过 key 进行正序排列
func SortAscKeyFloat64(i interface{}) (s []float64) {
	switch v := i.(type) {
	case map[float64]int:
		for k := range v {
			s = append(s, k)
		}
	case map[float64]int64:
		for k := range v {
			s = append(s, k)
		}
	case map[float64]int32:
		for k := range v {
			s = append(s, k)
		}
	case map[float64]int16:
		for k := range v {
			s = append(s, k)
		}
	case map[float64]int8:
		for k := range v {
			s = append(s, k)
		}
	case map[float64]uint:
		for k := range v {
			s = append(s, k)
		}
	case map[float64]uint64:
		for k := range v {
			s = append(s, k)
		}
	case map[float64]uint32:
		for k := range v {
			s = append(s, k)
		}
	case map[float64]uint16:
		for k := range v {
			s = append(s, k)
		}
	case map[float64]uint8:
		for k := range v {
			s = append(s, k)
		}
	case map[float64]float64:
		for k := range v {
			s = append(s, k)
		}
	case map[float64]float32:
		for k := range v {
			s = append(s, k)
		}
	case map[float64]bool:
		for k := range v {
			s = append(s, k)
		}
	case map[float64]string:
		for k := range v {
			s = append(s, k)
		}
	case map[float64][]string:
		for k := range v {
			s = append(s, k)
		}
	case map[float64]interface{}:
		for k := range v {
			s = append(s, k)
		}
	case map[interface{}]interface{}:
		for k := range v {
			s = append(s, cast.ToFloat64(k))
		}
	}
	sort.Float64s(s)
	return
}

// SortDescKeyFloat64 map 通过 key 进行倒序排列
func SortDescKeyFloat64(i interface{}) (s []float64) {
	switch v := i.(type) {
	case map[float64]int:
		for k := range v {
			s = append(s, k)
		}
	case map[float64]int64:
		for k := range v {
			s = append(s, k)
		}
	case map[float64]int32:
		for k := range v {
			s = append(s, k)
		}
	case map[float64]int16:
		for k := range v {
			s = append(s, k)
		}
	case map[float64]int8:
		for k := range v {
			s = append(s, k)
		}
	case map[float64]uint:
		for k := range v {
			s = append(s, k)
		}
	case map[float64]uint64:
		for k := range v {
			s = append(s, k)
		}
	case map[float64]uint32:
		for k := range v {
			s = append(s, k)
		}
	case map[float64]uint16:
		for k := range v {
			s = append(s, k)
		}
	case map[float64]uint8:
		for k := range v {
			s = append(s, k)
		}
	case map[float64]float64:
		for k := range v {
			s = append(s, k)
		}
	case map[float64]float32:
		for k := range v {
			s = append(s, k)
		}
	case map[float64]bool:
		for k := range v {
			s = append(s, k)
		}
	case map[float64]string:
		for k := range v {
			s = append(s, k)
		}
	case map[float64][]string:
		for k := range v {
			s = append(s, k)
		}
	case map[float64]interface{}:
		for k := range v {
			s = append(s, k)
		}
	case map[interface{}]interface{}:
		for k := range v {
			s = append(s, cast.ToFloat64(k))
		}
	}
	sort.Sort(sort.Reverse(sort.Float64Slice(s)))
	return
}
