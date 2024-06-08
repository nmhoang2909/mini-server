package main

import (
	"fmt"
	"net/http"
	"strings"
)

type PlayerStore interface {
	GetPlayerScore(name string) int
	RecordWin(name string)
}

type PlayerServer struct {
	stores PlayerStore
}

func (svc *PlayerServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	router := http.NewServeMux()

	router.Handle("/league", http.HandlerFunc(svc.leagueHandler))
	router.Handle("/players/", http.HandlerFunc(svc.playersHandler))

	router.ServeHTTP(w, r)
}

func (svc *PlayerServer) showScore(name string, w http.ResponseWriter) {
	score := svc.stores.GetPlayerScore(name)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, score)
}

func (svc *PlayerServer) processWin(name string, w http.ResponseWriter) {
	svc.stores.RecordWin(name)
	w.WriteHeader(http.StatusAccepted)
}

func (svc *PlayerServer) leagueHandler(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
}
func (svc *PlayerServer) playersHandler(w http.ResponseWriter, r *http.Request) {
	playerName := strings.TrimPrefix(r.URL.Path, "/players/")
	switch r.Method {
	case http.MethodPost:
		svc.processWin(playerName, w)
	case http.MethodGet:
		svc.showScore(playerName, w)
	}
}
