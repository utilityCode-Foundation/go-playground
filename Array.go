package main

import "fmt"

func main() {
	var nums [10]int

	for i := 0; i < 10; i++ {
		nums[i] = i + 1
	}
	PrintValue(nums[:])

}
func PrintValue(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i], " ")
	}
}
