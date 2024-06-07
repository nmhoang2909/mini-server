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
	method := r.Method
	player := strings.TrimPrefix(r.URL.Path, "/players/")
	if method == http.MethodPost {
		svc.stores.RecordWin(player)
		w.WriteHeader(http.StatusAccepted)
		return
	}
	score := svc.stores.GetPlayerScore(player)
	if score == 0 {
		w.WriteHeader(http.StatusNotFound)
	}
	fmt.Fprint(w, score)
}
