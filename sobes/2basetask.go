// Находим максимальное четное число
package main

import (
	"fmt"
	"sync"
)

func main() {
	var max int
	var wg sync.WaitGroup
	var mu sync.Mutex
	wg.Add(1000)

	for i := 1000; i > 0; i-- {
		go func() {
			defer wg.Done()
			mu.Lock()
			if i%2 == 0 && i > max {
				max = i
			}
			mu.Unlock()
		}()
	}
	wg.Wait()

	fmt.Printf("Maximum is %d", max)
}
