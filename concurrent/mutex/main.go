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

type Foo struct {
	mu    sync.Mutex
	count int
}

//3 通过defer可以实现 Lock/Unlock成对出现,不会遗漏或多调用
func (f *Foo) Bar() {
	f.mu.Lock()
	defer f.mu.Unlock()

	if f.count < 1000 {
		f.count += 3
		return
	}

	f.count++
	return
}

func cas(val *int32, old, new int32) bool
func semacquire(*int32)
func semrelease(*int32)

//互斥锁的结构，包含两个字段
type Mutex struct {
	key  int32 //锁是否被持有的标识
	sema int32 //信号量专用，用以阻塞/唤醒goroutine
}

//保证成功在val上增加delta值
func xadd(val *int32, delta int32) (new int32) {
	for {
		v := *val
		if cas(val, v, v+delta) {
			return v + delta
		}
	}
	panic("unreached")
}

//请求锁
func (m *Mutex) Lock() {
	if xadd(&m.key, 1) == 1 { //
		return
	}
}
