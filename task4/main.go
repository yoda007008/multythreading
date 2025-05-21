package main

import (
	"fmt"
	"sync"
)

func main() {
	counter := 0
	var mutex sync.Mutex
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			mutex.Lock()
			counter += 1
			mutex.Unlock()
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(counter)
}
