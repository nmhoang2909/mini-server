package main

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"

	"github.com/nmhoang2909/mini-server/stub"
)

func TestGETPlayer(t *testing.T) {
	svc := &PlayerServer{
		stores: &stub.StubPlayerStore{
			Scores: map[string]int{
				"Pepper": 20,
				"Floyd":  10,
			},
		},
	}
	t.Run("get player Pepper's score", func(t *testing.T) {
		want := "20"
		request, err := http.NewRequest(http.MethodGet, "/players/Pepper", nil)
		if err != nil {
			t.Fatal("failed to new http request")
		}
		response := httptest.NewRecorder()
		svc.ServeHTTP(response, request)
		got := response.Body.String()
		assert.Equal(t, want, got)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("get player Floyd's score", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Floyd", nil)
		response := httptest.NewRecorder()
		svc.ServeHTTP(response, request)
		want := "10"
		got := response.Body.String()
		assert.Equal(t, want, got)
		assert.Equal(t, http.StatusOK, response.Code)
	})

	t.Run("return 404 on missing players", func(t *testing.T) {
		request, _ := http.NewRequest(http.MethodGet, "/players/Apollo", nil)
		response := httptest.NewRecorder()
		svc.ServeHTTP(response, request)
		want := http.StatusNotFound
		assert.Equal(t, want, response.Code)
	})

}

func TestStoreWins(t *testing.T) {
	t.Run("it returns accepted on POST", func(t *testing.T) {
		store := stub.StubPlayerStore{
			Scores: map[string]int{},
		}
		server := &PlayerServer{stores: &store}

		request, _ := http.NewRequest(http.MethodPost, "/players/Pepper", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		want := http.StatusAccepted
		assert.Equal(t, want, response.Code)
	})

	t.Run("it records wins when POST", func(t *testing.T) {
		store := stub.StubPlayerStore{
			Scores:   map[string]int{},
			WinCalls: []string{},
		}
		server := &PlayerServer{stores: &store}

		request, _ := http.NewRequest(http.MethodPost, "/players/Pepper", nil)
		response := httptest.NewRecorder()
		server.ServeHTTP(response, request)
		want := http.StatusAccepted
		assert.Equal(t, want, response.Code)
		assert.Equal(t, 1, len(store.WinCalls))
		assert.Equal(t, "Pepper", store.WinCalls[0])
	})
}
