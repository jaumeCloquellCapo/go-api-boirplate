package provider

import (
	"ApiRest/config"
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
)

var REDIS_CTX = context.Background()

//InitializeCache
func InitializeCache(config config.Config) (rdb *redis.Client) {
	rdb = redis.NewClient(&redis.Options{
		Network:  "tcp",
		Addr:     fmt.Sprintf(config.Redis.Host + ":" + config.Redis.Port),
		Password: config.Redis.Password,
		DB:       config.Redis.DB,
	})
	return
}
