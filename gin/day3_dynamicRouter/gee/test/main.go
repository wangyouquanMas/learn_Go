package main

import (
	"learn_Go/gin/day3_dynamicRouter/gee/test/gee"
	"net/http"
)

//如何写函数体
//ctx如何实例话？

func main() {
	e := gee.New()
	e.Get("/conclusion/", func(c *gee.Context) {
		c.RespByJSON(200, "conclusion")
	})
	e.Get("/conclusion/abc", func(c *gee.Context) {
		c.RespByJSON(200, "conclusion/abc")
	})
	e.Get("/hello/:name", func(c *gee.Context) {
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
