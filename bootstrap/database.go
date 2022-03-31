package bootstrap

import (
	"fmt"
	"time"

	// GORM 的 MySQL 数据库驱动导入
	"gorm.io/driver/mysql"
	"gorm.io/gorm"

	"gin-biz-web-api/pkg/config"
	"gin-biz-web-api/pkg/console"
	"gin-biz-web-api/pkg/database"
	"gin-biz-web-api/pkg/logger"
)

// setupDB 初始化数据库和 ORM
func setupDB() {

	console.Info("init database ...")

	var dbConfig gorm.Dialector
	switch config.GetString("cfg.database.driver") {
	case "mysql":
		// 构建 dsn 信息。DSN 全称为 Data Source Name，表示【数据源信息】
		// user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
			config.GetString("cfg.database.mysql.username"),
			config.GetString("cfg.database.mysql.password"),
			config.GetString("cfg.database.mysql.host"),
			config.GetString("cfg.database.mysql.port"),
			config.GetString("cfg.database.mysql.database"),
			config.GetString("cfg.database.mysql.charset"),
		)
		dbConfig = mysql.New(mysql.Config{
			DSN: dsn,
		})
	default:
		console.Exit("database driver not supported")
	}

	// 连接数据库，并设置 gorm 的日志模式
	database.Connect(dbConfig, logger.NewGormLogger())

	// 设置最大连接数，0 表示无限制，默认为 0
	// 在高并发的情况下，将值设为大于 10，可以获得比设置为 1 接近六倍的性能提升。而设置为 10 跟设置为 0（也就是无限制），在高并发的情况下，性能差距不明显
	// 最大连接数不要大于数据库系统设置的最大连接数 show variables like 'max_connections';
	// 这个值是整个系统的，如有其他应用程序也在共享这个数据库，这个可以合理地控制小一点
	database.SQLDB.SetMaxOpenConns(config.GetInt("cfg.database.mysql.max_open_connections"))
	// 设置最大空闲连接数，0 表示不设置空闲连接数，默认为 2
	// 在高并发的情况下，将值设为大于 0，可以获得比设置为 0 超过 20 倍的性能提升
	// 这是因为设置为 0 的情况下，每一个 SQL 连接执行任务以后就销毁掉了，执行新任务时又需要重新建立连接。很明显，重新建立连接是很消耗资源的一个动作
	// 此值不能大于 SetMaxOpenConns 的值，大于的情况下 mysql 驱动会自动将其纠正
	database.SQLDB.SetMaxIdleConns(config.GetInt("cfg.database.mysql.max_idle_connections"))
	// 设置每个连接的过期时间
	// 设置连接池里每一个连接的过期时间，过期会自动关闭。理论上来讲，在并发的情况下，此值越小，连接就会越快被关闭，也意味着更多的连接会被创建。
	// 设置的值不应该超过 MySQL 的 wait_timeout 设置项（默认情况下是 8 个小时）
	// 此值也不宜设置过短，关闭和创建都是极耗系统资源的操作。
	// 设置此值时，需要特别注意 SetMaxIdleConns 空闲连接数的设置。假如设置了 100 个空闲连接，过期时间设置了 1 分钟，在没有任何应用的 SQL 操作情况下，数据库连接每 1.6 秒就销毁和新建一遍。
	// 这里的推荐，比较保守的做法是设置五分钟
	database.SQLDB.SetConnMaxLifetime(time.Duration(config.GetInt("cfg.database.mysql.max_life_seconds")) * time.Second)

}
