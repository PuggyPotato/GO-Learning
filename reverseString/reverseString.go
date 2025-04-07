package main

import "fmt"

func main(){
	var input string;
	fmt.Print("Enter A Word:")
	fmt.Scanln(&input)
	var result string= reverseString(input)
	fmt.Print("The Result Of The Reversed String:",result)
}

func reverseString(s string) string{
	runes := []rune(s)
	for i,j := 0,len(runes)-1; i < j;i,j=i+1,j-1{
		runes[i],runes[j] = runes[j],runes[i]
	}
	return string(runes)
}