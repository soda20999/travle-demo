package postgresql

import (
	discover "iam/internal/business/discover/model"
	"iam/pkg/config" // 替换为你实际的配置包路径

	"go.uber.org/zap"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var DB *gorm.DB

// Init 初始化 PostgreSQL 连接
func Init(cfg *config.PostgreSQLConfig) (err error) {
	// 直接使用你提供的 Neon DSN 字符串
	dsn := cfg.DSN

	DB, err = gorm.Open(postgres.Open(dsn), &gorm.Config{
		// 建议开发环境下开启 Info 级别日志，查看生成的 SQL
		Logger: logger.Default.LogMode(logger.Info),
	})

	if err != nil {
		return err
	}

	// 【低代码/无 DTO 关键】在此处注册模型，系统启动自动建表/同步字段
	// 以后新增模型只需往这里一塞，不需要写迁移脚本，也不用写 DTO
	err = DB.AutoMigrate(
		&discover.Province{},
		&discover.City{},
		&discover.Attraction{},
		// &models.User{},
		// &models.Message{},
	)

	return err
}

// Close 关闭数据库连接
func Close() {
	sqlDB, err := DB.DB()
	if err != nil {
		zap.L().Error("get sql.DB failed", zap.Error(err))
		return
	}
	sqlDB.Close()
}
