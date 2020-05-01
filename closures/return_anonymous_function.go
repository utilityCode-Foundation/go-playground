package main

import "fmt"

//func main()  {
//    //returnedFunction:=return_anonymous_fucntion()
//    //returnedFunction("Printing from a function that returns a function")
//	return_anonymous_fucntion()("Printing from a function that returns a function")
//}

func return_anonymous_fucntion() func(somethig string){
	return func(str string){
		fmt.Println(str)
	}
}
