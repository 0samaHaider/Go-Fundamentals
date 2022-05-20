package main 

import(
	"fmt"
)
func main(){
fmt.Println("Hello from switch ")

switch "OSAMA HAIDER"{

case "OSAMA":
	fmt.Println("CASE 1")

case "ALI":
	fmt.Println("CASE 2")

case"OSAMA HAIDER":
	fmt.Println("CASE 3 ")

default:
	fmt.Println("Nothing match ")


}
}