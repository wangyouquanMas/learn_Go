package main

import (
	"fmt"
	"net/http"
)

func test(w http.ResponseWriter, r *http.Request) {
	//CanonicalHeaderKey 用于将Header 中key-value对中 key首字母和 - 后字母大写
	res := http.CanonicalHeaderKey("accept-encoding")
	fmt.Println(res)
	//Error replies to the request with the specified error message and HTTP code
	http.Error(w, "无法发现页面", 404)
}

func main() {
	http.HandleFunc("/", test)
	http.ListenAndServe(":8080", nil)
}
