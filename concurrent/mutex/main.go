package main

import (
	"fmt"
	"sync"
)

func main() {
	var count = 0
	var mu sync.Mutex
	// sync.WaitGroup： 等待10个goroutine完成
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				mu.Lock()
				count++
				mu.Unlock()
			}
		}()
	}
	//等待10个goroutine完成
	wg.Wait()
	fmt.Println(count)
}
