package main

import (
	"fmt"
)

func main() {
	sum := add(10, 20)
	fmt.Println(sum)

}

func add(a, b int) int {
	return (a + b)
}
