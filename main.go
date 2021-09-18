package main

import (
	h "myapp/handler"
	"net/http"

	"github.com/gin-gonic/gin"
	stats "github.com/semihalev/gin-stats"
)

func main() {
	router := gin.Default()

	router.Use(stats.RequestStats())

	router.GET("/stats", func(c *gin.Context) {
		c.JSON(http.StatusOK, stats.Report())
	})

	router.GET("/", h.GetCounter)
	router.POST("/", h.SetCounter)
	router.Run(":8080")
}
