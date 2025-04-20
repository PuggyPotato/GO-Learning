	package main

	import (
		"fmt"
		"time"
	)

	func main(){

		for{
			fmt.Println("Timer Started!")
			timer :=time.NewTimer(5* time.Second)
			<- timer.C
			fmt.Println("5 Second Has Passed!")
			printSomething()

		}

	}

	func printSomething(){
		fmt.Println("Print This!")
	}