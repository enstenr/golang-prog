package main

import "fmt"
func main() {
name := "Rajesh"
fmt.Print(name)

pointer:=&name
fmt.Print(pointer)
*pointer="Raj1"
fmt.Print(pointer)
fmt.Print(*pointer)
}