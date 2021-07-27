package main

import (
	"fmt"
	"net/http"

	"github.com/gorilla/mux"
)

//MaxBytesReader : 通过限制请求体大小，来防止有意无意的发送大量请求消耗服务器资源
// n :表示限制body的 byte大小
func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.MaxBytesReader(w, r.Body, 10)
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/handle", maxBytes(PostHandler)).Methods("POST")
	http.ListenAndServe(":8080", r)
}

// Middleware to enforce the maximum post body size
func maxBytes(f http.HandlerFunc) http.HandlerFunc {
	return func(w http.ResponseWriter, r *http.Request) {
		// As an example, limit post body to 10 bytes
		r.Body = http.MaxBytesReader(w, r.Body, 10)
		f(w, r)
	}
}

func PostHandler(w http.ResponseWriter, r *http.Request) {
	// How do I know if the form data has been truncated?
	book := r.FormValue("email")
	fmt.Fprintf(w, "You've requested the book: %s\n", book)
}
