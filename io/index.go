package myio

import (
	"fmt"
)

func greet(){
	var name string
	fmt.Println("what is your name")
	fmt.Scanln(&name)
	fmt.Println("hello", name)
}
