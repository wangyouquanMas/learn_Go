package main

import (
	"io"
	"log"
	"net/http"
)

///*HandleFunc：pattern对应的hanler是一个function*/
// 将普通函数 转换为 handler

func main() {
	h1 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #1!\n")
	}
	h2 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #2!\n")
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/endpoint", h2)
	//ListenAndServe listens on the TCP network address addr and then calls Serve with handler to handle requests on incoming connections.
	// handler设置为nil，默认为多路处理器作为handler，默认多路处理器实现了ServerHttp方法
	log.Fatal(http.ListenAndServe(":8080", nil))
}
