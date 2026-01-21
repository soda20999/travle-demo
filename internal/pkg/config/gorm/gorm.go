// internal/pkg/config/gorm/gorm.go
package gorm

import (
	"errors"
	"fmt"
	"iam/pkg/config"
	"log"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

var (
	ErrRecordNotFound = errors.New("数据未找到")
)



// InitGorm 初始化 GORM 数据库连接
func InitGorm() (err error) {
	cfg := config.Conf.GormConfig

	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=%s&parseTime=True&loc=Local",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName,
		cfg.Charset,
	)

	var logLevel logger.LogLevel
	switch cfg.LogLevel {
	case "silent":
		logLevel = logger.Silent
	case "error":
		logLevel = logger.Error
	case "warn":
		logLevel = logger.Warn
	case "info":
		logLevel = logger.Info
	default:
		logLevel = logger.Info
	}

	Db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})

	if err != nil {
		log.Printf("open gorm mysql failed: %v", err)
		return err
	}

	// 设置连接池
	sqlDB, err := Db.DB()
	if err != nil {
		return err
	}

	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)

	return nil
}

// Close 关闭数据库连接
func Close() {
	if Db != nil {
		sqlDB, _ := Db.DB()
		if sqlDB != nil {
			_ = sqlDB.Close()
		}
	}
}

// Ping 测试数据库连接
func Ping() error {
	if Db == nil {
		return fmt.Errorf("gorm database is not initialized")
	}

	sqlDB, err := Db.DB()
	if err != nil {
		return err
	}

	return sqlDB.Ping()
}