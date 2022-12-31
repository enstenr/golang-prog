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
	input,_ := reader.ReadString('\n')
	fmt.Print(input)

	numinput,_ := reader.ReadString('\n')
	aFloat,err:=strconv.ParseFloat(strings.TrimSpace(numinput),60)
	if( err != nil ){
		fmt.Print(err)
	}else{
	fmt.Print(aFloat)
}
}