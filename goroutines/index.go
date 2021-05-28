package goroutines

import (
	"fmt"
	"time"
)

func greet(){
	fmt.Println("hello worlds")
}

// go run --race main.go

func main3() {
	go greet()

	var msg = "hello"
	go func ()  {
		fmt.Println(msg)	// can access it because of closures but the output will be => bye => race condition
	}()
	msg = "bye" // this var will be updated first then our above go routing will be called => after the end of main func

	var msg2 = "hello"
	go func (msg2 string)  {
		fmt.Println(msg2)	// hello
	}(msg2)	// copy saved
	msg2 = "bye"

	time.Sleep(100 * time.Millisecond)
}