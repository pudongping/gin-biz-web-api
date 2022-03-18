// 日志处理相关逻辑
package logger

import (
	"encoding/json"
	"os"
	"strings"
	"time"

	"go.uber.org/zap"
	"go.uber.org/zap/zapcore"
	"gopkg.in/natefinch/lumberjack.v2"

	"gin-biz-web-api/pkg/app"
	"gin-biz-web-api/pkg/console"
)

// Logger 全局 Logger 对象
var Logger *zap.Logger

// InitLogger 日志初始化
func InitLogger(filename string, maxSize, maxBackup, maxAge int, compress bool, logType string, level string) {
	// 获取日志写入介质
	writeSyncer := getLogWriter(filename, maxSize, maxBackup, maxAge, compress, logType)

	// 设置日志等级，具体请见 config/log.go 配置文件
	logLevel := new(zapcore.Level)
	if err := logLevel.UnmarshalText([]byte(level)); err != nil {
		console.Exit("日志初始化错误，日志级别设置有误。请修改 config/log.go 文件中的 log.level 配置项")
	}

	// 初始化 core
	core := zapcore.NewCore(getEncoder(), writeSyncer, logLevel)

	// 初始化 Logger
	Logger = zap.New(
		core,
		zap.AddCaller(),                   // 调用文件和行号，内部使用 runtime.Caller
		zap.AddCallerSkip(1),              // 封装了一层，调用文件去除一层(runtime.Caller(1))
		zap.AddStacktrace(zap.ErrorLevel), // 输出调用堆栈，Error 时才会显示 stacktrace
	)

	// 将自定义的 logger 替换为全局的 logger
	// zap.L() 相当于 *zap.Logger
	// eg：以下二者打印出来的内容完全一致
	// zap.L().Debug("Cache", zap.String("Flush", "danger!!!"), zap.Time("time", time.Now()))
	// logger.Debug("Cache", zap.String("Flush", "danger!!!"), zap.Time("time", time.Now()))
	zap.ReplaceGlobals(Logger)
}

// getEncoder 设置日志存储格式
func getEncoder() zapcore.Encoder {

	// 日志格式规则
	encoderConfig := zapcore.EncoderConfig{
		TimeKey:        "time",
		LevelKey:       "level", // 日志中级别的键名，默认为 level
		NameKey:        "logger",
		CallerKey:      "caller", // 代码调用输出文件名和行号，如 internal/controller/user_controller.go:88
		FunctionKey:    zapcore.OmitKey,
		MessageKey:     "message", // 日志中信息的键名，默认为 msg
		StacktraceKey:  "stacktrace",
		LineEnding:     zapcore.DefaultLineEnding,      // 每行日志的结尾添加 "\n"
		EncodeLevel:    zapcore.CapitalLevelEncoder,    // 日志中级别的格式，默认为小写，如 debug/info，这里采用日志级别名称大写，如 ERROR、INFO
		EncodeTime:     customTimeEncoder,              // 时间格式，我们自定义为 2006-01-02 15:04:05
		EncodeDuration: zapcore.SecondsDurationEncoder, // 执行时间，以秒为单位
		EncodeCaller:   zapcore.ShortCallerEncoder,     // Caller 短格式，如：internal/dao/user_dao.go:66，长格式为绝对路径
	}

	// debug 模式时的配置
	if app.IsDebug() {
		// 终端输出的关键词高亮
		encoderConfig.EncodeLevel = zapcore.CapitalColorLevelEncoder
		// 本地设置内置的 Console 解码器（支持 stacktrace 换行）
		return zapcore.NewConsoleEncoder(encoderConfig)
	}

	// 线上环境使用 JSON 编码器
	return zapcore.NewJSONEncoder(encoderConfig)

}

// customTimeEncoder 自定义友好的时间格式
func customTimeEncoder(t time.Time, enc zapcore.PrimitiveArrayEncoder) {
	enc.AppendString(t.Format("2006-01-02 15:04:05"))
}

// getLogWriter 日志记录介质（提供了标准输出和文件输出）
func getLogWriter(filename string, maxSize, maxBackup, maxAge int, compress bool, logType string) zapcore.WriteSyncer {

	// 如果配置了按照日期的方式记录日志文件
	if logType == "daily" {
		logName := time.Now().Format("2006-01-02.log")
		filename = strings.ReplaceAll(filename, "logs.log", logName)
	}

	// 滚动日志，详见 config/log.go 配置文件
	lumberJackLogger := &lumberjack.Logger{
		Filename:   filename,  // 文件路径
		MaxSize:    maxSize,   // 单个文件最大尺寸，默认单位为 M
		MaxAge:     maxAge,    // 最大时间，默认单位为：天 day
		MaxBackups: maxBackup, // 最多保存多少个备份文件
		Compress:   compress,  // 是否压缩
		LocalTime:  true,      // 使用本地时间
	}

	if app.IsDebug() {
		// debug 模式时，终端打印和记录文件
		return zapcore.NewMultiWriteSyncer(zapcore.AddSync(os.Stdout), zapcore.AddSync(lumberJackLogger))
	} else {
		// 非 debug 模式时，均只记录文件
		return zapcore.AddSync(lumberJackLogger)
	}

}

// Dump 方便调试时使用，会记录 warn 级别信息
// eg1：只有一个参数时
// 	logger.Dump(struct {
//		Name, Sex string
//		Age       int32
//	}{
//		Name: "alex",
//		Sex:  "m",
//		Age:  18,
//	})
//
// output：2022-03-18 01:18:19     WARN    cache/redis_driver.go:52        Dump    {"data": "{\"Name\":\"alex\",\"Sex\":\"m\",\"Age\":18}"}
// eg2：有两个参数时
// 	logger.Dump(struct {
//		Name, Sex string
//		Age       int32
//	}{
//		Name: "alex",
//		Sex:  "m",
//		Age:  18,
//	}, "个人信息")
// output：2022-03-18 01:20:43     WARN    cache/redis_driver.go:52        Dump    {"个人信息": "{\"Name\":\"alex\",\"Sex\":\"m\",\"Age\":18}"}
func Dump(value interface{}, msg ...string) {
	valueString := jsonString(value)
	// 判断第二个参数是否传参 msg
	if len(msg) > 0 {
		Logger.Warn("Dump", zap.String(msg[0], valueString))
	} else {
		Logger.Warn("Dump", zap.String("data", valueString))
	}
}

// LogErrorIf 当 err != nil 时记录 error 等级的日志（有调用堆栈信息）
// eg：logger.LogErrorIf(errors.New("没有权限"))
// output：2022-03-18 01:23:33     ERROR   cache/redis_driver.go:53        Error Occurred: {"error": "没有权限"}
func LogErrorIf(err error) {
	if err != nil {
		Logger.Error("Error Occurred:", zap.Error(err))
	}
}

// LogFatalIf 当 err != nil 时记录 fatal 等级的日志（有调用堆栈信息）
// eg：logger.LogFatalIf(errors.New("没有权限"))
// output：2022-03-18 01:23:33     FATAL   cache/redis_driver.go:53        Error Occurred: {"error": "没有权限"}
func LogFatalIf(err error) {
	if err != nil {
		Logger.Fatal("Error Occurred:", zap.Error(err))
	}
}

// LogWarnIf 当 err != nil 时记录 warning 等级的日志
// eg：logger.LogWarnIf(errors.New("没有权限"))
// output：2022-03-18 01:26:21     WARN    cache/redis_driver.go:53        Error Occurred: {"error": "没有权限"}
func LogWarnIf(err error) {
	if err != nil {
		Logger.Warn("Error Occurred:", zap.Error(err))
	}
}

// LogInfoIf 当 err != nil 时记录 info 等级的日志
// eg：logger.LogInfoIf(errors.New("没有权限"))
// output：2022-03-18 01:27:25     INFO    cache/redis_driver.go:53        Error Occurred: {"error": "没有权限"}
func LogInfoIf(err error) {
	if err != nil {
		Logger.Info("Error Occurred:", zap.Error(err))
	}
}

// Debug 调试类日志
// eg：logger.Debug("Cache", zap.String("Flush", "danger!!!"), zap.Time("time", time.Now().Local()))
// output：2022-03-18 00:14:29     DEBUG   cache/redis_driver.go:52        Cache   {"Flush": "danger!!!", "time": "2022-03-18 00:14:29"}
func Debug(moduleName string, fields ...zap.Field) {
	Logger.Debug(moduleName, fields...)
}

// Info 告知类日志
// eg：logger.Info("Cache", zap.String("Flush", "danger!!!"), zap.Time("time", time.Now()))
// output：2022-03-18 00:25:41     INFO    cache/redis_driver.go:52        Cache   {"Flush": "danger!!!", "time": "2022-03-18 00:25:41"}
func Info(moduleName string, fields ...zap.Field) {
	Logger.Info(moduleName, fields...)
}

// Warn 警告类
// eg：logger.Warn("Cache", zap.String("Flush", "danger!!!"), zap.Time("time", time.Now()))
// output：2022-03-18 00:27:08     WARN    cache/redis_driver.go:52        Cache   {"Flush": "danger!!!", "time": "2022-03-18 00:27:08"}
func Warn(moduleName string, fields ...zap.Field) {
	Logger.Warn(moduleName, fields...)
}

// Error 错误时记录（会打印出调用堆栈，但是不会退出程序）
// eg：logger.Error("Cache", zap.String("Flush", "danger!!!"), zap.Time("time", time.Now()))
// output：2022-03-18 00:28:36     ERROR   cache/redis_driver.go:52        Cache   {"Flush": "danger!!!", "time": "2022-03-18 00:28:36"}
func Error(moduleName string, fields ...zap.Field) {
	Logger.Error(moduleName, fields...)
}

// Fatal 致命时记录（会打印出调用堆栈，写完 log 后调用 os.Exit(1) 直接退出程序）
// eg：logger.Fatal("Cache", zap.String("Flush", "danger!!!"), zap.Time("time", time.Now()))
// output：2022-03-18 00:32:42     FATAL   cache/redis_driver.go:52        Cache   {"Flush": "danger!!!", "time": "2022-03-18 00:32:42"}
func Fatal(moduleName string, fields ...zap.Field) {
	Logger.Fatal(moduleName, fields...)
}

// DebugString
// eg：logger.DebugString("Cache", "Flush", "danger!!!")
// output：2022-03-18 00:38:09     DEBUG   cache/redis_driver.go:51        Cache   {"Flush": "danger!!!"}
func DebugString(moduleName, name, msg string) {
	Logger.Debug(moduleName, zap.String(name, msg))
}

// InfoString
// eg：logger.InfoString("Cache", "Flush", "danger!!!")
// output：2022-03-18 00:40:19     INFO    cache/redis_driver.go:51        Cache   {"Flush": "danger!!!"}
func InfoString(moduleName, name, msg string) {
	Logger.Info(moduleName, zap.String(name, msg))
}

// WarnString
// eg：logger.WarnString("Cache", "Flush", "danger!!!")
// output：2022-03-18 00:41:39     WARN    cache/redis_driver.go:51        Cache   {"Flush": "danger!!!"}
func WarnString(moduleName, name, msg string) {
	Logger.Warn(moduleName, zap.String(name, msg))
}

// ErrorString 错误时记录（会打印出调用堆栈，但是不会退出程序）
// eg：logger.ErrorString("Cache", "Flush", "danger!!!")
// output：2022-03-18 00:42:57     ERROR   cache/redis_driver.go:51        Cache   {"Flush": "danger!!!"}
func ErrorString(moduleName, name, msg string) {
	Logger.Error(moduleName, zap.String(name, msg))
}

// FatalString 记录致命错误日志（会打印出调用堆栈，写完 log 后调用 os.Exit(1) 直接退出程序）
// eg：logger.FatalString("Cache", "Flush", "danger!!!")
// output：2022-03-18 00:44:39     FATAL   cache/redis_driver.go:51        Cache   {"Flush": "danger!!!"}
func FatalString(moduleName, name, msg string) {
	Logger.Fatal(moduleName, zap.String(name, msg))
}

// DebugJSON
// eg1：logger.DebugJSON("Cache", "Flush", map[string][]string{"boys": {"alex", "bob"}, "sex":  {"f", "m"}})
// output：2022-03-18 00:52:32     DEBUG   cache/redis_driver.go:51        Cache   {"Flush": "{\"boys\":[\"alex\",\"bob\"],\"sex\":[\"f\",\"m\"]}"}
//
// eg2：
// 	logger.DebugJSON("Cache", "Flush", struct {
//		Name, Sex string
//		Age       int32
//	}{
//		Name: "alex",
//		Sex:  "m",
//		Age:  18,
//	})
// output：2022-03-18 01:10:10     DEBUG   cache/redis_driver.go:52        Cache   {"Flush": "{\"Name\":\"alex\",\"Sex\":\"m\",\"Age\":18}"}
func DebugJSON(moduleName, name string, value interface{}) {
	Logger.Debug(moduleName, zap.String(name, jsonString(value)))
}

// InfoJSON
// eg：logger.InfoJSON("Cache", "Flush", map[string][]string{"boys": {"alex", "bob"}, "sex":  {"f", "m"}})
// output：2022-03-18 00:54:39     INFO    cache/redis_driver.go:51        Cache   {"Flush": "{\"boys\":[\"alex\",\"bob\"],\"sex\":[\"f\",\"m\"]}"}
func InfoJSON(moduleName, name string, value interface{}) {
	Logger.Info(moduleName, zap.String(name, jsonString(value)))
}

// WarnJSON
// eg：logger.WarnJSON("Cache", "Flush", map[string][]string{"boys": {"alex", "bob"}, "sex":  {"f", "m"}})
// output：2022-03-18 00:55:38     WARN    cache/redis_driver.go:51        Cache   {"Flush": "{\"boys\":[\"alex\",\"bob\"],\"sex\":[\"f\",\"m\"]}"}
func WarnJSON(moduleName, name string, value interface{}) {
	Logger.Warn(moduleName, zap.String(name, jsonString(value)))
}

// ErrorJSON 错误时记录（会打印出调用堆栈，但是不会退出程序）
// eg：logger.ErrorJSON("Cache", "Flush", map[string][]string{"boys": {"alex", "bob"}, "sex":  {"f", "m"}})
// output：2022-03-18 00:56:43     ERROR   cache/redis_driver.go:51        Cache   {"Flush": "{\"boys\":[\"alex\",\"bob\"],\"sex\":[\"f\",\"m\"]}"}
func ErrorJSON(moduleName, name string, value interface{}) {
	Logger.Error(moduleName, zap.String(name, jsonString(value)))
}

// FatalJSON 记录致命错误日志（会打印出调用堆栈，写完 log 后调用 os.Exit(1) 直接退出程序）
// eg：logger.FatalJSON("Cache", "Flush", map[string][]string{"boys": {"alex", "bob"}, "sex":  {"f", "m"}})
// output：2022-03-18 00:58:06     FATAL   cache/redis_driver.go:51        Cache   {"Flush": "{\"boys\":[\"alex\",\"bob\"],\"sex\":[\"f\",\"m\"]}"}
func FatalJSON(moduleName, name string, value interface{}) {
	Logger.Fatal(moduleName, zap.String(name, jsonString(value)))
}

// jsonString 将数据转成 json 字符串
func jsonString(value interface{}) string {
	b, err := json.Marshal(value)
	if err != nil {
		Logger.Error("Logger", zap.String("JSON marshal error", err.Error()))
	}
	return string(b)
}
