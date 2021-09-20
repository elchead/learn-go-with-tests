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
	store, closeDb, err := poker.FileSystemPlayerStoreFromFile(dbFileName) //poker.NewFileSystemPlayerStore(db)
	defer closeDb()
	if err != nil {
		log.Fatalf("problem creating file system player store, %v ", err)
	}

	alerter := poker.BlindAlerterFunc(poker.Alerter)
	poker.NewCLI(poker.NewTexasHoldem(store, alerter), os.Stdin, os.Stdout).PlayPoker()
}
