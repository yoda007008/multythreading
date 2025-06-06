package main

import (
	"context"
	"fmt"
	"math/rand"
	"sync"
	"time"
)

func main() {
	fmt.Println(Do(context.Background(), []User{{"aaa"}, {"bbb"}, {"ccc"}, {"ddd"}, {"eee"}}))
}

type User struct {
	Name string
}

func fetchByName(ctx context.Context, userName string) (int, error) {
	// Тут происходит сетевой поход, который по userName возвращает userID
	time.Sleep(10 * time.Millisecond) // Имитация сетевого похода
	return rand.Int() % 100000, nil
}

func Do(ctx context.Context, users []User) (map[string]int, error) {
	collected := make(map[string]int)
	ch := make(chan error, 1)
	var wg = &sync.WaitGroup{}
	var mu = &sync.Mutex{}
	ctx, cancel := context.WithCancel(ctx)
	defer cancel()
	for _, u := range users {
		go func(user User) {
			defer wg.Done()
			select {
			case <-ctx.Done():
				return
			default:

			}

			userID, err := fetchByName(ctx, u.Name)
			if err != nil {
				select {
				case ch <- err:
					cancel()
				default:
				}
			}
			mu.Lock()
			collected[u.Name] = userID
			mu.Unlock()
		}(u)
	}

	done := make(chan struct{})
	go func() {
		wg.Wait()
		close(done)
	}()

	select {
	case <-done:
		return collected, nil
	case err := <-ch:
		return collected, err
	case <-ctx.Done():
		return collected, ctx.Err()
	}
}
