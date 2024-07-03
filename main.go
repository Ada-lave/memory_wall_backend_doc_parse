package main

import (
	memorywall "memory_wall/internal/http/memory_wall"

	"github.com/gin-gonic/gin"
)

func main() {
	server := gin.Default()
	memorywall.InitMemoryWallRouter(server)

	if err := server.Run(":8081"); err != nil {
		panic(err)
	}
}
