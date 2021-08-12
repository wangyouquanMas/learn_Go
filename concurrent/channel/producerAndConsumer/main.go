package main

import (
	"fmt"
	"time"
)

func calculateNextInt(prev int) int {
	time.Sleep(1 * time.Second)
	return prev + 1
}

func producer(data chan int) {

	time.Sleep(1 * time.Second)

	var i = 0
	for {
		i = calculateNextInt(i)
		fmt.Printf("producer produce s=%v\n", i)
		data <- i
	}
}

//使用 range 来操作 channel 的时候,一旦 channel 关闭，channel 内部数据读完之后循环自动结束。
func consumer(data chan int) {
	for s := range data {
		fmt.Printf("consumer consume s=%v\n", s)
	}
}

func main() {
	data := make(chan int)

	go producer(data)
	go consumer(data)

	time.Sleep(5 * time.Second)
}

//
//func main() {
//	data := make(chan int)
//
//	// producer
//	go func() {
//		var i = 0
//		for {
//			i = calculateNextInt(i)
//			data <- i
//		}
//	}()
//
//	// consumer
//	for i := range data {
//		fmt.Printf("i=%v\n", i)
//	}
//}

/*

	Q1 chan can be range ?
		Go提供了range关键字，将其使用在channel上时，会自动等待channel的动作一直到channel被关闭

	Q2


*/
