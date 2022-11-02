package main

import(
	"github.com/enstenr/common/connection"
)
func main() {
	
	connection.GetArangoDBConnection("dev")
	
}