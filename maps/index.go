package main

import "fmt"


func main(){
	// map => key value pairs => types of key/value has to be consistent
	// order of the map's kvs is not guranteed in the o/p
	chess := map[string]int{
		"a1": 1,
		"b1": 2,
		"c1": 3,
		"c4": 4,
		"c5": 5,
		"c6": 6,
	}
	// maps, slices and funcs not allowed as map keys
	fmt.Println(chess)

	makeChess := make(map[string]int)
	makeChess = map[string]int{
		"a1": 1,
		"b1": 2,
		"c1": 3,
		"c4": 4,
		"c5": 5,
		"c6": 6,
	}
	fmt.Println("a1 = ", makeChess["a1"])
	fmt.Println("a1 = ", makeChess["notpresent"])	// 0

	const key = "a1"

	// interrogate map like this:-
	value, ok := makeChess[key]

	if !ok {
		fmt.Println("value not found for the key ", key)
	} else {
		fmt.Printf("value for the key %s = %d\n", key, value)
	}

	fmt.Println(makeChess)	// new absent

	makeChess["new"] = 7	// adding to map
	fmt.Println(makeChess)

	delete(makeChess, "new") // deleting from the map
	fmt.Println(makeChess)

	mapLength := len(makeChess)
	fmt.Println("map length = ", mapLength)

	copiedMap := makeChess

	delete(copiedMap, "a2")	// will get deleted from the makeChess (parent) map too

	fmt.Println("makeChess = ", makeChess)
	fmt.Println("copiedMap = ", copiedMap)
	
	// makeChess := make(map[string]int, 6)
	// makeChess = map[string]int{
	// 	"a1": 1,
	// 	"b1": 2,
	// 	"c1": 3,
	// 	"c4": 4,
	// 	"c5": 5,
	// 	"c6": 6,
	// 	"c7": 6,
	// 	"c8": 6,
	// 	"c9": 6,
	// 	"c10": 6,
	// }
	// though overflowing the size works but
	// An effective hashing function requires the size of the hash table to be at least twice the number of elements. Reserving the size beforehand preempts the program initialising a larger map to ensure hashing efficiency.
}