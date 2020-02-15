package builder_pattern

import (
	"fmt"
	"strconv"
)

type Car interface {
	Drive() string
	Stop() string
}


type car struct {
	topSpeed int
	color    string
	carType string
}

func (c *car) Drive() string {
	return "Driving a "+c.carType+" at speed: " + strconv.Itoa(c.topSpeed)
}

func (c *car) Stop() string {
	return "Stopping a car: " + c.String()
}

func (c *car) String() string {
	return fmt.Sprintf("[%d, %d, %d]", c.carType, c.color, c.topSpeed)
}
