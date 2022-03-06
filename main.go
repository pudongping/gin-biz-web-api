package main

import (
	"gin-biz-web-api/bootstrap"
	"gin-biz-web-api/global"
	"gin-biz-web-api/pkg/console"
)

var (
	buildTime    string // 二进制文件编译时间
	buildVersion string // 二进制文件编译版本
	gitCommitID  string // 二进制文件编译时的 git 提交版本号
)

func main() {

	// 初始化加载命令行参数
	bootstrap.SetupFlag()

	if global.IsVersion {
		console.Info("build_time: %s", buildTime)
		console.Info("build_version: %s", buildVersion)
		console.Info("git_commit_id: %s", gitCommitID)
		return
	}

	// 初始化框架
	bootstrap.Initialize()

	// 启动服务
	bootstrap.RunServer()
}
