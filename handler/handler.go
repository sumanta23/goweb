package handler

import (
	"context"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/go-redis/redis/v8"
)

var ctx = context.Background()
var rdb = redis.NewClient(&redis.Options{
	Addr:     "localhost:6379",
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
	val := <- valCh
	c.String(http.StatusOK, "count is %s", val)
}

// SetCounter is Function get get counter via redis
func SetCounter(c *gin.Context) {
	err := rdb.Incr(ctx, "counter").Err()
	if err != nil {
		panic(err)
	}
	c.String(http.StatusOK, "Incrementaed the value")
}
