package main

import "fmt"

func stack () {
	stack := []int{1,2,3}
	stack = append(stack, 4)	// push
	stack = stack[1:]	// pop
	fmt.Println("stack = ", stack)
}