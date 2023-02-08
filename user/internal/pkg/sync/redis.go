package sync

import (
	"fmt"
	"github.com/go-redsync/redsync/v4"
	"github.com/go-redsync/redsync/v4/redis/goredis/v9"
	"github.com/redis/go-redis/v9"
	"user/internal/conf"
)

var rc *conf.Data_Redis

func InitRedis(c *conf.Data_Redis) {
	rc = c
}

func Mutex() *redsync.Mutex {
	client := redis.NewClient(&redis.Options{
		Addr: fmt.Sprintf("%v:%v", rc.Ip, rc.Port),
	})
	pool := goredis.NewPool(client)
	return redsync.New(pool).NewMutex(rc.LockKey)
}
