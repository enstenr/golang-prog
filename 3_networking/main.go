package main

import (
	"fmt"
	"github.com/showwin/speedtest-go/speedtest"
	 
)

func main(){
	user, _ := speedtest.FetchUserInfo()
	user.SetLocationByCity("Puducherry")
	user.SetLocation("Puducherry", 11.953,79.809)
	fmt.Println("Printing User")
	fmt.Println(user)
	// Get a list of servers near a specified location
	// user.SetLocationByCity("Tokyo")
	// user.SetLocation("Osaka", 34.6952, 135.5006)

	serverList, _ := speedtest.FetchServers(user)
	fmt.Println(serverList)
	targets, _ := serverList.FindServer([]int{})

	for _, s := range targets {
		s.PingTest()
		s.DownloadTest(true)
		s.UploadTest(true)
		fmt.Println(s)
		fmt.Printf("Latency: %s, Download: %f, Upload: %f\n", s.Latency, s.DLSpeed, s.ULSpeed)
	}
}