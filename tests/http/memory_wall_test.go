package http_test

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)


func SetupEngine() *gin.Engine {
	server := gin.Default()
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


