package main

import (
    "fmt"
    "time"
)

func worker(done chan bool) {
    fmt.Print("working...")
    time.Sleep(time.Second)
   

    done <- true
    done<-false
    done <- true
    done<-false
    done <- true
    done<-false
    fmt.Println("done")
}

func main(){

    done := make(chan bool)
    go worker(done)
    fmt.Print(<-done)
    fmt.Print(<-done)
     
}