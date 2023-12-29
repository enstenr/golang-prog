package main

import "fmt"
func main() {
name := "Rajesh"
fmt.Println(name)

pointer:=&name
fmt.Println(pointer)
*pointer="Raj1"
fmt.Println(pointer)
fmt.Println(*pointer)
}