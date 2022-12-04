package main

import ("fmt"
	"m_diff_package/math_package"

)
/**
********************************
go mod init <module_name>
go mod tidy
go run . 
**/
func main() {
fmt.Println("Welcome to golang")
result:=math_package.AddNumbers(1,2)
fmt.Print(result)
math_package.Log()
}