package database

import (
	"database/sql"
	"fmt"

	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"

	"gin-biz-web-api/pkg/config"
	"gin-biz-web-api/pkg/console"
)

// DB 对象
var DB *gorm.DB
var SQLDB *sql.DB

// Connect 连接数据库
func Connect(dbConfig gorm.Dialector, lg gormLogger.Interface) {

	// 使用 gorm.Open 连接数据库
	var err error
	// 这里需要注意：不能写成
	// 	DB, err := gorm.Open(dbConfig, &gorm.Config{
	//		Logger: lg,
	//	})
	// 因为 `:=` 会重新声明并创建了左侧的新局部变量，因此在其他包中调用 database.DB 变量时，它仍然是 nil
	// 因为根本就没有赋值到包全局变量 database.DB 上
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

// DropAllTables 删除所有表（其实是直接删库跑路，😊）
// most dangerous !!!
func DropAllTables() error {
	var err error
	console.Danger("Most dangerous!")

	switch config.GetString("database.driver") {
	case "mysql":
		err = dropMysqlDatabase()
	default:
		console.Exit("database driver not supported")
	}

	return err
}

// dropMysqlDatabase 删除数据库
func dropMysqlDatabase() error {
	dbname := CurrentDatabase()
	s := fmt.Sprintf("drop database %s;", dbname)
	if err := DB.Exec(s).Error; err != nil {
		return err
	}
	s = fmt.Sprintf("create database %s default charset utf8mb4 collate utf8mb4_general_ci;", dbname)
	if err := DB.Exec(s).Error; err != nil {
		return err
	}
	s = fmt.Sprintf("use %s;", dbname)
	if err := DB.Exec(s).Error; err != nil {
		return err
	}
	return nil
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
