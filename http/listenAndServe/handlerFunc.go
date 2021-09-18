package main

import (
	"fmt"
	"net/http"
)

/*
	http.ListenAndServe 的handler参数 会自己去执行实现的serveHTTP函数
*/

type HandlerFunc func(w http.ResponseWriter, r *http.Request)

type Engine struct {
	router map[string]HandlerFunc
}

//错误记录： 刚才初始化直接使用 new(Engine) ，返回的是type的 nil值，导致engine.router[key]异常；
// 所以应该手动实例化
func New() *Engine {
	return &Engine{router: make(map[string]HandlerFunc)}
}

func (engine *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {

	fmt.Println("ServeHttp")

	key := r.Method + "-" + r.URL.String()
	if handler, ok := engine.router[key]; ok {
		handler(w, r)
	} else {
		fmt.Println("wrong key")
	}

}

func (engine *Engine) addRouter(method, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	engine.router[key] = handler
}

func (engine *Engine) get(method, pattern string, handler HandlerFunc) {
	engine.addRouter(method, pattern, handler)
}

func main() {

	engine := New()

	engine.get("GET", "/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Println("engine.get func")
	})

	http.ListenAndServe(":8080", engine)
}
