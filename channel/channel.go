package main

import (
	"fmt"
	"time"
)


// channel 特点 可以实现goroutines同步  ，因为接收端和发送断会阻塞 在另一段没有ready前

func worker(id int , c chan int){


	for n := range c { // 用range来判断 通道是否结束，避免close后，一直接收空数据
		//fmt.Printf("Worker %d received %d\n",id ,<-c)
		fmt.Printf("Worker %d received %d\n", id, n)
	}

}

func createWorker(id int ) chan<- int {

	c := make(chan int)

	go func() {
		for{
			fmt.Printf("Worker %d received %c\n",id ,<-c)
		}
	}()

	return c

}


func chanDemo(){
	//var a chan int  //channel 声明

	var channels [10]chan<- int  // <-表示当前channel只能收
	for i :=0;i<10;i++{
		//channels[i] = make(chan int)
		//go worker(i,channels[i])
		channels[i] = createWorker(i)
	}

	for i:=0 ; i<10 ; i++{
		channels[i] <- 'a'+i
	}

	for i:=0 ; i<10 ; i++{
		channels[i] <- 'A'+i
	}
	//n := <-c
	time.Sleep(time.Minute)
}

func bufferedChannel(){
	//c := make(chan int)
	//
	//c <-1   //发数据必须有人收，但是可以加入缓冲区，就不用有人收

	c1 := make(chan int ,3)

	go worker(0,c1)

	c1 <-1
	c1 <-2
	c1 <-3
	c1 <-4
	c1 <-5

	time.Sleep(time.Minute)

}

// 发送方close
func channelClose(){
	c1 := make(chan int ,3)

	go worker(0,c1)

	c1 <-1
	c1 <-2
	c1 <-3
	c1 <-4
	c1 <-5

	close(c1)


}


func main(){
	//chanDemo()
	//bufferedChannel()
	channelClose()
}
