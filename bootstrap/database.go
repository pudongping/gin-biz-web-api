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

	switch config.GetString("cfg.database.driver") {
	case "mysql":
		setupDBMySQL()
	default:
		console.Exit("database driver not supported")
	}

}

func setupDBMySQL() {

	configs := config.Get("cfg.database.mysql")

	dbConfigs := make(map[string]*database.DBClientConfig)

	for group := range configs.(map[string]interface{}) {
		cfgPrefix := "cfg.database.mysql." + group + "."
		username := config.GetString(cfgPrefix + "username")
		password := config.GetString(cfgPrefix + "password")
		host := config.GetString(cfgPrefix + "host")
		port := config.GetString(cfgPrefix + "port")
		db := config.GetString(cfgPrefix + "database")
		charset := config.GetString(cfgPrefix + "charset")

		// 构建 dsn 信息。DSN 全称为 Data Source Name，表示【数据源信息】
		// user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
		dsn := fmt.Sprintf(
			"%s:%s@tcp(%s:%s)/%s?charset=%s&parseTime=True&loc=Local",
			username, password, host, port, db, charset)

		var dbConfig gorm.Dialector
		dbConfig = mysql.New(mysql.Config{
			DSN: dsn,
		})

		var cfg database.DBClientConfig
		cfg.DBConfig = dbConfig
		cfg.LG = logger.NewGormLogger()
		cfg.MaxOpenConns = config.GetInt(cfgPrefix + "max_open_connections")
		cfg.MaxIdleConns = config.GetInt(cfgPrefix + "max_idle_connections")
		cfg.ConnMaxLifetime = time.Duration(config.GetInt(cfgPrefix+"max_life_seconds")) * time.Second

		dbConfigs[group] = &cfg
	}

	database.ConnectMySQL(dbConfigs)
}
