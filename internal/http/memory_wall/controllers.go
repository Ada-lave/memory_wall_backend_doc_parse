package memorywall

import (
	"fmt"
	"net/http"

	"github.com/gin-gonic/gin"
)

type MemoryWallController struct {
	service *MemoryWallService
}

func NewMemoryWallController(service *MemoryWallService) *MemoryWallController {

	
	return &MemoryWallController{service: service}
}

func (MN *MemoryWallController) ParseDocx(c *gin.Context) {
	var request ParseDocxRequest
	if err := c.ShouldBind(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"validation error": err.Error()})
		return
	}

	response, err := MN.service.ParseDocx(*request.Files)

	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{
			"error": err.Error(),
		})
		fmt.Println(err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{
		"data": response,
	})
}

// func (MN *MemoryWallNET) parseAllDocxInStorage() {
// 	MN.router.Router.GET("docx/first-load", func (c *gin.Context) {
// 		names, err := MN.service.getAllDocxFileInfoFromStorage("/home/ada/Загрузки/ИНТЕРАКТИВНАЯ СТЕНА ПАМЯТИ/БЕССМЕРТНЫЙ ПОЛК. Стена памяти")
// 		if err != nil {
// 			c.JSON(http.StatusBadRequest, gin.H{
// 				"error": "falied load first data",
// 			})
// 			return
// 		}

// 		c.JSON(http.StatusOK, gin.H{
// 			"data": names,
// 		})
// 	})
// }
