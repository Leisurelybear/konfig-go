package server

import (
	"konfig-go/conf"
	"konfig-go/router"
)

func HttpRun(config conf.Config) {
	// Listen and Server in 0.0.0.0:8080
	err := router.Init().Run(config.AppConfig.HttpPort)
	if err != nil {
		return
	}
}
