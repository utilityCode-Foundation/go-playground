//package clousres
package main


import "fmt"

//func main(){
//  print_from_anynomous_function("anonymous function")
//}

func print_from_anynomous_function(something string){
 // this is anonymous function
	func (str string){
		fmt.Println("Hello from",str)
	}(something)
}