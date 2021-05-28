package main

import "fmt"

func main(){
	ids := []int{1,2,3}

	// When ranging over a slice, two values are returned for each iteration. The first is the index, and the second is a copy of the element at that index.
	for i, id := range ids{
		fmt.Println(i, id)
	}

	for _, id := range ids{
		fmt.Println(id)
	}

	sum := 0
	for _, id := range ids{
		sum += id
	}
	fmt.Println("sum = ",sum)

	// range with map
	mymap := map[string]int{"a": 1, "b": 2, "c": 3}
	for k, v := range mymap{
		fmt.Println("key, value", k, v)
		fmt.Printf("key = %s and value = %d\n", k, v)
	}
}
