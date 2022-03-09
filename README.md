# gin-biz-web-api
基于 gin 框架封装的业务 web api 脚手架，可直接拿过来上手开发业务逻辑代码。

## 关于项目目录

- configs -- 配置文件目录
- docs -- 项目文档相关目录
- global -- 全局变量
- internal -- 内部模块目录
    - dao -- 数据访问层（Database Access Object）
    - middleware -- HTTP 中间件
    - model -- 模型层
    - routers -- 路由
    - service -- 项目核心业务逻辑层
- pkg -- 项目相关的模块包
- storage -- 项目生成的临时文件
- scripts -- 各类构建、安装、分析等操作的脚本
- third_party -- 第三方的资源工具

## 编译及打包

- 将编译信息写入二进制文件中

```shell
# # 编译项目二进制文件时，可以通过 `ldflags` 工具，将一些编译相关的信息写入二进制文件中，方法日后查看二进制文件相关的信息
go build -ldflags \
"-X main.buildTime=`date +%Y-%m-%d,%H:%M:%S` -X main.buildVersion=1.0.0 -X main.gitCommitID=`git rev-parse HEAD`"
```

- 查看编译后的二进制文件和版本信息

```shell
./gin-biz-web-api -version  

# output is:
# build_time: 2022-02-27,15:15:42 
# build_version: 1.0.0 
# git_commit_id: 2e8304393d1a7830ec7f5e6e1fa21529ceeb84ed
```