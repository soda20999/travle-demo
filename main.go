package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"iam/internal/business/ar"
	"iam/internal/business/discover"
	"iam/internal/business/footprint"
	"iam/internal/business/preference"
	"iam/internal/business/user"
	cfggorm "iam/internal/pkg/config/gorm"
	"iam/internal/pkg/config/logger"
	postgresql "iam/internal/pkg/config/postsql"
	"iam/internal/pkg/config/redis"
	"iam/internal/pkg/route"
	"iam/internal/pkg/validator"
	"iam/pkg/config"
	"iam/pkg/snowflake"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	if err := config.Init(); err != nil {
		zap.L().Error("config init failed", zap.Error(err))
		return
	}

	if err := logger.Init(config.Conf.LogConfig, config.Conf.Mode); err != nil {
		zap.L().Error("logger init failed", zap.Error(err))
		return
	}
	defer zap.L().Sync()

	if err := postgresql.Init(config.Conf.PostgreSQLConfig); err != nil {
		fmt.Printf("loaded dsn: %s\n", config.Conf.PostgreSQLConfig.DSN)
		zap.L().Error("postgresql init failed", zap.Error(err))
		return
	}
	defer postgresql.Close()

	if err := redis.Init(config.Conf.RedisConfig); err != nil {
		zap.L().Error("redis init failed", zap.Error(err))
		return
	}
	defer redis.Close()

	if err := cfggorm.InitGorm(); err != nil {
		zap.L().Error("gorm init failed", zap.Error(err))
		return
	}
	defer cfggorm.Close()

	if err := postgresql.DB.AutoMigrate(
		&user.User{},
		&preference.TravelStyle{},
		&preference.UserTravelPreference{},
		&discover.Province{},
		&discover.City{},
		&discover.Attraction{},
		&footprint.Footprint{},
		&ar.ARScan{},
		&ar.ARScanResult{},
	); err != nil {
		zap.L().Error("postgresql automigrate failed", zap.Error(err))
		return
	}

	if err := cfggorm.Db.AutoMigrate(
		&user.User{},
		&preference.TravelStyle{},
		&preference.UserTravelPreference{},
		&discover.Province{},
		&discover.City{},
		&discover.Attraction{},
		&footprint.Footprint{},
		&ar.ARScan{},
		&ar.ARScanResult{},
	); err != nil {
		zap.L().Error("gorm automigrate failed", zap.Error(err))
		return
	}

	if err := snowflake.Init(config.Conf.StartTime, config.Conf.MachineId); err != nil {
		zap.L().Error("snowflake init failed", zap.Error(err))
		return
	}

	if err := validator.InitTrans("en"); err != nil {
		zap.L().Error("validator init failed", zap.Error(err))
		return
	}

	r := route.Setup()
	r.Run(fmt.Sprintf(":%d", config.Conf.Port))

	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("app.port")),
		Handler: r,
	}

	go func() {
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Error("listen failed", zap.Error(err))
		}
	}()

	quit := make(chan os.Signal, 1)
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM)
	<-quit

	zap.L().Info("Shutdown Server ...")
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()

	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Error("server shutdown failed", zap.Error(err))
	}

	zap.L().Info("Server exiting")
}
