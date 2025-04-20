package main


import(
	"fmt"
)

func main(){
	var integers = []int64{1,2,3,4}
	var reversed []int64

	for j:=len(integers)-1;j >=0;j--{
		reversed = append(reversed,integers[j])
	}
	fmt.Println(reversed)
}