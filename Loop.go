package main

import "fmt"

func main() {
	for i := 0; i <= 1000; i++ {
		if i == 850 {
			break
		}
		if i == 848 {
			continue
		}
		fmt.Println(i)
	}

}
