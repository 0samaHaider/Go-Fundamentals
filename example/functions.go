package main

import(
	"fmt"
)
func testingFunctionString(a string)string  {
	fmt.Println("Hello from testingFunction ",a)
return a 
}
func testingFunctionAdd(x int , y int )int {
	c:=x+y
	fmt.Println("Addition of two numbers are ",c)
	return c 
}


func main(){
	fmt.Println("Hello from main function ")
	testingFunctionString("xyz")
testingFunctionAdd(2,3)
}
