package main

import (
	"fmt"
	"net/http"
)

//ParseHTTPVersion parses an HTTP version string. "HTTP/1.0" returns (1, 0, true).

func main() {

	major, minor, ok := http.ParseHTTPVersion("HTTP/1.1")
	fmt.Println(major, minor, ok)

	//http.ListenAndServe(":8080", nil)
}
