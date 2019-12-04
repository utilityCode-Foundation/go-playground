package main

import "fmt"

func main() {
	fmt.Println(fibonacciNum(5))
	fmt.Println(factorial(5))
}

func fibonacciNum(num int) int {
	if num <= 1 {
		return 1
	}

	return fibonacciNum(num-1) + fibonacciNum(num-2)
}

func factorial(num int) int {
	if num <= 1 {
		return 1
	}
	return num * factorial(num-1)

}
