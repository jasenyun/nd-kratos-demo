package data

import (
	"github.com/go-kratos/kratos/v2/log"
	"github.com/go-redis/redis/v8"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"jasenServer/internal/conf"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewGreeterRepo)

// Data .
type Data struct {
	db  *gorm.DB
	rdb *redis.Client
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	d := &Data{}
	if c.Database.Driver == "mysql" {
		gormDB, err := gorm.Open(mysql.Open(c.Database.Source), &gorm.Config{})
		if err != nil {
			log.NewHelper(logger).Error("mysql init fail")
		}

		d.db = gormDB
	}

	if c.Redis != nil {
		rdb := redis.NewClient(&redis.Options{
			Addr: c.Redis.Addr,
		})
		if err := rdb.Close(); err != nil {
			log.NewHelper(logger).Error(err)
		}

		d.rdb = rdb
	}

	return d, cleanup, nil
}
