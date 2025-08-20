package main

import (
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

const numRequests = 10000

var count int32

func networkRequest() {
	time.Sleep(time.Millisecond) // Эмуляция сетевого запроса
	atomic.AddInt32(&count, 1)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(10000)

	for i := 0; i < numRequests; i++ {
		go func(n int) {
			defer wg.Done()
			networkRequest()
		}(i)
	}

	wg.Wait()
	fmt.Println(count)
}
