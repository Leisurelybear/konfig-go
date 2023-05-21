package mysql

import (
	"context"
	"fmt"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"konfig-go/common/logger"
	"konfig-go/conf"
)

var Cli *Client

type Client struct {
	DB *gorm.DB
}

func NewClient(config conf.Config) (*Client, error) {
	ctx := context.Background()
	cfg := config.MySQLConfig
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%d)/%s?charset=utf8mb4&parseTime=True&loc=Local",
		cfg.Username,
		cfg.Password,
		cfg.Host,
		cfg.Port,
		cfg.Database,
	)
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		logger.Logger.Error(ctx, "errorutils to connect DB,")
		return nil, err
	}

	return &Client{
		DB: db,
	}, nil
}

func (c *Client) Close() error {
	sqlDB, err := c.DB.DB()
	if err != nil {
		return err
	}

	return sqlDB.Close()
}
