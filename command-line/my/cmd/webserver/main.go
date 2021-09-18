package main

import (
	"log"
	"net/http"

	"github.com/elchead/poker"
)

const dbFileName = "game.db.json"

func main() {
	store, closeDb, err := poker.FileSystemPlayerStoreFromFile(dbFileName) //poker.NewFileSystemPlayerStore(db)
	if err != nil {
		log.Fatal(err)
	}
	defer closeDb()
	server := poker.NewPlayerServer(store)

	log.Fatal(http.ListenAndServe(":5000", server))
}
