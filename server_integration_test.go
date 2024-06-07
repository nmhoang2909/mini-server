package main

import (
	"fmt"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestRecordingWinsAndRetrievingThem(t *testing.T) {
	store := NewInMemoryStore()
	server := &PlayerServer{store}

	server.ServeHTTP(httptest.NewRecorder(), postWin("Pepper"))
	server.ServeHTTP(httptest.NewRecorder(), postWin("Pepper"))
	server.ServeHTTP(httptest.NewRecorder(), postWin("Pepper"))

	request, _ := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
	response := httptest.NewRecorder()
	server.ServeHTTP(response, request)

	assert.Equal(t, http.StatusOK, response.Code)
	want := "3"
	assert.Equal(t, want, response.Body.String())
}

func postWin(name string) *http.Request {
	request, _ := http.NewRequest(http.MethodPost, fmt.Sprintf("/players/%s", name), nil)
	return request
}
