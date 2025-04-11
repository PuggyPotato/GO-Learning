package main

import (
	"fmt"
	"math/rand/v2"
)

func main(){
	var randomNum = rand.IntN(100) + 1

	var userGuess int
	var count = 0
	var lowerBound int = 1
	var upperBound int = 100


	fmt.Print("Guess A Number Between 1-100:")
	for{
		fmt.Scan(&userGuess)
		count++
		if(userGuess == randomNum){
			fmt.Println("Congrats! You Guessed It!")
			break

		}else if userGuess > randomNum{
			upperBound = userGuess
			fmt.Println("The Number Is Too High! Guess Lower!")
		}else{
			lowerBound = userGuess
			fmt.Println("The Number Is Too Low!Guess Higher!")
		}
		fmt.Printf("Guess A Number Between %v-%v:",lowerBound,upperBound)
	}
	fmt.Print("You Guessed The Number In ",count," Attempt")

}