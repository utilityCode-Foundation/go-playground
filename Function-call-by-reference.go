package main

import "fmt"

func main() {

	var a, b uint8
	a = 10
	b = 12

	swap(&a, &b)

	fmt.Println("a: ", a, " ,b: ", b)
}

func swap(a, b *uint8) {
	x := *a
	*a = *b
	*b = x
}
