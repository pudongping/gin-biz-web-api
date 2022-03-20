// Package bootstrap 处理程序初始化逻辑
package bootstrap

import (
	"fmt"

	"gin-biz-web-api/pkg/console"
)

func Initialize() {

	fmt.Println(printProjectName())
	console.Info("Initializing ...")

	// 初始化配置文件信息
	setupConfig()

	// 初始化日志
	setupLogger()

	// 初始化数据库
	setupDB()

	// 初始化 redis
	setupRedis()

	// 初始化缓存 cache
	setupCache()
}

func printProjectName() string {
	return `
 ________  ___  ________           ________  ___  ________          ___       __   _______   ________          ________  ________  ___     
|\   ____\|\  \|\   ___  \        |\   __  \|\  \|\_____  \        |\  \     |\  \|\  ___ \ |\   __  \        |\   __  \|\   __  \|\  \    
\ \  \___|\ \  \ \  \\ \  \       \ \  \|\ /\ \  \\|___/  /|       \ \  \    \ \  \ \   __/|\ \  \|\ /_       \ \  \|\  \ \  \|\  \ \  \   
 \ \  \  __\ \  \ \  \\ \  \       \ \   __  \ \  \   /  / /        \ \  \  __\ \  \ \  \_|/_\ \   __  \       \ \   __  \ \   ____\ \  \  
  \ \  \|\  \ \  \ \  \\ \  \       \ \  \|\  \ \  \ /  /_/__        \ \  \|\__\_\  \ \  \_|\ \ \  \|\  \       \ \  \ \  \ \  \___|\ \  \ 
   \ \_______\ \__\ \__\\ \__\       \ \_______\ \__\\________\       \ \____________\ \_______\ \_______\       \ \__\ \__\ \__\    \ \__\
    \|_______|\|__|\|__| \|__|        \|_______|\|__|\|_______|        \|____________|\|_______|\|_______|        \|__|\|__|\|__|     \|__|
`
}
