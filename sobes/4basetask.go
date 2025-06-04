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
	ch := make(chan int, 1)
	wg := &sync.WaitGroup{}
	id := 1

	for i := 0; i < 100000; i++ {
		ch <- i
	}
	close(ch)

	for i := 0; i < 300; i++ {
		wg.Add(1)
		go func() {
			defer wg.Done()
			resp := getPurchasesByID(url, id)
			fmt.Println(resp)
		}()
	}
	wg.Wait()
}
