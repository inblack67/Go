package main

import "fmt"

func main(){
	a := 5
	b := &a

	fmt.Println(a,b)
	fmt.Printf("%T\n", a)
	fmt.Printf("%T\n", b) // *int
	fmt.Printf("%d\n", *b) // 5
	fmt.Printf("%d\n", *&a) // 5

	*b = 20
	fmt.Println(a)	// 20
}
