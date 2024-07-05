package http_test

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io"
	memorywall "memory_wall/internal/http/memory_wall"
	"memory_wall/internal/http/memory_wall/models"
	"mime/multipart"
	"net/http"
	"net/http/httptest"
	"os"
	"path"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

type serverResponse struct {
	Data []models.ParseDocxResponse `json:"data"`
}

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

	var body bytes.Buffer
	var writer *multipart.Writer = multipart.NewWriter(&body)
	var filesData map[string]io.Reader = make(map[string]io.Reader)

	err := loadFiles(&filesData, "../data/docs")
	if err != nil {
		panic(err)
	}
	// fmt.Printf("%#v\n", filesData)
	loadFilesToBody(writer, filesData)


	responseRecorder := httptest.NewRecorder()
	req, _ := http.NewRequest("POST", "/parse/docx", &body)

	req.Header.Set("Content-Type", writer.FormDataContentType())

	server.ServeHTTP(responseRecorder, req)

	
	var response serverResponse

	err = json.Unmarshal(responseRecorder.Body.Bytes(), &response)

	if err != nil {
		panic(err)
	}

	
	fmt.Printf("%#v\n", response.Data[0].Filename)
	assert.Equal(t, 200, responseRecorder.Code)
	// assert.Equal(t, "test.docx", response.Filename)
}

func loadFiles(data *map[string]io.Reader, filesDir string) error {
	files, err := os.ReadDir(filesDir)

	for _, file := range files {
		if !file.IsDir() {
			f, err := os.Open(path.Join(filesDir, "/", file.Name())) 
			if err != nil {
				return err
			}
			(*data)[file.Name()] = f
		}
	}
	if err != nil {
		return err
	}
	return nil
}

func loadFilesToBody(writer *multipart.Writer, data map[string]io.Reader) {
	for _, reader := range data {
		if x, ok := reader.(io.Closer); ok {
			defer x.Close()
		}

		if x, ok := reader.(*os.File); ok {
			// fmt.Printf("NAME: %v\n", x.Name())
			fw, err := writer.CreateFormFile("files", x.Name())
			if err != nil {
				panic(err)
			}
			_, err = io.Copy(fw, reader)
			if err != nil {
				panic(err)
			}

		}
	}

	writer.Close()
}
