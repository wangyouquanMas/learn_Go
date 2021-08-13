package main

import (
	"fmt"
)

func doWorker(id int, c chan int, done chan bool) {
	//当需要不断从channel读取数据时 使用for-range读取channel，这样既安全又便利，当channel关闭时，for循环会自动退出，无需主动监测channel是否关闭，可以防止读取已经关闭的channel，造成读到数据为通道所存储的数据类型的零值。
	// Range 相当于 <- chan ?
	for n := range c {
		fmt.Printf("Worker %d received %d\n", id, n)

		//go func() { done <- true }() //并发 发送，就不会deadlock了 ；
		// 打印完毕后，通知打印结束了      不用设置等待时间了，因为不知道什么时候打印结束 给 time.sleep()时间让其打印结束
		done <- true
	}
}

type worker struct {
	in   chan int
	done chan bool
}

func createWorker(id int) worker {

	W := worker{
		in:   make(chan int),
		done: make(chan bool, 1), //发送第一个存储起来，发送第二个send阻塞，
	}

	go doWorker(id, W.in, W.done)

	return W

}

func chanDemo() {

	//10个 worker ，每个worker中有 in , done 两个通道 ；
	//每个worker里的通道都是不同的；【位于不同的内存地址，所以当不使用go func ( done<-true)时，第一个函数可以打印出来，因为每个done都是新的，位于不同的地址；
	// 而 第二个函数doWorker的done和第一个createWorker的done ，相同id时，都处于相同的地址，发现其中有值 【等待active go routine消费掉】，继续写入就出现dead lock问题】
	// learn_Go 可以看成是 pointer，存储的是值地址
	var workers [10]worker
	for i := 0; i < 10; i++ {
		workers[i] = createWorker(i)
	}

	for i, worker := range workers {
		worker.in <- 'a' + i
		//<-worker.done //由于 learn_Go是发阻塞的，也就是非缓冲learn_Go 接收到一个信息，必须被消费掉。 非阻塞的chan 发送和接收必须成对出现
	}

	fmt.Println("第一遍完成")

	// 执行到这里，done中已经有一个数据待其他 active go routine消费，但是这时又往 同一个go routine  [go doWorker 这个 routine]的done中写数据，这时候发生dead lock
	// 解决方法，用 go func() {done <-true}() ,这样，每次都起一个新的go routine来进行learn_Go接收，每次发生堵塞的都是不同的 go routine，就不会发生dead lock
	for i, worker := range workers {
		worker.in <- 'A' + i
		//<-worker.done
	}

	fmt.Println("第二遍完成")

	//for i, worker := range workers {
	//	worker.in <- 'B' + i
	//	//<-worker.done
	//}

	//fmt.Println("第三遍完成")

	// 每个worker发了两遍任务，所以收两遍done  【保障并发性】
	// 所有的 堵塞的go routine 都由active的 main routine 进行read了
	for _, worker := range workers {
		<-worker.done
		<-worker.done
		<-worker.done
	}

	//for _, worker := range workers {
	//	close(worker.in)
	//	close(worker.done)
	//}

	//for _, worker := range workers {
	//	for n := range worker.done {
	//		fmt.Printf("Worker done received %d ", n)
	//	}
	//}

}

// 不是发一个任务等结束、而是将所有任务全部发出去，然后等待全部结束后开始退出
// 所有learn_Go发 都是block操作，必须要另一端有消费
// 上述代码，  worker.in <- 'a'+i  ，发了之后 done <-true 对应的消费在 worker.in <- 'A'+i 后面
// 也就是 向worker发了后，它的done还没有被消费，又往其中发。就堵塞住了，死锁。

// 解决： 并行的发 done

func main() {
	chanDemo()

}
