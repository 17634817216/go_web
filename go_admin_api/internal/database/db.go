package database

import (
	"fmt"
	"go_admin_api/global"
	"gorm.io/driver/mysql"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"time"
)

// IDatabase 数据库接口
type IDatabase interface {
	Connect() error
	GetDB() *gorm.DB
}

// DBFactory 数据库工厂
type Myinit struct{}

// CreateDB 创建数据库实例
func (f *Myinit) CreateDB(dbType string, cfg *global.Application) IDatabase {
	switch dbType {
	case "mysql":
		return NewMySQLDB(cfg)
	case "postersql":
		return NewPostgresDB(cfg)
	default:
		return NewMySQLDB(cfg) // 默认使用 MySQL
	}
}

// Database 基础数据库结构
type Database struct {
	db  *gorm.DB
	cfg *global.Application
}

// GetDB 获取数据库实例
func (d *Database) GetDB() *gorm.DB {
	return d.db
}

// MySQLDB MySQL 实现
type MySQLDB struct {
	Database
}

// NewMySQLDB 创建 MySQL 实例
func NewMySQLDB(cfg *global.Application) IDatabase {
	return &MySQLDB{
		Database: Database{cfg: cfg},
	}
}

// Connect MySQL 连接实现
func (m *MySQLDB) Connect() error {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		m.cfg.DatabaseConfig.DB_USER,
		m.cfg.DatabaseConfig.DB_PASSWORD,
		m.cfg.DatabaseConfig.DB_HOST,
		m.cfg.DatabaseConfig.DB_PORT,
		m.cfg.DatabaseConfig.DB_NAME,
	)
	fmt.Println(dsn)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		global.App.Log.Info(fmt.Sprintf("MYSQL 数据库连接实例错误: %v", err))
		//return err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	// 设置连接池参数
	sqlDB.SetMaxIdleConns(10)                  // 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxOpenConns(100)                 // 设置打开数据库连接的最大数量
	sqlDB.SetConnMaxLifetime(time.Hour)        // 设置了连接可复用的最大时间
	sqlDB.SetConnMaxIdleTime(time.Minute * 10) // 设置空闲连接最大存活时间
	//log.Fatal("MYSQL 数据库连接成功 %v ", dsn)
	global.App.Log.Info(fmt.Sprintf("MYSQL 数据库连接成功 %v ", dsn))

	m.db = db
	return nil
}

// PostgresDB PostgreSQL 实现
type PostgresDB struct {
	Database
}

// NewPostgresDB 创建 PostgreSQL 实例
func NewPostgresDB(cfg *global.Application) IDatabase {
	return &PostgresDB{
		Database: Database{cfg: cfg},
	}
}

// Connect PostgreSQL 连接实现
func (p *PostgresDB) Connect() error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%d sslmode=disable TimeZone=Asia/Shanghai",
		p.cfg.DatabaseConfig.DB_HOST,
		p.cfg.DatabaseConfig.DB_USER,
		p.cfg.DatabaseConfig.DB_PASSWORD,
		p.cfg.DatabaseConfig.DB_NAME,
		p.cfg.DatabaseConfig.DB_PORT,
	)

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		global.App.Log.Info(fmt.Sprintf("PostgresDB 数据库连接实例错误: %v", err))

		return err
	}
	sqlDB, err := db.DB()
	if err != nil {
		return err
	}

	// 设置连接池参数
	sqlDB.SetMaxIdleConns(10)                  // 设置空闲连接池中连接的最大数量
	sqlDB.SetMaxOpenConns(100)                 // 设置打开数据库连接的最大数量
	sqlDB.SetConnMaxLifetime(time.Hour)        // 设置了连接可复用的最大时间
	sqlDB.SetConnMaxIdleTime(time.Minute * 10) // 设置空闲连接最大存活时间
	p.db = db
	//log.Fatal("PostgresDB 数据库连接成功 %v ", dsn)
	global.App.Log.Info(fmt.Sprintf("PostgresDB 数据库连接成功 %v ", dsn))

	return nil
}

func InitDatabase() error {
	dbInit := &Myinit{}

	// 创建数据库连接
	db := dbInit.CreateDB(global.App.DatabaseConfig.DB_DRVICE, global.App)

	// 连接数据库
	err := db.Connect()
	if err != nil {
		return err
	}

	// 设置全局数据库实例
	global.App.DB = db.GetDB()

	return nil
}
