package main

import (
	h "myapp/src/handler"
	"net/http"

	"github.com/gin-gonic/gin"
	stats "github.com/semihalev/gin-stats"
)

func main() {
	router := gin.Default()

	router.Use(stats.RequestStats())
	router.Use(gin.Logger())

	router.LoadHTMLGlob("src/static/templates/*.tmpl.html")
	router.Static("/static", "src/static/")

	router.GET("/stats", func(c *gin.Context) {
		c.JSON(http.StatusOK, stats.Report())
	})

	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})

	router.GET("/count", h.GetCounter)
	router.POST("/count", h.SetCounter)
	router.Run(":8080")
}
