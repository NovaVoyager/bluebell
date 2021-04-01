package redis

import (
	"fmt"

	"github.com/miaogu-go/bluebell/settings"

	"github.com/go-redis/redis"
)

var rdb *redis.Client

func Init(redisConf *settings.RedisConf) {
	rdb = redis.NewClient(&redis.Options{
		Addr:     fmt.Sprintf("%s:%d", redisConf.Host, redisConf.Port),
		Password: redisConf.Password,
		DB:       redisConf.Db,
		PoolSize: redisConf.PoolSize,
	})
	_, err := rdb.Ping().Result()
	if err != nil {
		panic(err)
	}
}

func Close() {
	_ = rdb.Close()
}
