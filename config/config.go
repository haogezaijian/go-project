package config

import (
	"github.com/spf13/viper"
	"log"
	// 其他可能的导入包
)

type Config struct {
	App struct {
		Name string
		Port string
	}

	Database struct {
		Dsn          string
		MaxIdleConns int
		MaxOpenconns int
	}
}

var AppConfig *Config

func InitConfig() {
	viper.SetConfigName("config")
	viper.SetConfigType("yml")
	viper.AddConfigPath("./config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatal("读取配置文件出错: %v", err)
	}

	AppConfig = &Config{}

	if err := viper.Unmarshal(AppConfig); err != nil {
		log.Fatal("读取配置文件到结构出错: %v", err)
	}

	initDB()
	InitRedis()
}
