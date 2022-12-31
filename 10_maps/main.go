package main

import "fmt"

func main() {
mapobj:=make(map[string]int)
mapobj["wa"]=1
mapobj["ca"]=2
mapobj["ny"]=3
fmt.Println(mapobj)

for k,v := range mapobj{
	fmt.Print(k,v)
}


}