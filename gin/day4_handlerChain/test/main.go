package main

import (
	"learn_Go/gin/day4_handlerChain/test/gee"
	"net/http"
)

//如何写函数体
//ctx如何实例话？

func main() {
	e := gee.New()
	e.Get("/conclusion", func(c *gee.Context) {
		c.RespByJSON(200, "map[key]value")
	})
	v := e.Group("/hello")
	v1 := v.Group("/v1")
	v1.Get("/:name", func(c *gee.Context) {
		// expect /hello/geektutu
		c.RespByString(http.StatusOK, "hello %s, you're at %s\n", c.Param("name"), c.Path)
	})
	e.POST("/testPost", func(c *gee.Context) {
		c.RespByByte(200, []byte("hello world"))
	})

	e.Get("/assets/*", func(c *gee.Context) {
		c.RespByJSON(http.StatusOK, gee.H{"filepath": c.Param("")})
	})

	e.Run(":8000")
}
