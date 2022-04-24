package database

import (
	"database/sql"
	"sync"
	"time"

	"gorm.io/gorm"
	gormLogger "gorm.io/gorm/logger"

	"gin-biz-web-api/pkg/config"
	"gin-biz-web-api/pkg/console"
)

type DBClientConfig struct {
	DBConfig        gorm.Dialector
	LG              gormLogger.Interface
	MaxOpenConns    int
	MaxIdleConns    int
	ConnMaxLifetime time.Duration
}

type MySQLClient struct {
	DB    *gorm.DB
	SQLDB *sql.DB
}

var (
	once             sync.Once
	mysqlCollections map[string]*MySQLClient
	DB               *gorm.DB // 默认 mysql 连接的 DB 对象
	SQLDB            *sql.DB  // 默认 mysql 连接中的 database/sql 包里的 *sql.DB 对象
)

func Instance(group ...string) *MySQLClient {
	if len(group) > 0 {
		if client, ok := mysqlCollections[group[0]]; ok {
			return client
		}
		console.Exit("The MySQL instance object named [%s] group could not be found!", group[0])
	}

	return mysqlCollections["default"]
}

func ConnectMySQL(configs map[string]*DBClientConfig) {
	once.Do(func() {

		if mysqlCollections == nil {
			mysqlCollections = make(map[string]*MySQLClient)
		}

		for group, cfg := range configs {
			var client = NewMysqlClient(cfg.DBConfig, cfg.LG)

			// ================ 连接池设置 =================
			// 设置最大连接数，0 表示无限制，默认为 0
			// 在高并发的情况下，将值设为大于 10，可以获得比设置为 1 接近六倍的性能提升。而设置为 10 跟设置为 0（也就是无限制），在高并发的情况下，性能差距不明显
			// 最大连接数不要大于数据库系统设置的最大连接数 show variables like 'max_connections';
			// 这个值是整个系统的，如有其他应用程序也在共享这个数据库，这个可以合理地控制小一点
			client.SQLDB.SetMaxOpenConns(cfg.MaxOpenConns)
			// 设置最大空闲连接数，0 表示不设置空闲连接数，默认为 2
			// 在高并发的情况下，将值设为大于 0，可以获得比设置为 0 超过 20 倍的性能提升
			// 这是因为设置为 0 的情况下，每一个 SQL 连接执行任务以后就销毁掉了，执行新任务时又需要重新建立连接。很明显，重新建立连接是很消耗资源的一个动作
			// 此值不能大于 SetMaxOpenConns 的值，大于的情况下 mysql 驱动会自动将其纠正
			client.SQLDB.SetMaxIdleConns(cfg.MaxIdleConns)
			// 设置每个连接的过期时间
			// 设置连接池里每一个连接的过期时间，过期会自动关闭。理论上来讲，在并发的情况下，此值越小，连接就会越快被关闭，也意味着更多的连接会被创建。
			// 设置的值不应该超过 MySQL 的 wait_timeout 设置项（默认情况下是 8 个小时）
			// 此值也不宜设置过短，关闭和创建都是极耗系统资源的操作。
			// 设置此值时，需要特别注意 SetMaxIdleConns 空闲连接数的设置。假如设置了 100 个空闲连接，过期时间设置了 1 分钟，在没有任何应用的 SQL 操作情况下，数据库连接每 1.6 秒就销毁和新建一遍。
			// 这里的推荐，比较保守的做法是设置五分钟
			client.SQLDB.SetConnMaxLifetime(cfg.ConnMaxLifetime)

			mysqlCollections[group] = client
		}

		setSimpleHelper()
	})
}

func NewMysqlClient(dbConfig gorm.Dialector, lg gormLogger.Interface) *MySQLClient {
	mysql := &MySQLClient{}

	var err error
	mysql.DB, err = gorm.Open(dbConfig, &gorm.Config{Logger: lg})
	console.ExitIf(err)

	// 获取底层的 sqlDB
	// *gorm.DB 对象的 DB() 方法，可以直接获取到 database/sql 包里的 *sql.DB 对象
	mysql.SQLDB, err = mysql.DB.DB()
	console.ExitIf(err)

	return mysql
}

func setSimpleHelper() {
	defaultInstance := Instance()
	DB = defaultInstance.DB
	SQLDB = defaultInstance.SQLDB
}

// DropAllTables 删除所有表（其实是直接删库跑路，😊）
// most dangerous !!!
func DropAllTables(group ...string) error {
	var err error
	console.Danger("Most dangerous!")

	switch config.GetString("cfg.database.driver") {
	case "mysql":
		err = dropMysqlDatabase(group...)
	default:
		console.Exit("database driver not supported")
	}

	return err
}

// dropMysqlDatabase 删除数据表
func dropMysqlDatabase(group ...string) error {
	dbname := CurrentDatabase(group...)
	db := Instance(group...).DB
	var tables []string

	// 读取所有数据表
	err := db.Table("information_schema.tables").
		Where("table_schema = ?", dbname).
		Pluck("table_name", &tables).
		Error
	if err != nil {
		return err
	}

	// 暂时关闭外键检测
	db.Exec("SET foreign_key_checks = 0;")

	// 删除所有表
	for _, table := range tables {
		if err := db.Migrator().DropTable(table); err != nil {
			return err
		}
	}

	// 开启 MySQL 外键检测
	db.Exec("SET foreign_key_checks = 1;")
	return nil
}

// CurrentDatabase 返回当前数据库名称
func CurrentDatabase(group ...string) string {
	return Instance(group...).DB.Migrator().CurrentDatabase()
}

// TableName 获取当前对象的表名称
// eg：database.TableName(&model.User{})
// output: "users"
func TableName(obj interface{}, group ...string) string {
	db := Instance(group...).DB
	stmt := &gorm.Statement{DB: db}
	_ = stmt.Parse(obj)
	return stmt.Schema.Table
}
