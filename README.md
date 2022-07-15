# gin-biz-web-api
基于 gin 框架封装的业务 web api 脚手架，可直接拿过来上手开发业务逻辑代码。

## 关于项目目录

- bootstrap —— 包初始化
- cmd —— 命令行  
- config —— 配置文件目录
- constant —— 常量目录  
- crontab —— 计划任务目录  
- docs —— 项目文档相关目录
- etc —— 配置文件  
- global —— 全局变量
- internal —— 内部模块目录
    - controller —— 控制器层
    - dao —— 数据访问层（Database Access Object）
    - middleware —— HTTP 中间件
    - requests —— 验证器
    - service —— 项目核心业务逻辑层
- job —— 异步任务目录
- model —— 模型层
- pkg —— 项目相关的模块包
- public —— 静态资源目录  
- routers —— 路由  
- scripts —— 各类构建、安装、分析等操作的脚本  
- storage —— 项目生成的临时文件
  - logs —— 日志文件夹
- third_party —— 第三方的资源工具

## 运行项目

> 本人开发时本地环境：`go version go1.16.3 darwin/amd64`

- 下载项目

```shell
git clone https://github.com/pudongping/gin-biz-web-api.git
```

- 下载项目相关依赖

> 此项目使用 `Go Modules` 进行依赖包管理，请注意首先得开启 `Go Modules`

```shell
cd <your-path>/gin-biz-web-api && go mod tidy
```

- 修改配置

将根目录下的 `/etc/config.yaml.example` 配置文件复制成 `/etc/config.yaml` 然后将 `/etc/config.yaml` 文件中的配置信息修改成你自己的配置。
如果启动项目时设置了 `--env` 参数，那么则会走对应的环境配置信息。  
比如启动项目时，执行了 `go run main.go --env=prod` 命令，那么则会使用 `/etc/prod_config.yaml` 文件中的配置信息，如果对应文件不存在，请  
将 `cp ./etc/config.yaml ./etc/prod_config.yaml` 复制一份。

- 启动项目

```shell
go run main.go
```

## 编译及打包

- 将编译信息写入二进制文件中

```shell
# 本地编译打包，在项目根目录下执行
make build-local
```

- 查看编译后的二进制文件和版本信息

```shell
./gin-biz-web-api -v 或者 ./gin-biz-web-api --version

# output is:
# Build Time: 2022-03-28,00:34:53
# Build Version: 1.0.0
# Build Go Version: go version go1.16.3 darwin/amd64
# Build Git Commit Hash ID: 5f112956c4c51c763f46a35eff3e767ead53abe4
```

## 第三方依赖

使用到的开源库：

- [gin-gonic/gin](https://github.com/gin-gonic/gin) —— http 框架、路由、路由组、中间件
- [gin-contrib/zap](https://github.com/gin-contrib/zap) —— 高性能日志方案
- [natefinch/lumberjack.v2](https://gopkg.in/natefinch/lumberjack.v2) —— 滚动日志
- [davecgh/go-spew](https://github.com/davecgh/go-spew) —— 漂亮的打印调试工具
- [spf13/viper](https://github.com/spf13/viper) —— 配置信息
- [spf13/cobra](https://github.com/spf13/cobra) —— 命令行结构
- [spf13/cast](https://github.com/spf13/cast) —— 类型转换
- [go-gorm/gorm](https://github.com/go-gorm/gorm) —— ORM 数据操作
- [go-redis/redis](https://github.com/go-redis/redis/v8) —— Redis 操作
- [mojocn/base64Captcha](https://github.com/mojocn/base64Captcha) —— 图片验证码
- [thedevsaddam/govalidator](https://github.com/thedevsaddam/govalidator) —— 请求验证器
- [iancoleman/strcase](https://github.com/iancoleman/strcase) —— 字符串大小写操作
- [gertd/go-pluralize](https://github.com/gertd/go-pluralize) —— 英文字符单数复数处理
- [golang-jwt/jwt](https://github.com/golang-jwt/jwt) —— JWT 认证
- [go-gomail/gomail](https://github.com/go-gomail/gomail) —— SMTP 邮件发送
- [disintegration/imaging](https://github.com/disintegration/imaging) —— 图片处理、裁剪用户头像
- [ulule/limiter](https://github.com/ulule/limiter/v3) —— 接口频率限流器
- [juju/ratelimit](https://github.com/juju/ratelimit) —— 令牌桶限流器
- [shirou/gopsutil](https://github.com/shirou/gopsutil) —— 查看系统信息
- [fsnotify/fsnotify](https://github.com/fsnotify/fsnotify) —— 配置热更新
- [robfig/cron](https://github.com/robfig/cron) —— 定时计划任务
- [hibiken/asynq](https://github.com/hibiken/asynq) —— 异步队列

## 其他

- 不想使用 `/etc` 目录下的配置文件，如何更换？

```shell
# 启动项目时可指定 `--config_path` 或者 `-c` 参数进行更换读取配置文件的目录
# 例如：
go run main.go -c=configs/
# 如果想读取多个目录下的配置文件信息，则可以
go run main.go -c=etc/,configs/
```

- 多环境下，如何区分配置信息？

```shell
# 启动项目时可指定 `--env` 或者 `-e` 参数进行多环境运行，注意：修改配置文件中的 `App.Env` 并不会自动切换读取配置文件
# 例如：
go run main.go -e=local
# 则使用的是 `/etc/local_config.yaml` 配置文件中的配置
# 目前支持 local、dev、test、prod

# 如果对应的配置文件不存在，请复制 `/etc/config.yaml` 文件并修改对应的文件名，文件不存在时会有对应的报错信息。
```

- 如何更换 `gin` 框架的启动模式？

```shell
# 项目启动时可指定 `--mode` 参数或者直接修改配置文件中的 `App.GinRunMode` 参数即可，如果同时设置，会优先于命令行参数 `--mode`，详见 `config/app.go` 文件。
# 例如：
go run main.go --mode=release
# gin 框架支持 debug、release、test 三种模式
```

- 如何更换 http 服务的启动端口？

```shell
# 项目启动时可指定 `--port` 参数或者 `-p` 参数或者直接修改配置文件中的 `App.Port` 参数即可，如果同时设置，会优先于命令行参数 `--port`，详见 `config/app.go` 文件。
# 例如：
go run main.go -p=8081
```

## 命令行

本项目使用 cobra 命令行启动 http web 服务，执行 `go run main.go` 其实默认执行了 `go run main.go server` 命令。

### 目前支持的命令

#### cache

可通过 `go run main.go cache -h` 查看更多详情信息。

命令 | 含义 | 示例
--- | --- | ---
go run main.go cache clear | 清空所有缓存 | go run main.go cache clear
go run main.go cache forget * | 删除某个缓存 key | go run main.go cache forget abc
go run main.go cache get * | 获取某个缓存 key 的值 | go run main.go cache get abc

#### generate

可通过 `go run main.go generate -h` 查看更多详情信息。

命令 | 含义 | 示例
--- | --- | ---
go run main.go generate * | 生成 jwt 的密钥 | go run main.go generate jwt-key

#### server

启动 gin http web 服务，没有参数。

命令 | 含义 | 示例
--- | --- | ---
go run main.go server | 启动 http 服务 | go run main.go server

#### make

可通过 `go run main.go make -h` 查看更多详情信息。

命令 | 含义 | 示例
--- | --- | ---
go run main.go make model * | 将 mysql 数据表生成对应的结构体 | go run main.go make model users

## 自定义包

包 | 用处
--- | ---
app | 和系统相关的方法
auth | 授权
cache | 缓存
captcha | 图片验证码
config | 配置
console | 控制台打印工具
crontab | 定时任务
database | 数据库
email | 邮件工具
errcode | 自定义错误码
file | 文件操作
hash | hash 处理
helper | 助手函数
job | 异步队列任务
jwt | JWT 授权验证
limiter | 接口访问频率控制
logger | 日志操作
paginator | 分页处理
redis | redis 缓存操作
responses | 统一数据返回
upload | 文件上传
validator | 验证器
verifycode | 验证码工具

## 文档

> 接口请求示例，可查看 `routers/api.go` 文件，多数功能都有示例接口。

- [异步队列任务](https://github.com/pudongping/gin-biz-web-api/blob/main/docs/guide/async_queue_job.md)
- [缓存](https://github.com/pudongping/gin-biz-web-api/blob/main/docs/guide/cache.md)
- [定时任务](https://github.com/pudongping/gin-biz-web-api/blob/main/docs/guide/crontab.md)
- [项目部署](https://github.com/pudongping/gin-biz-web-api/blob/main/docs/guide/deploy.md)
- [邮件](https://github.com/pudongping/gin-biz-web-api/blob/main/docs/guide/email.md)
- [日志](https://github.com/pudongping/gin-biz-web-api/blob/main/docs/guide/log.md)
- [redis](https://github.com/pudongping/gin-biz-web-api/blob/main/docs/guide/redis.md)
- [接口示例，带分页功能](https://github.com/pudongping/gin-biz-web-api/blob/main/docs/guide/document.md)