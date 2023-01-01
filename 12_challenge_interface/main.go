package main

import (
	"fmt"
	"tempService/printer"
	"tempService/models"
)


func main() {

	fmt.Println(" Welcome to Temperature Service ")
	p:=printer.New()
	defer p.CleanUp()
	p.CityHeader()

	lon:=models.NewCity("London",7.5,false,true)
	ams:=models.NewCity("Amsterdam",11,true,false)
	nyc:=models.NewCity("New York",-3,true,false)

	p.CityDetails(lon)
	p.CityDetails(ams)
	p.CityDetails(nyc)
}