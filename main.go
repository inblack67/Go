package main

import "fmt"

func main (){
	var name string
	fmt.Println("what is your name")
	fmt.Scanln(&name)
	fmt.Println("hello", name)

}
