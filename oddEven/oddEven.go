package main

import "fmt"



func main(){
	var number int

	fmt.Print("Enter A Number:")
	fmt.Scan(&number)
	if number % 2 == 0{
		fmt.Print("Even")
		return
	}
	fmt.Print("Odd")
}