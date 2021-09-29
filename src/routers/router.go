package routers

import (
	h "myapp/src/handler"
	wsmanager "myapp/src/wsmanager"
	"net/http"

	"github.com/gin-gonic/gin"
)

func AddRoutes(router *gin.Engine, hub *wsmanager.Hub) {
	router.GET("/count", h.GetCounter)
	router.POST("/count", h.SetCounter(hub))
}

func AddWS(router *gin.Engine, hub *wsmanager.Hub) {
	router.GET("/getwsurl", wsmanager.GetWSURL)
	router.GET("/ws", wsmanager.Upgrade(hub))
}

func AddStaticRoutes(router *gin.Engine) {
	router.LoadHTMLGlob("src/static/templates/*.tmpl.html")
	router.Static("/static", "src/static/")
	router.GET("/", func(c *gin.Context) {
		c.HTML(http.StatusOK, "index.tmpl.html", nil)
	})
}
