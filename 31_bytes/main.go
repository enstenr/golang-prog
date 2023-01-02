package main

import (
	"fmt"
	"unicode/utf8"
)

func main() {

	bs := []byte{226 ,151, 186} 
	fmt.Printf("%s", bs) 
	bs = []byte("◺")
	fmt.Println(bs) // Output: [226 151 186]
	bs=[]byte("汉")
	fmt.Println(bs)
	bs = []byte{226 ,151 ,186} 
	fmt.Printf("%s", bs) 
	s:=string(bs)
	fmt.Println(utf8.RuneCountInString(s))
	fmt.Println()
	bs = []byte("汉") 
	s = string(bs)
	fmt.Println(bs)
	fmt.Println(s)
	fmt.Println(utf8.RuneCountInString(s)) // Output: 1
	fmt.Println()

	name:="Rajesh"
	for _,b:=range name{
		fmt.Println(b)

	}
	ba:=[]byte{name[2],name[5]}
	fmt.Print(string(ba))

	fmt.Print(name[2])
}