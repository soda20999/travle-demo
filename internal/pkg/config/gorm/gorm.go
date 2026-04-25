package gorm

import (
	"errors"
	"iam/pkg/config"
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var Db *gorm.DB

var (
	ErrRecordNotFound = errors.New("record not found")
)

func InitGorm() error {
	dsn := config.Conf.PostgreSQLConfig.DSN
	cfg := config.Conf.PostgreSQLConfig

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

	var err error
	Db, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logLevel),
	})
	if err != nil {
		log.Printf("connect PostgreSQL failed: %v", err)
		return err
	}

	sqlDB, err := Db.DB()
	if err != nil {
		return err
	}
	sqlDB.SetMaxOpenConns(cfg.MaxOpenConns)
	sqlDB.SetMaxIdleConns(cfg.MaxIdleConns)

	log.Println("PostgreSQL connected")
	return nil
}

func Close() {
	sqlDB, err := Db.DB()
	if err != nil {
		log.Printf("get sql.DB failed: %v", err)
		return
	}
	sqlDB.Close()
}
