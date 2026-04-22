package config

import (
	"fmt"

	"github.com/fsnotify/fsnotify"
	"github.com/spf13/viper"
)

// 用来保存程序的所以配置信息的全局变量---其他包可用的全局变量
// Conf 全局变量
var Conf = new(AppConfig)

type AppConfig struct {
	Name      string `mapstructure:"name"`
	Mode      string `mapstructure:"mode"`
	Port      int    `mapstructure:"port"`
	Version   string `mapstructure:"version"`
	StartTime string `mapstructure:"start_time"`
	MachineId int64  `mapstructure:"machine_id"`
	

	*LogConfig      `mapstructure:"log"`
	*PostgreSQLConfig `mapstructure:"postgres"`
	*RedisConfig    `mapstructure:"redis"`
}

type LogConfig struct {
	Level      string `mapstructure:"level"`
	Filename   string `mapstructure:"filename"`
	MaxSize    int    `mapstructure:"max_size"`
	MaxAge     int    `mapstructure:"max_age"`
	MaxBackups int    `mapstructure:"max_backups"`
}

type PostgreSQLConfig struct {
	DSN string `mapstructure:"dsn"`
	MaxOpenConns int    `mapstructure:"max_open_conns"`
    MaxIdleConns int    `mapstructure:"max_idle_conns"`
    LogLevel     string `mapstructure:"log_level"`
}

type RedisConfig struct {
	Host       string `mapstructure:"host"`
	Port       int    `mapstructure:"port"`
	DB         int    `mapstructure:"db"`
	PoolSize   int    `mapstructure:"pool_size"`
	MaxRetries int    `mapstructure:"max_retries"`
}

func Init() (err error) {
	//------------下面查找配置文件的方法就使用这个---------------------
	// 设置配置文件名（不需要扩展名，Viper会根据扩展名自动识别）
	viper.SetConfigFile("./configs/config.yaml")
	//用于从etcd等从远程获取配置文件，不是对上面的配置文件限制
	viper.SetConfigType("yml")
	// 添加配置文件搜索路径 - 关键在这里
	viper.AddConfigPath("./configs/") // 相对于执行目录

	err = viper.ReadInConfig() // 读取配置信息
	if err != nil {
		fmt.Println("viper.ReadInConfig() error config file: %s\n", err)
		return
	}

	// 把读取到的配置反序列化到结构体当中---可以在程序使用Conf获取配置信息
	if err := viper.Unmarshal(Conf); err != nil {
		fmt.Println("viper.Unmarshal(Conf) err failed : %s\n", err)
	}

	viper.WatchConfig()
	viper.OnConfigChange(func(in fsnotify.Event) {
		fmt.Println("配置文件被修改了...")
		if err := viper.Unmarshal(Conf); err != nil { // 核心实现逻辑就是配置文件改变被识别后，直接再次把新配置属性反序列化给结构体即可
			fmt.Println("viper.Unmarshal(Conf) err failed : %s\n", err)
		}
	})

	return
}
