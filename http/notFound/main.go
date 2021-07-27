package main

import "net/http"

func ServeHTTP(w http.ResponseWriter, r *http.Request) {
	http.NotFound(w, r)
}
func main() {

	http.HandleFunc("/", ServeHTTP)
	http.ListenAndServe(":8080", nil)
}
