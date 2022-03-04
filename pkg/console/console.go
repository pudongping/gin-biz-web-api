// 命令行辅助方法
package console

import (
	"fmt"
	"os"

	"github.com/mgutz/ansi"
)

// Success 打印一条成功消息，绿色输出
func Success(format string, v ...interface{}) {
	colorOut("green", format, v...)
}

// Error 打印一条报错消息，红色输出
func Error(format string, v ...interface{}) {
	colorOut("red", format, v...)
}

// Warning 打印一条提示消息，黄色输出
func Warning(format string, v ...interface{}) {
	colorOut("yellow", format, v...)
}

// Info 打印一条普通消息，蓝色输出
func Info(format string, v ...interface{}) {
	colorOut("blue", format, v...)
}

// Exit 打印一条报错消息，并退出 os.Exit(1)
func Exit(format string, v ...interface{}) {
	Error(format, v...)
	os.Exit(1)
}

// ExitIf 语法糖，自带 err != nil 判断
func ExitIf(err error) {
	if err != nil {
		Exit(err.Error())
	}
}

// colorOut 内部使用，设置高亮颜色
func colorOut(color, message string, v ...interface{}) {
	_, _ = fmt.Fprintln(os.Stdout, ansi.Color(fmt.Sprintf(message, v...), color))
}
