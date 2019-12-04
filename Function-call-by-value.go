package main

import "fmt"

func main() {
	var a, b uint8
	a = 10
	b = 12

	x, y := swap(a, b)

	fmt.Println("a: ", x, " ,b: ", y)
}

func swap(a, b uint8) (uint8, uint8) {
	return b, a
}
