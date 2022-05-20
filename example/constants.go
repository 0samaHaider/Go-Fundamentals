package main

import(
	"fmt"
)
func main(){
	fmt.Println("Hello from constants file")
	//we can not update the value of 23
	const age=23
	fmt.Println("My age is ",age)

}