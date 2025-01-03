package configs

import (
	"log"

	"github.com/spf13/viper"
)

type Config struct {
	Port string
	DB   DataBase
	Log  Log
}
type DataBase struct {
	Name     string
	Host     string
	Port     string
	User     string
	Password string
}
type Log struct {
	Level      string
	MaxSize    int
	MaxBackups int
	MaxAge     int
}

var AppConfig *Config

func InitConfig() {
	// Search config in home directory with name ".config" (without extension).
	viper.AddConfigPath("./configs")
	viper.SetConfigType("yaml")
	viper.SetConfigName("config")

	if err := viper.ReadInConfig(); err != nil {
		log.Fatalf("fatal error: read config file: %v", err)
	}
	if err := viper.Unmarshal(&AppConfig); err != nil {
		log.Fatalf("set config file error: %v", err)
	}
}
