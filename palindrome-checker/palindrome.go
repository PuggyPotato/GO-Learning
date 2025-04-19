package main

import(
	"fmt"
	"strings"
	"os"
	"bufio"
)

func main(){
	var userInput string

	reader := bufio.NewReader(os.Stdin)

	fmt.Print("Enter A Word Or A Line:")
	userInput,_ = reader.ReadString('\n')
	userInput = strings.TrimSpace(userInput)

	chars := []rune(userInput)
	var reversedString  = []rune(userInput)

	j :=0

	for i:=len(chars)-1;i >=0;i--{
			reversedString[j] = chars[i]
			j++
		}
 

	if string(reversedString) == string(chars){
		fmt.Println("They Are Palindrome")
	}else{
		fmt.Println("Not Palindrome")
	}


}