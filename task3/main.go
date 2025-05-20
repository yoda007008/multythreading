package main

import "fmt"

func GetToChannel(ch chan int) {
	for i := 0; i < 5; i++ {
		ch <- i
	}
}

func main() {
	ch := make(chan int)

	go GetToChannel(ch)

	sum := 0
	for num := 0; num < 5; num++ {
		sum += num
	}

	fmt.Printf("Сумма всех значений %d\n", sum)
}
