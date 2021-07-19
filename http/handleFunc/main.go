package main

import (
	"io"
	"log"
	"net/http"
)

/*HandleFunc：pattern对应的hanler是一个function*/

func main() {
	h1 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #1!\n")
	}
	h2 := func(w http.ResponseWriter, _ *http.Request) {
		io.WriteString(w, "Hello from a HandleFunc #2!\n")
	}

	http.HandleFunc("/", h1)
	http.HandleFunc("/endpoint", h2)
	//ListenAndServe listens on the TCP network address addr and then calls Serve with handler to handle requests on incoming connections.
	log.Fatal(http.ListenAndServe(":8080", nil))
}

