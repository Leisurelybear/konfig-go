package router

import (
	"konfig-go/handler/collection"
	"net/http"

	"github.com/gin-gonic/gin"
)

func Init() *gin.Engine {
	// Disable Console Color
	// gin.DisableConsoleColor()
	r := gin.Default()

	// Ping test
	r.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	collectionHandler := r.Group("/collection")
	{
		collectionHandler.GET("/list", collection.List)
	}

	return r
}
