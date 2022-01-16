package main

import (
	"learn_Go/gin/day2_Context/test/gee"
)

//如何写函数体
//ctx如何实例话？

func main() {
	e := gee.New()
	e.Get("/conclusion", func(c *gee.Context) {
		c.RespByJSON(200, "map[key]value")
	})
	e.POST("/testPost", func(c *gee.Context) {
		c.RespByByte(200, []byte("hello world"))
	})
	e.Run(":8000")
}
