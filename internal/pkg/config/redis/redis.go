package redis

import (
	"context"
	"fmt"

	"iam/pkg/config"

	"github.com/redis/go-redis/v9" // 推荐使用go-redis
	"go.uber.org/zap"
)

// redis连接池，目前暂定这样
var rdb *redis.Client //创建redis客户端（go-redis 使用 Client 管理连接池）

func Init(cfg *config.RedisConfig) (err error) {
	rdb = redis.NewClient(&redis.Options{ //实例化一个客户端（内部含连接池）
		Addr: fmt.Sprintf("%s:%d",
			cfg.Host,
			cfg.Port), //要连接的redis数据库
		Password: "",           //密码，没有则为空
		DB:       cfg.DB,       //设置数据库
		PoolSize: cfg.PoolSize, //最大连接数（相当于原来的 MaxIdle/MaxActive 的概念，按需设置）
		// MinIdleConns: 10,        //可选：最小空闲连接数（go-redis v9支持）
		MaxRetries: cfg.MaxRetries, //重试次数，0为不重试
	})

	ctx := context.Background()
	_, err = rdb.Ping(ctx).Result()
	if err != nil {
		zap.L().Error("redis ping error failed: %v\n", zap.Error(err))
	}
	return
}

func Close() {
	_ = rdb.Close()
}
