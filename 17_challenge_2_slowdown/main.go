package main

import (
	"log"
	"strings"
	"time"
)

const delay = 7 * time.Second
func main() {
	msg := "Time to learn about Go strings"
	slowdown(msg)
}

func slowdown(msg string) {
	words:=strings.Split(msg, " ")
	for _,w  := range words {
		var pw []string
		for i,c:=range w{
			 
			rb:=strings.Repeat(string(c),i+1)
			pw=append(pw,rb)
		}
	 
		print(strings.Join(pw,""))
	
	}
}

func print(msg string) {
	log.Println(msg)
	time.Sleep(delay)
}