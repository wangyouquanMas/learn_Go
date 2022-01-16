package main

import (
	"fmt"
	"time"
)

func main() {
	//1 创建多个goroutine时，后面的goroutine不会因为前面的goroutine而被阻塞。
	go DelayPrint() // 第一个goroutine
	go HelloWorld() // 第二个goroutine
	time.Sleep(10 * time.Second)
	fmt.Println("gee func")
}
func DelayPrint() {
	for i := 1; i <= 3; i++ {
		fmt.Println(i)
	}
}
func HelloWorld() {
	fmt.Println("Hello goroutine")
}
