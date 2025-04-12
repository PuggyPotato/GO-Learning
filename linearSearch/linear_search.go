package main

import (
	"fmt"
)

func main(){
	var studentName [5]string
	var marks [5]int 
	var searchMark int
	var found = false

	for i:=0;i<5;i++{
		fmt.Print("Enter Student Name:")
		fmt.Scanln(&studentName[i])
		fmt.Print("Enter Student Mark:")
		fmt.Scanln(&marks[i])
	}

	for i:=0;i<5;i++{
		fmt.Printf("\n%v",studentName[i])
		fmt.Printf("\t %v",marks[i])
	}
	fmt.Print("\nEnter Your Desired Mark:")
	fmt.Scanln(&searchMark)

	for i:=0;i<5;i++{
		if(searchMark == marks[i]){
			fmt.Printf("\nFound at index[%v]",i)
			found = true
		}
	}
	if(found == false){
		fmt.Printf("Marks Not Found")
	}
}