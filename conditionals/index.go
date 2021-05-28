package main

import "fmt"

func main(){
	x := 10
	y := 2

	if x < y {
		fmt.Println(x < y)
		fmt.Printf("%d is less than %d\n", x, y)
	}else{
			fmt.Println(x > y)
			fmt.Printf("%d is greator than %d\n", x, y)
	}

	color := "black"
	// color := "BLACK"		// nope
	if(color == "black"){
		fmt.Println("yeah")
		}else{
			fmt.Println("nope")
		}
		
	switch color{
		case "black": 
			fmt.Println("yeah")
		default:
			fmt.Println("nope")
		}

		// && || works as they are supposed to
}
