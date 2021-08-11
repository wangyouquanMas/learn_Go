package main

import (
	"fmt"
	"golang.org/x/sync/errgroup"
	"log"
	"time"
)

func main() {
	//sync.WaitGroup 使用于协同场景

	//	var wg sync.WaitGroup
	//
	//	for i := 0; i < 100; i++ {
	//		//main协程通过调用 wg.Add(delta int) 设置worker协程的个数，然后创建worker协程；
	//		wg.Add(1)
	//		//创建 worker协程，并启动任务。
	//		go func(i int) {
	//			time.Sleep(2 * time.Second)
	//			fmt.Println("End:", i)
	//			//worker协程执行结束以后，都要调用 wg.Done()；
	//			wg.Done()
	//		}(i)
	//	}
	////main协程调用 wg.Wait() 且被block，直到所有worker协程全部执行结束后返回。
	//	wg.Wait()

	//woker goroutine的执行过程中遇到错误想要通知在检查点等待的协程处理该怎么办呢？WaitGroup并没有提供传播错误的功能。
	//	Go语言在扩展库提供的ErrorGroup并发原语正好适合在这种场景下使用，它在WaitGroup的功能基础上还提供了，错误传播以及上下文取消的功能。

	//errorgroup.Group 提供方法

	//Go方法，接收类型为func() error 的函数作为子任务函数，如果任务执行成功，就返回nil，否则就返回 error，并且会cancel 那个新的Context。
	//func WithContext(ctx context.Context) (*Group, context.Context)

	//Go方法，接收类型为func() error 的函数作为子任务函数，如果任务执行成功，就返回nil，否则就返回 error，并且会cancel 那个新的Context。
	//func (g *Group) Go(f func() error)

	//Wait方法，类似WaitGroup的 Wait 方法，调用后会阻塞地等待所有的子任务都完成，它才会返回。
	//如果有多个子任务返回错误，它只会返回第一个出现的错误，如果所有的子任务都执行成功，就返回nil。
	//func (g *Group) Wait() error

	var eg errgroup.Group
	for i := 0; i < 100; i++ {
		i := i
		eg.Go(func() error {
			time.Sleep(2 * time.Second)
			if i > 90 {
				fmt.Println("Error:", i)
				return fmt.Errorf("Error occurred: %d", i)
			}
			fmt.Println("End:", i)
			return nil
		})
	}
	if err := eg.Wait(); err != nil {
		log.Fatal(err)
	}

}
