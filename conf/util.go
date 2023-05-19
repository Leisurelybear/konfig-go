package conf

import (
	"github.com/pelletier/go-toml/v2"
	"log"
	"os"
)

func LoadConf() Config {
	config := defaultConfig
	// 读取配置文件
	configFile, _ := os.Open("./conf/config.toml")
	decoder := toml.NewDecoder(configFile)
	if err := decoder.Decode(&config); err != nil {
		log.Fatal("Failed to decode config file:", err)
	}
	return config
}
