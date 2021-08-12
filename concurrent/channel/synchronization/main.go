package main

import "time"

/*
  A1 ： chan synchronization


*/

func mission(done chan bool) {
	time.Sleep(time.Second)
	//通知任务完成
	done <- true
}

func main() {
	done := make(chan bool, 1)
	go mission(done)
	//等待mission任务完成
	<-done
}
