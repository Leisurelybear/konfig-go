package main

import (
	"konfig-go/conf"
	"konfig-go/router"
)

func Serve(conf *conf.Config) {
	// Listen and Server in 0.0.0.0:8080
	err := router.Init().Run(conf.AppConfig.Port)
	if err != nil {
		return
	}
}

func main() {
	config := conf.LoadConf()
	Serve(&config)
}
