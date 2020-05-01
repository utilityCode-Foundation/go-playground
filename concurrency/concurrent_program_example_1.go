package main

import (
	"fmt"
	"strconv"
	"time"
)

func main() {

	go printStr("zero")
	printStr("msi")

}

func printStr(str string) {
	for i := 0; i < 50; i++ {
		fmt.Println(strconv.Itoa(i) + " " + str)
		time.Sleep(time.Second * 1)
	}

}
