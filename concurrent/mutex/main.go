package main

import (
	"fmt"
	"sync"
)

func main() {
	var count = 0
	var counter Counter
	var mu sync.Mutex
	// 1 sync.WaitGroup： 等待10个goroutine完成
	var wg sync.WaitGroup
	wg.Add(10)

	for i := 0; i < 10; i++ {
		go func() {
			defer wg.Done()
			for j := 0; j < 100000; j++ {
				counter.Incr()
				counter.Lock()
				counter.count++
				mu.Lock()
				count++
				mu.Unlock()
				counter.Unlock()
			}
		}()
	}
	//等待10个goroutine完成
	wg.Wait()
	fmt.Println(count)
}

//2 嵌入字段可以直接在struct上调用lock
type Counter struct {
	//一般Mutex要控制的字段，放在Mutex后面
	sync.Mutex
	count uint64
}

//也可以将加锁过程封装成方法
func (c *Counter) Incr() {
	c.Lock()
	c.count++
	c.Unlock()
}

//3
