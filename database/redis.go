package database

import (
	"context"
	"fmt"

	"github.com/go-redis/redis/v9"
)

var ctx = context.Background()

func GetRedis() *redis.Client {

	rdb := redis.NewClient(&redis.Options{
		Addr:     "redis:6379",
		Password: "",
		DB:       0,
	})

	pong, err := rdb.Ping(ctx).Result()
	fmt.Println("🔴 Redis ping: ", pong, err)

	return rdb
}
