package main

import (
	"fmt"
	"math"
)

func isPrime (num int) bool {
	if num <= 2 {
		return true;
	}
	
	upto := math.Floor(math.Sqrt(float64(num)))
	for i := 3; i < int(upto); i++ {
		if(num % i == 0){
			return false
		}
	}
	return false;
}

func main (){
	res := isPrime(29355126551);
	fmt.Println("5 isPrime => ", res)
}
