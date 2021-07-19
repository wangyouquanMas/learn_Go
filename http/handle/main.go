package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

type countHandler struct{
	mu sync.Mutex
	n int
}

//ServeHTTP: write reply headers and data to ResponseWriter
func (h *countHandler) ServeHTTP(w http.ResponseWriter,r * http.Request){
	h.mu.Lock()
	defer h.mu.Unlock()
	h.n++
	fmt.Fprintf(w,"count is %d\n",h.n)
}


func main(){
	// new 会带动结构体内部元素的初始化？ 被结构体拥有的方法也会被初始化？
	http.Handle("/count",new(countHandler))
	log.Fatal(http.ListenAndServe(":8080",nil))
}