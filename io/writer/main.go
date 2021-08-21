package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"sync"
)

// write接口将 buffer内容写入到文件流 {只要实现了io.writer接口就可以}中

func main() {

	file, _ := os.Create("test.txt")
	fmt.Fprintln(file, "ABC")

	http.HandleFunc("/endpoint", ServeHTTP)
	http.Handle("/test", new(countHandler))
	//ListenAndServe listens on the TCP network address addr and then calls Serve with handler to handle requests on incoming connections.
	// handler设置为nil，默认为多路处理器作为handler，默认多路处理器实现了ServerHttp方法
	log.Fatal(http.ListenAndServe(":8080", nil))

}

func ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(rw, "Hello, World")
}

func (c *countHandler) ServeHTTP(rw http.ResponseWriter, req *http.Request) {
	fmt.Fprintln(rw, "Hello, World")
}

type countHandler struct {
	mu sync.Mutex
	n  int
}
