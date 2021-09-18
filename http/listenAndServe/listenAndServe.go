package main

import (
	"fmt"
	"net/http"
)

/*
	1 功能介绍
      listenAndServe ：
		监听端口，接收到请求后，将连接交给实现了Handler接口的 handler进行处理

	2 拓展
		可以作为统一的http请求控制入口，自由定义路由映射规则，统一添加处理逻辑。
*/

type Engine struct {
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	switch r.URL.Path {
	case "/":
		fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)

	case "/hello":
		for k, v := range r.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	default:
		fmt.Fprintf(w, "404 NOT FOUND: %s\n", r.URL)
	}
}

func main() {
	engine := new(Engine)
	http.ListenAndServe(":9000", engine)
}
