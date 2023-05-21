package main

import (
	"konfig-go/common/storage/mysql"
	"konfig-go/conf"
	"konfig-go/server/grpc"
	"konfig-go/server/http"
)

func Serve(conf conf.Config) {
	go grpc.Run(conf)
	http.Run(conf)
}

func main() {
	config := conf.LoadConf()

	// init mysql
	mysql.Cli, _ = mysql.NewClient(config)

	defer Close()
	Serve(config)
}

func Close() {
	err := mysql.Cli.Close()
	if err != nil {

	}
}
