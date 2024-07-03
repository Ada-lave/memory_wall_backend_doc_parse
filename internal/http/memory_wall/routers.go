package memorywall

import "github.com/gin-gonic/gin"



func InitMemoryWallRouter(e *gin.Engine) {

	group := e.Group("parse")
	service := newMemoryWallService()
	controller := NewMemoryWallController(service)
	
	group.POST("/docx", controller.ParseDocx)
} 

