package main

import (
	"net/http"
)

func main() {
	//host := flag.String("host", "127.0.0.1", "listen host")
	//port := flag.String("port", "80", "listen port")

	http.HandleFunc("/hello", Hello)

	err := http.ListenAndServe(":80", nil)

	if err != nil {
		panic(err)
	}
}

func Hello(w http.ResponseWriter, req *http.Request) {
	w.Write([]byte("Hello World"))
}
