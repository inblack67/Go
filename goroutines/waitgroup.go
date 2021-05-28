package goroutiness

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{} // sync multiple go routines together

func main2() {
	var msg = "hello"
	wg.Add(1)	// added one wait group
	go func (msg string)  {
		fmt.Println(msg)
		wg.Done() // done with this group => decrement
	}(msg)
	msg = "bye"
	wg.Wait()	// exit
}