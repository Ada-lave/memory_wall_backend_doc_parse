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
	testCases := []struct {
		Name string
		File string
		Want models.ParseDocxResponse
	}{
		{
			Name: "Base test",
			File: "../data/docs/test.docx",
			Want: models.ParseDocxResponse{
				Filename: "test.docx",
				HumanInfo: models.HumanInfo{
					FirstName:    "ВИКТОР",
					LastName:     "ДУБРОВСКИХ",
					MiddleName:   "ЕГОРОВИЧ",
					Birthday:     "1922-06-12 00:00:00 +0000 UTC",
					Deathday:     "1993-07-04 00:00:00 +0000 UTC",
					MilitaryRank: "старший лейтенант; авиационный моторист, старший писарь по учету самолетов и моторов (после тяжелого ранения).",
				},
			},
		},
	}

	for _, tc := range testCases {
		t.Run(tc.Name, func(t *testing.T) {
			server := SetupEngine()
			var body bytes.Buffer
			var writer *multipart.Writer = multipart.NewWriter(&body)
			var filesData map[string]io.Reader = make(map[string]io.Reader)

			err := loadFile(&filesData, tc.File)
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
			assert.Equal(t, tc.Want.HumanInfo.FirstName, response.Data[0].FirstName)
			assert.Equal(t, tc.Want.HumanInfo.LastName, response.Data[0].LastName)
			assert.Equal(t, tc.Want.HumanInfo.MiddleName, response.Data[0].MiddleName)
			assert.Equal(t, tc.Want.HumanInfo.Birthday, response.Data[0].Birthday)
			assert.Equal(t, tc.Want.HumanInfo.Deathday, response.Data[0].Deathday)
			assert.Equal(t, tc.Want.HumanInfo.MilitaryRank, response.Data[0].MilitaryRank)
		})
	}
}

func loadFile(data *map[string]io.Reader, file string) error {

	f, err := os.Open(path.Join(file))
	if err != nil {
		return err
	}
	(*data)[file] = f

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
