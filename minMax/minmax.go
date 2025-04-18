package main

import (
	"fmt"
)


func main(){
	var arrNum []int64
	var currentNum int64

	for{
		fmt.Print("Enter Number:")
		fmt.Scan(&currentNum)
		fmt.Println("To Quit Enter -99")
		if currentNum == -99{
			break;
		}
		arrNum = append(arrNum, currentNum)
	}
	if len(arrNum) !=0{
		var max =arrNum[0]
		var min =arrNum[0]
		for _,num:=range arrNum{
			if num > max{
				max = num
			}
			if num < min{
				min = num
			}
		}
		fmt.Println("The Max Is:",max)
		fmt.Println("The Min Is:",min)
	}
	
}