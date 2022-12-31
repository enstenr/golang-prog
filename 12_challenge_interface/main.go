package main

import (
	"fmt"
	"tempService/printer"
)


func main() {

	fmt.Printf(" Welcome to Temperature Service ")
	p:=printer.New()
	defer p.CleanUp()
	p.CityHeader()
}