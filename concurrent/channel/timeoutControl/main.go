package main

import (
	"fmt"
	"time"
)

//select如果没有case需要处理，就会一直阻塞， select可以实现超时控制
func main() {
	chstr := make(chan string, 1)
	go func() {
		time.Sleep(time.Second * 2)
		chstr <- "retuslt"
	}()

	select {
	case res := <-chstr:
		fmt.Println(res)
	case <-time.After(time.Second * 1):
		fmt.Println("timeout")
	}
}
