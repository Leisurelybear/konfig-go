package conf

import (
	"github.com/gin-gonic/gin"
)

type Config struct {
	MySQLConfig mySQLConfig `toml:"mysql"`
	RedisConfig redisConfig `toml:"redis"`
	AppConfig   appConfig   `toml:"app"`
}

type mySQLConfig struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Username string `toml:"username"`
	Password string `toml:"password"`
	Database string `toml:"database"`
}

type redisConfig struct {
	Host     string `toml:"host"`
	Port     int    `toml:"port"`
	Password string `toml:"password"`
	DB       int    `toml:"db"`
}

type appConfig struct {
	Mode string `toml:"mode"`
	Port string `toml:"port"`
}

var defaultConfig = Config{
	MySQLConfig: mySQLConfig{
		Host:     "localhost",
		Port:     3306,
		Username: "root",
		Password: "password",
		Database: "mydb",
	},
	RedisConfig: redisConfig{
		Host:     "localhost",
		Port:     6379,
		Password: "password",
		DB:       0,
	},
	AppConfig: appConfig{
		Port: ":8080",
		Mode: gin.ReleaseMode,
	},
}
