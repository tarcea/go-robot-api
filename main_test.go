package main

import (
	"io"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func SetUpRouter() *gin.Engine {
	router := gin.Default()
	return router
}

func TestGetHandler(t *testing.T) {

	mockResponse := `{"Report":{"x":1,"y":3,"orientation":"E"}}`

	r := SetUpRouter()

	r.GET("/game", GetHandler)
	// test 1
	reqOk, _ := http.NewRequest("GET", "/game?width=5&deep=5&orientation=N&x=1&y=2&command=RFRFFRFRFR", nil)
	w := httptest.NewRecorder()
	r.ServeHTTP(w, reqOk)

	responseData, _ := io.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusOK, w.Code)

	// test 2
	mockResponse = `{"message":"value for 'width' not present or wrong query provided"}`
	reqNoWidth, _ := http.NewRequest("GET", "/game?widt=5&deep=5&orientation=N&x=1&y=2&command=RFRFFRFRFR", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, reqNoWidth)

	responseData, _ = io.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// test 3
	mockResponse = `{"message":"value for 'deep' not present or wrong query provided"}`
	reqNoDeep, _ := http.NewRequest("GET", "/game?width=5&dee=5&orientation=N&x=1&y=2&command=RFRFFRFRFR", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, reqNoDeep)

	responseData, _ = io.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// test 4
	mockResponse = `{"message":"value for 'orientation' not present or wrong query provided"}`
	reqNoOrientation, _ := http.NewRequest("GET", "/game?width=5&deep=5&orientatio=N&x=1&y=2&command=RFRFFRFRFR", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, reqNoOrientation)

	responseData, _ = io.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusBadRequest, w.Code)

	// test 5
	mockResponse = `{"message":"value for 'x' not present or wrong query provided"}`
	reqNoX, _ := http.NewRequest("GET", "/game?width=5&deep=5&orientation=N&x=&y=2&command=RFRFFRFRFR", nil)
	w = httptest.NewRecorder()
	r.ServeHTTP(w, reqNoX)

	responseData, _ = io.ReadAll(w.Body)
	assert.Equal(t, mockResponse, string(responseData))
	assert.Equal(t, http.StatusBadRequest, w.Code)
}
