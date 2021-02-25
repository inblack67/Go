package main

import "fmt"

// Hello ...
type Hello struct {
	reply string
}

func main () {
	a := 12
	b := a	// copy
	c := &a
	*c = 69
	fmt.Println(a, b, *c)	// 69, 12, 69


	// go does not allow pointer arithmetic
	arr := [3]int{1,2,3}
	p1 := &arr[0]
	p2 := &arr[1]
	p3 := &arr[2]
	fmt.Printf("p1 = %p, p2 = %p, p3 = %p\n", p1, p2, p3)
	// sub := p3 - p1	// not allowed

	// default val for pointers => nil

	greet := new(Hello)
	// (*greet).reply = "worlds"	// original way
	greet.reply = "worlds"	// syntactic sugar
	fmt.Println("greet.reply = ", greet.reply)	// worlds

	// arrays are value types => maps and slices are not.
}