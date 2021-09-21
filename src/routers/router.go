package routers

import (
	h "myapp/src/handler"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddRoutes(router *gin.Engine) {
	router.GET("/count", h.GetCounter)
	router.POST("/count", h.SetCounter)
}

func AddStaticRoutes(router *gin.Engine) {
	router.LoadHTMLGlob("src/static/templates/*.tmpl.html")
	router.Static("/static", "src/static/")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})
}
