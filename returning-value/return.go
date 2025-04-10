package main

import (
	"fmt"
)

func oddSum(num1 int64) bool{
	if(num1 % 2 ==0){
		return true
	}else{
		return false
	}
}

func main(){
	var number int64
	var isEven bool
	fmt.Print("Enter A Number:")
	fmt.Scan(&number)
	isEven = oddSum(number)
	if(isEven == true){
		fmt.Println(number,"Is Even")
	}else{
		fmt.Println(number,"Is Odd")
	}
}