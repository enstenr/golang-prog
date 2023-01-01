package main

import (
	 
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"os"
	"time"
)

const path = "entries.json"

// raffleEntry is the struct we unmarshal raffle entries into
type raffleEntry struct {
		ID      string `json:"id"`
		Name    string `json:"name"`
		Country string `json:"country,omitempty"`
}

// importData reads the raffle entries from file and creates the entries slice.
func importData() []raffleEntry {
	var data []raffleEntry
	jsonFile, err := os.ReadFile(path)
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	 
	if err != nil {
		// handle error
	}
	 
	if err := json.Unmarshal(jsonFile, &data); err != nil {
		//handle error
	}
	return data

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