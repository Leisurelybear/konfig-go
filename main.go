package main

import (
	"konfig-go/conf"
	"konfig-go/server"
)

func Serve(conf conf.Config) {
	go server.GrpcRun(conf)
	server.HttpRun(conf)
}

func main() {
	config := conf.LoadConf()
	Serve(config)
}
