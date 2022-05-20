package main

import(
	"fmt"
)

func main(){
	type TestingStruct struct {
			Name string
	Age int 
	City string
	}

	valuesInstruct:= TestingStruct{
		Name:"Osama Haider",
		Age:23,
		City:"Islamabad",
	}
fmt.Println("Name = ",valuesInstruct.Name)
fmt.Println(valuesInstruct)
}