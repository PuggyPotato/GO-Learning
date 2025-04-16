package main

import (
	"fmt"
)

type Student struct{
	Name string
	ID int64
	Grades []float64

}

func main(){	
	var student Student

	student.Name = "Potato"
	student.ID = 123456
	student.Grades = append(student.Grades, 1,2,3,4,5,6)
	var average = student.Average()


	fmt.Println("Student Name:",student.Name)
	fmt.Println("Student ID:",student.ID)
	fmt.Println("Student Grades:",student.Grades)
	fmt.Println("Student Grades Average:",average)
}

func(s Student) Average() float64{
	var total float64
	var average float64
	for _,grade := range s.Grades{
		total += grade
	}

	average = total / float64(len(s.Grades))
	return average
}