package main

import "fmt"
/**
********************************
go mod init <module_name>
go mod tidy
go run . 
**/
func main() {
fmt.Println("Welcome to golang")
result:=addNumbers(1,2)
fmt.Print(result)
}