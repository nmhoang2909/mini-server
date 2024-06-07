package main

import (
	"log"
	"net/http"
)

func main() {
	store := NewInMemoryStore()
	svc := &PlayerServer{store}
	log.Fatal(http.ListenAndServe(":8989", svc))
}
