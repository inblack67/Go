package main

import "fmt"

func main(){
	// [keyType]valueType
	emails := make(map[string]int)
	emails["test"] = 1

	fmt.Println(emails)  // map[test:1]
	fmt.Println(len(emails))
	
	delete(emails, "test");
	fmt.Println(emails)
	fmt.Println(len(emails))

	emails2 := map[string]string{"ok":"1", "key":"val"}
	fmt.Println(emails2)
}
