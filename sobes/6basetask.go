package main

import (
	"fmt"
	"sync"
)

func main() {
	ch1 := make(chan int, 10)
	ch2 := make(chan int, 20)

	ch1 <- 1
	ch2 <- 2
	ch2 <- 4

	close(ch1)
	close(ch2)

	ch3 := asyncMerge[int](ch1, ch2)

	for val := range ch3 {
		fmt.Println(val)
	}
}

// todo async Merge
func asyncMerge[T any](chans ...chan T) chan T {
	resch := make(chan T)
	wg := &sync.WaitGroup{}

	for _, ch := range chans {
		wg.Add(1)

		ch := ch // захватываем текущий канал в замыкании

		go func() {
			defer wg.Done() // гарантируем вызов даже при панике

			for a := range ch {
				resch <- a // отправляем значение, а не канал
			}
		}()
	}

	go func() {
		wg.Wait()
		close(resch) // закрываем результат после завершения всех горутин
	}()

	return resch
}
