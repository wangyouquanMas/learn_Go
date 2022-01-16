package main

import (
	"github.com/pkg/errors"
	"log"
	"net"
	"net/http"
	rpc "net/rpc"
	"time"
)

//定义传入参数
type Args struct {
	A, B int
}

type Quotient struct {
	Quo, Rem int
}

//定义返回对象
type Arith int

//实现这个类型的两个方法

func (t *Arith) Multiply(args *Args, reply *int) error {
	*reply = args.A * args.B
	return nil
}
func (t *Arith) Divide(args *Args, quo *Quotient) error {
	if args.B == 0 {
		return errors.New("divide by zero")
	}
	quo.Quo = args.A / args.B
	quo.Rem = args.A % args.B
	return nil
}

// 实现rpc服务器
func main() {
	//1生成Arith对象
	arith := new(Arith)
	//2注册该服务
	rpc.Register(arith)
	//3通过http暴露出来 ，客户端可以看到服务Arith,以及他的两个方法
	rpc.HandleHTTP()
	l, e := net.Listen("tcp", ":1234")
	if e != nil {
		log.Fatalf("listen error:", e)
	}
	//http.ListenAndServe()

	go http.Serve(l, nil)

	time.Sleep(1 * time.Minute)
}
