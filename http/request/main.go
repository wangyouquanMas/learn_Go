package main

import (
	"channel/http/request/handler"
	"net/http"
)

//request struct 包含字段的使用

func main() {

	http.HandleFunc("/net/http/test", handler.UploadFile)

	http.ListenAndServe(":8080", nil)
}
