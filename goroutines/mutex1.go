package goroutines 

import (
	"fmt"
	"sync"
)

var wg21 = sync.WaitGroup{}
var count = 0
var mute = sync.RWMutex{}	// r/w => one can write multi can read but writer can write only after all readers are done

func hello(){
	mute.RLock() // write lock
	fmt.Println("hello = ", count) // random res => no synch our routines
	mute.RUnlock()
	wg12.Done()
}

func increment(){
	mute.Lock() // write lock
	count++
	mute.Unlock()
	wg12.Done()
}

func main(){
	// 20 go routines
	for i := 0; i < 20; i++ {
		wg2.Add(2)
		go hello()
		go increment()	
	}
	wg2.Wait()
}