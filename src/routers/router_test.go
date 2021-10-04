package routers

import (
	wsmanager "myapp/src/wsmanager"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestAddRoutes(t *testing.T) {

	router := gin.Default()
	hub := wsmanager.NewHub()

	AddRoutes(router, hub)
	routesInfo := router.Routes()

	sizeOfSlice := len(routesInfo)

	assert.Equal(t, sizeOfSlice, 3)
}

func TestAddWS(t *testing.T) {

	router := gin.Default()
	hub := wsmanager.NewHub()

	AddWS(router, hub)
	routesInfo := router.Routes()

	sizeOfSlice := len(routesInfo)

	assert.Equal(t, sizeOfSlice, 2)
}
