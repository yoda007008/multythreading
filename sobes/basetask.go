package main

import (
	"fmt"
	"sync"
)

var counter int

func main() {
	ch := make(chan int)
	wg := sync.WaitGroup{}
	mu := sync.Mutex{}

	for i := 0; i < 3; i++ {
		wg.Add(1)
		go func(i int) {
			defer wg.Done()
			ch <- i
			mu.Lock()
			counter++
			mu.Unlock()
		}(i)
	}

	go func() {
		wg.Wait()
		close(ch)
	}()

	for val := range ch {
		fmt.Println("Got:", val)
	}

	fmt.Println("Final counter:", counter)
}
