package main

import (
	"fmt"
	"log"
	"net/http"
	"sync"
)

//Handle: 需要实现handle接口的handle

type countHandler struct {
	mu sync.Mutex
	n  int
}

//ServeHTTP: write reply headers and data to ResponseWriter
func (h *countHandler) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	h.mu.Lock()
	defer h.mu.Unlock()
	h.n++
	fmt.Fprintf(w, "count is %d\n", h.n)
}

func main() {
	http.Handle("/count", new(countHandler))
	log.Fatal(http.ListenAndServe(":8080", nil))

}
