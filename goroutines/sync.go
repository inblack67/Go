package goroutinesn

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}
var count = 0
var mute = sync.RWMutex{}	// r/w => one can write multi can read but writer can write only after all readers are done

func hello(){
	// mute.RLock() // read lock
	fmt.Println("hello = ", count) // if above line exec => random res => no synch our routines
	mute.RUnlock()
	wg.Done()
}

func increment(){
	// mute.Lock() // write lock
	count++
	mute.Unlock()
	wg.Done()
}

func main(){
	// 20 go routines
	for i := 0; i < 20; i++ {
		wg.Add(2)
		mute.RLock()
		go hello()
		mute.Lock()
		go increment()	
	}
	wg.Wait()
}