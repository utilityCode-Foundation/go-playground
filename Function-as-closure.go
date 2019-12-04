package main

import "fmt"

func main() {
	sum := getSum(10, 10)
	fmt.Println(sum())
	fmt.Println(sum())

}

func getSum(a, b int) func() int {
	i := a + b

	return func() int {
		i++
		return i
	}

}
