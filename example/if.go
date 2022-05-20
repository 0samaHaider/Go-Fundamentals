package main

import(
	"fmt"
)

func main (){
	fmt.Println("Hello from if.go")
	a:=10
	b:=10
	if a>b {
	fmt.Println(a ," is greater than ", b )
} else if  a<b {
	fmt.Println(a ," is less than ", b )
	} else {
		fmt.Println(a,"  is equal to the ", b)
	}

}