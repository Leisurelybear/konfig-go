package main

import (
	"konfig-go/conf"
	"konfig-go/server"
)

func main() {
	conf := conf.LoadConf()
	server.Serve(&conf)
}
