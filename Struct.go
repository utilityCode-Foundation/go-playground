package main

import "fmt"

func main() {

	var student1 student
	student1.id = 1
	student1.name = "Shahidul"
	student1.address = "Kalikair,Gazipur"

	PrintStudentInfo(student1)
	Modify(&student1)
	PrintStudentInfo(student1)
}

type student struct {
	id      uint8
	name    string
	address string
}

func PrintStudentInfo(stdn student) {
	fmt.Println(stdn)
}

func Modify(stdn *student) {
	stdn.name = "zeromsi"
}
