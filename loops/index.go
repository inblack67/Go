package main

import "fmt"

func main(){
	for i := 0; i < 10; i++ {
		fmt.Println(i)		// 0-9
	}

	// fizbuzz
	for i := 1; i < 100; i++ {
		if(i % 3 == 0 && i % 5 == 0){
			fmt.Println("fizbuzz")
		}else if(i % 3 == 0){
			fmt.Println("fizz")
		}else if(i % 5 == 0){
			fmt.Println("buzz")
		}
	}
}
