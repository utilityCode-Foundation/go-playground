package main

import (
	"fmt"
)

func main() {
	sum := getSum(10, 20)
	fmt.Println(sum)

}

func getSum(a, b int) int {
	return (a + b)
}
