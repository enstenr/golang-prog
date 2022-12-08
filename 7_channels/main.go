package main

import (
	"fmt"
	"time"
)
/**
	The send / receive will wait for the message to be received
	Send a value into a channel using the channel <- syntax.
**/
func main() {

    messages := make(chan string)

    go func() { 
		time.Sleep(5*time.Second)
		messages <- "ping" }()

    msg := <-messages
    fmt.Println(msg)
}