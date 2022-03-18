# 日志

> 使用了 [zap](https://github.com/gin-contrib/zap) 包和 [lumberjack.v2](https://gopkg.in/natefinch/lumberjack.v2) 作为日志驱动，更多详情请见：`pkg/logger/logger.go` 文件。
> 
> 配置文件详见：`config/log.go`

## 使用封装好的方法

- 使用 `Log*If` 系列方法时

> 当 `err != nil` 时才会记录日志文件

```go

logger.LogErrorIf(errors.New("没有权限"))
// output：2022-03-18 01:23:33     ERROR   cache/redis_driver.go:53        Error Occurred: {"error": "没有权限"}

logger.LogFatalIf(errors.New("没有权限"))
// output：2022-03-18 01:23:33     FATAL   cache/redis_driver.go:53        Error Occurred: {"error": "没有权限"}

logger.LogWarnIf(errors.New("没有权限"))
// output：2022-03-18 01:26:21     WARN    cache/redis_driver.go:53        Error Occurred: {"error": "没有权限"}

logger.LogInfoIf(errors.New("没有权限"))
// output：2022-03-18 01:27:25     INFO    cache/redis_driver.go:53        Error Occurred: {"error": "没有权限"}

```

- 使用日志级别系列方法时

```go

logger.Debug("Cache", zap.String("Flush", "danger!!!"), zap.Time("time", time.Now().Local()))
// output：2022-03-18 00:14:29     DEBUG   cache/redis_driver.go:52        Cache   {"Flush": "danger!!!", "time": "2022-03-18 00:14:29"}

logger.Info("Cache", zap.String("Flush", "danger!!!"), zap.Time("time", time.Now()))
// output：2022-03-18 00:25:41     INFO    cache/redis_driver.go:52        Cache   {"Flush": "danger!!!", "time": "2022-03-18 00:25:41"}

logger.Warn("Cache", zap.String("Flush", "danger!!!"), zap.Time("time", time.Now()))
// output：2022-03-18 00:27:08     WARN    cache/redis_driver.go:52        Cache   {"Flush": "danger!!!", "time": "2022-03-18 00:27:08"}

// 会打印出调用堆栈，但是不会退出程序
logger.Error("Cache", zap.String("Flush", "danger!!!"), zap.Time("time", time.Now()))
// output：2022-03-18 00:28:36     ERROR   cache/redis_driver.go:52        Cache   {"Flush": "danger!!!", "time": "2022-03-18 00:28:36"}

// 会打印出调用堆栈，写完 log 后调用 os.Exit(1) 直接退出程序
logger.Fatal("Cache", zap.String("Flush", "danger!!!"), zap.Time("time", time.Now()))
// output：2022-03-18 00:32:42     FATAL   cache/redis_driver.go:52        Cache   {"Flush": "danger!!!", "time": "2022-03-18 00:32:42"}

```

- 使用 `*String` 系列方法时

```go

logger.DebugString("Cache", "Flush", "danger!!!")
// output：2022-03-18 00:38:09     DEBUG   cache/redis_driver.go:51        Cache   {"Flush": "danger!!!"}

logger.InfoString("Cache", "Flush", "danger!!!")
// output：2022-03-18 00:40:19     INFO    cache/redis_driver.go:51        Cache   {"Flush": "danger!!!"}

logger.WarnString("Cache", "Flush", "danger!!!")
// output：2022-03-18 00:41:39     WARN    cache/redis_driver.go:51        Cache   {"Flush": "danger!!!"}

// 会打印出调用堆栈，但是不会退出程序
logger.ErrorString("Cache", "Flush", "danger!!!")
// output：2022-03-18 00:42:57     ERROR   cache/redis_driver.go:51        Cache   {"Flush": "danger!!!"}

// 会打印出调用堆栈，写完 log 后调用 os.Exit(1) 直接退出程序
logger.FatalString("Cache", "Flush", "danger!!!")
// output：2022-03-18 00:44:39     FATAL   cache/redis_driver.go:51        Cache   {"Flush": "danger!!!"}

```

- 使用 `*JSON` 系列方法时

```go

logger.DebugJSON("Cache", "Flush", map[string][]string{"boys": {"alex", "bob"}, "sex":  {"f", "m"}})
// output：2022-03-18 00:52:32     DEBUG   cache/redis_driver.go:51        Cache   {"Flush": "{\"boys\":[\"alex\",\"bob\"],\"sex\":[\"f\",\"m\"]}"}

logger.DebugJSON("Cache", "Flush", struct {
    Name, Sex string
    Age       int32
}{
    Name: "alex",
    Sex:  "m",
    Age:  18,
})
// output：2022-03-18 01:10:10     DEBUG   cache/redis_driver.go:52        Cache   {"Flush": "{\"Name\":\"alex\",\"Sex\":\"m\",\"Age\":18}"}

logger.InfoJSON("Cache", "Flush", map[string][]string{"boys": {"alex", "bob"}, "sex":  {"f", "m"}})
// output：2022-03-18 00:54:39     INFO    cache/redis_driver.go:51        Cache   {"Flush": "{\"boys\":[\"alex\",\"bob\"],\"sex\":[\"f\",\"m\"]}"}

logger.WarnJSON("Cache", "Flush", map[string][]string{"boys": {"alex", "bob"}, "sex":  {"f", "m"}})
// output：2022-03-18 00:55:38     WARN    cache/redis_driver.go:51        Cache   {"Flush": "{\"boys\":[\"alex\",\"bob\"],\"sex\":[\"f\",\"m\"]}"}

// 会打印出调用堆栈，但是不会退出程序
logger.ErrorJSON("Cache", "Flush", map[string][]string{"boys": {"alex", "bob"}, "sex":  {"f", "m"}})
// output：2022-03-18 00:56:43     ERROR   cache/redis_driver.go:51        Cache   {"Flush": "{\"boys\":[\"alex\",\"bob\"],\"sex\":[\"f\",\"m\"]}"}

// 会打印出调用堆栈，写完 log 后调用 os.Exit(1) 直接退出程序
logger.FatalJSON("Cache", "Flush", map[string][]string{"boys": {"alex", "bob"}, "sex":  {"f", "m"}})
// output：2022-03-18 00:58:06     FATAL   cache/redis_driver.go:51        Cache   {"Flush": "{\"boys\":[\"alex\",\"bob\"],\"sex\":[\"f\",\"m\"]}"}

```

## 需要直接使用 `zap` 包提供的方法时

```go

// 比如要使用 `Panic()` 方法时，可以直接使用
zap.L().Panic("Cache", zap.String("Flush", "danger!!!"), zap.Time("time", time.Now()))


// 因为已经将自定义的 logger 替换成了全局的 logger 因此，以下二者打印出来的内容结构将完全一致
zap.L().Debug("Cache", zap.String("Flush", "danger!!!"), zap.Time("time", time.Now()))
logger.Debug("Cache", zap.String("Flush", "danger!!!"), zap.Time("time", time.Now()))

```