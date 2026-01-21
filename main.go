package main

import (
	"context"
	"fmt"
	"net/http"
	"os"
	"os/signal"
	"syscall"
	"time"

	"iam/internal/pkg/config/gorm"
	"iam/internal/pkg/config/logger"
	"iam/internal/pkg/config/mysql"
	"iam/internal/pkg/config/redis"
	"iam/internal/pkg/route"
	"iam/internal/pkg/validator"
	"iam/pkg/config"

	"iam/pkg/snowflake"

	"github.com/spf13/viper"
	"go.uber.org/zap"
)

func main() {
	//1.加载配置
	if err := config.Init(); err != nil {
		zap.L().Error("config Init() error config failed :&v\n ", zap.Error(err))
		return
	}

	//初始化日志文件
	if err := logger.Init(config.Conf.LogConfig, config.Conf.Mode); err != nil {
		zap.L().Error("logger Init() error logger failed :&v\n ", zap.Error(err))
		return
	}
	defer zap.L().Sync()
	zap.L().Debug("logger Init() success") //全局使用logger记录日志格式，zap.L()即可

	//初始化mysql
	if err := mysql.Init(config.Conf.MySQLConfig); err != nil {
		zap.L().Error("mysql Init() error mysql failed :&v\n ", zap.Error(err))
		return
	}
	defer mysql.Close()

	//初始化redis
	if err := redis.Init(config.Conf.RedisConfig); err != nil {
		zap.L().Error("redis Init() error redis failed :&v\n ", zap.Error(err))
		return
	}
	defer redis.Close()

	//初始化gorm
	if err := gorm.InitGorm(); err != nil {
		zap.L().Error("gorm Init() error gorm failed :&v\n ", zap.Error(err))
		return
	}
	defer gorm.Close()

	if err := snowflake.Init(config.Conf.StartTime, config.Conf.MachineId); err != nil {
		zap.L().Error("snowflake Init() error snowflake failed :&v\n ", zap.Error(err))
		return
	}
	//初始化翻译器validator
	if err := validator.InitTrans("en"); err != nil {
		zap.L().Error("validator InitTrans() error validator failed :&v\n ", zap.Error(err))
		return
	}

	//注册路由
	r := route.Setup()
	r.Run(fmt.Sprintf(":%d", config.Conf.Port))

	//优雅关机
	srv := &http.Server{
		Addr:    fmt.Sprintf(":%d", viper.GetInt("app.port")),
		Handler: r,
	}

	go func() { //开一个go是为了让下面的listen函数不会无限循环
		// 开启一个goroutine启动服务
		if err := srv.ListenAndServe(); err != nil && err != http.ErrServerClosed {
			zap.L().Error("listen: %s\n", zap.Error(err))
		}
	}()

	// 等待中断信号来优雅地关闭服务器，为关闭服务器操作设置一个5秒的超时
	quit := make(chan os.Signal, 1) // 创建一个接收信号的通道
	// kill 默认会发送 syscall.SIGTERM 信号
	// kill -2 发送 syscall.SIGINT 信号，我们常用的Ctrl+C就是触发系统SIGINT信号
	// kill -9 发送 syscall.SIGKILL 信号，但是不能被捕获，所以不需要添加它
	// signal.Notify把收到的 syscall.SIGINT或syscall.SIGTERM 信号转发给quit
	signal.Notify(quit, syscall.SIGINT, syscall.SIGTERM) // 此处不会阻塞
	<-quit                                               // 阻塞在此，当接收到上述两种信号时才会往下执行
	zap.L().Info("Shutdown Server ...")
	// 创建一个5秒超时的context
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	// 5秒内优雅关闭服务（将未处理完的请求处理完再关闭服务），超过5秒就超时退出
	if err := srv.Shutdown(ctx); err != nil {
		zap.L().Error("Server Shutdown: ", zap.Error(err))
	}

	zap.L().Info("Server exiting")
}
