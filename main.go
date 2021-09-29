package main

import (
	"net/http"

	routers "myapp/src/routers"
	wsmanager "myapp/src/wsmanager"

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

	hub := wsmanager.NewHub()
	go hub.Run()

	routers.AddStaticRoutes(router)

	routers.AddRoutes(router, hub)

	routers.AddWS(router, hub)

	router.Run(":5000")
}
