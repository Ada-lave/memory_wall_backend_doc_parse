package logs

import (
	"io"
	"os"
	"github.com/gin-gonic/gin"
)

func InitLogger() {
	var file *os.File
	var logFile string = "logs/gin.log"
	_, err := os.Stat("logs/gin.log")
	if os.IsNotExist(err) {
		file, err = os.Create(logFile)
		if err != nil {
			panic(err)
		}
	} else {
		file, err = os.OpenFile(logFile, os.O_APPEND|os.O_WRONLY, 0644)
		if err != nil {
			panic(err)
		}
		file.Write([]byte("\n\n\n"))
	}

	gin.DefaultWriter = io.MultiWriter(file, os.Stdin)
}