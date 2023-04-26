package main

import (
	"fmt"
	 
)

func main() {
	input := []int{5, 3, 9, 1, 6, 7, 2, 3}
	fmt.Print(len(input))

	for   outerIndex :=0; outerIndex<len(input); outerIndex++ {
	for   index :=0; index<len(input)-1; index++ {
		 
		if input[index] > input[index+1]{
			temp:=input[index+1]
			input[index+1]=input[index]
			input[index]=temp
			
		 
	}
}
	}
	fmt.Println()

	input = []int{5, 3, 9, 1, 6, 7, 2, 3}
	fmt.Print(len(input))

	middle:=len(input)/2
	
	for   outerIndex :=0; outerIndex<len(input); outerIndex++ {
	for   index :=0; index<len(input)-1; index++ {
		 
		if input[index] > input[index+1]{
			temp:=input[index+1]
			input[index+1]=input[index]
			input[index]=temp
			
		 
	}
}
	
	fmt.Print(input)
}
}