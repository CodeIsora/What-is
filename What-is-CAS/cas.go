package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

var (
	wg        sync.WaitGroup
	threadNum int
	sharedNum int64
)

func main() {

	threadNum = 1000
	sharedNum = 0

	wg.Add(threadNum)

	for i := 0; i < threadNum; i++ {
		go change(i)
	}

	wg.Wait()
}

func change(index int) {
	defer wg.Done()
	if atomic.CompareAndSwapInt64(&sharedNum, sharedNum, sharedNum+1) {
		// fmt.Printf("%d goroutine change success, num is: %d \n", index, *sharedNum)
	} else {
		fmt.Printf("%d goroutine change failed, num is: %d \n", index, sharedNum)
	}
}
