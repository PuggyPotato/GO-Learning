package main


import (
	"time"
	"fmt"
)

func main(){
	var userTimer float64


	fmt.Print("How long of a timer do you want(second):")
	fmt.Scan(&userTimer)

	timer := time.NewTimer(time.Duration(userTimer) * time.Second)
	fmt.Println("Timer Started!")
	<- timer.C
	fmt.Println("Timer Ended!")
}