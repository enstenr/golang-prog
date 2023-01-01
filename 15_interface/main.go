package main

import "fmt"

type student struct {
	name       string
	rollNumber string
}
type IStudent interface {
	wish()
}

func NewStudent(name string) IStudent {
	return &student{
		name: name,
	}
}
func (s student) wish() {
	fmt.Printf("good morning %v",s.name)
}
func (s student) GetName() string {
	return s.name
}
func main() {
std:=NewStudent("S Rajesh")
std.wish()
 
var stud student
stud.name="S Rajesh"
stud.rollNumber="23357143"
stud.wish()
}
