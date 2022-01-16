package main

import (
	"fmt"
	"learn_Go/http/request/handler"
	"net/http"
)

//request struct 包含字段的使用

func main() {
	// 静态资源处理
	http.Handle("/static/",
		http.StripPrefix("/static/",
			http.FileServer(http.Dir("./static"))))
	http.HandleFunc("/net/http/conclusion", handler.UploadFile)

	// 监听端口
	fmt.Println("上传服务正在启动, 监听端口:8080...")
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		fmt.Printf("Failed to start server, err:%s", err.Error())
	}
	//http.ListenAndServe(":8080", nil)
}
