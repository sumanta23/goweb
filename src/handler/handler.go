package handler

import (
	"fmt"
	redis "myapp/src/io/redis"
	"myapp/src/wsmanager"
	"net/http"

	"github.com/gin-gonic/gin"
)

// GetCounter is Function get get counter via redis
func GetCounter(c *gin.Context) {
	valCh := make(chan string)
	go redis.GetCounter(valCh)
	val := <-valCh
	mapD := map[string]string{"count": val}
	c.JSON(http.StatusOK, mapD)
}

// SetCounter is Function get get counter via redis
func SetCounter(hub *wsmanager.Hub) func(c *gin.Context) {
	return func(c *gin.Context) {
		val := redis.IncrCounter()
		hub.Broadcast(fmt.Sprint(val))
		mapD := map[string]bool{"success": true}
		c.JSON(http.StatusOK, mapD)
	}
}

func ResetCounter(hub *wsmanager.Hub) func(c *gin.Context) {
	return func(c *gin.Context) {
		val := redis.DeleteCounter()
		hub.Broadcast(fmt.Sprint(val))
		mapD := map[string]bool{"success": true}
		c.JSON(http.StatusOK, mapD)
	}
}
