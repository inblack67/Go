package main

import "fmt"

func main(){
	// string
	// bool
	// int
	// int  int8  int16  int32  int64
	// uint uint8 uint16 uint32 uint64 uintptr
	// byte - alias for uint8
	// rune - alias for int32
	// float32 float64
	// complex64 complex128

	var name string = "some name"
	// var name = "some name"
	fmt.Println(name)
	
	var age = 20
	// var age int = 20;
	fmt.Println(age)
	fmt.Printf("%T\n", name) // string
	fmt.Printf("%T\n", age) // int
	
	var age2 int32 = 0
	fmt.Printf("%T\n", age2) // int32

	var valid = true
	// var valid bool = true
	valid = false
	fmt.Println(valid)

	const isInvalid = true
	// isInvalid = false // error
	fmt.Println("isInvalid =", isInvalid)

	// Shorthand var => only works inside a function
	test := 1
	test = 2
	fmt.Println(test)

	size := 1.3
	fmt.Println(size)
	fmt.Printf("%T\n", size)	// float64
	
	var size2 float32 = 0.1
	fmt.Println(size2)
	fmt.Printf("%T\n", size2)	// float32

	// myname := "test"
	// email := "test@test.com"
	// password := "none"

	// shorthand
	myname, email, password := "test", "test@gmail.com", "none"

	fmt.Println(myname, email, password)
}
