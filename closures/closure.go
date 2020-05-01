package main

import "fmt"

func main(){

	clouser1:=closure_example()

	fmt.Println(clouser1())
	fmt.Println(clouser1())

}

func closure_example() func() int{
	var i=0
	return func() int{
		i++;
		return i
	}
}
