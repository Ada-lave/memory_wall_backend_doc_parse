package main

import (
	memorywall "memory_wall/internal/http/memory_wall"

	"github.com/gin-gonic/gin"
)

type MemoryWallServer struct {
	ginServer *gin.Engine

}

func (MS *MemoryWallServer) New() {
	MS.ginServer = gin.Default()
}

func (MS *MemoryWallServer) Start() {
	if err := MS.ginServer.Run(); err != nil {
		panic(err)
	}
}

func main() {
	var server MemoryWallServer
	var memoryWallNET memorywall.MemoryWallNET
	server.New()
	memoryWallNET.Start(server.ginServer)
	server.Start()
}