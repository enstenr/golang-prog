package main

import "fmt"

func main() {
	//initialize a slice of strings
var colors=[]string{"Red", "Green", "Yellow", "Blue", "Magenta", "Pink"}
colors=append(colors, "Orange")

colors=append(colors[0:len(colors)])
fmt.Println(colors)
}