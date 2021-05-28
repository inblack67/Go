package main

import "fmt"

func main(){
	// declaring size at first is imp
	var myArr [2] string

	myArr[0] = "apple"
	myArr[1] = "grapes"

	fmt.Println(myArr)

	arr2 := [2]int{1,2}
	fmt.Println(len(arr2))

	fmt.Println(arr2[0:2])	// range => o/p = 1,2
}
