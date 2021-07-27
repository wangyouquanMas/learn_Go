package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	"net/url"
)

//解析request

//URL *url.URL : 请求行中url

func main() {
	h1 := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #1!\n")
		fmt.Fprintln(w, r.URL.RawQuery)
		fmt.Fprintln(w, r.URL.Host)
		fmt.Fprintln(w, r.URL.Path)
		fmt.Fprintln(w, "通过handleFunc启动一个http服务")
		rawQuery := r.URL.RawQuery
		//ParseQuery :解析请求返回 map 类型 的key-value对
		va, _ := url.ParseQuery(rawQuery)
		fmt.Fprintln(w, "获取url中key为name值", va.Get("name"))
	}
	h2 := func(w http.ResponseWriter, r *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #2!\n")
		fmt.Fprintln(w, r.URL.RawQuery)
		fmt.Fprintln(w, r.URL.Host)
		fmt.Fprintln(w, r.URL.Path)
		fmt.Fprintln(w, "通过handleFunc启动一个http服务")
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/endpoint", h2)
	//ListenAndServe listens on the TCP network address addr and then calls Serve with handler to handle requests on incoming connections.
	// handler设置为nil，默认为多路处理器作为handler，默认多路处理器实现了ServerHttp方法
	log.Fatal(http.ListenAndServe(":8080", nil))
}
