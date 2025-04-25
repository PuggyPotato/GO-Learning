package main

import (
	"fmt"
	"bufio"
	"os"
	//"strings"
)

func main(){
	var petNames []string
	var currentPetName string
	var userChoice int64 = 0
	scanner := bufio.NewScanner(os.Stdin)
	fmt.Println("Welcome To Whatever Pet Thinggy")
	fmt.Println("Choose What You Want To Do! \n1. Add Pet \n2. List Pet \n3. Exit")
	fmt.Print("Enter Choice:")
	fmt.Scan(&userChoice)
	for(userChoice != 3){
		if userChoice == 1{
			fmt.Print("Enter Pet Name:")
			scanner.Scan()
			scanner.Scan()
			currentPetName = scanner.Text()
			//currentPetName = strings.TrimSpace(currentPetName)
			petNames = append(petNames, currentPetName)
		}
		if userChoice == 2 {
			if len(petNames) == 0{
				fmt.Println("No Pet Name Now.")
			}
			for i:=0;i< len(petNames);i++{
				fmt.Println(i+1,".",petNames[i])
			}
		}
		if userChoice != 1 && userChoice != 2 && userChoice != 3{
			fmt.Println("Enter A Valid Choice:")
		}

		fmt.Print("Enter Choice:")
		fmt.Scan(&userChoice)

	}
	fmt.Print("GoodBye!")

}