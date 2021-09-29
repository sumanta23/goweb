package wsmanager

import (
	"log"
	"net/http"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
	"github.com/gorilla/websocket"
)

func checkOrigin(*http.Request) bool {
	return false
}

var upgrader = websocket.Upgrader{
	ReadBufferSize:    1024,
	WriteBufferSize:   1024,
	EnableCompression: true,
	CheckOrigin: func(r *http.Request) bool {
		origin := r.Header.Get("Origin")
		log.Println("origin: %v", origin)
		return true
	},
}

func Upgrade(hub *Hub) func(c *gin.Context) {
	return func(c *gin.Context) {
		ip, _ := c.RemoteIP()
		log.Println("upgrading conn: %v", ip)
		conn, err := upgrader.Upgrade(c.Writer, c.Request, nil)
		if err != nil {
			log.Println("Failed to set websocket upgrade: %v", err)
			return
		}

		client := &Client{hub: hub, conn: conn, send: make(chan []byte, 256), cid: uuid.NewString()}
		client.hub.register <- client

		// Allow collection of memory referenced by the caller by doing all work in
		// new goroutines.
		go client.writePump()
		go client.readPump()
	}
}

func GetWSURL(c *gin.Context) {
	mapD := map[string]string{"ws": "/ws"}
	c.JSON(http.StatusOK, mapD)
}
