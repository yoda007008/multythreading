package main

import (
	"fmt"
	"sync"
)

func getPurchasesByID(url string, id int) bool {
	// todo простая логика
	if url == "" || id < 1 {
		return false
	}
	return true
}

func main() {
	url := "https://www.wildberries.ru/getPurchasesByID?id=%d"
	wg := &sync.WaitGroup{}
	ch := make(chan int, 1)

	for id := range 100_000 {
		ch <- id
	}
	close(ch)

	for i := 0; i < 300; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			resp := getPurchasesByID(url, 1)
			fmt.Println(resp)
		}()
	}
	wg.Wait()
}
