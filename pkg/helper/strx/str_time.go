// 时间转字符串
// 1秒(s) ＝1000毫秒(ms)
// 1毫秒(ms)＝1000微秒 (us) ==> Milliseconds ==> 毫秒
// 1微秒(us)＝1000纳秒 (ns)  ==> Microseconds  ==> 微秒
// 1纳秒(ns)＝1000皮秒 (ps)  ==> Nanoseconds  ==> 纳秒
package strx

import (
	"fmt"
	"time"
)

// StrMicroseconds 将 time.Duration 类型（time.Duration 类型以纳秒[nano seconds] 为单位）
// 输出为小数点后 3 位的毫秒[ms] 为单位
func StrMicroseconds(elapsed time.Duration) string {
	return fmt.Sprintf("%.3fms", float64(elapsed.Milliseconds()))
}
