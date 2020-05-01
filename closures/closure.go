package main

import (
	"fmt"
	"strconv"
)

func main(){

	clouser1:=closure_example()

	fmt.Println(clouser1())
	fmt.Println(clouser1())

	// function as param
	fmt.Println(myFunc(testFunc,123))

}


func closure_example() func() int{
	var i=0
	return func() int{
		i++;
		return i
	}
}

// function as param
type convert func(int) string

func testFunc(val int) string{
	return "The value is:"+ strconv.Itoa(val)
}

func myFunc(fn convert, val int) string{
	return fn(val)
}

func temp2() {

}