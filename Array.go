package main

import "fmt"

func main() {
	// Approach 1: inserting with index number
	// Only applicable for fixed sized array
	var fixedNums [10]int
	for i := 0; i < 10; i++ {
		fixedNums[i] = i + 1
	}

	// For fixed sized array,
	// we must pass like below
	fmt.Println("============= Fixed size array =============")
	PrintValue(fixedNums[:])

	// Approach 2: With append(...)
	var nums []int
	for i := 0; i < 10; i++ {
		nums = append(nums, i+1)
	}

	// For dynamic sized array,
	// we can also pass the array like below
	fmt.Println("============= Dynamic size array =============")
	PrintValue(nums)

}
func PrintValue(arr []int) {
	for i := 0; i < len(arr); i++ {
		fmt.Println(arr[i], " ")
	}
}
