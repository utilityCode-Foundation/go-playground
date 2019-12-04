package main

import "fmt"

func main() {
	x := 10
	if x != 10 {
		fmt.Println("X is not equal to 10")
	} else if x == 10 {
		fmt.Println("X is equal to 10")
	}
}
