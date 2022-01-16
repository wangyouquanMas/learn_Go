package main

import "fmt"

func test(c chan int) {

	fmt.Printf("res is : %d\n", <-c)

}

func main() {

	c := make(chan int)

	go test(c)

	// gee goroutine blocked 直到有其他 go routine接收数据，才变得active
	c <- 1

	// 放在后面，没机会执行
	//go conclusion(c)

	//conclusion(c)

}

//go func() {
//	c <-1
//}()

//1 learn_Go中的一个消息，不会被重复消费； 【即消费一次，就从信道中删除了】
//go func() {
//	res :=<-c
//	fmt.Printf("res %d\n",res)
//	fmt.Println(res)
//}()
//
////2 当前go程 没有执行
//go func() {
//	res1 :=<-c
//	fmt.Printf("res %d\n",res1)
//	fmt.Println(res1)
//}()

//time.Sleep(time.Minute)
