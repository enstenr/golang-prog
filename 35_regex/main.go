package main

import (
	"fmt"
	"regexp"
)

func main() {
	matched, err := regexp.Match(`foo.*`,[]byte(`seafood`))
	if err!= nil {
        panic(err)
    }
	if matched {
        fmt.Println("seafood")
    }


}