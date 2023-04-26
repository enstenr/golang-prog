package main

import (
	"bufio"
	"fmt"
	"os"
 
	 
)

func main (){
	print("bufio package ")
	rd:=bufio.NewReader(os.Stdin)
	bytes,err:=rd.ReadBytes('\n')
	if err != nil{
		fmt.Print(err.Error())
	}
	 
	fmt.Print(bytes)
 
	for i := 1; i <= 1000; i++ {
	//fmt.Println(i,string(i),byte(i))
	fmt.Printf("%08b\n", i)
	}

}