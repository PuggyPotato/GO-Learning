package main

import (
	"fmt"
)

func main(){
	var itemName [] string
	var itemPrice [] float64

	for {
		var a float64
		var b string
		var proceed string

		fmt.Print("Continue?(Y/N):")
		fmt.Scanln(&proceed)

		if(proceed == "N"){
			break;
		}

		fmt.Print("Enter The Item Name:")
		fmt.Scanln(&b)
		fmt.Print("Enter The Item Price:")
		fmt.Scanln(&a)
		itemName = append(itemName, b)
		itemPrice = append(itemPrice, a)

	}
	for i := 0;i < len(itemPrice);i++{
		fmt.Println(itemName[i],":",itemPrice[i],"$")
	}

	
}