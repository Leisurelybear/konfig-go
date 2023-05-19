package server

import (
	"konfig-go/conf"
	"konfig-go/router"
)

func Serve(conf *conf.Config) {
	// Listen and Server in 0.0.0.0:8080
	err := router.Router().Run(conf.AppConfig.Port)
	if err != nil {
		return
	}
}
