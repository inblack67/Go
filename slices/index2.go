package main

import "fmt"

func index2(){
	arr := []int{1,2,3}

	arr = append(arr, 4)	// push

	arr = arr[1:]	// shift => pop from the first
	arr = arr[: len(arr) - 1]	// pop => from last

	fmt.Println("arr = ", arr)

	arr2 := []int{1,2,3,4,5,6}
	arr2 = append(arr2[:2] /* exclusive */, arr2[3:]...)	// removing 3 (at 2nd index)
	fmt.Println("arr2 = ", arr2)
}