package main

import (
	"io"
	"net/http"
	"sync"
)

func FetchURLs(urls []string) map[string]string {
	results := make(map[string]string)
	var wg sync.WaitGroup
	var mu sync.Mutex

	for _, url := range urls {
		wg.Add(1)

		go func(u string) {
			defer wg.Done()
			resp, err := http.Get(u)
			if err != nil {
				mu.Lock()
				results[u] = "error"
				mu.Unlock()
				return
			}
			defer resp.Body.Close()
			body, err := io.ReadAll(resp.Body)
			if err != nil {
				mu.Lock()
				results[u] = "error"
				mu.Unlock()
				return
			}

			bodyStr := string(body)
			if len(bodyStr) > 100 {
				bodyStr = bodyStr[:100]
			}

			mu.Lock()
			results[u] = bodyStr
			mu.Unlock()
		}(url)
	}

	wg.Wait()

	return results
}

func main() {
	adress := []string{
		"https://www.example.com",
		"http://example.org",
		"https://www.google.com",
		"https://www.youtube.com",
	}
	FetchURLs(adress)
}
