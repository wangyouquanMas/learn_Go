package gee

import "net/http"

type Handler func(ctx *Context)

type Router struct {
	mapPathHandlers map[string]Handler
}

//包级别引用
func newRouter() *Router {
	return &Router{
		mapPathHandlers: make(map[string]Handler, 0),
	}
}

//路由映射
func (r *Router) addRouter(method, url string, handler Handler) {
	key := method + "-" + url
	r.mapPathHandlers[key] = handler
}

//这个业务handler就是开发者自己预先写好了的
func (r *Router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := r.mapPathHandlers[key]; ok {
		handler(c)
	} else {
		c.RespByString(http.StatusNotFound, "404 NOT FOUND :%s\n", c.Path)
	}
}
