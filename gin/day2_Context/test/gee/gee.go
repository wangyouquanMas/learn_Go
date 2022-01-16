package gee

import "net/http"

type Engine struct {
	router *Router
}

func New() *Engine {
	return &Engine{router: newRouter()}
}

//封装原则1
func (e *Engine) AddRouter(method, path string, handler Handler) {
	e.router.addRouter(method, path, handler)
}

func (e *Engine) Get(path string, handler Handler) {
	e.AddRouter("GET", path, handler)
}

func (e *Engine) POST(path string, handler Handler) {
	e.AddRouter("POST", path, handler)
}

//engine实现了handler接口中的servehttp，所以也是handler
func (e *Engine) Run(addr string) error {
	return http.ListenAndServe(addr, e)
}

func (e *Engine) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	context := NewContext(w, r)
	e.router.handle(context)
}
