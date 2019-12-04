package main

import (
	"fmt"
	"math"
)

func main() {

	circle := Circle{5.4}
	fmt.Println(getArea(circle))

	rectangle := Rectangle{10, 20}

	fmt.Println(getArea(rectangle))
}

type Shape interface {
	area() float64
}

type Circle struct {
	radious float64
}

type Rectangle struct {
	height, width float64
}

func (circle Circle) area() float64 {
	return 3.1416 * math.Sqrt(circle.radious)
}

func (rectangle Rectangle) area() float64 {
	return rectangle.height * rectangle.width
}

func getArea(shape Shape) float64 {
	return shape.area()
}
