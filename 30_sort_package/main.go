package main

import (
	"encoding/json"
	"fmt"
	"log"
	"os"
	"sort"
	"strings"
)


const path = "items.json"
// SaleItem represents the item part of the big sale.
type SaleItem struct {
	Name           string  `json:"name"`
	OriginalPrice  float64 `json:"originalPrice"`
	ReducedPrice   float64 `json:"reducedPrice"`
	SalePercentage float64
}

func main() {
	items := importData()
	fmt.Println(items)
	
	var strArray=[]string{"one", "two", "three", "four", "five"}
	for _	,value :=range(strArray){
		target:="five"
		fmt.Print(value)
		i, found := sort.Find(len(value), func(i int) int {
			return strings.Compare(value, target)
		})
				if found {
			fmt.Printf("found %s at entry %d\n", target, i)
		} else {
			fmt.Printf("%s not found, would insert at %d", target, i)
		}
		fmt.Println("")
	} 
	
}

// importData reads the raffle entries from file and 
// creates the entries slice.
func importData() []SaleItem {
	file, err := os.ReadFile(path)
	if err != nil {
		log.Fatal(err)
	}

	var data []SaleItem
	err = json.Unmarshal(file, &data)
	if err != nil {
		log.Fatal(err)
	}

	return data
}