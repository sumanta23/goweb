package main

import (
	"net/http"

	routers "myapp/src/routers"

	"github.com/gin-gonic/gin"
	stats "github.com/semihalev/gin-stats"
)

func main() {
	router := gin.Default()

	router.Use(stats.RequestStats())
	router.Use(gin.Logger())

	router.GET("/stats", func(c *gin.Context) {
		c.JSON(http.StatusOK, stats.Report())
	})

	routers.AddStaticRoutes(router)

	routers.AddRoutes(router)

	router.Run(":8080")
}
