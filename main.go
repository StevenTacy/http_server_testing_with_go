package main

import (
	"httpserver/internal/http_server"
	"log"
	"net/http"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	// O_RDWR read & write flags | O_CREATE create the file if not exist
	// 0666 set permission on the created file
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)
	if err != nil {
		log.Fatalf("problem opening %s %v", dbFileName, err)
	}

	store, err := httpserver.NewFileSystemPlayerStore(db)
	if err != nil {
		log.Fatalf("problem creating file system player store %v", err)
	}
	server := httpserver.NewPlayerServer(store)
	if err := http.ListenAndServe(":5000", server); err != nil {
		log.Fatalf("could not listen on port 5000 %v", err)
	}
}
