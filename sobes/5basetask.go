package main

import (
	"fmt"
	"net/http"
	"sync"
)

func main() {
	googleURL := "https://www.google.com" // todo сайт гугл
	avitoURL := "https://www.avito.ru"    // todo сайт авито
	wg := &sync.WaitGroup{}
	wg.Add(2)

	go func() {
		defer wg.Done()
		resp, err := http.Get(googleURL)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Google URL status%d\n:", resp.Status)
	}()

	go func() {
		defer wg.Done()
		resp, err := http.Get(avitoURL)
		if err != nil {
			fmt.Println(err)
			return
		}
		fmt.Printf("Avito URL status%d\n:", resp.Status)
	}()

	wg.Wait()
}
