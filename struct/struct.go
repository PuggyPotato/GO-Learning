package main



import(
	"fmt"
)

type Student struct{
	name string
	marks int64
}

func main(){
	var students []Student
	var name string
	var marks int64
	var searchedMark int64
	var flag = false

	for i:=0;i<5;i++{
		fmt.Print("Enter Student Name:")
		fmt.Scan(&name)
		fmt.Print("Enter Student Marks:")
		fmt.Scan(&marks)
		students = append(students, Student{name:name,marks:marks})
	}
	fmt.Printf("Name \t Marks")
	for i:=0;i<5;i++{
		fmt.Printf("\n %v \t %v",students[i].name,students[i].marks)
	}

	fmt.Print("\nEnter The Mark You Want To Search:")
	fmt.Scan(&searchedMark)

	for i:=0;i<5;i++{
		if(searchedMark == students[i].marks){
			fmt.Printf("Found Matching At Index[%v] \n",i)
			flag = true
		}
	}
	if(flag == false){
		fmt.Print("No Match Found")
	}
}