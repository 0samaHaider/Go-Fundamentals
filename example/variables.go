package main

import (
	"fmt"
	"os"
	"reflect"
)

var (
	name, phone, city string
	age, dob          int
)

func main() {
	fmt.Println("The value of name variable for now is ", name)
	name := "Osama Haider"
	fmt.Println("The value of name variable for now is ", name)
	fmt.Println("The type of variable name is ", reflect.TypeOf(name))
	//get environment variables 
	currentUserName := os.Getenv("USERNAME")
	fmt.Println( currentUserName)

}
