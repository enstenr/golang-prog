package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"os"
)

type Output struct {
    Name   string `json:"name"`
    Type   string `json:"type"`
    Age    int    `json:"Age"`
    Tree Tree `json:"Tree"`
}

// Social struct which contains a
// list of links
type Tree struct {
    Name string `json:"Name"`
}

func main() {
	jsonFile, err := os.Open("duplicate_sku.json")
	if err != nil {
		fmt.Println(err)
	}
	fmt.Println("Successfully Opened users.json")
	// defer the closing of our jsonFile so that we can parse it later on
	defer jsonFile.Close()

	if err != nil {
		// handle error
	}
	byteValue, _ := ioutil.ReadAll(jsonFile)
	var users Output

	// we unmarshal our byteArray which contains our
	// jsonFile's content into 'users' which we defined above
	json.Unmarshal(byteValue, &users)
	fmt.Print(users)

 
}
