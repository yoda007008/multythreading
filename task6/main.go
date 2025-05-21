package main

import (
	"context"
	"fmt"
	"time"
)

func StartBatchProcessor(ctx context.Context, input <-chan int) {
	const maxBatchSize = 5
	const batchTimeout = 2 * time.Second

	batch := make([]int, 0, maxBatchSize)
	timer := time.NewTimer(batchTimeout)

	for {
		select {
		case <-ctx.Done():
			fmt.Println("Context canceled")
			return
		case val := <-input:
			batch = append(batch, val)
			if len(batch) == maxBatchSize {
				fmt.Println("Processed batch")
				batch = batch[:0]
				if !timer.Stop() {
					<-timer.C
				}
				timer.Reset(batchTimeout)
			}
		}
	}
}

func main() {
	input := make(chan int)
	ctx, cancel := context.WithTimeout(context.Background(), 10*time.Second)
	defer cancel()

	go StartBatchProcessor(ctx, input)

	go func() {
		for i := 0; i < 20; i++ {
			input <- i
			time.Sleep(300 * time.Millisecond)
		}
	}()

	<-ctx.Done()
	fmt.Println("Function main stopped")

}
