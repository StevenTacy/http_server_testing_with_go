package main

import (
	"httpserver/internal/http_server"
	"log"
	"net/http"
)

const dbFileName = "game.db.json"

func main() {
	// O_RDWR read & write flags | O_CREATE create the file if not exist
	// 0666 set permission on the created file
	store, closeFunc, err := httpserver.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer closeFunc()

	server := httpserver.NewPlayerServer(store)
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
