package main

import (
	"fmt"
	"bufio"
	"os"
)

func printName(name string){
	fmt.Printf("Hello %v",name)
}

func main(){
	scanner := bufio.NewScanner(os.Stdin)
	var name string
	fmt.Print("Enter Your Name:")
	scanner.Scan()
	name = scanner.Text()
	printName(name)

}