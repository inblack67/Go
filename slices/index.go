package main

import "fmt"

func main(){

	// dynamic size
	arr2 := []int{1,2}
	fmt.Println(arr2)
	fmt.Println(len(arr2))

	fmt.Println(arr2[0:2])	// range => o/p = 1,2

	// slice => array with a dynamic size
	slice := []int{1,2,3}
	fmt.Println("slice = ", slice)
	fmt.Println("slice length = ", len(slice))

	// The cap built-in function returns the capacity of v, according to its type:
	// Array: the number of elements in v (same as len(v)).
	// Pointer to array: the number of elements in *v (same as len(v)).
	// Slice: the maximum length the slice can reach when resliced;
	// if v is nil, cap(v) is zero.
	// Channel: the channel buffer capacity, in units of elements;
	// if v is nil, cap(v) is zero.
	fmt.Println("capacity length = ", cap(slice))

	// slices are reference types
	a := []int{1,2,3}
	b := a
	b[0] = 0
	fmt.Println("slice a = ", a)	// 0
	fmt.Println("slice b = ", b)	// 0

	// slicing operations => can also work with arrays
	a1 := []int{1,2,3,4,5,6,7,8,9}
	b1 := a1[:]	// slice of all elements
	c1 := a1[3:]	// slice from the 4th element
	d1 := a1[:6]	// slice of first 6 elements
	e1 := a1[3:6]	// slice of 4th, 5th and 6th elements

	a1[3] = 69	// reflect in all

	// 3:6 => first is inclusive and second is exclusive

	fmt.Println("slice of all elements = ", b1)
	fmt.Println("slice from the 4th element = ", c1)
	fmt.Println("slice of first 6 elements = ", d1)
	fmt.Println("slice of 4th, 5th and 6th elements = ", e1)

	// make
	bakedSlice := make([]int, 3, 5)	// cap of 100, size of 3
	// bakedSlice[2] = 11	// under length - 1 => works
	bakedSlice = append(bakedSlice, 1)	// costly => recreation of array to a bigger size => copy so go with make 
	bakedSlice = append(bakedSlice, 2,3,4,5)	// appending multiple elements
	// when adding more element go creates a new array of double the size => cache logic?
	fmt.Println("baked array = ", bakedSlice)	// [0,0,0]
	fmt.Println("baked slice length  = ", len(bakedSlice))
	fmt.Println("baked slice capacity  = ", cap(bakedSlice))

	someSlice := []int{1,2,3}
	spreadMe := []int{4,5,6}

	someSlice = append(someSlice, spreadMe...)

	fmt.Println("spread the slice into the parent", someSlice)
}
