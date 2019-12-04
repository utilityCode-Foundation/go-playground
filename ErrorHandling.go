package main

import (
	"errors"
	"fmt"
)

func main() {
	height := 5
	width := 5
	res, err := getAreaOfARectangle(height, width)

	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println(res)
	}

}
func getAreaOfARectangle(height int, width int) (int, error) {
	if height <= 0 || width <= 0 {
		return 0, errors.New("Height or width is less than 1")
	}
	return (height * width), nil
}
