package second

import (
	"encoding/gob"
	"fmt"
	"net"
	"net/http"
	"net/rpc"
)

type Arith int

func (arith *Arith) Mult(a, b int) {
	fmt.Println(a * b)
}

func main() {

	arith := new(Arith)
	rpc.Register(arith)

	//支持http 连接，转为tcp处理请求
	//Server.HandleHTTP｀设置rpc的上下文路径，｀rpc.HandleHTTP｀使用默认的上下文路径｀DefaultRPCPath｀、 DefaultDebugPath。
	//这样，当你启动一个http server的时候 ｀http.ListenAndServe｀，上面设置的上下文将用作RPC传输，这个上下文的请求会教给ServeHTTP来处理。
	rpc.HandleHTTP()

	//直接tcp连接 处理请求
	l, _ := net.Listen("tcp", ":1234")
	go http.Serve(l, nil)
}
