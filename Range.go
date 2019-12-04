package main

import "fmt"

func main() {

	nums := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}

	for i := range nums {
		fmt.Print(nums[i], " ")
	}
	fmt.Println()

	country := map[string]string{"BD": "Bangladesh", "USA": "America", "IND": "India"}

	for code := range country {
		fmt.Print(code, ":", country[code], " || ")
	}
	fmt.Println()

	for code, name := range country {
		fmt.Print(code, ":", name, " || ")
	}
}
