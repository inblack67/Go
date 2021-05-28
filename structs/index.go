package structs

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


type extinct struct {
	isExtinct bool
}

type deno struct {
	extinct	// composition => deno has a char of extinct
	carnivore bool
}

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

	// value types => unlike maps and slices => but you can always use & for ref

	st1 := Person{
		name: "somename",
		email: "someemail",
		username: "someusername",
		password: "somepassword",
	}

	st2 := st1

	st2.username = "hacked"

	fmt.Println("st1.username = ", st1.username)
	fmt.Println("st2.username = ", st2.username)
	fmt.Println("st1.username = ", st1.username)

	p1 := person{name:"test",email:"test@test.com", username:"someuser",password:"123"}

	fmt.Println(p1.email)
	p1.name = "owned"
	fmt.Println(p1.name)

	greetRes := p1.greet()
	fmt.Println(greetRes)

	fmt.Println(strconv.Itoa(2))	// to string

	fmt.Println(p1.own())

	// no inheritence in go but composition
	trex := deno{
		extinct: extinct{
			isExtinct: true,
		},
		carnivore: true,
	}
	fmt.Println("trex = ", trex)
}
