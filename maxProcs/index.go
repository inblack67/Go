package maxProcs

import (
	"fmt"
	"runtime"
)

func main(){
	// single threaded
	runtime.GOMAXPROCS(1)	// truly concurrent application with no parallelism at all
	numberOfAvailThreads := runtime.GOMAXPROCS(-1)
	fmt.Println(numberOfAvailThreads)	// 4
} 