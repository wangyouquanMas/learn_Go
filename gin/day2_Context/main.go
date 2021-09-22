package main

import (
	"learn_Go/gin/day2_Context/gee"
	"net/http"
)

/*
	功能 ：
		1 将路由独立出来，方便扩展
		2 context封装 net/http ,

	优势： context封装 增强复用性
		1：原先 header 、 编解码都需要自己编写

*/
func main() {
	r := gee.New()
	r.GET("/", func(c *gee.Context) {
		c.HTML(http.StatusOK, "<h1>Hello Gee</h1>")
	})
	r.GET("/hello", func(c *gee.Context) {
		// expect /hello?name=geektutu
		c.String(http.StatusOK, "hello %s, you're at %s\n", c.Query("name"), c.Path)
	})

	r.POST("/login", func(c *gee.Context) {
		c.JSON(http.StatusOK, gee.H{
			"username": c.PostForm("username"),
			"password": c.PostForm("password"),
		})
	})

	r.Run(":9999")
}
