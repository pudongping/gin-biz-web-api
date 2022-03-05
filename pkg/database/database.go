package database

import (
	"database/sql"

	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"

	"gin-biz-web-api/pkg/console"
)

// DB 对象
var DB *gorm.DB
var SQLDB *sql.DB

// Connect 连接数据库
func Connect(dbConfig gorm.Dialector, lg gormLogger.Interface) {

	// 使用 gorm.Open 连接数据库
	var err error
	DB, err = gorm.Open(dbConfig, &gorm.Config{
		Logger: lg,
	})
	// 处理错误，要是有错误会直接退出程序
	console.ExitIf(err)

	// 获取底层的 sqlDB
	// *gorm.DB 对象的 DB() 方法，可以直接获取到 database/sql 包里的 *sql.DB 对象
	SQLDB, err = DB.DB()
	console.ExitIf(err)

}

// CurrentDatabase 返回当前数据库名称
func CurrentDatabase() string {
	return DB.Migrator().CurrentDatabase()
}

// TableName 获取当前对象的表名称
// eg：database.TableName(&user_model.User{})
// output: "users"
func TableName(obj interface{}) string {
	stmt := &gorm.Statement{DB: DB}
	_ = stmt.Parse(obj)
	return stmt.Schema.Table
}
