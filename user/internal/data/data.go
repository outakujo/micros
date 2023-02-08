package data

import (
	"fmt"
	"user/internal/conf"
	"user/internal/pkg/sync"

	"github.com/go-kratos/kratos/v2/log"
	"github.com/google/wire"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

// ProviderSet is data providers.
var ProviderSet = wire.NewSet(NewData, NewUserRepo)

// Data .
type Data struct {
	*gorm.DB
}

// NewData .
func NewData(c *conf.Data, logger log.Logger) (*Data, func(), error) {
	sync.InitRedis(c.Redis)
	cleanup := func() {
		log.NewHelper(logger).Info("closing the data resources")
	}
	dsn := "%s:%s@tcp(%s)/%s?charset=utf8mb4&parseTime=True&loc=Local"
	database := c.Database
	db, err := gorm.Open(mysql.Open(fmt.Sprintf(dsn, database.User,
		database.Pass, database.Addr, database.Db)), &gorm.Config{})
	if err != nil {
		return nil, nil, err
	}
	return &Data{db}, cleanup, nil
}
