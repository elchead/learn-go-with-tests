package main

import (
	"fmt"
	"log"
	"os"

	"github.com/elchead/poker"
)

const dbFileName = "game.db.json"

func main() {
	fmt.Println("Lets play poker")
	db, err := os.OpenFile(dbFileName, os.O_RDWR|os.O_CREATE, 0666)

	store, err := poker.NewFileSystemPlayerStore(db)

	if err != nil {
		log.Fatalf("problem creating file system player store, %v ", err)
	}

	game := poker.CLI{store, os.Stdin}
	game.PlayPoker()
}
