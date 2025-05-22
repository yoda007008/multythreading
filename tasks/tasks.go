package fulltasks

import (
	"context"
	"fmt"
	"sync"
	"sync/atomic"
	"time"
)

/*
Задача: Напишите функцию, которая запускает горутину, выполняющую fmt.Println("Hello from goroutine!"),
и использует sync.WaitGroup для ожидания её завершения.
*/
func RunningGoroutine() {
	var wg sync.WaitGroup

	wg.Add(1)

	go func() {
		defer wg.Done()
		fmt.Println("Hello from goroutine")
	}()

	wg.Wait()
}

/*
Задача: Напишите программу, которая запускает 5 горутин, каждая из которых печатает свой номер (от 1 до 5),
и использует sync.WaitGroup для их синхронизации(нужно подождать их выполнения).
*/

func RunFiveGoroutine() {
	var wg sync.WaitGroup

	for i := 0; i < 5; i++ {
		wg.Add(1)
		go func(n int) {
			defer wg.Done()
			fmt.Printf("Hello from goroutine %d\n", n)
		}(i)
	}

	wg.Wait()

}

/*
Задача: Напишите функцию, которая создает горутину,
отправляющую числа от 1 до 5 в канал, а затем в main извлекает их и складывает, результат выводит в консоль.
*/
func GetToChannel(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
}

func Run() {
	ch := make(chan int)

	go GetToChannel(ch)

	sum := 0
	for num := 0; num < 5; num++ {
		sum += num
	}

	fmt.Printf("Сумма всех значений %d\n", sum)
}

/*
Задача: Напишите программу, где 10 горутин инкрементируют один счётчик, защищая его sync.Mutex.
*/
func MutexCounter() {
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

/*
Задача: Напишите программу, где 10 горутин инкрементируют один счётчик без использования мютексов, через атомики.
*/
func AtomicCounter() {
	var counter int64
	var wg sync.WaitGroup

	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func() {
			atomic.AddInt64(&counter, 1)
			wg.Done()
		}()
	}
	wg.Wait()

	fmt.Println(atomic.LoadInt64(&counter))
}

/*
Реализуй функцию StartBatchProcessor(ctx context.Context, input <-chan int), которая:
Собирает числа из канала input в батчи по максимум 5 элементов.
Если в течение 2 секунд батч не собран — обрабатывает то, что есть.
Обработка батча — это просто fmt.Println("Processed batch:", batch).
Выход из функции должен происходить при отмене контекста (ctx.Done()).
*/
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

func Start() {
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
