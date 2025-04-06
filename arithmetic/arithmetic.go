package main

import "fmt"


func main(){
	var num1 float64
	var num2 float64
	var symbol string
	var answer float64
	fmt.Print("Enter Number 1:")
	fmt.Scan(&num1)
	fmt.Print("Enter Number 2:")
	fmt.Scan(&num2)
	fmt.Print("Enter The Operation You Want To Do(+,-,*,/):")
	fmt.Scan(&symbol)
	if(symbol == "+"){
		answer = num1 + num2
	}else if(symbol == "-"){
		answer = num1 - num2
	}else if(symbol == "*"){
		answer = num1 * num2
	}else if(symbol == "/"){
		answer = num1 / num2
	}else{
		fmt.Println("Error")
	}

	fmt.Println("The Answer Is ", answer)

}