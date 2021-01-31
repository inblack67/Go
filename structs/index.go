package main

import (
	"fmt"
	"strconv"
)

// Person is
type Person struct{
	name string
	email string
	username string
	password string
}

// type person struct{
// 	name string
// 	email string
// 	username string
// 	password string
// }

// shorthand
type person struct{
	name, email, password, username string
}

// similar to this
func (p person) greet() string{
	return "hello " + p.username
}

// changes defaults
func (p *person) own() string{
	p.username = "owned"
	return "your username is " +  p.username
}

func main(){
	p1 := person{name:"test",email:"test@test.com", username:"someuser",password:"123"}

	fmt.Println(p1.email)
	p1.name = "owned"
	fmt.Println(p1.name)

	greetRes := p1.greet()
	fmt.Println(greetRes)

	fmt.Println(strconv.Itoa(2))	// to string

	fmt.Println(p1.own())
}
