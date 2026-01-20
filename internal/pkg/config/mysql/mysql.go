package mysql

import (
	"fmt"

	"iam/pkg/config"

	_ "github.com/go-sql-driver/mysql"
	"github.com/jmoiron/sqlx" // sqlx 扩展库
	"github.com/spf13/viper"
	"go.uber.org/zap"
)

var Db *sqlx.DB

func Init(cfg *config.MySQLConfig) (err error) {
	dsn := fmt.Sprintf("%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.DBName)

	Db, err = sqlx.Connect("mysql", dsn)
	if err != nil {
		zap.L().Error("open mysql failed", zap.Error(err))
		return err
	}

	Db.SetMaxOpenConns(viper.GetInt("mysql.max_open_conns"))
	Db.SetMaxIdleConns(viper.GetInt("mysql.max_idle_conns"))

	return
}

// 小技巧，config内部封装一个close方法对外关闭，无需暴露mysql实例
func Close() {
	_ = Db.Close()
}
