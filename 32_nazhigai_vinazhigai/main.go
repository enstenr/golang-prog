package main

import (
	"fmt"
	"math"
	"time"
)

func main(){

	starttime:=time.Date(time.Now().Year(),time.Now().Month(),(time.Now().Day()),int(06),int(30),int(30),int(30),time.Now().Location())
	
	//fetching current time
	 currentTime := time.Now()
	 
	 
	 fmt.Println("Past Time: ", starttime)
	 fmt.Println("Current Time: ", currentTime)
	 //differnce between pastdate and current date
	 diff := currentTime.Sub(starttime)
	 fmt.Printf("time difference is %v or %v in minutes\n", diff, diff.Minutes())
	 fmt.Print(" 5 Nazhigai = 2 Hours")
	 totalMinutes:=diff.Minutes()
	 totalviNazhigai:=math.Round(5*totalMinutes/2*100)/100
	 totalNazhigai:=math.Round((totalviNazhigai/60)*100)/100
 

	  
	 fmt.Println("")
	 fmt.Printf("Total Nazhigai is %v and Total Vinazhigai is  %v \n", totalNazhigai,totalviNazhigai)
}