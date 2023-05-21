package mysql

import (
	"konfig-go/common/storage/mysql"
	"konfig-go/database/mysql/collection"
	"konfig-go/database/mysql/config"
)

var (
	ConfigDAO     *config.ConfigDAO
	CollectionDAO *collection.CollectionDAO
)

func Init() {
	ConfigDAO = config.NewConfigDAO(mysql.Cli.DB)
	CollectionDAO = collection.NewCollectionDAO(mysql.Cli.DB)
}
