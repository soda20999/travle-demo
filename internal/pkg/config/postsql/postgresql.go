package postgresql

import (
	"iam/pkg/config"

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

func Init(cfg *config.PostgreSQLConfig) error {
	dsn := cfg.DSN

	var err error
	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return err
	}

	return nil
}

func Close() {
	sqlDB, err := DB.DB()
	if err != nil {
		zap.L().Error("get sql.DB failed", zap.Error(err))
		return
	}
	sqlDB.Close()
}
