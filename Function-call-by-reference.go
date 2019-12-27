package main

import "fmt"

func main() {

	var a, b uint8
	a = 10
	b = 12

	swapValue(&a, &b)

	fmt.Println("a: ", a, " ,b: ", b)
}

func swapValue(a, b *uint8) {
	x := *a
	*a = *b
	*b = x
}
