package main

import (
	"fmt"
	bp "go-playground/design-pattern/creational/builder-pattern"
)

	func main() {
		builder := bp.New()
		car := builder.TopSpeed(50).Paint(bp.BLUE).Type(bp.BUS).Build()
		fmt.Println(car.Drive())
		fmt.Println(car.Stop())
	}


