

package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)


var (
	counter int64
	wg sync.WaitGroup
)

func incCounter(id int){
	defer wg.Done()

	for count := 0; count < 2; count ++ {
		//value := counter
		//runtime.Gosched()

		//value ++

		//counter = value
		atomic.AddInt64(&counter, 1)
		runtime.Gosched()
	}
}


func main(){

	wg.Add(2)

	go incCounter(1)
	go incCounter(2)

	wg.Wait()

	fmt.Println("Final counter:", counter)
}