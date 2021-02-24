package main

import "fmt"

func main(){
	fmt.Println("hello")
	// exec just before the main fn returns => close pg conn or something
	defer fmt.Println("defered worlds")	// printed at the last
	fmt.Println("worlds")

	// all => LIFO order => 3, 2, 1
	defer fmt.Println("1")
	defer fmt.Println("2")
	defer fmt.Println("3")

	start := true
	defer fmt.Println("start = ", start)	// true => takes val defined at the time defer is called => not at the time deferred fn is exec
	start = false
}
