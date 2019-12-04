package main

import "fmt"

func main() {
	country := make(map[string]string)
	country["BD"] = "Bangladesh"
	country["USA"] = "America"
	country["IND"] = "India"

	fmt.Println(country["BD"])
	fmt.Println(country)

	delete(country, "IND")
	fmt.Println(country)
}
