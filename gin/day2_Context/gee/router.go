package gee

import (
	"net/http"
)

/*
	1 功能介绍
		addRoute 实现了路由映射表；提供了用户静态注册能力【用户使用http请求url时，会调用注册函数，实现路由注册】

    将路由部分单独提取出来
*/
type router struct {
	handlers map[string]HandlerFunc
}

func newRouter() *router {
	return &router{handlers: make(map[string]HandlerFunc)}
}

func (router *router) addRoute(method, pattern string, handler HandlerFunc) {
	key := method + "-" + pattern
	router.handlers[key] = handler
}

func (engine *router) handle(c *Context) {
	key := c.Method + "-" + c.Path
	if handler, ok := engine.handlers[key]; ok {
		handler(c)
	} else {
		c.String(http.StatusNotFound, "404 NOT FOUND : %s\n", c.Path)
	}
}
