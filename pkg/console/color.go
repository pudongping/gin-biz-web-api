package console

import (
	"fmt"
)

// 前景 背景 颜色
// ---------------------------------------
// 30  40  黑色
// 31  41  红色
// 32  42  绿色
// 33  43  橙黄色
// 34  44  蓝色
// 35  45  紫红色（品红色）
// 36  46  青蓝色
// 37  47  浅灰色
// 39  49  系统发行版默认值
//
// 代码 输出样式
// -------------------------
//  0  终端默认设置
//  1  粗体
//  4  使用下划线
//  5  闪烁
//  7  反白显示
//  8  暗色

// fmt.Printf("\033[1;31;44m%s\033[0m\n","Hello Alex!")
// 以上表示：粗体、蓝底红字输出 "Hello Alex!"

// 前景色的取值范围，也就是文字的颜色取值范围
const (
	TextBlack = iota + 30
	TextRed
	TextGreen
	TextYellow
	TextBlue
	TextMagenta
	TextCyan
	TextWhite
	TextDefault = 39
)

const (
	BgBlack = iota + 40
	BgRed
	BgGreen
	BgYellow
	BgBlue
	BgMagenta
	BgCyan
	BgWhite
	BgDefault = 49
)

const (
	OutStyleDefault   = 0
	OutStyleBold      = 1
	OutStyleUnderline = 4
	OutStyleHighlight = 5
	OutStyleFlip      = 7
	OutStyleGray      = 8
)

// Black 黑色
func Black(msg string) string {
	return SetColor(msg, OutStyleDefault, TextBlack, BgDefault)
}

// Red 红色
func Red(msg string) string {
	return SetColor(msg, OutStyleDefault, TextRed, BgDefault)
}

// Green 绿色
func Green(msg string) string {
	return SetColor(msg, OutStyleDefault, TextGreen, BgDefault)
}

// Yellow 橙黄色
func Yellow(msg string) string {
	return SetColor(msg, OutStyleDefault, TextYellow, BgDefault)
}

// Blue 蓝色
func Blue(msg string) string {
	return SetColor(msg, OutStyleDefault, TextBlue, BgDefault)
}

// Magenta 品红色
func Magenta(msg string) string {
	return SetColor(msg, OutStyleDefault, TextMagenta, BgDefault)
}

// Cyan 青蓝色
func Cyan(msg string) string {
	return SetColor(msg, OutStyleDefault, TextCyan, BgDefault)
}

// White 浅灰色
func White(msg string) string {
	return SetColor(msg, OutStyleDefault, TextWhite, BgDefault)
}

func SetColor(msg string, cfg, fg, bg int) string {
	// 0x1B 标记变换颜色的起始标记
	// cfg 输出方式
	// fg 前景色（文字颜色）
	// bg 背景色
	// msg 文本内容
	// 0m 表示取消颜色设置
	return fmt.Sprintf("%c[%d;%d;%dm%s%c[0m", 0x1B, cfg, fg, bg, msg, 0x1B)
}
