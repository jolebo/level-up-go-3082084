package main

import (
	"embed"
	"encoding/json"
	"log"
	"math/rand"
	"time"
)

const path = "entries.json"

// raffleEntry is the struct we unmarshal raffle entries into
//
//go:embed entries.json
var files embed.FS

type raffleEntry struct {
	Id,Name,Country string
}

// importData reads the raffle entries from file and creates the entries slice.
func importData() []raffleEntry {
	
	data, _ := files.ReadFile(path)
	dt := []raffleEntry{}
	err := json.Unmarshal(data,&dt)
	if err != nil {
		panic("Error when unmarshaling json")
	}
	return dt
}

// getWinner returns a random winner from a slice of raffle entries.
func getWinner(entries []raffleEntry) raffleEntry {
	rand.Seed(time.Now().Unix())
	wi := rand.Intn(len(entries))
	return entries[wi]
}

func main() {
	entries := importData()
	log.Println("And... the raffle winning entry is...")
	winner := getWinner(entries)
	time.Sleep(500 * time.Millisecond)
	log.Println(winner)
}
