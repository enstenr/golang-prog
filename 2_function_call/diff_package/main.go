package main

import ("fmt"
	"m_diff_package/math_pkg"

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
fmt.Println(result)
math_package.Log()
}