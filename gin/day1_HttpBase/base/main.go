package main

import (
	"fmt"
	"learn_Go/gin/day1_HttpBase/base/gee"
	"net/http"
)

func main() {
	r := gee.New()
	//在访问路径时，就进行路由注册
	r.GET("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
	})

	r.GET("/hello", func(w http.ResponseWriter, r *http.Request) {
		for k, v := range r.Header {
			fmt.Fprintf(w, "Header[%q] = %q\n", k, v)
		}
	})
	//根据路由获取handler执行 serveHttp
	r.Run()
}
