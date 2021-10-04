package redis

import (
	"context"
	"os"

	"github.com/go-redis/redis/v8"
)

func getRedisHostName() string {
	var redisServer = os.Getenv("REDIS")
	if redisServer == "" {
		redisServer = "localhost"
	}
	return redisServer
}

var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
	Addr:     getRedisHostName() + ":6379",
	Password: "", // no password set
	DB:       0,  // use default DB
})

func GetCounter(ch chan string) {
	val, err := rdb.Get(ctx, "counter").Result()
	if err != nil {
		ch <- "1"
	}
	ch <- val
}

func IncrCounter() int64 {
	val, err := rdb.Incr(ctx, "counter").Result()
	if err != nil {
		panic(err)
	}
	return val
}

func DeleteCounter() int64 {
	val, err := rdb.Del(ctx, "counter").Result()
	if err != nil {
		panic(err)
	}
	return val
}
