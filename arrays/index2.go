package main

import "fmt"

func index2(){

	// arrays => size has to be known at the compile time

	grades := [3]int{1, 2 ,3}		// array of 3 int 	
	dynamicGrades := [...]int{1,2,3,4,5}	// dynamic size
	fmt.Println(grades)
	fmt.Println(dynamicGrades)
	var students [3]string
	students[0] = "A"
	students[1] = "B"
	students[2] = "C"
	fmt.Println(students)
	gradesLength := len(grades)
	studentsLength := len(students)
	fmt.Println(gradesLength)
	fmt.Println(studentsLength)
	
	var matrix [3][3]int
	matrix[0] = [3]int{1,2,3}
	matrix[1] = [3]int{1,2,3}
	matrix[2] = [3]int{1,2,3}
	fmt.Println("Matrix = ", matrix)

	a := [...]int{1,2,3}
	b := a	// not a pointer but a copy from the original
	b[1] = 67
	fmt.Println("original = ", a)	// [1,2,3]
	fmt.Println("copy = ", b)	// [1,67,3]

	pointerToArray := &a
	pointerToArray[0] = 69
	fmt.Println("Array a = ", a)
	fmt.Println("Pointer to Array a = ", pointerToArray)
}