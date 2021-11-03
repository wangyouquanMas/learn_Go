package main

import (
	"fmt"
	"time"
)

/*
	使用select +  time.after实现超时控制
*/

func main() {

	ch := make(chan string)

	go func() {
		time.Sleep(time.Second * 2)
		ch <- "result"
	}()

	select {
	case res := <-ch:
		fmt.Println(res)
		//time.After(time.Second * 1)：表示time.Duration长的时候后返回一条time.Time类型的通道消息。
	case <-time.After(time.Second * 1):
		fmt.Println("timeout")
	}
}
