package main

import (
	"fmt"
	"reflect"
)

// Person is
type Person struct{
	Name string	`json:"name" gorm:"default: null"`	
	Email string
	Username string
	Password string
}

func main(){
	structType := reflect.TypeOf(Person{})
	fmt.Println(structType)
	field, _ := structType.FieldByName("Name")
	tags := field.Tag
	fmt.Println(tags)
}
