package memorywall

import "github.com/gin-gonic/gin"

type MemoryWallRouter struct {
	Router *gin.RouterGroup
}

func (MR *MemoryWallRouter) New(e *gin.Engine) {
	MR.Router = e.Group("parse")
} 