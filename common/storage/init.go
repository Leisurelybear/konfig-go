package storage

import (
	"konfig-go/common/storage/mysql"
	"konfig-go/conf"
)

func Init(config conf.Config) {
	mysql.Cli, _ = mysql.NewClient(config)
}
