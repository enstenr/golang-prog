package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"
)
func main()  {
	
	reader:= bufio.NewReader(os.Stdin)
	fmt.Print("Enter Text: ")
	input,err := reader.ReadString('a')
	if err != nil {
		fmt.Println("Error reading input:", err)
		return
	}
	fmt.Print(input)

	// Discard the remaining newline character
	_, err = reader.ReadString('\n')
	if err != nil {
		fmt.Println("Error discarding newline character:", err)
		return
	}

	numinput,_ := reader.ReadString('\n')
	fmt.Println(numinput)
	aFloat,err:=strconv.ParseFloat(strings.TrimSpace(numinput),64)
	if( err != nil ){
		fmt.Print(err)
	}else{
	fmt.Print(aFloat)
}
}