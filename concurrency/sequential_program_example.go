package main

import (
	"fmt"
	"time"
)

func main() {
	print("zero")
	print("msi")
}

func print(str string) {

	for i := 0; i < 10; i++ {
		fmt.Println(str)
		time.Sleep(time.Microsecond*1000)
	}
}
