package goroutines 

import (
	"fmt"
	"sync"
)

var wg2 = sync.WaitGroup{}
var count = 0

func hello(){
	fmt.Println("hello = ", count) // random res => no synch our routines
	wg.Done()
}

func increment(){
	count++
	wg.Done()
}

func one(){
	// 20 go routines
	for i := 0; i < 20; i++ {
		wg2.Add(2)
		go hello()
		go increment()	
	}
	wg.Wait()
}