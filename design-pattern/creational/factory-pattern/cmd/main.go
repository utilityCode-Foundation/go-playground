package main

import (
	"fmt"
	factory "go-playground/design-pattern/creational/factory-pattern"
)

func main() {
	paymentMethod,err:=  factory.GetPaymentMethod(factory.CASH)
	if err!=nil{
		fmt.Println(err)
	}
	fmt.Println(paymentMethod.Pay(100.00))

}