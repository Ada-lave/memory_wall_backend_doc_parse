package http_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	memorywall "memory_wall/internal/http/memory_wall"
	"memory_wall/internal/http/memory_wall/models"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetupEngine() *gin.Engine {
	server := gin.Default()
	memorywall.InitMemoryWallRouter(server)
	server.GET("/ping", func(c *gin.Context) {
		c.String(http.StatusOK, "pong")
	})

	return server
}

func TestPingRoute(t *testing.T) {
	server := SetupEngine()

	w := httptest.NewRecorder()

	req, _ := http.NewRequest("GET", "/ping", nil)

	server.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	assert.Equal(t, "pong", w.Body.String())
}

func TestParseDocx(t *testing.T) {
	server := SetupEngine()

	body := &bytes.Buffer{}

	writer := multipart.NewWriter(body)
	formDataPart, err := writer.CreateFormFile("files", "test.docx")
	if err != nil {
		panic(err)
	}

	testFile, err := os.ReadFile("../data/docs/test.docx")
	if err != nil {
		panic(err)
	}

	_, err = formDataPart.Write(testFile)

	if err != nil {
		panic(err)
	}

	writer.Close()

	responseRecorder := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/parse/docx", body)

	req.Header.Set("Content-Type", writer.FormDataContentType())

	server.ServeHTTP(responseRecorder, req)

	var response models.ParseDocxResponse

	err = json.Unmarshal(responseRecorder.Body.Bytes(), &response)

	if err != nil {
		panic(err)
	}

	fmt.Println(responseRecorder.Body)
	assert.Equal(t, 200, responseRecorder.Code)
	assert.Equal(t, "test.docx", response.Filename)
}
