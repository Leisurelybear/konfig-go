package main

import (
	"konfig-go/common/messagequeue"
	"konfig-go/common/storage"
	"konfig-go/conf"
	"konfig-go/database/mysql"
	"konfig-go/server/grpc"
	"konfig-go/server/http"
)

func Serve(conf conf.Config) {
	go grpc.Run(conf)
	http.Run(conf)
}

func main() {
	defer Close()

	config := conf.LoadConf()

	// init mysql
	storage.Init(config)

	// init dao
	mysql.Init()

	// initProducer
	messagequeue.Init()

	Serve(config)
}

func Close() {

}
