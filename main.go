package main

import (
	memorywall "memory_wall/internal/http/memory_wall"
	"net/http"

	"github.com/gin-gonic/gin"
)

func setupEngine() *gin.Engine {
	server := gin.Default()
	server.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	return server
}

func main() {
	server := setupEngine()
	memorywall.InitMemoryWallRouter(server)

	if err := server.Run(":8081"); err != nil {
		panic(err)
	}
}
