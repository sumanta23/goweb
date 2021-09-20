package handler

import (
	"context"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
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

func getCounterFromRedis(ch chan string) {
	val, err := rdb.Get(ctx, "counter").Result()
	if err != nil {
		ch <- "1"
	}
	ch <- val
}

// GetCounter is Function get get counter via redis
func GetCounter(c *gin.Context) {
	valCh := make(chan string)
	go getCounterFromRedis(valCh)
	val := <-valCh
	mapD := map[string]string{"count": val}
	c.JSON(http.StatusOK, mapD)
}

// SetCounter is Function get get counter via redis
func SetCounter(c *gin.Context) {
	err := rdb.Incr(ctx, "counter").Err()
	if err != nil {
		panic(err)
	}
	mapD := map[string]bool{"success": true}
	c.JSON(http.StatusOK, mapD)
}
