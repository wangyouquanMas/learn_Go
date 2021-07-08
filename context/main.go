package main

import (
	"context"
	"fmt"
	"time"
)

//实现功能 ： 用来设置截止日期、同步信号，传递请求相关值的结构体。
//包含4个方法

type Context interface {

	// 返回 context.Context 被取消的时间，也就是完成工作的截止日期；
	Deadline() (deadline time.Time, ok bool)
	//Done — 返回一个 Channel，这个 Channel 会在当前工作完成或者上下文被取消后关闭，多次调用 Done 方法会返回同一个 Channel；
	Done() <-chan struct{}
	//返回 context.Context 结束的原因，它只会在 Done 方法对应的 Channel 关闭时返回非空的值；
	//如果 context.Context 被取消，会返回 Canceled 错误；
	//如果 context.Context 超时，会返回 DeadlineExceeded 错误；
	Err() error
	//从 context.Context 中获取键对应的值，对于同一个上下文来说，多次调用 Value 并传入相同的 Key 会返回相同的结果，该方法可以用来传递请求特定的数据；
	Value(key interface{}) interface{}
}

// context 同步信号功能 ：是在不同 Goroutine 之间同步请求特定数据、取消信号以及处理请求的截止日期。
func main() {
	//创建1s过期的context
ctx, cancel := context.WithTimeout(context.Background(), 1*time.Second)
defer cancel()

    //goroutine1
go handle(ctx, 500*time.Millisecond)

   //goroutine2  [main 协程]
select {

//once the Done channel is closed, the case <-ctx.Done(): is selected.
case <-ctx.Done():
fmt.Println("main", ctx.Err())
}
}

func handle(ctx context.Context, duration time.Duration) {
	select {
	case <-ctx.Done():
		fmt.Println("handle", ctx.Err())
	case <-time.After(duration):
		fmt.Println("process request with", duration)
	}
}