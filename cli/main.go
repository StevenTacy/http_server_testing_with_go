package main

import (
	"fmt"
	httpserver "httpserver/internal/http_server"
	"log"
	"os"
)

const dbFileName = "game.db.json"

func main() {
	fmt.Println("Let's play poker")
	fmt.Println("Type {Name} wins to record a win")

	store, closeFunc, err := httpserver.FileSystemPlayerStoreFromFile(dbFileName)
	if err != nil {
		log.Fatal(err)
	}
	defer closeFunc()

	game := httpserver.NewCLI(store, os.Stdin)
	game.PlayPoker()
}
