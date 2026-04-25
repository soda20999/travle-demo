// internal/pkg/config/gorm/gorm.go
package gorm

import (
	"errors"
	ar_model "iam/internal/business/ar/model"
	"iam/internal/business/discover/model"
	foot_model "iam/internal/business/footprint/model"
	pref_model "iam/internal/business/preference/model"
	user_model "iam/internal/business/user/model"
	"iam/pkg/config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

var (
	ErrRecordNotFound = errors.New("数据未找到")
)

// InitGorm 初始化 GORM 数据库连接
// InitGorm 初始化 GORM 数据库连接
func InitGorm() (err error) {
	// 1. 直接获取配置文件中的 DSN 字符串
	// 建议在 config.go 的 PostgresConfig 结构体中增加一个 DSN 字段
	dsn := config.Conf.PostgreSQLConfig.DSN

	// 如果你还没改 config 结构体，可以暂时先硬编码测试：
	// dsn := "postgresql://neondb_owner:npg_bfF67qShuMZv@ep-lingering-term-a4zdc2i2-pooler.us-east-1.aws.neon.tech/vchat?sslmode=require"

	cfg := config.Conf.PostgreSQLConfig

	// 2. 日志级别逻辑保持不变
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

	// 3. 使用 DSN 打开连接
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})

	if err != nil {
		log.Printf("连接 PostgreSQL 失败: %v", err)
		return err
	}

	// 4. 设置连接池（这一步依然重要，防止 Neon 连接数炸裂）
	sqlDB, err := Db.DB()
	if err != nil {
		return err
	}
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)

	// 5. 自动同步（你的无 DTO 核心逻辑）
	if err := Db.AutoMigrate(
		&user_model.User{},
		&pref_model.TravelStyle{},
		&pref_model.UserTravelPreference{},
		&ar_model.ARScan{},
		&ar_model.ARScanResult{},
		&model.Province{},
		&model.City{},
		&model.Attraction{},
		&foot_model.Footprint{},
	); err != nil {
		log.Printf("AutoMigrate 失败: %v", err)
		return err
	}

	log.Println("✅ PostgreSQL (Neon) 连接并同步成功!")
	return nil
}

// Close 关闭数据库连接
func Close() {
	sqlDB, err := Db.DB()
	if err != nil {
		log.Printf("获取 sql.DB 失败: %v", err)
		return
	}
	sqlDB.Close()
}
