package test

import (
	"github.com/google/wire"
	"github.com/katsuokaisao/gorm/config"
	"github.com/katsuokaisao/gorm/infra/rdb"
)

var dbSet = wire.NewSet(
	rdb.NewRDB,
	provideRDBConfig,
)

func provideRDBConfig(cfg *config.Config) rdb.Config {
	return rdb.Config{
		Driver:   cfg.DB.Driver,
		Address:  cfg.DB.Address,
		Username: cfg.DB.Username,
		Password: cfg.DB.Password,
		Database: cfg.DB.Database,
		Debug:    cfg.DB.Debug,
	}
}
