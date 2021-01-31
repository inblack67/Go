package main

import "fmt"

func main(){

	// dynamic size
	arr2 := []int{1,2}
	fmt.Println(arr2)
	fmt.Println(len(arr2))

	fmt.Println(arr2[0:2])	// range => o/p = 1,2
}
