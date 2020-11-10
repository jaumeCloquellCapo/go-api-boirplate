package provider

import (
	"context"
	"fmt"
	"github.com/go-redis/redis/v8"
	"os"
	"strconv"
)

var REDIS_CTX = context.Background()

//InitializeCache
func InitializeCache() (rdb *redis.Client) {
	db, _ := strconv.Atoi(os.Getenv("REDIS_DATABASE"))
	rdb = redis.NewClient(&redis.Options{
		Network:  "tcp",
		Addr:     fmt.Sprintf(os.Getenv("REDIS_HOST") + ":" + os.Getenv("REDIS_PORT")),
		Password: os.Getenv("REDIS_PASSWORD"),
		DB:       db,
	})
	return rdb
}
