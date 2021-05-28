package main

import (
	"fmt"
	"log"
)

func panicker () {
	fmt.Println("about to panic")
	defer func () {
		if err := recover(); err != nil {
			log.Println("recovery err = ", err)
		}
	}()
	panic("packied!!")
}

func main(){
	// panic => exceptions
	// fmt.Println("start")
	// defer fmt.Println("end")	// panic happens after the deferred statements are exec 
	// panic("panicking")

	fmt.Println("start")
	panicker()
	fmt.Println("end")
}
