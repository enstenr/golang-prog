package main

import (
	"bufio"
	"fmt"
	"os"
 
)
func main()  {
	
	reader:= bufio.NewReader(os.Stdin)
	fmt.Print("Enter Text: ")
	input,err := reader.ReadString('s')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	fmt.Println(input)
  

	_,err = reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	 

	fmt.Print("Enter second Text: ")
	input,err = reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	fmt.Println(input)

 
}