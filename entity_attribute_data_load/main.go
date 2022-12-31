package main

import (
	// import our encoding/json package
	"github.com/enstenr/common/connection"

	"os"
)

type ValueTypeToUnmarshalTo struct {
	Key     string `json:"_key"`
	From    string `json:"_from"`
	To      string `json:"_to"`
	Created string `json:"created"`
}

func main() {
	env, flag := os.LookupEnv("env")
	if !flag {
		env = "dev"
	}
	connection.ParseJson(env)

}
