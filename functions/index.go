package main

import "fmt"

func greet(name string) string {
	return "hello " + name
}

func getSum(num1 int, num2 int) int {
	return num1 + num2
}

func main(){
	greetRes := greet("test")
	fmt.Println(greetRes)
	
	getSumRes := getSum(2,2)
	fmt.Println(getSumRes)
}
