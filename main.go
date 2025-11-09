package main

import (
	"httpserver/internal/http_server"
	"log"
	"net/http"
)

func main() {
	server := &httpserver.PlayerServer{httpserver.NewInmemoryPlayerStore()}
	log.Fatal(http.ListenAndServe(":5000", server))
}
